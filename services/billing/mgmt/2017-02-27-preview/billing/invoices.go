package billing

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// InvoicesClient is the billing client provides access to billing resources for Azure Web-Direct subscriptions. Other
// subscription types which were not purchased directly through the Azure web portal are not supported through this
// preview API.
type InvoicesClient struct {
	ManagementClient
}

// NewInvoicesClient creates an instance of the InvoicesClient client.
func NewInvoicesClient(subscriptionID string) InvoicesClient {
	return NewInvoicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoicesClientWithBaseURI creates an instance of the InvoicesClient client.
func NewInvoicesClientWithBaseURI(baseURI string, subscriptionID string) InvoicesClient {
	return InvoicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a named invoice resource. When getting a single invoice, the downloadUrl property is expanded
// automatically.
//
// invoiceName is the name of an invoice resource.
func (client InvoicesClient) Get(invoiceName string) (result Invoice, err error) {
	req, err := client.GetPreparer(invoiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InvoicesClient) GetPreparer(invoiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"invoiceName":    autorest.Encode("path", invoiceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-27-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/invoices/{invoiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InvoicesClient) GetResponder(resp *http.Response) (result Invoice, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLatest gets the most recent invoice. When getting a single invoice, the downloadUrl property is expanded
// automatically.
func (client InvoicesClient) GetLatest() (result Invoice, err error) {
	req, err := client.GetLatestPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "GetLatest", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLatestSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "GetLatest", resp, "Failure sending request")
		return
	}

	result, err = client.GetLatestResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "GetLatest", resp, "Failure responding to request")
	}

	return
}

// GetLatestPreparer prepares the GetLatest request.
func (client InvoicesClient) GetLatestPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-27-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/invoices/latest", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetLatestSender sends the GetLatest request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicesClient) GetLatestSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetLatestResponder handles the response to the GetLatest request. The method always
// closes the http.Response Body.
func (client InvoicesClient) GetLatestResponder(resp *http.Response) (result Invoice, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the available invoices for a subscription in reverse chronological order beginning with the most recent
// invoice. In preview, invoices are available via this API only for invoice periods which end December 1, 2016 or
// later
//
// expand is may be used to expand the downloadUrl property within a list of invoices. This enables download links to
// be generated for multiple invoices at once. By default, downloadURLs are not included when listing invoices. filter
// is may be used to filter invoices by invoicePeriodEndDate. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and
// 'and'. It does not currently support 'ne', 'or', or 'not' skiptoken is skiptoken is only used if a previous
// operation returned a partial result. If a previous response contains a nextLink element, the value of the nextLink
// element will include a skiptoken parameter that specifies a starting point to use for subsequent calls. top is may
// be used to limit the number of results to the most recent N invoices.
func (client InvoicesClient) List(expand string, filter string, skiptoken string, top *int32) (result InvoicesListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: 100, Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "billing.InvoicesClient", "List")
	}

	req, err := client.ListPreparer(expand, filter, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InvoicesClient) ListPreparer(expand string, filter string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-02-27-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/invoices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InvoicesClient) ListResponder(resp *http.Response) (result InvoicesListResult, err error) {
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
func (client InvoicesClient) ListNextResults(lastResults InvoicesListResult) (result InvoicesListResult, err error) {
	req, err := lastResults.InvoicesListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client InvoicesClient) ListComplete(expand string, filter string, skiptoken string, top *int32, cancel <-chan struct{}) (<-chan Invoice, <-chan error) {
	resultChan := make(chan Invoice)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(expand, filter, skiptoken, top)
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
