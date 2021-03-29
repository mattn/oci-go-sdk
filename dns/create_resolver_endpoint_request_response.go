// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v38/common"
	"net/http"
)

// CreateResolverEndpointRequest wrapper for the CreateResolverEndpoint operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateResolverEndpoint.go.html to see an example of how to use CreateResolverEndpointRequest.
type CreateResolverEndpointRequest struct {

	// The OCID of the target resolver.
	ResolverId *string `mandatory:"true" contributesTo:"path" name:"resolverId"`

	// Details for creating a new resolver endpoint.
	CreateResolverEndpointDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case
	// of a timeout or server error without risk of executing that same action
	// again. Retry tokens expire after 24 hours, but can be invalidated before
	// then due to conflicting operations (for example, if a resource has been
	// deleted and purged from the system, then a retry of the original creation
	// request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope CreateResolverEndpointScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateResolverEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateResolverEndpointRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateResolverEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateResolverEndpointResponse wrapper for the CreateResolverEndpoint operation
type CreateResolverEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ResolverEndpoint instance
	ResolverEndpoint `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	Etag *string `presentIn:"header" name:"etag"`

	// The full URI of the resource related to the request.
	Location *string `presentIn:"header" name:"location"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response CreateResolverEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateResolverEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateResolverEndpointScopeEnum Enum with underlying type: string
type CreateResolverEndpointScopeEnum string

// Set of constants representing the allowable values for CreateResolverEndpointScopeEnum
const (
	CreateResolverEndpointScopeGlobal  CreateResolverEndpointScopeEnum = "GLOBAL"
	CreateResolverEndpointScopePrivate CreateResolverEndpointScopeEnum = "PRIVATE"
)

var mappingCreateResolverEndpointScope = map[string]CreateResolverEndpointScopeEnum{
	"GLOBAL":  CreateResolverEndpointScopeGlobal,
	"PRIVATE": CreateResolverEndpointScopePrivate,
}

// GetCreateResolverEndpointScopeEnumValues Enumerates the set of values for CreateResolverEndpointScopeEnum
func GetCreateResolverEndpointScopeEnumValues() []CreateResolverEndpointScopeEnum {
	values := make([]CreateResolverEndpointScopeEnum, 0)
	for _, v := range mappingCreateResolverEndpointScope {
		values = append(values, v)
	}
	return values
}
