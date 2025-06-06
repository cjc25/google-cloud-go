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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.7
// source: google/cloud/dialogflow/cx/v3beta1/changelog.proto

package cxpb

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for
// [Changelogs.ListChangelogs][google.cloud.dialogflow.cx.v3beta1.Changelogs.ListChangelogs].
type ListChangelogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The agent containing the changelogs.
	// Format: `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The filter string. Supports filter by user_email, resource, type and
	// create_time. Some examples:
	//  1. By user email:
	//     user_email = "someone@google.com"
	//  2. By resource name:
	//     resource = "projects/123/locations/global/agents/456/flows/789"
	//  3. By resource display name:
	//     display_name = "my agent"
	//  4. By action:
	//     action = "Create"
	//  5. By type:
	//     type = "flows"
	//  6. By create time. Currently predicates on `create_time` and
	//     `create_time_epoch_seconds` are supported:
	//     create_time_epoch_seconds > 1551790877 AND create_time <=
	//     2017-01-15T01:30:15.01Z
	//  7. Combination of above filters:
	//     resource = "projects/123/locations/global/agents/456/flows/789"
	//     AND user_email = "someone@google.com"
	//     AND create_time <= 2017-01-15T01:30:15.01Z
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// The maximum number of items to return in a single page. By default 100 and
	// at most 1000.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous list request.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListChangelogsRequest) Reset() {
	*x = ListChangelogsRequest{}
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChangelogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangelogsRequest) ProtoMessage() {}

func (x *ListChangelogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangelogsRequest.ProtoReflect.Descriptor instead.
func (*ListChangelogsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescGZIP(), []int{0}
}

func (x *ListChangelogsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListChangelogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListChangelogsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListChangelogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response message for
// [Changelogs.ListChangelogs][google.cloud.dialogflow.cx.v3beta1.Changelogs.ListChangelogs].
type ListChangelogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of changelogs. There will be a maximum number of items returned
	// based on the page_size field in the request. The changelogs will be ordered
	// by timestamp.
	Changelogs []*Changelog `protobuf:"bytes,1,rep,name=changelogs,proto3" json:"changelogs,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListChangelogsResponse) Reset() {
	*x = ListChangelogsResponse{}
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChangelogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChangelogsResponse) ProtoMessage() {}

func (x *ListChangelogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChangelogsResponse.ProtoReflect.Descriptor instead.
func (*ListChangelogsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescGZIP(), []int{1}
}

func (x *ListChangelogsResponse) GetChangelogs() []*Changelog {
	if x != nil {
		return x.Changelogs
	}
	return nil
}

func (x *ListChangelogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request message for
// [Changelogs.GetChangelog][google.cloud.dialogflow.cx.v3beta1.Changelogs.GetChangelog].
type GetChangelogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the changelog to get.
	// Format:
	// `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/changelogs/<ChangelogID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetChangelogRequest) Reset() {
	*x = GetChangelogRequest{}
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChangelogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChangelogRequest) ProtoMessage() {}

func (x *GetChangelogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChangelogRequest.ProtoReflect.Descriptor instead.
func (*GetChangelogRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescGZIP(), []int{2}
}

func (x *GetChangelogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Changelogs represents a change made to a given agent.
type Changelog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of the changelog.
	// Format:
	// `projects/<ProjectID>/locations/<LocationID>/agents/<AgentID>/changelogs/<ChangelogID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Email address of the authenticated user.
	UserEmail string `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	// The affected resource display name of the change.
	DisplayName string `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The action of the change.
	Action string `protobuf:"bytes,11,opt,name=action,proto3" json:"action,omitempty"`
	// The affected resource type.
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	// The affected resource name of the change.
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// The timestamp of the change.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The affected language code of the change.
	LanguageCode string `protobuf:"bytes,14,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *Changelog) Reset() {
	*x = Changelog{}
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Changelog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Changelog) ProtoMessage() {}

func (x *Changelog) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Changelog.ProtoReflect.Descriptor instead.
func (*Changelog) Descriptor() ([]byte, []int) {
	return file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescGZIP(), []int{3}
}

func (x *Changelog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Changelog) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *Changelog) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Changelog) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Changelog) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Changelog) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Changelog) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Changelog) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

var File_google_cloud_dialogflow_cx_v3beta1_changelog_proto protoreflect.FileDescriptor

var file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78,
	0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x12, 0x23,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x6f, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x8f, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x52, 0x0a, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x25, 0x0a, 0x23,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x6f, 0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x84, 0x03, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x77, 0xea, 0x41, 0x74, 0x0a, 0x23, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x12, 0x4d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x7b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x7d, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x7d,
	0x32, 0xa5, 0x04, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x12,
	0xd6, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x73, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e,
	0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0xda, 0x41, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x12, 0x3c, 0x2f, 0x76, 0x33, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0xc3, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e,
	0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x22, 0x4b, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e,
	0x12, 0x3c, 0x2f, 0x76, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x1a, 0x78,
	0xca, 0x41, 0x19, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x59, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0xc5, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x78, 0x2e, 0x76, 0x33, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x42, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x33, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x63, 0x78, 0x70, 0x62, 0x3b, 0x63, 0x78, 0x70, 0x62, 0xa2, 0x02, 0x02,
	0x44, 0x46, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x78, 0x2e,
	0x56, 0x33, 0x42, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x66,
	0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x58, 0x3a, 0x3a, 0x56, 0x33, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescOnce sync.Once
	file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescData = file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDesc
)

func file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescGZIP() []byte {
	file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescOnce.Do(func() {
		file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescData)
	})
	return file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDescData
}

var file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_goTypes = []any{
	(*ListChangelogsRequest)(nil),  // 0: google.cloud.dialogflow.cx.v3beta1.ListChangelogsRequest
	(*ListChangelogsResponse)(nil), // 1: google.cloud.dialogflow.cx.v3beta1.ListChangelogsResponse
	(*GetChangelogRequest)(nil),    // 2: google.cloud.dialogflow.cx.v3beta1.GetChangelogRequest
	(*Changelog)(nil),              // 3: google.cloud.dialogflow.cx.v3beta1.Changelog
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_depIdxs = []int32{
	3, // 0: google.cloud.dialogflow.cx.v3beta1.ListChangelogsResponse.changelogs:type_name -> google.cloud.dialogflow.cx.v3beta1.Changelog
	4, // 1: google.cloud.dialogflow.cx.v3beta1.Changelog.create_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.cloud.dialogflow.cx.v3beta1.Changelogs.ListChangelogs:input_type -> google.cloud.dialogflow.cx.v3beta1.ListChangelogsRequest
	2, // 3: google.cloud.dialogflow.cx.v3beta1.Changelogs.GetChangelog:input_type -> google.cloud.dialogflow.cx.v3beta1.GetChangelogRequest
	1, // 4: google.cloud.dialogflow.cx.v3beta1.Changelogs.ListChangelogs:output_type -> google.cloud.dialogflow.cx.v3beta1.ListChangelogsResponse
	3, // 5: google.cloud.dialogflow.cx.v3beta1.Changelogs.GetChangelog:output_type -> google.cloud.dialogflow.cx.v3beta1.Changelog
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_init() }
func file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_init() {
	if File_google_cloud_dialogflow_cx_v3beta1_changelog_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_goTypes,
		DependencyIndexes: file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_depIdxs,
		MessageInfos:      file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_msgTypes,
	}.Build()
	File_google_cloud_dialogflow_cx_v3beta1_changelog_proto = out.File
	file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_rawDesc = nil
	file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_goTypes = nil
	file_google_cloud_dialogflow_cx_v3beta1_changelog_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChangelogsClient is the client API for Changelogs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChangelogsClient interface {
	// Returns the list of Changelogs.
	ListChangelogs(ctx context.Context, in *ListChangelogsRequest, opts ...grpc.CallOption) (*ListChangelogsResponse, error)
	// Retrieves the specified Changelog.
	GetChangelog(ctx context.Context, in *GetChangelogRequest, opts ...grpc.CallOption) (*Changelog, error)
}

type changelogsClient struct {
	cc grpc.ClientConnInterface
}

func NewChangelogsClient(cc grpc.ClientConnInterface) ChangelogsClient {
	return &changelogsClient{cc}
}

func (c *changelogsClient) ListChangelogs(ctx context.Context, in *ListChangelogsRequest, opts ...grpc.CallOption) (*ListChangelogsResponse, error) {
	out := new(ListChangelogsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.cx.v3beta1.Changelogs/ListChangelogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *changelogsClient) GetChangelog(ctx context.Context, in *GetChangelogRequest, opts ...grpc.CallOption) (*Changelog, error) {
	out := new(Changelog)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.cx.v3beta1.Changelogs/GetChangelog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangelogsServer is the server API for Changelogs service.
type ChangelogsServer interface {
	// Returns the list of Changelogs.
	ListChangelogs(context.Context, *ListChangelogsRequest) (*ListChangelogsResponse, error)
	// Retrieves the specified Changelog.
	GetChangelog(context.Context, *GetChangelogRequest) (*Changelog, error)
}

// UnimplementedChangelogsServer can be embedded to have forward compatible implementations.
type UnimplementedChangelogsServer struct {
}

func (*UnimplementedChangelogsServer) ListChangelogs(context.Context, *ListChangelogsRequest) (*ListChangelogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChangelogs not implemented")
}
func (*UnimplementedChangelogsServer) GetChangelog(context.Context, *GetChangelogRequest) (*Changelog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangelog not implemented")
}

func RegisterChangelogsServer(s *grpc.Server, srv ChangelogsServer) {
	s.RegisterService(&_Changelogs_serviceDesc, srv)
}

func _Changelogs_ListChangelogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChangelogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangelogsServer).ListChangelogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.cx.v3beta1.Changelogs/ListChangelogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangelogsServer).ListChangelogs(ctx, req.(*ListChangelogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Changelogs_GetChangelog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangelogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangelogsServer).GetChangelog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.cx.v3beta1.Changelogs/GetChangelog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangelogsServer).GetChangelog(ctx, req.(*GetChangelogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Changelogs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.cx.v3beta1.Changelogs",
	HandlerType: (*ChangelogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListChangelogs",
			Handler:    _Changelogs_ListChangelogs_Handler,
		},
		{
			MethodName: "GetChangelog",
			Handler:    _Changelogs_GetChangelog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/cx/v3beta1/changelog.proto",
}
