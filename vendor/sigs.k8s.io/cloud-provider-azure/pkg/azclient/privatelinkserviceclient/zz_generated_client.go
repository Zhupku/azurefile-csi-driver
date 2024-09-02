// /*
// Copyright The Kubernetes Authors.
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
// */

// Code generated by client-gen. DO NOT EDIT.
package privatelinkserviceclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/metrics"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

type Client struct {
	*armnetwork.PrivateLinkServicesClient
	subscriptionID string
	tracer         tracing.Tracer
}

func New(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (Interface, error) {
	if options == nil {
		options = utils.GetDefaultOption()
	}
	tr := options.TracingProvider.NewTracer(utils.ModuleName, utils.ModuleVersion)

	client, err := armnetwork.NewPrivateLinkServicesClient(subscriptionID, credential, options)
	if err != nil {
		return nil, err
	}
	return &Client{
		PrivateLinkServicesClient: client,
		subscriptionID:            subscriptionID,
		tracer:                    tr,
	}, nil
}

const GetOperationName = "PrivateLinkServicesClient.Get"

// Get gets the PrivateLinkService
func (client *Client) Get(ctx context.Context, resourceGroupName string, resourceName string, expand *string) (result *armnetwork.PrivateLinkService, err error) {
	var ops *armnetwork.PrivateLinkServicesClientGetOptions
	if expand != nil {
		ops = &armnetwork.PrivateLinkServicesClientGetOptions{Expand: expand}
	}
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "PrivateLinkService", "get")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, GetOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := client.PrivateLinkServicesClient.Get(ctx, resourceGroupName, resourceName, ops)
	if err != nil {
		return nil, err
	}
	//handle statuscode
	return &resp.PrivateLinkService, nil
}

const CreateOrUpdateOperationName = "PrivateLinkServicesClient.Create"

// CreateOrUpdate creates or updates a PrivateLinkService.
func (client *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, resource armnetwork.PrivateLinkService) (result *armnetwork.PrivateLinkService, err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "PrivateLinkService", "create_or_update")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, CreateOrUpdateOperationName, client.tracer, nil)
	defer endSpan(err)
	resp, err := utils.NewPollerWrapper(client.PrivateLinkServicesClient.BeginCreateOrUpdate(ctx, resourceGroupName, resourceName, resource, nil)).WaitforPollerResp(ctx)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		return &resp.PrivateLinkService, nil
	}
	return nil, nil
}

const DeleteOperationName = "PrivateLinkServicesClient.Delete"

// Delete deletes a PrivateLinkService by name.
func (client *Client) Delete(ctx context.Context, resourceGroupName string, resourceName string) (err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "PrivateLinkService", "delete")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, DeleteOperationName, client.tracer, nil)
	defer endSpan(err)
	_, err = utils.NewPollerWrapper(client.BeginDelete(ctx, resourceGroupName, resourceName, nil)).WaitforPollerResp(ctx)
	return err
}

const ListOperationName = "PrivateLinkServicesClient.List"

// List gets a list of PrivateLinkService in the resource group.
func (client *Client) List(ctx context.Context, resourceGroupName string) (result []*armnetwork.PrivateLinkService, err error) {
	metricsCtx := metrics.BeginARMRequest(client.subscriptionID, resourceGroupName, "PrivateLinkService", "list")
	defer metricsCtx.Observe(ctx, err)
	ctx, endSpan := runtime.StartSpan(ctx, ListOperationName, client.tracer, nil)
	defer endSpan(err)
	pager := client.PrivateLinkServicesClient.NewListPager(resourceGroupName, nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, nextResult.Value...)
	}
	return result, nil
}
