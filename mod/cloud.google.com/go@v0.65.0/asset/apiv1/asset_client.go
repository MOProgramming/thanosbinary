// Copyright 2020 Google LLC
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

package asset

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ExportAssets            []gax.CallOption
	BatchGetAssetsHistory   []gax.CallOption
	CreateFeed              []gax.CallOption
	GetFeed                 []gax.CallOption
	ListFeeds               []gax.CallOption
	UpdateFeed              []gax.CallOption
	DeleteFeed              []gax.CallOption
	SearchAllResources      []gax.CallOption
	SearchAllIamPolicies    []gax.CallOption
	AnalyzeIamPolicy        []gax.CallOption
	ExportIamPolicyAnalysis []gax.CallOption
}

func defaultClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("cloudasset.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ExportAssets: []gax.CallOption{},
		BatchGetAssetsHistory: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateFeed: []gax.CallOption{},
		GetFeed: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListFeeds: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateFeed: []gax.CallOption{},
		DeleteFeed: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchAllResources: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchAllIamPolicies: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		AnalyzeIamPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ExportIamPolicyAnalysis: []gax.CallOption{},
	}
}

// Client is a client for interacting with Cloud Asset API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type Client struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	client assetpb.AssetServiceClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new asset service client.
//
// Asset service definition.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultClientOptions()

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
	c := &Client{
		connPool:    connPool,
		CallOptions: defaultCallOptions(),

		client: assetpb.NewAssetServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	c.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ExportAssets exports assets with time and resource types to a given Cloud Storage
// location/BigQuery table. For Cloud Storage location destinations, the
// output format is newline-delimited JSON. Each line represents a
// google.cloud.asset.v1.Asset in the JSON format; for BigQuery table
// destinations, the output table stores the fields in asset proto as columns.
// This API implements the google.longrunning.Operation API
// , which allows you to keep track of the export. We recommend intervals of
// at least 2 seconds with exponential retry to poll the export operation
// result. For regular-size resource parent, the export operation usually
// finishes within 5 minutes.
func (c *Client) ExportAssets(ctx context.Context, req *assetpb.ExportAssetsRequest, opts ...gax.CallOption) (*ExportAssetsOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ExportAssets[0:len(c.CallOptions.ExportAssets):len(c.CallOptions.ExportAssets)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ExportAssets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportAssetsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// BatchGetAssetsHistory batch gets the update history of assets that overlap a time window.
// For IAM_POLICY content, this API outputs history when the asset and its
// attached IAM POLICY both exist. This can create gaps in the output history.
// Otherwise, this API outputs history with asset in both non-delete or
// deleted status.
// If a specified asset does not exist, this API returns an INVALID_ARGUMENT
// error.
func (c *Client) BatchGetAssetsHistory(ctx context.Context, req *assetpb.BatchGetAssetsHistoryRequest, opts ...gax.CallOption) (*assetpb.BatchGetAssetsHistoryResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchGetAssetsHistory[0:len(c.CallOptions.BatchGetAssetsHistory):len(c.CallOptions.BatchGetAssetsHistory)], opts...)
	var resp *assetpb.BatchGetAssetsHistoryResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.BatchGetAssetsHistory(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateFeed creates a feed in a parent project/folder/organization to listen to its
// asset updates.
func (c *Client) CreateFeed(ctx context.Context, req *assetpb.CreateFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateFeed[0:len(c.CallOptions.CreateFeed):len(c.CallOptions.CreateFeed)], opts...)
	var resp *assetpb.Feed
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFeed gets details about an asset feed.
func (c *Client) GetFeed(ctx context.Context, req *assetpb.GetFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetFeed[0:len(c.CallOptions.GetFeed):len(c.CallOptions.GetFeed)], opts...)
	var resp *assetpb.Feed
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListFeeds lists all asset feeds in a parent project/folder/organization.
func (c *Client) ListFeeds(ctx context.Context, req *assetpb.ListFeedsRequest, opts ...gax.CallOption) (*assetpb.ListFeedsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListFeeds[0:len(c.CallOptions.ListFeeds):len(c.CallOptions.ListFeeds)], opts...)
	var resp *assetpb.ListFeedsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ListFeeds(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateFeed updates an asset feed configuration.
func (c *Client) UpdateFeed(ctx context.Context, req *assetpb.UpdateFeedRequest, opts ...gax.CallOption) (*assetpb.Feed, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "feed.name", url.QueryEscape(req.GetFeed().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateFeed[0:len(c.CallOptions.UpdateFeed):len(c.CallOptions.UpdateFeed)], opts...)
	var resp *assetpb.Feed
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteFeed deletes an asset feed.
func (c *Client) DeleteFeed(ctx context.Context, req *assetpb.DeleteFeedRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteFeed[0:len(c.CallOptions.DeleteFeed):len(c.CallOptions.DeleteFeed)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteFeed(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SearchAllResources searches all Cloud resources within the specified scope, such as a project,
// folder, or organization. The caller must be granted the
// cloudasset.assets.searchAllResources permission on the desired scope,
// otherwise the request will be rejected.
func (c *Client) SearchAllResources(ctx context.Context, req *assetpb.SearchAllResourcesRequest, opts ...gax.CallOption) *ResourceSearchResultIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "scope", url.QueryEscape(req.GetScope())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SearchAllResources[0:len(c.CallOptions.SearchAllResources):len(c.CallOptions.SearchAllResources)], opts...)
	it := &ResourceSearchResultIterator{}
	req = proto.Clone(req).(*assetpb.SearchAllResourcesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*assetpb.ResourceSearchResult, string, error) {
		var resp *assetpb.SearchAllResourcesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.SearchAllResources(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

// SearchAllIamPolicies searches all IAM policies within the specified scope, such as a project,
// folder, or organization. The caller must be granted the
// cloudasset.assets.searchAllIamPolicies permission on the desired scope,
// otherwise the request will be rejected.
func (c *Client) SearchAllIamPolicies(ctx context.Context, req *assetpb.SearchAllIamPoliciesRequest, opts ...gax.CallOption) *IamPolicySearchResultIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "scope", url.QueryEscape(req.GetScope())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SearchAllIamPolicies[0:len(c.CallOptions.SearchAllIamPolicies):len(c.CallOptions.SearchAllIamPolicies)], opts...)
	it := &IamPolicySearchResultIterator{}
	req = proto.Clone(req).(*assetpb.SearchAllIamPoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*assetpb.IamPolicySearchResult, string, error) {
		var resp *assetpb.SearchAllIamPoliciesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.SearchAllIamPolicies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetResults(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

// AnalyzeIamPolicy analyzes IAM policies to answer which identities have what accesses on
// which resources.
func (c *Client) AnalyzeIamPolicy(ctx context.Context, req *assetpb.AnalyzeIamPolicyRequest, opts ...gax.CallOption) (*assetpb.AnalyzeIamPolicyResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "analysis_query.scope", url.QueryEscape(req.GetAnalysisQuery().GetScope())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.AnalyzeIamPolicy[0:len(c.CallOptions.AnalyzeIamPolicy):len(c.CallOptions.AnalyzeIamPolicy)], opts...)
	var resp *assetpb.AnalyzeIamPolicyResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.AnalyzeIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ExportIamPolicyAnalysis exports the answers of which identities have what accesses on which
// resources to a Google Cloud Storage or a BigQuery destination. For Cloud
// Storage destination, the output format is the JSON format that represents a
// google.cloud.asset.v1.AnalyzeIamPolicyResponse.
// This method implements the
// google.longrunning.Operation, which allows
// you to track the export status. We recommend intervals of at least 2
// seconds with exponential retry to poll the export operation result. The
// metadata contains the request to help callers to map responses to requests.
func (c *Client) ExportIamPolicyAnalysis(ctx context.Context, req *assetpb.ExportIamPolicyAnalysisRequest, opts ...gax.CallOption) (*ExportIamPolicyAnalysisOperation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "analysis_query.scope", url.QueryEscape(req.GetAnalysisQuery().GetScope())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ExportIamPolicyAnalysis[0:len(c.CallOptions.ExportIamPolicyAnalysis):len(c.CallOptions.ExportIamPolicyAnalysis)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ExportIamPolicyAnalysis(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportIamPolicyAnalysisOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// ExportAssetsOperation manages a long-running operation from ExportAssets.
type ExportAssetsOperation struct {
	lro *longrunning.Operation
}

// ExportAssetsOperation returns a new ExportAssetsOperation from a given name.
// The name must be that of a previously created ExportAssetsOperation, possibly from a different process.
func (c *Client) ExportAssetsOperation(name string) *ExportAssetsOperation {
	return &ExportAssetsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportAssetsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*assetpb.ExportAssetsResponse, error) {
	var resp assetpb.ExportAssetsResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ExportAssetsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*assetpb.ExportAssetsResponse, error) {
	var resp assetpb.ExportAssetsResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ExportAssetsOperation) Metadata() (*assetpb.ExportAssetsRequest, error) {
	var meta assetpb.ExportAssetsRequest
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportAssetsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportAssetsOperation) Name() string {
	return op.lro.Name()
}

// ExportIamPolicyAnalysisOperation manages a long-running operation from ExportIamPolicyAnalysis.
type ExportIamPolicyAnalysisOperation struct {
	lro *longrunning.Operation
}

// ExportIamPolicyAnalysisOperation returns a new ExportIamPolicyAnalysisOperation from a given name.
// The name must be that of a previously created ExportIamPolicyAnalysisOperation, possibly from a different process.
func (c *Client) ExportIamPolicyAnalysisOperation(name string) *ExportIamPolicyAnalysisOperation {
	return &ExportIamPolicyAnalysisOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportIamPolicyAnalysisOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*assetpb.ExportIamPolicyAnalysisResponse, error) {
	var resp assetpb.ExportIamPolicyAnalysisResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ExportIamPolicyAnalysisOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*assetpb.ExportIamPolicyAnalysisResponse, error) {
	var resp assetpb.ExportIamPolicyAnalysisResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ExportIamPolicyAnalysisOperation) Metadata() (*assetpb.ExportIamPolicyAnalysisRequest, error) {
	var meta assetpb.ExportIamPolicyAnalysisRequest
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportIamPolicyAnalysisOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportIamPolicyAnalysisOperation) Name() string {
	return op.lro.Name()
}

// IamPolicySearchResultIterator manages a stream of *assetpb.IamPolicySearchResult.
type IamPolicySearchResultIterator struct {
	items    []*assetpb.IamPolicySearchResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.IamPolicySearchResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IamPolicySearchResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IamPolicySearchResultIterator) Next() (*assetpb.IamPolicySearchResult, error) {
	var item *assetpb.IamPolicySearchResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IamPolicySearchResultIterator) bufLen() int {
	return len(it.items)
}

func (it *IamPolicySearchResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ResourceSearchResultIterator manages a stream of *assetpb.ResourceSearchResult.
type ResourceSearchResultIterator struct {
	items    []*assetpb.ResourceSearchResult
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*assetpb.ResourceSearchResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ResourceSearchResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ResourceSearchResultIterator) Next() (*assetpb.ResourceSearchResult, error) {
	var item *assetpb.ResourceSearchResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ResourceSearchResultIterator) bufLen() int {
	return len(it.items)
}

func (it *ResourceSearchResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
