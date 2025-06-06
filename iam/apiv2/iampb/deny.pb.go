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
// source: google/iam/v2/deny.proto

package iampb

import (
	expr "google.golang.org/genproto/googleapis/type/expr"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A deny rule in an IAM deny policy.
type DenyRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The identities that are prevented from using one or more permissions on
	// Google Cloud resources. This field can contain the following values:
	//
	//   - `principalSet://goog/public:all`: A special identifier that represents
	//     any principal that is on the internet, even if they do not have a Google
	//     Account or are not logged in.
	//
	//   - `principal://goog/subject/{email_id}`: A specific Google Account.
	//     Includes Gmail, Cloud Identity, and Google Workspace user accounts. For
	//     example, `principal://goog/subject/alice@example.com`.
	//
	//   - `deleted:principal://goog/subject/{email_id}?uid={uid}`: A specific
	//     Google Account that was deleted recently. For example,
	//     `deleted:principal://goog/subject/alice@example.com?uid=1234567890`. If
	//     the Google Account is recovered, this identifier reverts to the standard
	//     identifier for a Google Account.
	//
	//   - `principalSet://goog/group/{group_id}`: A Google group. For example,
	//     `principalSet://goog/group/admins@example.com`.
	//
	//   - `deleted:principalSet://goog/group/{group_id}?uid={uid}`: A Google group
	//     that was deleted recently. For example,
	//     `deleted:principalSet://goog/group/admins@example.com?uid=1234567890`. If
	//     the Google group is restored, this identifier reverts to the standard
	//     identifier for a Google group.
	//
	//   - `principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}`:
	//     A Google Cloud service account. For example,
	//     `principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com`.
	//
	//   - `deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/{service_account_id}?uid={uid}`:
	//     A Google Cloud service account that was deleted recently. For example,
	//     `deleted:principal://iam.googleapis.com/projects/-/serviceAccounts/my-service-account@iam.gserviceaccount.com?uid=1234567890`.
	//     If the service account is undeleted, this identifier reverts to the
	//     standard identifier for a service account.
	//
	//   - `principalSet://goog/cloudIdentityCustomerId/{customer_id}`: All of the
	//     principals associated with the specified Google Workspace or Cloud
	//     Identity customer ID. For example,
	//     `principalSet://goog/cloudIdentityCustomerId/C01Abc35`.
	DeniedPrincipals []string `protobuf:"bytes,1,rep,name=denied_principals,json=deniedPrincipals,proto3" json:"denied_principals,omitempty"`
	// The identities that are excluded from the deny rule, even if they are
	// listed in the `denied_principals`. For example, you could add a Google
	// group to the `denied_principals`, then exclude specific users who belong to
	// that group.
	//
	// This field can contain the same values as the `denied_principals` field,
	// excluding `principalSet://goog/public:all`, which represents all users on
	// the internet.
	ExceptionPrincipals []string `protobuf:"bytes,2,rep,name=exception_principals,json=exceptionPrincipals,proto3" json:"exception_principals,omitempty"`
	// The permissions that are explicitly denied by this rule. Each permission
	// uses the format `{service_fqdn}/{resource}.{verb}`, where `{service_fqdn}`
	// is the fully qualified domain name for the service. For example,
	// `iam.googleapis.com/roles.list`.
	DeniedPermissions []string `protobuf:"bytes,3,rep,name=denied_permissions,json=deniedPermissions,proto3" json:"denied_permissions,omitempty"`
	// Specifies the permissions that this rule excludes from the set of denied
	// permissions given by `denied_permissions`. If a permission appears in
	// `denied_permissions` _and_ in `exception_permissions` then it will _not_ be
	// denied.
	//
	// The excluded permissions can be specified using the same syntax as
	// `denied_permissions`.
	ExceptionPermissions []string `protobuf:"bytes,4,rep,name=exception_permissions,json=exceptionPermissions,proto3" json:"exception_permissions,omitempty"`
	// The condition that determines whether this deny rule applies to a request.
	// If the condition expression evaluates to `true`, then the deny rule is
	// applied; otherwise, the deny rule is not applied.
	//
	// Each deny rule is evaluated independently. If this deny rule does not apply
	// to a request, other deny rules might still apply.
	//
	// The condition can use CEL functions that evaluate
	// [resource
	// tags](https://cloud.google.com/iam/help/conditions/resource-tags). Other
	// functions and operators are not supported.
	DenialCondition *expr.Expr `protobuf:"bytes,5,opt,name=denial_condition,json=denialCondition,proto3" json:"denial_condition,omitempty"`
}

func (x *DenyRule) Reset() {
	*x = DenyRule{}
	mi := &file_google_iam_v2_deny_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DenyRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenyRule) ProtoMessage() {}

func (x *DenyRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_iam_v2_deny_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenyRule.ProtoReflect.Descriptor instead.
func (*DenyRule) Descriptor() ([]byte, []int) {
	return file_google_iam_v2_deny_proto_rawDescGZIP(), []int{0}
}

func (x *DenyRule) GetDeniedPrincipals() []string {
	if x != nil {
		return x.DeniedPrincipals
	}
	return nil
}

func (x *DenyRule) GetExceptionPrincipals() []string {
	if x != nil {
		return x.ExceptionPrincipals
	}
	return nil
}

func (x *DenyRule) GetDeniedPermissions() []string {
	if x != nil {
		return x.DeniedPermissions
	}
	return nil
}

func (x *DenyRule) GetExceptionPermissions() []string {
	if x != nil {
		return x.ExceptionPermissions
	}
	return nil
}

func (x *DenyRule) GetDenialCondition() *expr.Expr {
	if x != nil {
		return x.DenialCondition
	}
	return nil
}

var File_google_iam_v2_deny_proto protoreflect.FileDescriptor

var file_google_iam_v2_deny_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f,
	0x64, 0x65, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x02, 0x0a, 0x08, 0x44, 0x65, 0x6e, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x6e, 0x69, 0x65,
	0x64, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x65,
	0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x65, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x12, 0x2d,
	0x0a, 0x12, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x6e, 0x69,
	0x65, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a,
	0x15, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78,
	0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x10, 0x64, 0x65, 0x6e, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52,
	0x0f, 0x64, 0x65, 0x6e, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x7b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x44, 0x65, 0x6e, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x69, 0x61, 0x6d, 0x70, 0x62, 0x3b, 0x69, 0x61, 0x6d, 0x70,
	0x62, 0xaa, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x49, 0x61, 0x6d, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x56, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_iam_v2_deny_proto_rawDescOnce sync.Once
	file_google_iam_v2_deny_proto_rawDescData = file_google_iam_v2_deny_proto_rawDesc
)

func file_google_iam_v2_deny_proto_rawDescGZIP() []byte {
	file_google_iam_v2_deny_proto_rawDescOnce.Do(func() {
		file_google_iam_v2_deny_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_iam_v2_deny_proto_rawDescData)
	})
	return file_google_iam_v2_deny_proto_rawDescData
}

var file_google_iam_v2_deny_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_iam_v2_deny_proto_goTypes = []any{
	(*DenyRule)(nil),  // 0: google.iam.v2.DenyRule
	(*expr.Expr)(nil), // 1: google.type.Expr
}
var file_google_iam_v2_deny_proto_depIdxs = []int32{
	1, // 0: google.iam.v2.DenyRule.denial_condition:type_name -> google.type.Expr
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_iam_v2_deny_proto_init() }
func file_google_iam_v2_deny_proto_init() {
	if File_google_iam_v2_deny_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_iam_v2_deny_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_iam_v2_deny_proto_goTypes,
		DependencyIndexes: file_google_iam_v2_deny_proto_depIdxs,
		MessageInfos:      file_google_iam_v2_deny_proto_msgTypes,
	}.Build()
	File_google_iam_v2_deny_proto = out.File
	file_google_iam_v2_deny_proto_rawDesc = nil
	file_google_iam_v2_deny_proto_goTypes = nil
	file_google_iam_v2_deny_proto_depIdxs = nil
}
