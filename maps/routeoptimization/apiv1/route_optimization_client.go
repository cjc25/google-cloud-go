// Copyright 2025 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package routeoptimization

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	routeoptimizationpb "cloud.google.com/go/maps/routeoptimization/apiv1/routeoptimizationpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	OptimizeTours      []gax.CallOption
	BatchOptimizeTours []gax.CallOption
	GetOperation       []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("routeoptimization.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("routeoptimization.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("routeoptimization.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://routeoptimization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		OptimizeTours: []gax.CallOption{
			gax.WithTimeout(3600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchOptimizeTours: []gax.CallOption{},
		GetOperation:       []gax.CallOption{},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		OptimizeTours: []gax.CallOption{
			gax.WithTimeout(3600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		BatchOptimizeTours: []gax.CallOption{},
		GetOperation:       []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Route Optimization API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	OptimizeTours(context.Context, *routeoptimizationpb.OptimizeToursRequest, ...gax.CallOption) (*routeoptimizationpb.OptimizeToursResponse, error)
	BatchOptimizeTours(context.Context, *routeoptimizationpb.BatchOptimizeToursRequest, ...gax.CallOption) (*BatchOptimizeToursOperation, error)
	BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// Client is a client for interacting with Route Optimization API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for optimizing vehicle tours.
//
// Validity of certain types of fields:
//
//	google.protobuf.Timestamp
//
//	  Times are in Unix time: seconds since 1970-01-01T00:00:00+00:00.
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.protobuf.Duration
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.type.LatLng
//
//	  latitude must be in [-90.0, 90.0].
//
//	  longitude must be in [-180.0, 180.0].
//
//	  at least one of latitude and longitude must be non-zero.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// OptimizeTours sends an OptimizeToursRequest containing a ShipmentModel and returns an
// OptimizeToursResponse containing ShipmentRoutes, which are a set of
// routes to be performed by vehicles minimizing the overall cost.
//
// A ShipmentModel model consists mainly of Shipments that need to be
// carried out and Vehicles that can be used to transport the Shipments.
// The ShipmentRoutes assign Shipments to Vehicles. More specifically,
// they assign a series of Visits to each vehicle, where a Visit
// corresponds to a VisitRequest, which is a pickup or delivery for a
// Shipment.
//
// The goal is to provide an assignment of ShipmentRoutes to Vehicles that
// minimizes the total cost where cost has many components defined in the
// ShipmentModel.
func (c *Client) OptimizeTours(ctx context.Context, req *routeoptimizationpb.OptimizeToursRequest, opts ...gax.CallOption) (*routeoptimizationpb.OptimizeToursResponse, error) {
	return c.internalClient.OptimizeTours(ctx, req, opts...)
}

// BatchOptimizeTours optimizes vehicle tours for one or more OptimizeToursRequest
// messages as a batch.
//
// This method is a Long Running Operation (LRO). The inputs for optimization
// (OptimizeToursRequest messages) and outputs (OptimizeToursResponse
// messages) are read from and written to Cloud Storage in user-specified
// format. Like the OptimizeTours method, each OptimizeToursRequest
// contains a ShipmentModel and returns an OptimizeToursResponse
// containing ShipmentRoute fields, which are a set of routes to be
// performed by vehicles minimizing the overall cost.
//
// The user can poll operations.get to check the status of the LRO:
//
// If the LRO done field is false, then at least one request is still
// being processed. Other requests may have completed successfully and their
// results are available in Cloud Storage.
//
// If the LRO’s done field is true, then all requests have been processed.
// Any successfully processed requests will have their results available in
// Cloud Storage. Any requests that failed will not have their results
// available in Cloud Storage. If the LRO’s error field is set, then it
// contains the error from one of the failed requests.
func (c *Client) BatchOptimizeTours(ctx context.Context, req *routeoptimizationpb.BatchOptimizeToursRequest, opts ...gax.CallOption) (*BatchOptimizeToursOperation, error) {
	return c.internalClient.BatchOptimizeTours(ctx, req, opts...)
}

// BatchOptimizeToursOperation returns a new BatchOptimizeToursOperation from a given name.
// The name must be that of a previously created BatchOptimizeToursOperation, possibly from a different process.
func (c *Client) BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation {
	return c.internalClient.BatchOptimizeToursOperation(name)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *Client) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Route Optimization API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client routeoptimizationpb.RouteOptimizationClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewClient creates a new route optimization client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for optimizing vehicle tours.
//
// Validity of certain types of fields:
//
//	google.protobuf.Timestamp
//
//	  Times are in Unix time: seconds since 1970-01-01T00:00:00+00:00.
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.protobuf.Duration
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.type.LatLng
//
//	  latitude must be in [-90.0, 90.0].
//
//	  longitude must be in [-180.0, 180.0].
//
//	  at least one of latitude and longitude must be non-zero.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		client:           routeoptimizationpb.NewRouteOptimizationClient(connPool),
		CallOptions:      &client.CallOptions,
		logger:           internaloption.GetLogger(opts),
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version, "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	logger *slog.Logger
}

// NewRESTClient creates a new route optimization rest client.
//
// A service for optimizing vehicle tours.
//
// Validity of certain types of fields:
//
//	google.protobuf.Timestamp
//
//	  Times are in Unix time: seconds since 1970-01-01T00:00:00+00:00.
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.protobuf.Duration
//
//	  seconds must be in [0, 253402300799],
//	  i.e. in [1970-01-01T00:00:00+00:00, 9999-12-31T23:59:59+00:00].
//
//	  nanos must be unset or set to 0.
//
//	google.type.LatLng
//
//	  latitude must be in [-90.0, 90.0].
//
//	  longitude must be in [-180.0, 180.0].
//
//	  at least one of latitude and longitude must be non-zero.
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://routeoptimization.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://routeoptimization.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://routeoptimization.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://routeoptimization.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN", "pb", protoVersion)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) OptimizeTours(ctx context.Context, req *routeoptimizationpb.OptimizeToursRequest, opts ...gax.CallOption) (*routeoptimizationpb.OptimizeToursResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).OptimizeTours[0:len((*c.CallOptions).OptimizeTours):len((*c.CallOptions).OptimizeTours)], opts...)
	var resp *routeoptimizationpb.OptimizeToursResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.OptimizeTours, req, settings.GRPC, c.logger, "OptimizeTours")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) BatchOptimizeTours(ctx context.Context, req *routeoptimizationpb.BatchOptimizeToursRequest, opts ...gax.CallOption) (*BatchOptimizeToursOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).BatchOptimizeTours[0:len((*c.CallOptions).BatchOptimizeTours):len((*c.CallOptions).BatchOptimizeTours)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.BatchOptimizeTours, req, settings.GRPC, c.logger, "BatchOptimizeTours")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchOptimizeToursOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// OptimizeTours sends an OptimizeToursRequest containing a ShipmentModel and returns an
// OptimizeToursResponse containing ShipmentRoutes, which are a set of
// routes to be performed by vehicles minimizing the overall cost.
//
// A ShipmentModel model consists mainly of Shipments that need to be
// carried out and Vehicles that can be used to transport the Shipments.
// The ShipmentRoutes assign Shipments to Vehicles. More specifically,
// they assign a series of Visits to each vehicle, where a Visit
// corresponds to a VisitRequest, which is a pickup or delivery for a
// Shipment.
//
// The goal is to provide an assignment of ShipmentRoutes to Vehicles that
// minimizes the total cost where cost has many components defined in the
// ShipmentModel.
func (c *restClient) OptimizeTours(ctx context.Context, req *routeoptimizationpb.OptimizeToursRequest, opts ...gax.CallOption) (*routeoptimizationpb.OptimizeToursResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:optimizeTours", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).OptimizeTours[0:len((*c.CallOptions).OptimizeTours):len((*c.CallOptions).OptimizeTours)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &routeoptimizationpb.OptimizeToursResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "OptimizeTours")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// BatchOptimizeTours optimizes vehicle tours for one or more OptimizeToursRequest
// messages as a batch.
//
// This method is a Long Running Operation (LRO). The inputs for optimization
// (OptimizeToursRequest messages) and outputs (OptimizeToursResponse
// messages) are read from and written to Cloud Storage in user-specified
// format. Like the OptimizeTours method, each OptimizeToursRequest
// contains a ShipmentModel and returns an OptimizeToursResponse
// containing ShipmentRoute fields, which are a set of routes to be
// performed by vehicles minimizing the overall cost.
//
// The user can poll operations.get to check the status of the LRO:
//
// If the LRO done field is false, then at least one request is still
// being processed. Other requests may have completed successfully and their
// results are available in Cloud Storage.
//
// If the LRO’s done field is true, then all requests have been processed.
// Any successfully processed requests will have their results available in
// Cloud Storage. Any requests that failed will not have their results
// available in Cloud Storage. If the LRO’s error field is set, then it
// contains the error from one of the failed requests.
func (c *restClient) BatchOptimizeTours(ctx context.Context, req *routeoptimizationpb.BatchOptimizeToursRequest, opts ...gax.CallOption) (*BatchOptimizeToursOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:batchOptimizeTours", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "BatchOptimizeTours")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &BatchOptimizeToursOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *restClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetOperation")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// BatchOptimizeToursOperation returns a new BatchOptimizeToursOperation from a given name.
// The name must be that of a previously created BatchOptimizeToursOperation, possibly from a different process.
func (c *gRPCClient) BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation {
	return &BatchOptimizeToursOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// BatchOptimizeToursOperation returns a new BatchOptimizeToursOperation from a given name.
// The name must be that of a previously created BatchOptimizeToursOperation, possibly from a different process.
func (c *restClient) BatchOptimizeToursOperation(name string) *BatchOptimizeToursOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &BatchOptimizeToursOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}
