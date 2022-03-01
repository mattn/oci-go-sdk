// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// VisionService API
//
// A description of the VisionService API.
//

package aivision

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v60/common"
	"strings"
)

// ObjectStorageDataset Dataset that resides in OCI Object Storage.
type ObjectStorageDataset struct {

	// The namespace name of the Object Storage bucket that contains the input data file.
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// The name of the Object Storage bucket that contains the input data file.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The object name of the input data file.
	ObjectName *string `mandatory:"false" json:"objectName"`
}

func (m ObjectStorageDataset) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageDataset) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ObjectStorageDataset) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageDataset ObjectStorageDataset
	s := struct {
		DiscriminatorParam string `json:"datasetType"`
		MarshalTypeObjectStorageDataset
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageDataset)(m),
	}

	return json.Marshal(&s)
}
