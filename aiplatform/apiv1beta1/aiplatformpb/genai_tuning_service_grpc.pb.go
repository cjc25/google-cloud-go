// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.7
// source: google/cloud/aiplatform/v1beta1/genai_tuning_service.proto

package aiplatformpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GenAiTuningService_CreateTuningJob_FullMethodName  = "/google.cloud.aiplatform.v1beta1.GenAiTuningService/CreateTuningJob"
	GenAiTuningService_GetTuningJob_FullMethodName     = "/google.cloud.aiplatform.v1beta1.GenAiTuningService/GetTuningJob"
	GenAiTuningService_ListTuningJobs_FullMethodName   = "/google.cloud.aiplatform.v1beta1.GenAiTuningService/ListTuningJobs"
	GenAiTuningService_CancelTuningJob_FullMethodName  = "/google.cloud.aiplatform.v1beta1.GenAiTuningService/CancelTuningJob"
	GenAiTuningService_RebaseTunedModel_FullMethodName = "/google.cloud.aiplatform.v1beta1.GenAiTuningService/RebaseTunedModel"
)

// GenAiTuningServiceClient is the client API for GenAiTuningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenAiTuningServiceClient interface {
	// Creates a TuningJob. A created TuningJob right away will be attempted to
	// be run.
	CreateTuningJob(ctx context.Context, in *CreateTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error)
	// Gets a TuningJob.
	GetTuningJob(ctx context.Context, in *GetTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error)
	// Lists TuningJobs in a Location.
	ListTuningJobs(ctx context.Context, in *ListTuningJobsRequest, opts ...grpc.CallOption) (*ListTuningJobsResponse, error)
	// Cancels a TuningJob.
	// Starts asynchronous cancellation on the TuningJob. The server makes a best
	// effort to cancel the job, but success is not guaranteed. Clients can use
	// [GenAiTuningService.GetTuningJob][google.cloud.aiplatform.v1beta1.GenAiTuningService.GetTuningJob]
	// or other methods to check whether the cancellation succeeded or whether the
	// job completed despite cancellation. On successful cancellation, the
	// TuningJob is not deleted; instead it becomes a job with a
	// [TuningJob.error][google.cloud.aiplatform.v1beta1.TuningJob.error] value
	// with a [google.rpc.Status.code][google.rpc.Status.code] of 1, corresponding
	// to `Code.CANCELLED`, and
	// [TuningJob.state][google.cloud.aiplatform.v1beta1.TuningJob.state] is set
	// to `CANCELLED`.
	CancelTuningJob(ctx context.Context, in *CancelTuningJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Rebase a TunedModel.
	// Creates a LongRunningOperation that takes a legacy Tuned GenAI model
	// Reference and creates a TuningJob based on newly available model.
	RebaseTunedModel(ctx context.Context, in *RebaseTunedModelRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type genAiTuningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenAiTuningServiceClient(cc grpc.ClientConnInterface) GenAiTuningServiceClient {
	return &genAiTuningServiceClient{cc}
}

func (c *genAiTuningServiceClient) CreateTuningJob(ctx context.Context, in *CreateTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error) {
	out := new(TuningJob)
	err := c.cc.Invoke(ctx, GenAiTuningService_CreateTuningJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genAiTuningServiceClient) GetTuningJob(ctx context.Context, in *GetTuningJobRequest, opts ...grpc.CallOption) (*TuningJob, error) {
	out := new(TuningJob)
	err := c.cc.Invoke(ctx, GenAiTuningService_GetTuningJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genAiTuningServiceClient) ListTuningJobs(ctx context.Context, in *ListTuningJobsRequest, opts ...grpc.CallOption) (*ListTuningJobsResponse, error) {
	out := new(ListTuningJobsResponse)
	err := c.cc.Invoke(ctx, GenAiTuningService_ListTuningJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genAiTuningServiceClient) CancelTuningJob(ctx context.Context, in *CancelTuningJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GenAiTuningService_CancelTuningJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genAiTuningServiceClient) RebaseTunedModel(ctx context.Context, in *RebaseTunedModelRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, GenAiTuningService_RebaseTunedModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenAiTuningServiceServer is the server API for GenAiTuningService service.
// All implementations should embed UnimplementedGenAiTuningServiceServer
// for forward compatibility
type GenAiTuningServiceServer interface {
	// Creates a TuningJob. A created TuningJob right away will be attempted to
	// be run.
	CreateTuningJob(context.Context, *CreateTuningJobRequest) (*TuningJob, error)
	// Gets a TuningJob.
	GetTuningJob(context.Context, *GetTuningJobRequest) (*TuningJob, error)
	// Lists TuningJobs in a Location.
	ListTuningJobs(context.Context, *ListTuningJobsRequest) (*ListTuningJobsResponse, error)
	// Cancels a TuningJob.
	// Starts asynchronous cancellation on the TuningJob. The server makes a best
	// effort to cancel the job, but success is not guaranteed. Clients can use
	// [GenAiTuningService.GetTuningJob][google.cloud.aiplatform.v1beta1.GenAiTuningService.GetTuningJob]
	// or other methods to check whether the cancellation succeeded or whether the
	// job completed despite cancellation. On successful cancellation, the
	// TuningJob is not deleted; instead it becomes a job with a
	// [TuningJob.error][google.cloud.aiplatform.v1beta1.TuningJob.error] value
	// with a [google.rpc.Status.code][google.rpc.Status.code] of 1, corresponding
	// to `Code.CANCELLED`, and
	// [TuningJob.state][google.cloud.aiplatform.v1beta1.TuningJob.state] is set
	// to `CANCELLED`.
	CancelTuningJob(context.Context, *CancelTuningJobRequest) (*emptypb.Empty, error)
	// Rebase a TunedModel.
	// Creates a LongRunningOperation that takes a legacy Tuned GenAI model
	// Reference and creates a TuningJob based on newly available model.
	RebaseTunedModel(context.Context, *RebaseTunedModelRequest) (*longrunningpb.Operation, error)
}

// UnimplementedGenAiTuningServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGenAiTuningServiceServer struct {
}

func (UnimplementedGenAiTuningServiceServer) CreateTuningJob(context.Context, *CreateTuningJobRequest) (*TuningJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTuningJob not implemented")
}
func (UnimplementedGenAiTuningServiceServer) GetTuningJob(context.Context, *GetTuningJobRequest) (*TuningJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTuningJob not implemented")
}
func (UnimplementedGenAiTuningServiceServer) ListTuningJobs(context.Context, *ListTuningJobsRequest) (*ListTuningJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTuningJobs not implemented")
}
func (UnimplementedGenAiTuningServiceServer) CancelTuningJob(context.Context, *CancelTuningJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTuningJob not implemented")
}
func (UnimplementedGenAiTuningServiceServer) RebaseTunedModel(context.Context, *RebaseTunedModelRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebaseTunedModel not implemented")
}

// UnsafeGenAiTuningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenAiTuningServiceServer will
// result in compilation errors.
type UnsafeGenAiTuningServiceServer interface {
	mustEmbedUnimplementedGenAiTuningServiceServer()
}

func RegisterGenAiTuningServiceServer(s grpc.ServiceRegistrar, srv GenAiTuningServiceServer) {
	s.RegisterService(&GenAiTuningService_ServiceDesc, srv)
}

func _GenAiTuningService_CreateTuningJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTuningJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).CreateTuningJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenAiTuningService_CreateTuningJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).CreateTuningJob(ctx, req.(*CreateTuningJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenAiTuningService_GetTuningJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTuningJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).GetTuningJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenAiTuningService_GetTuningJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).GetTuningJob(ctx, req.(*GetTuningJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenAiTuningService_ListTuningJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTuningJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).ListTuningJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenAiTuningService_ListTuningJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).ListTuningJobs(ctx, req.(*ListTuningJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenAiTuningService_CancelTuningJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTuningJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).CancelTuningJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenAiTuningService_CancelTuningJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).CancelTuningJob(ctx, req.(*CancelTuningJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenAiTuningService_RebaseTunedModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RebaseTunedModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenAiTuningServiceServer).RebaseTunedModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenAiTuningService_RebaseTunedModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenAiTuningServiceServer).RebaseTunedModel(ctx, req.(*RebaseTunedModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenAiTuningService_ServiceDesc is the grpc.ServiceDesc for GenAiTuningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenAiTuningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1beta1.GenAiTuningService",
	HandlerType: (*GenAiTuningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTuningJob",
			Handler:    _GenAiTuningService_CreateTuningJob_Handler,
		},
		{
			MethodName: "GetTuningJob",
			Handler:    _GenAiTuningService_GetTuningJob_Handler,
		},
		{
			MethodName: "ListTuningJobs",
			Handler:    _GenAiTuningService_ListTuningJobs_Handler,
		},
		{
			MethodName: "CancelTuningJob",
			Handler:    _GenAiTuningService_CancelTuningJob_Handler,
		},
		{
			MethodName: "RebaseTunedModel",
			Handler:    _GenAiTuningService_RebaseTunedModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1beta1/genai_tuning_service.proto",
}
