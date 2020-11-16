// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v29/common"
	"net/http"
)

// UpdateRRSetRequest wrapper for the UpdateRRSet operation
type UpdateRRSetRequest struct {

	// The name or OCID of the target zone.
	ZoneNameOrId *string `mandatory:"true" contributesTo:"path" name:"zoneNameOrId"`

	// The target fully-qualified domain name (FQDN) within the target zone.
	Domain *string `mandatory:"true" contributesTo:"path" name:"domain"`

	// The type of the target RRSet within the target zone.
	Rtype *string `mandatory:"true" contributesTo:"path" name:"rtype"`

	// A full list of records for the RRSet.
	UpdateRrSetDetails `contributesTo:"body"`

	// The `If-Match` header field makes the request method conditional on the
	// existence of at least one current representation of the target resource,
	// when the field-value is `*`, or having a current representation of the
	// target resource that has an entity-tag matching a member of the list of
	// entity-tags provided in the field-value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"If-Match"`

	// The `If-Unmodified-Since` header field makes the request method
	// conditional on the selected representation's last modification date being
	// earlier than or equal to the date provided in the field-value.  This
	// field accomplishes the same purpose as If-Match for cases where the user
	// agent does not have an entity-tag for the representation.
	IfUnmodifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Unmodified-Since"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope UpdateRRSetScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the resource is associated with.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateRRSetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateRRSetRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateRRSetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateRRSetResponse wrapper for the UpdateRRSet operation
type UpdateRRSetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RecordCollection instance
	RecordCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current version of the record collection, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response UpdateRRSetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateRRSetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UpdateRRSetScopeEnum Enum with underlying type: string
type UpdateRRSetScopeEnum string

// Set of constants representing the allowable values for UpdateRRSetScopeEnum
const (
	UpdateRRSetScopeGlobal  UpdateRRSetScopeEnum = "GLOBAL"
	UpdateRRSetScopePrivate UpdateRRSetScopeEnum = "PRIVATE"
)

var mappingUpdateRRSetScope = map[string]UpdateRRSetScopeEnum{
	"GLOBAL":  UpdateRRSetScopeGlobal,
	"PRIVATE": UpdateRRSetScopePrivate,
}

// GetUpdateRRSetScopeEnumValues Enumerates the set of values for UpdateRRSetScopeEnum
func GetUpdateRRSetScopeEnumValues() []UpdateRRSetScopeEnum {
	values := make([]UpdateRRSetScopeEnum, 0)
	for _, v := range mappingUpdateRRSetScope {
		values = append(values, v)
	}
	return values
}
