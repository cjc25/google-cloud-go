// Copyright 2024 Google LLC
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

//go:build go1.23

package apigeeregistry_test

import (
	"context"

	apigeeregistry "cloud.google.com/go/apigeeregistry/apiv1"
	apigeeregistrypb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

func ExampleRegistryClient_ListApiDeploymentRevisions_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigeeregistrypb.ListApiDeploymentRevisionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb#ListApiDeploymentRevisionsRequest.
	}
	for resp, err := range c.ListApiDeploymentRevisions(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListApiDeployments_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigeeregistrypb.ListApiDeploymentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb#ListApiDeploymentsRequest.
	}
	for resp, err := range c.ListApiDeployments(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListApiSpecRevisions_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigeeregistrypb.ListApiSpecRevisionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb#ListApiSpecRevisionsRequest.
	}
	for resp, err := range c.ListApiSpecRevisions(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListApiSpecs_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigeeregistrypb.ListApiSpecsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb#ListApiSpecsRequest.
	}
	for resp, err := range c.ListApiSpecs(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListApiVersions_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigeeregistrypb.ListApiVersionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb#ListApiVersionsRequest.
	}
	for resp, err := range c.ListApiVersions(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListApis_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigeeregistrypb.ListApisRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb#ListApisRequest.
	}
	for resp, err := range c.ListApis(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListArtifacts_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &apigeeregistrypb.ListArtifactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb#ListArtifactsRequest.
	}
	for resp, err := range c.ListArtifacts(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListLocations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	for resp, err := range c.ListLocations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_ListOperations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := apigeeregistry.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	}
	for resp, err := range c.ListOperations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}