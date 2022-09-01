// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OperationSummaryFromProcedure The operation object.
type OperationSummaryFromProcedure struct {
	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The model version of the object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// The operation name.
	Name *string `mandatory:"false" json:"name"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key of the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The status of an object that can be set to value 1 for shallow reference across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

//GetMetadata returns Metadata
func (m OperationSummaryFromProcedure) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m OperationSummaryFromProcedure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperationSummaryFromProcedure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OperationSummaryFromProcedure) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOperationSummaryFromProcedure OperationSummaryFromProcedure
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOperationSummaryFromProcedure
	}{
		"PROCEDURE",
		(MarshalTypeOperationSummaryFromProcedure)(m),
	}

	return json.Marshal(&s)
}
