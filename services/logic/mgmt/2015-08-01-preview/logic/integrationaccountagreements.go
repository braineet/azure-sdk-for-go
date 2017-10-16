package logic

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// IntegrationAccountAgreementsClient is the REST API for Azure Logic Apps.
type IntegrationAccountAgreementsClient struct {
	ManagementClient
}

// NewIntegrationAccountAgreementsClient creates an instance of the IntegrationAccountAgreementsClient client.
func NewIntegrationAccountAgreementsClient(subscriptionID string) IntegrationAccountAgreementsClient {
	return NewIntegrationAccountAgreementsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationAccountAgreementsClientWithBaseURI creates an instance of the IntegrationAccountAgreementsClient
// client.
func NewIntegrationAccountAgreementsClientWithBaseURI(baseURI string, subscriptionID string) IntegrationAccountAgreementsClient {
	return IntegrationAccountAgreementsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an integration account agreement.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. agreementName
// is the integration account agreement name. agreement is the integration account agreement.
func (client IntegrationAccountAgreementsClient) CreateOrUpdate(resourceGroupName string, integrationAccountName string, agreementName string, agreement IntegrationAccountAgreement) (result IntegrationAccountAgreement, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, integrationAccountName, agreementName, agreement)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client IntegrationAccountAgreementsClient) CreateOrUpdatePreparer(resourceGroupName string, integrationAccountName string, agreementName string, agreement IntegrationAccountAgreement) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agreementName":          autorest.Encode("path", agreementName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", pathParameters),
		autorest.WithJSON(agreement),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAgreementsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAgreementsClient) CreateOrUpdateResponder(resp *http.Response) (result IntegrationAccountAgreement, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an integration account agreement.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. agreementName
// is the integration account agreement name.
func (client IntegrationAccountAgreementsClient) Delete(resourceGroupName string, integrationAccountName string, agreementName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, integrationAccountName, agreementName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationAccountAgreementsClient) DeletePreparer(resourceGroupName string, integrationAccountName string, agreementName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agreementName":          autorest.Encode("path", agreementName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAgreementsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAgreementsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets an integration account agreement.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. agreementName
// is the integration account agreement name.
func (client IntegrationAccountAgreementsClient) Get(resourceGroupName string, integrationAccountName string, agreementName string) (result IntegrationAccountAgreement, err error) {
	req, err := client.GetPreparer(resourceGroupName, integrationAccountName, agreementName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client IntegrationAccountAgreementsClient) GetPreparer(resourceGroupName string, integrationAccountName string, agreementName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"agreementName":          autorest.Encode("path", agreementName),
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements/{agreementName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAgreementsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAgreementsClient) GetResponder(resp *http.Response) (result IntegrationAccountAgreement, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of integration account agreements.
//
// resourceGroupName is the resource group name. integrationAccountName is the integration account name. top is the
// number of items to be included in the result. filter is the filter to apply on the operation.
func (client IntegrationAccountAgreementsClient) List(resourceGroupName string, integrationAccountName string, top *int32, filter string) (result IntegrationAccountAgreementListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, integrationAccountName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client IntegrationAccountAgreementsClient) ListPreparer(resourceGroupName string, integrationAccountName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"integrationAccountName": autorest.Encode("path", integrationAccountName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/agreements", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationAccountAgreementsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client IntegrationAccountAgreementsClient) ListResponder(resp *http.Response) (result IntegrationAccountAgreementListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client IntegrationAccountAgreementsClient) ListNextResults(lastResults IntegrationAccountAgreementListResult) (result IntegrationAccountAgreementListResult, err error) {
	req, err := lastResults.IntegrationAccountAgreementListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic.IntegrationAccountAgreementsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client IntegrationAccountAgreementsClient) ListComplete(resourceGroupName string, integrationAccountName string, top *int32, filter string, cancel <-chan struct{}) (<-chan IntegrationAccountAgreement, <-chan error) {
	resultChan := make(chan IntegrationAccountAgreement)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, integrationAccountName, top, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
