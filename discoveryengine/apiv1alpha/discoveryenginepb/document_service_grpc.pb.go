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
// source: google/cloud/discoveryengine/v1alpha/document_service.proto

package discoveryenginepb

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
	DocumentService_GetDocument_FullMethodName               = "/google.cloud.discoveryengine.v1alpha.DocumentService/GetDocument"
	DocumentService_ListDocuments_FullMethodName             = "/google.cloud.discoveryengine.v1alpha.DocumentService/ListDocuments"
	DocumentService_CreateDocument_FullMethodName            = "/google.cloud.discoveryengine.v1alpha.DocumentService/CreateDocument"
	DocumentService_UpdateDocument_FullMethodName            = "/google.cloud.discoveryengine.v1alpha.DocumentService/UpdateDocument"
	DocumentService_DeleteDocument_FullMethodName            = "/google.cloud.discoveryengine.v1alpha.DocumentService/DeleteDocument"
	DocumentService_ImportDocuments_FullMethodName           = "/google.cloud.discoveryengine.v1alpha.DocumentService/ImportDocuments"
	DocumentService_PurgeDocuments_FullMethodName            = "/google.cloud.discoveryengine.v1alpha.DocumentService/PurgeDocuments"
	DocumentService_GetProcessedDocument_FullMethodName      = "/google.cloud.discoveryengine.v1alpha.DocumentService/GetProcessedDocument"
	DocumentService_BatchGetDocumentsMetadata_FullMethodName = "/google.cloud.discoveryengine.v1alpha.DocumentService/BatchGetDocumentsMetadata"
)

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	// Gets a [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Document, error)
	// Gets a list of [Document][google.cloud.discoveryengine.v1alpha.Document]s.
	ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error)
	// Creates a [Document][google.cloud.discoveryengine.v1alpha.Document].
	CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*Document, error)
	// Updates a [Document][google.cloud.discoveryengine.v1alpha.Document].
	UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*Document, error)
	// Deletes a [Document][google.cloud.discoveryengine.v1alpha.Document].
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Bulk import of multiple
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s. Request
	// processing may be synchronous. Non-existing items are created.
	//
	// Note: It is possible for a subset of the
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s to be
	// successfully updated.
	ImportDocuments(ctx context.Context, in *ImportDocumentsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Permanently deletes all selected
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s in a branch.
	//
	// This process is asynchronous. Depending on the number of
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s to be deleted,
	// this operation can take hours to complete. Before the delete operation
	// completes, some [Document][google.cloud.discoveryengine.v1alpha.Document]s
	// might still be returned by
	// [DocumentService.GetDocument][google.cloud.discoveryengine.v1alpha.DocumentService.GetDocument]
	// or
	// [DocumentService.ListDocuments][google.cloud.discoveryengine.v1alpha.DocumentService.ListDocuments].
	//
	// To get a list of the
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s to be deleted,
	// set
	// [PurgeDocumentsRequest.force][google.cloud.discoveryengine.v1alpha.PurgeDocumentsRequest.force]
	// to false.
	PurgeDocuments(ctx context.Context, in *PurgeDocumentsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Gets the parsed layout information for a
	// [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetProcessedDocument(ctx context.Context, in *GetProcessedDocumentRequest, opts ...grpc.CallOption) (*ProcessedDocument, error)
	// Gets index freshness metadata for
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s. Supported for
	// website search only.
	BatchGetDocumentsMetadata(ctx context.Context, in *BatchGetDocumentsMetadataRequest, opts ...grpc.CallOption) (*BatchGetDocumentsMetadataResponse, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, DocumentService_GetDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error) {
	out := new(ListDocumentsResponse)
	err := c.cc.Invoke(ctx, DocumentService_ListDocuments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, DocumentService_CreateDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, DocumentService_UpdateDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DocumentService_DeleteDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) ImportDocuments(ctx context.Context, in *ImportDocumentsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, DocumentService_ImportDocuments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) PurgeDocuments(ctx context.Context, in *PurgeDocumentsRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, DocumentService_PurgeDocuments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) GetProcessedDocument(ctx context.Context, in *GetProcessedDocumentRequest, opts ...grpc.CallOption) (*ProcessedDocument, error) {
	out := new(ProcessedDocument)
	err := c.cc.Invoke(ctx, DocumentService_GetProcessedDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) BatchGetDocumentsMetadata(ctx context.Context, in *BatchGetDocumentsMetadataRequest, opts ...grpc.CallOption) (*BatchGetDocumentsMetadataResponse, error) {
	out := new(BatchGetDocumentsMetadataResponse)
	err := c.cc.Invoke(ctx, DocumentService_BatchGetDocumentsMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
// All implementations should embed UnimplementedDocumentServiceServer
// for forward compatibility
type DocumentServiceServer interface {
	// Gets a [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetDocument(context.Context, *GetDocumentRequest) (*Document, error)
	// Gets a list of [Document][google.cloud.discoveryengine.v1alpha.Document]s.
	ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error)
	// Creates a [Document][google.cloud.discoveryengine.v1alpha.Document].
	CreateDocument(context.Context, *CreateDocumentRequest) (*Document, error)
	// Updates a [Document][google.cloud.discoveryengine.v1alpha.Document].
	UpdateDocument(context.Context, *UpdateDocumentRequest) (*Document, error)
	// Deletes a [Document][google.cloud.discoveryengine.v1alpha.Document].
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*emptypb.Empty, error)
	// Bulk import of multiple
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s. Request
	// processing may be synchronous. Non-existing items are created.
	//
	// Note: It is possible for a subset of the
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s to be
	// successfully updated.
	ImportDocuments(context.Context, *ImportDocumentsRequest) (*longrunningpb.Operation, error)
	// Permanently deletes all selected
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s in a branch.
	//
	// This process is asynchronous. Depending on the number of
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s to be deleted,
	// this operation can take hours to complete. Before the delete operation
	// completes, some [Document][google.cloud.discoveryengine.v1alpha.Document]s
	// might still be returned by
	// [DocumentService.GetDocument][google.cloud.discoveryengine.v1alpha.DocumentService.GetDocument]
	// or
	// [DocumentService.ListDocuments][google.cloud.discoveryengine.v1alpha.DocumentService.ListDocuments].
	//
	// To get a list of the
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s to be deleted,
	// set
	// [PurgeDocumentsRequest.force][google.cloud.discoveryengine.v1alpha.PurgeDocumentsRequest.force]
	// to false.
	PurgeDocuments(context.Context, *PurgeDocumentsRequest) (*longrunningpb.Operation, error)
	// Gets the parsed layout information for a
	// [Document][google.cloud.discoveryengine.v1alpha.Document].
	GetProcessedDocument(context.Context, *GetProcessedDocumentRequest) (*ProcessedDocument, error)
	// Gets index freshness metadata for
	// [Document][google.cloud.discoveryengine.v1alpha.Document]s. Supported for
	// website search only.
	BatchGetDocumentsMetadata(context.Context, *BatchGetDocumentsMetadataRequest) (*BatchGetDocumentsMetadataResponse, error)
}

// UnimplementedDocumentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (UnimplementedDocumentServiceServer) GetDocument(context.Context, *GetDocumentRequest) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedDocumentServiceServer) ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}
func (UnimplementedDocumentServiceServer) CreateDocument(context.Context, *CreateDocumentRequest) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocument not implemented")
}
func (UnimplementedDocumentServiceServer) UpdateDocument(context.Context, *UpdateDocumentRequest) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDocument not implemented")
}
func (UnimplementedDocumentServiceServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedDocumentServiceServer) ImportDocuments(context.Context, *ImportDocumentsRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportDocuments not implemented")
}
func (UnimplementedDocumentServiceServer) PurgeDocuments(context.Context, *PurgeDocumentsRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeDocuments not implemented")
}
func (UnimplementedDocumentServiceServer) GetProcessedDocument(context.Context, *GetProcessedDocumentRequest) (*ProcessedDocument, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProcessedDocument not implemented")
}
func (UnimplementedDocumentServiceServer) BatchGetDocumentsMetadata(context.Context, *BatchGetDocumentsMetadataRequest) (*BatchGetDocumentsMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetDocumentsMetadata not implemented")
}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_GetDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_ListDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).ListDocuments(ctx, req.(*ListDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_CreateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).CreateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_CreateDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).CreateDocument(ctx, req.(*CreateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_UpdateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).UpdateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_UpdateDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).UpdateDocument(ctx, req.(*UpdateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_DeleteDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_ImportDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).ImportDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_ImportDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).ImportDocuments(ctx, req.(*ImportDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_PurgeDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).PurgeDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_PurgeDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).PurgeDocuments(ctx, req.(*PurgeDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_GetProcessedDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessedDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetProcessedDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_GetProcessedDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetProcessedDocument(ctx, req.(*GetProcessedDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_BatchGetDocumentsMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetDocumentsMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).BatchGetDocumentsMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_BatchGetDocumentsMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).BatchGetDocumentsMetadata(ctx, req.(*BatchGetDocumentsMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentService_ServiceDesc is the grpc.ServiceDesc for DocumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDocument",
			Handler:    _DocumentService_GetDocument_Handler,
		},
		{
			MethodName: "ListDocuments",
			Handler:    _DocumentService_ListDocuments_Handler,
		},
		{
			MethodName: "CreateDocument",
			Handler:    _DocumentService_CreateDocument_Handler,
		},
		{
			MethodName: "UpdateDocument",
			Handler:    _DocumentService_UpdateDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _DocumentService_DeleteDocument_Handler,
		},
		{
			MethodName: "ImportDocuments",
			Handler:    _DocumentService_ImportDocuments_Handler,
		},
		{
			MethodName: "PurgeDocuments",
			Handler:    _DocumentService_PurgeDocuments_Handler,
		},
		{
			MethodName: "GetProcessedDocument",
			Handler:    _DocumentService_GetProcessedDocument_Handler,
		},
		{
			MethodName: "BatchGetDocumentsMetadata",
			Handler:    _DocumentService_BatchGetDocumentsMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/document_service.proto",
}
