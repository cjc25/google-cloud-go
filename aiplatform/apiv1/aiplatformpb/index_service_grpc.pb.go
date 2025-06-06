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
// source: google/cloud/aiplatform/v1/index_service.proto

package aiplatformpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IndexService_CreateIndex_FullMethodName      = "/google.cloud.aiplatform.v1.IndexService/CreateIndex"
	IndexService_GetIndex_FullMethodName         = "/google.cloud.aiplatform.v1.IndexService/GetIndex"
	IndexService_ListIndexes_FullMethodName      = "/google.cloud.aiplatform.v1.IndexService/ListIndexes"
	IndexService_UpdateIndex_FullMethodName      = "/google.cloud.aiplatform.v1.IndexService/UpdateIndex"
	IndexService_DeleteIndex_FullMethodName      = "/google.cloud.aiplatform.v1.IndexService/DeleteIndex"
	IndexService_UpsertDatapoints_FullMethodName = "/google.cloud.aiplatform.v1.IndexService/UpsertDatapoints"
	IndexService_RemoveDatapoints_FullMethodName = "/google.cloud.aiplatform.v1.IndexService/RemoveDatapoints"
)

// IndexServiceClient is the client API for IndexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexServiceClient interface {
	// Creates an Index.
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Gets an Index.
	GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*Index, error)
	// Lists Indexes in a Location.
	ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error)
	// Updates an Index.
	UpdateIndex(ctx context.Context, in *UpdateIndexRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes an Index.
	// An Index can only be deleted when all its
	// [DeployedIndexes][google.cloud.aiplatform.v1.Index.deployed_indexes] had
	// been undeployed.
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Add/update Datapoints into an Index.
	UpsertDatapoints(ctx context.Context, in *UpsertDatapointsRequest, opts ...grpc.CallOption) (*UpsertDatapointsResponse, error)
	// Remove Datapoints from an Index.
	RemoveDatapoints(ctx context.Context, in *RemoveDatapointsRequest, opts ...grpc.CallOption) (*RemoveDatapointsResponse, error)
}

type indexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexServiceClient(cc grpc.ClientConnInterface) IndexServiceClient {
	return &indexServiceClient{cc}
}

func (c *indexServiceClient) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, IndexService_CreateIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, IndexService_GetIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error) {
	out := new(ListIndexesResponse)
	err := c.cc.Invoke(ctx, IndexService_ListIndexes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) UpdateIndex(ctx context.Context, in *UpdateIndexRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, IndexService_UpdateIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, IndexService_DeleteIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) UpsertDatapoints(ctx context.Context, in *UpsertDatapointsRequest, opts ...grpc.CallOption) (*UpsertDatapointsResponse, error) {
	out := new(UpsertDatapointsResponse)
	err := c.cc.Invoke(ctx, IndexService_UpsertDatapoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *indexServiceClient) RemoveDatapoints(ctx context.Context, in *RemoveDatapointsRequest, opts ...grpc.CallOption) (*RemoveDatapointsResponse, error) {
	out := new(RemoveDatapointsResponse)
	err := c.cc.Invoke(ctx, IndexService_RemoveDatapoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexServiceServer is the server API for IndexService service.
// All implementations should embed UnimplementedIndexServiceServer
// for forward compatibility
type IndexServiceServer interface {
	// Creates an Index.
	CreateIndex(context.Context, *CreateIndexRequest) (*longrunningpb.Operation, error)
	// Gets an Index.
	GetIndex(context.Context, *GetIndexRequest) (*Index, error)
	// Lists Indexes in a Location.
	ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error)
	// Updates an Index.
	UpdateIndex(context.Context, *UpdateIndexRequest) (*longrunningpb.Operation, error)
	// Deletes an Index.
	// An Index can only be deleted when all its
	// [DeployedIndexes][google.cloud.aiplatform.v1.Index.deployed_indexes] had
	// been undeployed.
	DeleteIndex(context.Context, *DeleteIndexRequest) (*longrunningpb.Operation, error)
	// Add/update Datapoints into an Index.
	UpsertDatapoints(context.Context, *UpsertDatapointsRequest) (*UpsertDatapointsResponse, error)
	// Remove Datapoints from an Index.
	RemoveDatapoints(context.Context, *RemoveDatapointsRequest) (*RemoveDatapointsResponse, error)
}

// UnimplementedIndexServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIndexServiceServer struct {
}

func (UnimplementedIndexServiceServer) CreateIndex(context.Context, *CreateIndexRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndex not implemented")
}
func (UnimplementedIndexServiceServer) GetIndex(context.Context, *GetIndexRequest) (*Index, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndex not implemented")
}
func (UnimplementedIndexServiceServer) ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIndexes not implemented")
}
func (UnimplementedIndexServiceServer) UpdateIndex(context.Context, *UpdateIndexRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIndex not implemented")
}
func (UnimplementedIndexServiceServer) DeleteIndex(context.Context, *DeleteIndexRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedIndexServiceServer) UpsertDatapoints(context.Context, *UpsertDatapointsRequest) (*UpsertDatapointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDatapoints not implemented")
}
func (UnimplementedIndexServiceServer) RemoveDatapoints(context.Context, *RemoveDatapointsRequest) (*RemoveDatapointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDatapoints not implemented")
}

// UnsafeIndexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexServiceServer will
// result in compilation errors.
type UnsafeIndexServiceServer interface {
	mustEmbedUnimplementedIndexServiceServer()
}

func RegisterIndexServiceServer(s grpc.ServiceRegistrar, srv IndexServiceServer) {
	s.RegisterService(&IndexService_ServiceDesc, srv)
}

func _IndexService_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_CreateIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).CreateIndex(ctx, req.(*CreateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_GetIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).GetIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_GetIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).GetIndex(ctx, req.(*GetIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_ListIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).ListIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_ListIndexes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).ListIndexes(ctx, req.(*ListIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_UpdateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).UpdateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_UpdateIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).UpdateIndex(ctx, req.(*UpdateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_DeleteIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_UpsertDatapoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDatapointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).UpsertDatapoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_UpsertDatapoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).UpsertDatapoints(ctx, req.(*UpsertDatapointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IndexService_RemoveDatapoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDatapointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexServiceServer).RemoveDatapoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IndexService_RemoveDatapoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexServiceServer).RemoveDatapoints(ctx, req.(*RemoveDatapointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IndexService_ServiceDesc is the grpc.ServiceDesc for IndexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.aiplatform.v1.IndexService",
	HandlerType: (*IndexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIndex",
			Handler:    _IndexService_CreateIndex_Handler,
		},
		{
			MethodName: "GetIndex",
			Handler:    _IndexService_GetIndex_Handler,
		},
		{
			MethodName: "ListIndexes",
			Handler:    _IndexService_ListIndexes_Handler,
		},
		{
			MethodName: "UpdateIndex",
			Handler:    _IndexService_UpdateIndex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _IndexService_DeleteIndex_Handler,
		},
		{
			MethodName: "UpsertDatapoints",
			Handler:    _IndexService_UpsertDatapoints_Handler,
		},
		{
			MethodName: "RemoveDatapoints",
			Handler:    _IndexService_RemoveDatapoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/aiplatform/v1/index_service.proto",
}
