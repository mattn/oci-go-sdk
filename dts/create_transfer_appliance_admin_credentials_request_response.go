// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dts

import (
	"github.com/oracle/oci-go-sdk/v54/common"
	"net/http"
)

// CreateTransferApplianceAdminCredentialsRequest wrapper for the CreateTransferApplianceAdminCredentials operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dts/CreateTransferApplianceAdminCredentials.go.html to see an example of how to use CreateTransferApplianceAdminCredentialsRequest.
type CreateTransferApplianceAdminCredentialsRequest struct {

	// ID of the Transfer Job
	Id *string `mandatory:"true" contributesTo:"path" name:"id"`

	// Label of the Transfer Appliance
	TransferApplianceLabel *string `mandatory:"true" contributesTo:"path" name:"transferApplianceLabel"`

	AdminPublicKey TransferAppliancePublicKey `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateTransferApplianceAdminCredentialsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateTransferApplianceAdminCredentialsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateTransferApplianceAdminCredentialsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateTransferApplianceAdminCredentialsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateTransferApplianceAdminCredentialsResponse wrapper for the CreateTransferApplianceAdminCredentials operation
type CreateTransferApplianceAdminCredentialsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TransferApplianceCertificate instance
	TransferApplianceCertificate `presentIn:"body"`

	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateTransferApplianceAdminCredentialsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateTransferApplianceAdminCredentialsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
