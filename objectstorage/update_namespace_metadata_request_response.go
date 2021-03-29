// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v38/common"
	"net/http"
)

// UpdateNamespaceMetadataRequest wrapper for the UpdateNamespaceMetadata operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/UpdateNamespaceMetadata.go.html to see an example of how to use UpdateNamespaceMetadataRequest.
type UpdateNamespaceMetadataRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Request object for update NamespaceMetadata.
	UpdateNamespaceMetadataDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateNamespaceMetadataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateNamespaceMetadataRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateNamespaceMetadataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateNamespaceMetadataResponse wrapper for the UpdateNamespaceMetadata operation
type UpdateNamespaceMetadataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The NamespaceMetadata instance
	NamespaceMetadata `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateNamespaceMetadataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateNamespaceMetadataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
