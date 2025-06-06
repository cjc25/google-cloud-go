// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/spanner/apiv1/spannerpb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// KvMeta is the Metadata for mocked KV table.
func KvMeta() *spannerpb.ResultSetMetadata {
	return &spannerpb.ResultSetMetadata{
		RowType: &spannerpb.StructType{
			Fields: []*spannerpb.StructType_Field{
				{
					Name: "Key",
					Type: &spannerpb.Type{Code: spannerpb.TypeCode_STRING},
				},
				{
					Name: "Value",
					Type: &spannerpb.Type{Code: spannerpb.TypeCode_STRING},
				},
			},
		},
	}
}

// StatementResultType indicates the type of result returned by a SQL
// statement.
type StatementResultType int

const (
	// StatementResultError indicates that the sql statement returns an error.
	StatementResultError StatementResultType = 0
	// StatementResultResultSet indicates that the sql statement returns a
	// result set.
	StatementResultResultSet StatementResultType = 1
	// StatementResultUpdateCount indicates that the sql statement returns an
	// update count.
	StatementResultUpdateCount StatementResultType = 2
	// MaxRowsPerPartialResultSet is the maximum number of rows returned in
	// each PartialResultSet. This number is deliberately set to a low value to
	// ensure that most queries return more than one PartialResultSet.
	MaxRowsPerPartialResultSet = 1
)

// The method names that can be used to register execution times and errors.
const (
	MethodBeginTransaction    string = "BEGIN_TRANSACTION"
	MethodCommitTransaction   string = "COMMIT_TRANSACTION"
	MethodBatchCreateSession  string = "BATCH_CREATE_SESSION"
	MethodCreateSession       string = "CREATE_SESSION"
	MethodDeleteSession       string = "DELETE_SESSION"
	MethodGetSession          string = "GET_SESSION"
	MethodExecuteSql          string = "EXECUTE_SQL"
	MethodExecuteStreamingSql string = "EXECUTE_STREAMING_SQL"
	MethodExecuteBatchDml     string = "EXECUTE_BATCH_DML"
	MethodStreamingRead       string = "EXECUTE_STREAMING_READ"
	MethodBatchWrite          string = "BATCH_WRITE"
	MethodPartitionQuery      string = "PARTITION_QUERY"
)

// StatementResult represents a mocked result on the test server. The result is
// either of: a ResultSet, an update count or an error.
type StatementResult struct {
	Type         StatementResultType
	Err          error
	ResultSet    *spannerpb.ResultSet
	UpdateCount  int64
	ResumeTokens [][]byte
	SetLastFlag  bool
}

// PartialResultSetExecutionTime represents execution times and errors that
// should be used when a PartialResult at the specified resume token is to
// be returned.
type PartialResultSetExecutionTime struct {
	ResumeToken   []byte
	ExecutionTime time.Duration
	Err           error
}

// ToPartialResultSets converts a ResultSet to a PartialResultSet. This method
// is used to convert a mocked result to a PartialResultSet when one of the
// streaming methods are called.
func (s *StatementResult) ToPartialResultSets(resumeToken []byte) (result []*spannerpb.PartialResultSet, err error) {
	var startIndex uint64
	if len(resumeToken) > 0 {
		if startIndex, err = DecodeResumeToken(resumeToken); err != nil {
			return nil, err
		}
	}

	totalRows := uint64(len(s.ResultSet.Rows))
	if totalRows > 0 {
		for {
			rowCount := min(totalRows-startIndex, uint64(MaxRowsPerPartialResultSet))
			rows := s.ResultSet.Rows[startIndex : startIndex+rowCount]
			values := make([]*structpb.Value,
				len(rows)*len(s.ResultSet.Metadata.RowType.Fields))
			var idx int
			for _, row := range rows {
				for colIdx := range s.ResultSet.Metadata.RowType.Fields {
					values[idx] = row.Values[colIdx]
					idx++
				}
			}
			var rt []byte
			if len(s.ResumeTokens) == 0 {
				rt = EncodeResumeToken(startIndex + rowCount)
			} else {
				rt = s.ResumeTokens[startIndex]
			}
			startIndex += rowCount
			result = append(result, &spannerpb.PartialResultSet{
				Metadata:    s.ResultSet.Metadata,
				Values:      values,
				ResumeToken: rt,
				// set the last flag only for last PartialResultSet
				Last: s.SetLastFlag && startIndex == totalRows,
			})

			if startIndex == totalRows {
				break
			}
		}
	} else {
		result = append(result, &spannerpb.PartialResultSet{
			Metadata: s.ResultSet.Metadata,
			Last:     s.SetLastFlag,
		})
	}
	return result, nil
}

func min(x, y uint64) uint64 {
	if x > y {
		return y
	}
	return x
}

func (s *StatementResult) updateCountToPartialResultSet(exact bool) *spannerpb.PartialResultSet {
	return &spannerpb.PartialResultSet{
		Stats: s.convertUpdateCountToResultSet(exact).Stats,
	}
}

// Converts an update count to a ResultSet, as DML statements also return the
// update count as the statistics of a ResultSet.
func (s *StatementResult) convertUpdateCountToResultSet(exact bool) *spannerpb.ResultSet {
	rs := &spannerpb.ResultSet{
		Stats: &spannerpb.ResultSetStats{
			RowCount: &spannerpb.ResultSetStats_RowCountLowerBound{
				RowCountLowerBound: s.UpdateCount,
			},
		},
	}
	if exact {
		rs = &spannerpb.ResultSet{
			Stats: &spannerpb.ResultSetStats{
				RowCount: &spannerpb.ResultSetStats_RowCountExact{
					RowCountExact: s.UpdateCount,
				},
			},
		}
	}
	if s.ResultSet != nil && s.ResultSet.Metadata != nil && s.ResultSet.Metadata.Transaction != nil {
		rs.Metadata = &spannerpb.ResultSetMetadata{
			Transaction: s.ResultSet.Metadata.Transaction,
		}
	}
	return rs
}

func (s StatementResult) getResultSetWithTransactionSet(selector *spannerpb.TransactionSelector, tx []byte) *StatementResult {
	res := &StatementResult{
		Type:         s.Type,
		Err:          s.Err,
		UpdateCount:  s.UpdateCount,
		ResumeTokens: s.ResumeTokens,
		SetLastFlag:  s.SetLastFlag,
	}
	if s.ResultSet != nil {
		p, err := deepCopy(s.ResultSet)
		if err != nil {
			// panic here as this should never happen.
			panic(err)
		}
		res.ResultSet = p.(*spannerpb.ResultSet)
	}
	if _, ok := selector.GetSelector().(*spannerpb.TransactionSelector_Begin); ok {
		if res.ResultSet == nil {
			res.ResultSet = &spannerpb.ResultSet{}
		}
		if res.ResultSet.Metadata == nil {
			res.ResultSet.Metadata = &spannerpb.ResultSetMetadata{}
		}
		res.ResultSet.Metadata.Transaction = &spannerpb.Transaction{Id: tx}
	}
	return res
}

func deepCopy(v interface{}) (interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	vptr := reflect.New(reflect.TypeOf(v))
	err = json.Unmarshal(data, vptr.Interface())
	if err != nil {
		return nil, err
	}
	return vptr.Elem().Interface(), err
}

// SimulatedExecutionTime represents the time the execution of a method
// should take, and any errors that should be returned by the method.
type SimulatedExecutionTime struct {
	MinimumExecutionTime time.Duration
	RandomExecutionTime  time.Duration
	Errors               []error
	Responses            []interface{}
	// Keep error after execution. The error will continue to be returned until
	// it is cleared.
	KeepError bool
}

// InMemSpannerServer contains the SpannerServer interface plus a couple
// of specific methods for adding mocked results and resetting the server.
type InMemSpannerServer interface {
	spannerpb.SpannerServer

	// Stops this server.
	Stop()

	// Resets the in-mem server to its default state, deleting all sessions and
	// transactions that have been created on the server. Mocked results are
	// not deleted.
	Reset()

	// Sets an error that will be returned by the next server call. The server
	// call will also automatically clear the error.
	SetError(err error)

	// Puts a mocked result on the server for a specific sql statement. The
	// server does not parse the SQL string in any way, it is merely used as
	// a key to the mocked result. The result will be used for all methods that
	// expect a SQL statement, including (batch) DML methods.
	PutStatementResult(sql string, result *StatementResult) error

	// Puts a mocked result on the server for a specific partition token. The
	// result will only be used for query requests that specify a partition
	// token.
	PutPartitionResult(partitionToken []byte, result *StatementResult) error

	// Adds a PartialResultSetExecutionTime to the server that should be returned
	// for the specified SQL string.
	AddPartialResultSetError(sql string, err PartialResultSetExecutionTime)

	// Removes a mocked result on the server for a specific sql statement.
	RemoveStatementResult(sql string)

	// Aborts the specified transaction . This method can be used to test
	// transaction retry logic.
	AbortTransaction(id []byte)

	// Puts a simulated execution time for one of the Spanner methods.
	PutExecutionTime(method string, executionTime SimulatedExecutionTime)
	// Freeze stalls all requests.
	Freeze()
	// Unfreeze restores processing requests.
	Unfreeze()

	TotalSessionsCreated() uint
	TotalSessionsDeleted() uint
	SetMaxSessionsReturnedByServerPerBatchRequest(sessionCount int32)
	SetMaxSessionsReturnedByServerInTotal(sessionCount int32)

	ReceivedRequests() chan interface{}
	DumpSessions() map[string]bool
	ClearPings()
	DumpPings() []string
}

type inMemSpannerServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	spannerpb.SpannerServer

	mu sync.Mutex
	// Set to true when this server been stopped. This is the end state of a
	// server, a stopped server cannot be restarted.
	stopped bool
	// If set, all calls return this error.
	err error
	// The mock server creates session IDs using this counter.
	sessionCounter uint64
	// The sessions that have been created on this mock server.
	sessions map[string]*spannerpb.Session
	// Last use times per session.
	sessionLastUseTime map[string]time.Time
	// The mock server creates transaction IDs per session using these
	// counters.
	transactionCounters map[string]*uint64
	// The transactions that have been created on this mock server.
	transactions                          map[string]*spannerpb.Transaction
	multiplexedSessionTransactionsToSeqNo map[string]*atomic.Int32
	// The transactions that have been (manually) aborted on the server.
	abortedTransactions map[string]bool
	// The transactions that are marked as PartitionedDMLTransaction
	partitionedDmlTransactions map[string]bool
	// The mocked results for this server.
	statementResults map[string]*StatementResult
	partitionResults map[string]*StatementResult
	// The simulated execution times per method.
	executionTimes map[string]*SimulatedExecutionTime
	// The simulated errors for partial result sets
	partialResultSetErrors map[string][]*PartialResultSetExecutionTime

	totalSessionsCreated uint
	totalSessionsDeleted uint
	// The maximum number of sessions that will be created per batch request.
	maxSessionsReturnedByServerPerBatchRequest int32
	maxSessionsReturnedByServerInTotal         int32
	receivedRequests                           chan interface{}
	// Session ping history.
	pings []string

	// Server will stall on any requests.
	freezed chan struct{}
}

// NewInMemSpannerServer creates a new in-mem test server.
func NewInMemSpannerServer() InMemSpannerServer {
	res := &inMemSpannerServer{}
	res.initDefaults()
	res.statementResults = make(map[string]*StatementResult)
	res.partitionResults = make(map[string]*StatementResult)
	res.executionTimes = make(map[string]*SimulatedExecutionTime)
	res.partialResultSetErrors = make(map[string][]*PartialResultSetExecutionTime)
	res.receivedRequests = make(chan interface{}, 1000000)
	// Produce a closed channel, so the default action of ready is to not block.
	res.Freeze()
	res.Unfreeze()
	return res
}

func (s *inMemSpannerServer) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stopped = true
	close(s.receivedRequests)
}

// Resets the test server to its initial state, deleting all sessions and
// transactions that have been created on the server. This method will not
// remove mocked results.
func (s *inMemSpannerServer) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	close(s.receivedRequests)
	s.receivedRequests = make(chan interface{}, 1000000)
	s.initDefaults()
}

func (s *inMemSpannerServer) SetError(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.err = err
}

// Registers a mocked result for a SQL statement on the server.
func (s *inMemSpannerServer) PutStatementResult(sql string, result *StatementResult) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.statementResults[sql] = result
	return nil
}

func (s *inMemSpannerServer) RemoveStatementResult(sql string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.statementResults, sql)
}

// Registers a mocked result for a partition token on the server.
func (s *inMemSpannerServer) PutPartitionResult(partitionToken []byte, result *StatementResult) error {
	tokenString := string(partitionToken)
	s.mu.Lock()
	defer s.mu.Unlock()
	s.partitionResults[tokenString] = result
	return nil
}

func (s *inMemSpannerServer) AbortTransaction(id []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.abortedTransactions[string(id)] = true
}

func (s *inMemSpannerServer) PutExecutionTime(method string, executionTime SimulatedExecutionTime) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.executionTimes[method] = &executionTime
}

func (s *inMemSpannerServer) AddPartialResultSetError(sql string, partialResultSetError PartialResultSetExecutionTime) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.partialResultSetErrors[sql] = append(s.partialResultSetErrors[sql], &partialResultSetError)
}

// Freeze stalls all requests.
func (s *inMemSpannerServer) Freeze() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.freezed = make(chan struct{})
}

// Unfreeze restores processing requests.
func (s *inMemSpannerServer) Unfreeze() {
	s.mu.Lock()
	defer s.mu.Unlock()
	close(s.freezed)
}

// ready checks conditions before executing requests
func (s *inMemSpannerServer) ready() {
	s.mu.Lock()
	freezed := s.freezed
	s.mu.Unlock()
	// check if server should be freezed
	<-freezed
}

func (s *inMemSpannerServer) TotalSessionsCreated() uint {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.totalSessionsCreated
}

func (s *inMemSpannerServer) TotalSessionsDeleted() uint {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.totalSessionsDeleted
}

func (s *inMemSpannerServer) SetMaxSessionsReturnedByServerPerBatchRequest(sessionCount int32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.maxSessionsReturnedByServerPerBatchRequest = sessionCount
}

func (s *inMemSpannerServer) SetMaxSessionsReturnedByServerInTotal(sessionCount int32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.maxSessionsReturnedByServerInTotal = sessionCount
}

func (s *inMemSpannerServer) ReceivedRequests() chan interface{} {
	return s.receivedRequests
}

// ClearPings clears the ping history from the server.
func (s *inMemSpannerServer) ClearPings() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pings = nil
}

// DumpPings dumps the ping history.
func (s *inMemSpannerServer) DumpPings() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]string(nil), s.pings...)
}

// DumpSessions dumps the internal session table.
func (s *inMemSpannerServer) DumpSessions() map[string]bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	st := map[string]bool{}
	for s := range s.sessions {
		st[s] = true
	}
	return st
}

func (s *inMemSpannerServer) initDefaults() {
	s.sessionCounter = 0
	s.maxSessionsReturnedByServerPerBatchRequest = 100
	s.sessions = make(map[string]*spannerpb.Session)
	s.sessionLastUseTime = make(map[string]time.Time)
	s.transactions = make(map[string]*spannerpb.Transaction)
	s.multiplexedSessionTransactionsToSeqNo = make(map[string]*atomic.Int32)
	s.abortedTransactions = make(map[string]bool)
	s.partitionedDmlTransactions = make(map[string]bool)
	s.transactionCounters = make(map[string]*uint64)
}

func (s *inMemSpannerServer) getPreCommitToken(transactionID, operation string) *spannerpb.MultiplexedSessionPrecommitToken {
	s.mu.Lock()
	defer s.mu.Unlock()
	sequence, ok := s.multiplexedSessionTransactionsToSeqNo[transactionID]
	if !ok {
		return nil
	}
	return &spannerpb.MultiplexedSessionPrecommitToken{
		SeqNum:         sequence.Add(1),
		PrecommitToken: []byte(fmt.Sprintf("precommit-token-%v-%v", operation, sequence.Load())),
	}
}

func (s *inMemSpannerServer) generateSessionNameLocked(database string, isMultiplexed bool) string {
	s.sessionCounter++
	if isMultiplexed {
		return fmt.Sprintf("%s/sessions/multiplexed-%d", database, s.sessionCounter)
	}
	return fmt.Sprintf("%s/sessions/%d", database, s.sessionCounter)
}

func (s *inMemSpannerServer) findSession(name string) (*spannerpb.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	session := s.sessions[name]
	if session == nil {
		return nil, newSessionNotFoundError(name)
	}
	return session, nil
}

// sessionResourceType is the type name of Spanner sessions.
const sessionResourceType = "type.googleapis.com/google.spanner.v1.Session"

func newSessionNotFoundError(name string) error {
	s := gstatus.Newf(codes.NotFound, "Session not found: Session with id %s not found", name)
	s, _ = s.WithDetails(&errdetails.ResourceInfo{ResourceType: sessionResourceType, ResourceName: name})
	return s.Err()
}

func (s *inMemSpannerServer) updateSessionLastUseTime(session string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessionLastUseTime[session] = time.Now()
}

func getCurrentTimestamp() *timestamppb.Timestamp {
	t := time.Now()
	return &timestamppb.Timestamp{Seconds: t.Unix(), Nanos: int32(t.Nanosecond())}
}

// Gets the transaction id from the transaction selector. If the selector
// specifies that a new transaction should be started, this method will start
// a new transaction and return the id of that transaction.
func (s *inMemSpannerServer) getTransactionID(session *spannerpb.Session, txSelector *spannerpb.TransactionSelector) []byte {
	var res []byte
	if txSelector.GetBegin() != nil {
		// Start a new transaction.
		res = s.beginTransaction(session, txSelector.GetBegin()).Id
	} else if txSelector.GetId() != nil {
		res = txSelector.GetId()
	}
	return res
}

func (s *inMemSpannerServer) generateTransactionName(session string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	counter, ok := s.transactionCounters[session]
	if !ok {
		counter = new(uint64)
		s.transactionCounters[session] = counter
	}
	*counter++
	return fmt.Sprintf("%s/transactions/%d", session, *counter)
}

func (s *inMemSpannerServer) beginTransaction(session *spannerpb.Session, options *spannerpb.TransactionOptions) *spannerpb.Transaction {
	id := s.generateTransactionName(session.Name)
	res := &spannerpb.Transaction{
		Id:            []byte(id),
		ReadTimestamp: getCurrentTimestamp(),
	}
	s.mu.Lock()
	if options.GetReadWrite() != nil && session.Multiplexed {
		s.multiplexedSessionTransactionsToSeqNo[id] = new(atomic.Int32)
	}
	s.transactions[id] = res
	s.partitionedDmlTransactions[id] = options.GetPartitionedDml() != nil
	s.mu.Unlock()
	return res
}

func (s *inMemSpannerServer) getTransactionByID(session *spannerpb.Session, id []byte) (*spannerpb.Transaction, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	tx, ok := s.transactions[string(id)]
	if !ok {
		return nil, gstatus.Error(codes.NotFound, "Transaction not found")
	}
	if !strings.HasPrefix(string(id), session.Name) {
		return nil, gstatus.Error(codes.InvalidArgument, "Transaction was started in a different session.")
	}

	aborted, ok := s.abortedTransactions[string(id)]
	if ok && aborted {
		return nil, newAbortedErrorWithMinimalRetryDelay()
	}
	return tx, nil
}

func newAbortedErrorWithMinimalRetryDelay() error {
	st := gstatus.New(codes.Aborted, "Transaction has been aborted")
	retry := &errdetails.RetryInfo{
		RetryDelay: durationpb.New(time.Nanosecond),
	}
	st, _ = st.WithDetails(retry)
	return st.Err()
}

func (s *inMemSpannerServer) removeTransaction(tx *spannerpb.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.transactions, string(tx.Id))
	delete(s.multiplexedSessionTransactionsToSeqNo, string(tx.Id))
	delete(s.partitionedDmlTransactions, string(tx.Id))
}

func (s *inMemSpannerServer) getPartitionResult(partitionToken []byte) (*StatementResult, error) {
	tokenString := string(partitionToken)
	s.mu.Lock()
	defer s.mu.Unlock()
	result, ok := s.partitionResults[tokenString]
	if !ok {
		return nil, gstatus.Error(codes.Internal, fmt.Sprintf("No result found for partition token %v", tokenString))
	}
	return result, nil
}

func (s *inMemSpannerServer) getStatementResult(sql string) (*StatementResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	result, ok := s.statementResults[sql]
	if !ok {
		return nil, gstatus.Error(codes.Internal, fmt.Sprintf("No result found for statement %v", sql))
	}
	return result, nil
}

func (s *inMemSpannerServer) simulateExecutionTime(method string, req interface{}) (interface{}, error) {
	s.mu.Lock()

	// Check if the server is stopped
	if s.stopped {
		s.mu.Unlock()
		return nil, gstatus.Error(codes.Unavailable, "server has been stopped")
	}

	// Send the request to the receivedRequests channel
	s.receivedRequests <- req
	s.mu.Unlock()
	s.ready()
	s.mu.Lock()

	// Check for a simulated error
	if s.err != nil {
		err := s.err
		s.err = nil
		s.mu.Unlock()
		return nil, err
	}

	// Check for a simulated execution time
	executionTime, ok := s.executionTimes[method]
	s.mu.Unlock()
	if ok {
		var randTime int64
		if executionTime.RandomExecutionTime > 0 {
			randTime = rand.Int63n(int64(executionTime.RandomExecutionTime))
		}
		totalExecutionTime := time.Duration(int64(executionTime.MinimumExecutionTime) + randTime)
		<-time.After(totalExecutionTime)
		s.mu.Lock()

		// Check for errors in the execution time
		if len(executionTime.Errors) > 0 {
			err := executionTime.Errors[0]
			if !executionTime.KeepError {
				executionTime.Errors = executionTime.Errors[1:]
			}
			s.mu.Unlock()
			return nil, err
		}

		// Check for responses in the execution time
		if len(executionTime.Responses) > 0 {
			response := executionTime.Responses[0]
			executionTime.Responses = executionTime.Responses[1:]
			s.mu.Unlock()
			return response, nil
		}
		s.mu.Unlock()
	}

	return nil, nil
}

func (s *inMemSpannerServer) CreateSession(ctx context.Context, req *spannerpb.CreateSessionRequest) (*spannerpb.Session, error) {
	if _, err := s.simulateExecutionTime(MethodCreateSession, req); err != nil {
		return nil, err
	}
	if req.Database == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing database")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.maxSessionsReturnedByServerInTotal > int32(0) && int32(len(s.sessions)) == s.maxSessionsReturnedByServerInTotal {
		return nil, gstatus.Error(codes.ResourceExhausted, "No more sessions available")
	}
	ts := getCurrentTimestamp()
	var (
		creatorRole   string
		isMultiplexed bool
	)
	if req.Session != nil {
		creatorRole = req.Session.CreatorRole
		isMultiplexed = req.Session.Multiplexed
	}
	sessionName := s.generateSessionNameLocked(req.Database, isMultiplexed)
	header := metadata.New(map[string]string{"server-timing": "gfet4t7; dur=123"})
	if err := grpc.SendHeader(ctx, header); err != nil {
		return nil, gstatus.Errorf(codes.Internal, "unable to send 'server-timing' header")
	}
	session := &spannerpb.Session{Name: sessionName, CreateTime: ts, ApproximateLastUseTime: ts, CreatorRole: creatorRole, Multiplexed: isMultiplexed}
	s.totalSessionsCreated++
	s.sessions[sessionName] = session
	return session, nil
}

func (s *inMemSpannerServer) BatchCreateSessions(ctx context.Context, req *spannerpb.BatchCreateSessionsRequest) (*spannerpb.BatchCreateSessionsResponse, error) {
	if _, err := s.simulateExecutionTime(MethodBatchCreateSession, req); err != nil {
		return nil, err
	}
	if req.Database == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing database")
	}
	if req.SessionCount <= 0 {
		return nil, gstatus.Error(codes.InvalidArgument, "Session count must be >= 0")
	}
	sessionsToCreate := req.SessionCount
	s.mu.Lock()
	defer s.mu.Unlock()
	// return a non-retryable error if the limit is reached
	if s.maxSessionsReturnedByServerInTotal > int32(0) && int32(len(s.sessions)) >= s.maxSessionsReturnedByServerInTotal {
		return nil, gstatus.Error(codes.OutOfRange, "No more sessions available")
	}
	if sessionsToCreate > s.maxSessionsReturnedByServerPerBatchRequest {
		sessionsToCreate = s.maxSessionsReturnedByServerPerBatchRequest
	}
	if s.maxSessionsReturnedByServerInTotal > int32(0) && (sessionsToCreate+int32(len(s.sessions))) > s.maxSessionsReturnedByServerInTotal {
		sessionsToCreate = s.maxSessionsReturnedByServerInTotal - int32(len(s.sessions))
	}
	sessions := make([]*spannerpb.Session, sessionsToCreate)
	for i := int32(0); i < sessionsToCreate; i++ {
		sessionName := s.generateSessionNameLocked(req.Database, false)
		ts := getCurrentTimestamp()
		var creatorRole string
		if req.SessionTemplate != nil {
			creatorRole = req.SessionTemplate.CreatorRole
		}
		sessions[i] = &spannerpb.Session{Name: sessionName, CreateTime: ts, ApproximateLastUseTime: ts, CreatorRole: creatorRole}
		s.totalSessionsCreated++
		s.sessions[sessionName] = sessions[i]
	}
	header := metadata.New(map[string]string{"server-timing": "gfet4t7; dur=123"})
	if err := grpc.SendHeader(ctx, header); err != nil {
		return nil, gstatus.Errorf(codes.Internal, "unable to send 'server-timing' header")
	}
	return &spannerpb.BatchCreateSessionsResponse{Session: sessions}, nil
}

func (s *inMemSpannerServer) GetSession(ctx context.Context, req *spannerpb.GetSessionRequest) (*spannerpb.Session, error) {
	if _, err := s.simulateExecutionTime(MethodGetSession, req); err != nil {
		return nil, err
	}
	if req.Name == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Name)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *inMemSpannerServer) ListSessions(ctx context.Context, req *spannerpb.ListSessionsRequest) (*spannerpb.ListSessionsResponse, error) {
	s.mu.Lock()
	if s.stopped {
		s.mu.Unlock()
		return nil, gstatus.Error(codes.Unavailable, "server has been stopped")
	}
	s.receivedRequests <- req
	s.mu.Unlock()
	if req.Database == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing database")
	}
	expectedSessionName := req.Database + "/sessions/"
	var sessions []*spannerpb.Session
	s.mu.Lock()
	for _, session := range s.sessions {
		if strings.Index(session.Name, expectedSessionName) == 0 {
			sessions = append(sessions, session)
		}
	}
	s.mu.Unlock()
	sort.Slice(sessions[:], func(i, j int) bool {
		return sessions[i].Name < sessions[j].Name
	})
	res := &spannerpb.ListSessionsResponse{Sessions: sessions}
	return res, nil
}

func (s *inMemSpannerServer) DeleteSession(ctx context.Context, req *spannerpb.DeleteSessionRequest) (*emptypb.Empty, error) {
	if _, err := s.simulateExecutionTime(MethodDeleteSession, req); err != nil {
		return nil, err
	}
	if req.Name == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	if _, err := s.findSession(req.Name); err != nil {
		return nil, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.totalSessionsDeleted++
	delete(s.sessions, req.Name)
	return &emptypb.Empty{}, nil
}

func (s *inMemSpannerServer) ExecuteSql(ctx context.Context, req *spannerpb.ExecuteSqlRequest) (*spannerpb.ResultSet, error) {
	if _, err := s.simulateExecutionTime(MethodExecuteSql, req); err != nil {
		return nil, err
	}
	if req.Sql == "SELECT 1" {
		s.mu.Lock()
		s.pings = append(s.pings, req.Session)
		s.mu.Unlock()
	}
	if req.Session == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return nil, err
	}
	var id []byte
	s.updateSessionLastUseTime(session.Name)
	if id = s.getTransactionID(session, req.Transaction); id != nil {
		_, err = s.getTransactionByID(session, id)
		if err != nil {
			return nil, err
		}
	}
	var statementResult *StatementResult
	if req.PartitionToken != nil {
		statementResult, err = s.getPartitionResult(req.PartitionToken)
	} else {
		statementResult, err = s.getStatementResult(req.Sql)
	}
	if err != nil {
		return nil, err
	}
	s.mu.Lock()
	statementResult = statementResult.getResultSetWithTransactionSet(req.GetTransaction(), id)
	isPartitionedDml := s.partitionedDmlTransactions[string(id)]
	s.mu.Unlock()
	switch statementResult.Type {
	case StatementResultError:
		return nil, statementResult.Err
	case StatementResultResultSet:

		// if request's session is multiplexed and transaction is Read/Write then add Pre-commit Token in Metadata
		if statementResult.ResultSet != nil {
			statementResult.ResultSet.PrecommitToken = s.getPreCommitToken(string(id), "ResultSetPrecommitToken")
		}
		return statementResult.ResultSet, nil
	case StatementResultUpdateCount:
		res := statementResult.convertUpdateCountToResultSet(!isPartitionedDml)
		res.PrecommitToken = s.getPreCommitToken(string(id), "ResultSetPrecommitToken")
		return res, nil
	}
	return nil, gstatus.Error(codes.Internal, "Unknown result type")
}

func (s *inMemSpannerServer) ExecuteStreamingSql(req *spannerpb.ExecuteSqlRequest, stream spannerpb.Spanner_ExecuteStreamingSqlServer) error {
	if _, err := s.simulateExecutionTime(MethodExecuteStreamingSql, req); err != nil {
		return err
	}
	return s.executeStreamingSQL(req, stream)
}

func (s *inMemSpannerServer) executeStreamingSQL(req *spannerpb.ExecuteSqlRequest, stream spannerpb.Spanner_ExecuteStreamingSqlServer) error {
	if req.Session == "" {
		return gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return err
	}
	s.updateSessionLastUseTime(session.Name)
	var id []byte
	if id = s.getTransactionID(session, req.Transaction); id != nil {
		_, err = s.getTransactionByID(session, id)
		if err != nil {
			return err
		}
	}
	var statementResult *StatementResult
	if req.PartitionToken != nil {
		statementResult, err = s.getPartitionResult(req.PartitionToken)
	} else {
		statementResult, err = s.getStatementResult(req.Sql)
	}
	if err != nil {
		return err
	}
	s.mu.Lock()
	statementResult = statementResult.getResultSetWithTransactionSet(req.GetTransaction(), id)
	isPartitionedDml := s.partitionedDmlTransactions[string(id)]
	s.mu.Unlock()
	switch statementResult.Type {
	case StatementResultError:
		return statementResult.Err
	case StatementResultResultSet:
		parts, err := statementResult.ToPartialResultSets(req.ResumeToken)
		if err != nil {
			return err
		}
		var nextPartialResultSetError *PartialResultSetExecutionTime
		s.mu.Lock()
		pErrors := s.partialResultSetErrors[req.Sql]
		if len(pErrors) > 0 {
			nextPartialResultSetError = pErrors[0]
			s.partialResultSetErrors[req.Sql] = pErrors[1:]
		}
		s.mu.Unlock()
		for _, part := range parts {
			if nextPartialResultSetError != nil && bytes.Equal(part.ResumeToken, nextPartialResultSetError.ResumeToken) {
				if nextPartialResultSetError.ExecutionTime > 0 {
					<-time.After(nextPartialResultSetError.ExecutionTime)
				}
				if nextPartialResultSetError.Err != nil {
					return nextPartialResultSetError.Err
				}
			}
			// For every PartialResultSet, if request's session is multiplexed and transaction is Read/Write then add Pre-commit Token in Metadata
			// and increment the sequence number
			part.PrecommitToken = s.getPreCommitToken(string(id), "PartialResultSetPrecommitToken")
			if err := stream.Send(part); err != nil {
				return err
			}
		}
		return nil
	case StatementResultUpdateCount:
		part := statementResult.updateCountToPartialResultSet(!isPartitionedDml)
		if err := stream.Send(part); err != nil {
			return err
		}
		return nil
	}
	return gstatus.Error(codes.Internal, "Unknown result type")
}

func (s *inMemSpannerServer) ExecuteBatchDml(ctx context.Context, req *spannerpb.ExecuteBatchDmlRequest) (*spannerpb.ExecuteBatchDmlResponse, error) {
	if _, err := s.simulateExecutionTime(MethodExecuteBatchDml, req); err != nil {
		return nil, err
	}
	if req.Session == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return nil, err
	}
	s.updateSessionLastUseTime(session.Name)
	var id []byte
	if id = s.getTransactionID(session, req.Transaction); id != nil {
		_, err = s.getTransactionByID(session, id)
		if err != nil {
			return nil, err
		}
	}
	s.mu.Lock()
	isPartitionedDml := s.partitionedDmlTransactions[string(id)]
	s.mu.Unlock()
	resp := &spannerpb.ExecuteBatchDmlResponse{}
	resp.ResultSets = make([]*spannerpb.ResultSet, len(req.Statements))
	resp.Status = &status.Status{Code: int32(codes.OK)}
	for idx, batchStatement := range req.Statements {
		statementResult, err := s.getStatementResult(batchStatement.Sql)
		if err != nil {
			return nil, err
		}
		s.mu.Lock()
		statementResult = statementResult.getResultSetWithTransactionSet(req.GetTransaction(), id)
		s.mu.Unlock()
		switch statementResult.Type {
		case StatementResultError:
			resp.Status = &status.Status{Code: int32(gstatus.Code(statementResult.Err)), Message: statementResult.Err.Error()}
			resp.ResultSets = resp.ResultSets[:idx]
			return resp, nil
		case StatementResultResultSet:
			return nil, gstatus.Error(codes.InvalidArgument, fmt.Sprintf("Not an update statement: %v", batchStatement.Sql))
		case StatementResultUpdateCount:
			resp.ResultSets[idx] = statementResult.convertUpdateCountToResultSet(!isPartitionedDml)
		}
	}
	resp.PrecommitToken = s.getPreCommitToken(string(id), "ExecuteBatchDmlResponsePrecommitToken")
	return resp, nil
}

func (s *inMemSpannerServer) Read(ctx context.Context, req *spannerpb.ReadRequest) (*spannerpb.ResultSet, error) {
	s.mu.Lock()
	if s.stopped {
		s.mu.Unlock()
		return nil, gstatus.Error(codes.Unavailable, "server has been stopped")
	}
	s.receivedRequests <- req
	s.mu.Unlock()
	header := metadata.New(map[string]string{"server-timing": "gfet4t7; dur=123"})
	if err := grpc.SendHeader(ctx, header); err != nil {
		return nil, gstatus.Errorf(codes.Internal, "unable to send 'server-timing' header")
	}
	return nil, gstatus.Error(codes.Unimplemented, "Method not yet implemented")
}

func (s *inMemSpannerServer) StreamingRead(req *spannerpb.ReadRequest, stream spannerpb.Spanner_StreamingReadServer) error {
	if _, err := s.simulateExecutionTime(MethodStreamingRead, req); err != nil {
		return err
	}
	sqlReq := &spannerpb.ExecuteSqlRequest{
		Session:        req.Session,
		Transaction:    req.Transaction,
		PartitionToken: req.PartitionToken,
		ResumeToken:    req.ResumeToken,
		// KeySet is currently ignored.
		Sql: fmt.Sprintf(
			"SELECT %s FROM %s",
			strings.Join(req.Columns, ", "),
			req.Table,
		),
	}
	header := metadata.New(map[string]string{"server-timing": "gfet4t7; dur=123"})
	if err := grpc.SendHeader(stream.Context(), header); err != nil {
		return gstatus.Errorf(codes.Internal, "unable to send 'server-timing' header")
	}
	return s.executeStreamingSQL(sqlReq, stream)
}

func (s *inMemSpannerServer) BeginTransaction(ctx context.Context, req *spannerpb.BeginTransactionRequest) (*spannerpb.Transaction, error) {
	if _, err := s.simulateExecutionTime(MethodBeginTransaction, req); err != nil {
		return nil, err
	}
	if req.Session == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return nil, err
	}
	s.updateSessionLastUseTime(session.Name)
	tx := s.beginTransaction(session, req.Options)
	if session.Multiplexed && req.MutationKey != nil {
		tx.PrecommitToken = s.getPreCommitToken(string(tx.Id), "TransactionPrecommitToken")
	}
	return tx, nil
}

func (s *inMemSpannerServer) Commit(ctx context.Context, req *spannerpb.CommitRequest) (*spannerpb.CommitResponse, error) {
	mockResponse, err := s.simulateExecutionTime(MethodCommitTransaction, req)
	if err != nil {
		return nil, err
	}
	if req.Session == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return nil, err
	}
	s.updateSessionLastUseTime(session.Name)
	var tx *spannerpb.Transaction
	if req.GetSingleUseTransaction() != nil {
		tx = s.beginTransaction(session, req.GetSingleUseTransaction())
	} else if req.GetTransactionId() != nil {
		tx, err = s.getTransactionByID(session, req.GetTransactionId())
		if err != nil {
			return nil, err
		}
	} else {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing transaction in commit request")
	}
	resp, ok := mockResponse.(*spannerpb.CommitResponse)
	if !ok {
		resp = &spannerpb.CommitResponse{CommitTimestamp: getCurrentTimestamp()}
		s.removeTransaction(tx)
	}
	if req.ReturnCommitStats {
		resp.CommitStats = &spannerpb.CommitResponse_CommitStats{
			MutationCount: int64(1),
		}
	}
	return resp, nil
}

func (s *inMemSpannerServer) Rollback(ctx context.Context, req *spannerpb.RollbackRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	if s.stopped {
		s.mu.Unlock()
		return nil, gstatus.Error(codes.Unavailable, "server has been stopped")
	}
	s.receivedRequests <- req
	s.mu.Unlock()
	if req.Session == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return nil, err
	}
	s.updateSessionLastUseTime(session.Name)
	tx, err := s.getTransactionByID(session, req.TransactionId)
	if err != nil {
		return nil, err
	}
	s.removeTransaction(tx)
	return &emptypb.Empty{}, nil
}

func (s *inMemSpannerServer) PartitionQuery(ctx context.Context, req *spannerpb.PartitionQueryRequest) (*spannerpb.PartitionResponse, error) {
	if _, err := s.simulateExecutionTime(MethodPartitionQuery, req); err != nil {
		return nil, err
	}
	if req.Session == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return nil, err
	}
	var id []byte
	var tx *spannerpb.Transaction
	s.updateSessionLastUseTime(session.Name)
	if id = s.getTransactionID(session, req.Transaction); id != nil {
		tx, err = s.getTransactionByID(session, id)
		if err != nil {
			return nil, err
		}
	}
	var partitions []*spannerpb.Partition
	for i := int64(0); i < req.PartitionOptions.MaxPartitions; i++ {
		token := make([]byte, 10)
		_, err := rand.Read(token)
		if err != nil {
			return nil, gstatus.Error(codes.Internal, "failed to generate random partition token")
		}
		partitions = append(partitions, &spannerpb.Partition{PartitionToken: token})
	}
	return &spannerpb.PartitionResponse{
		Partitions:  partitions,
		Transaction: tx,
	}, nil
}

func (s *inMemSpannerServer) PartitionRead(ctx context.Context, req *spannerpb.PartitionReadRequest) (*spannerpb.PartitionResponse, error) {
	s.mu.Lock()
	if s.stopped {
		s.mu.Unlock()
		return nil, gstatus.Error(codes.Unavailable, "server has been stopped")
	}
	s.receivedRequests <- req
	s.mu.Unlock()
	if req.Session == "" {
		return nil, gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return nil, err
	}
	var id []byte
	var tx *spannerpb.Transaction
	s.updateSessionLastUseTime(session.Name)
	if id = s.getTransactionID(session, req.Transaction); id != nil {
		tx, err = s.getTransactionByID(session, id)
		if err != nil {
			return nil, err
		}
	}
	var partitions []*spannerpb.Partition
	for i := int64(0); i < req.PartitionOptions.MaxPartitions; i++ {
		token := make([]byte, 10)
		_, err := rand.Read(token)
		if err != nil {
			return nil, gstatus.Error(codes.Internal, "failed to generate random partition token")
		}
		partitions = append(partitions, &spannerpb.Partition{PartitionToken: token})
	}
	return &spannerpb.PartitionResponse{
		Partitions:  partitions,
		Transaction: tx,
	}, nil
}

// EncodeResumeToken return mock resume token encoding for an uint64 integer.
func EncodeResumeToken(t uint64) []byte {
	rt := make([]byte, 16)
	binary.PutUvarint(rt, t)
	return rt
}

// DecodeResumeToken decodes a mock resume token into an uint64 integer.
func DecodeResumeToken(t []byte) (uint64, error) {
	s, n := binary.Uvarint(t)
	if n <= 0 {
		return 0, fmt.Errorf("invalid resume token: %v", t)
	}
	return s, nil
}

func (s *inMemSpannerServer) BatchWrite(req *spannerpb.BatchWriteRequest, stream spannerpb.Spanner_BatchWriteServer) error {
	if _, err := s.simulateExecutionTime(MethodBatchWrite, req); err != nil {
		return err
	}
	return s.batchWrite(req, stream)
}

func (s *inMemSpannerServer) batchWrite(req *spannerpb.BatchWriteRequest, stream spannerpb.Spanner_BatchWriteServer) error {
	if req.Session == "" {
		return gstatus.Error(codes.InvalidArgument, "Missing session name")
	}
	session, err := s.findSession(req.Session)
	if err != nil {
		return err
	}
	s.updateSessionLastUseTime(session.Name)
	if len(req.GetMutationGroups()) == 0 {
		return gstatus.Error(codes.InvalidArgument, "No mutations in Batch Write")
	}
	// For each MutationGroup, write a BatchWriteResponse to the response stream
	for idx := range req.GetMutationGroups() {
		res := &spannerpb.BatchWriteResponse{
			Indexes:         []int32{int32(idx)},
			CommitTimestamp: getCurrentTimestamp(),
			Status:          &status.Status{},
		}
		if err = stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}
