// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// AnnotationCollection Results of a annotations search. Contains boh AnnotationSummary items and other information, such as metadata.
type AnnotationCollection struct {

	// List of annotations.
	Items []AnnotationSummary `mandatory:"true" json:"items"`
}

func (m AnnotationCollection) String() string {
	return common.PointerString(m)
}
