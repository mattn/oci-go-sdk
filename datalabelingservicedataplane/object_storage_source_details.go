// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DlsDataPlane API
//
// A description of the DlsDataPlane API.
//

package datalabelingservicedataplane

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// ObjectStorageSourceDetails Object Storage Source Details.
type ObjectStorageSourceDetails struct {

	// The path relative to the prefix specified in the dataset source details (file name).
	RelativePath *string `mandatory:"true" json:"relativePath"`

	// The full path of the file this record belongs to.
	Path *string `mandatory:"true" json:"path"`

	// The offset into the file containing the content.
	Offset *float32 `mandatory:"false" json:"offset"`

	// The length from offset into the file containing the content.
	Length *float32 `mandatory:"false" json:"length"`
}

func (m ObjectStorageSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageSourceDetails ObjectStorageSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeObjectStorageSourceDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageSourceDetails)(m),
	}

	return json.Marshal(&s)
}
