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
// source: google/cloud/aiplatform/v1/pipeline_failure_policy.proto

package aiplatformpb

import (
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

// Represents the failure policy of a pipeline. Currently, the default of a
// pipeline is that the pipeline will continue to run until no more tasks can be
// executed, also known as PIPELINE_FAILURE_POLICY_FAIL_SLOW. However, if a
// pipeline is set to PIPELINE_FAILURE_POLICY_FAIL_FAST, it will stop scheduling
// any new tasks when a task has failed. Any scheduled tasks will continue to
// completion.
type PipelineFailurePolicy int32

const (
	// Default value, and follows fail slow behavior.
	PipelineFailurePolicy_PIPELINE_FAILURE_POLICY_UNSPECIFIED PipelineFailurePolicy = 0
	// Indicates that the pipeline should continue to run until all possible
	// tasks have been scheduled and completed.
	PipelineFailurePolicy_PIPELINE_FAILURE_POLICY_FAIL_SLOW PipelineFailurePolicy = 1
	// Indicates that the pipeline should stop scheduling new tasks after a task
	// has failed.
	PipelineFailurePolicy_PIPELINE_FAILURE_POLICY_FAIL_FAST PipelineFailurePolicy = 2
)

// Enum value maps for PipelineFailurePolicy.
var (
	PipelineFailurePolicy_name = map[int32]string{
		0: "PIPELINE_FAILURE_POLICY_UNSPECIFIED",
		1: "PIPELINE_FAILURE_POLICY_FAIL_SLOW",
		2: "PIPELINE_FAILURE_POLICY_FAIL_FAST",
	}
	PipelineFailurePolicy_value = map[string]int32{
		"PIPELINE_FAILURE_POLICY_UNSPECIFIED": 0,
		"PIPELINE_FAILURE_POLICY_FAIL_SLOW":   1,
		"PIPELINE_FAILURE_POLICY_FAIL_FAST":   2,
	}
)

func (x PipelineFailurePolicy) Enum() *PipelineFailurePolicy {
	p := new(PipelineFailurePolicy)
	*p = x
	return p
}

func (x PipelineFailurePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PipelineFailurePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_enumTypes[0].Descriptor()
}

func (PipelineFailurePolicy) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_enumTypes[0]
}

func (x PipelineFailurePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PipelineFailurePolicy.Descriptor instead.
func (PipelineFailurePolicy) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_aiplatform_v1_pipeline_failure_policy_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2a, 0x8e, 0x01, 0x0a, 0x15, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x27, 0x0a, 0x23, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x49, 0x50,
	0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x5f, 0x50, 0x4f,
	0x4c, 0x49, 0x43, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x53, 0x4c, 0x4f, 0x57, 0x10, 0x01,
	0x12, 0x25, 0x0a, 0x21, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x5f, 0x46, 0x41, 0x53, 0x54, 0x10, 0x02, 0x42, 0xd8, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x1a, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescData = file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_goTypes = []any{
	(PipelineFailurePolicy)(0), // 0: google.cloud.aiplatform.v1.PipelineFailurePolicy
}
var file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_init() }
func file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_init() {
	if File_google_cloud_aiplatform_v1_pipeline_failure_policy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_enumTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_pipeline_failure_policy_proto = out.File
	file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_pipeline_failure_policy_proto_depIdxs = nil
}
