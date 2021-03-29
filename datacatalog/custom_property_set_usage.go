// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v38/common"
)

// CustomPropertySetUsage Details of a single custom property.
type CustomPropertySetUsage struct {

	// Unique Identifier of the attribute which is ID
	Key *string `mandatory:"false" json:"key"`

	// Name of the custom property
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The custom property value
	Value *string `mandatory:"false" json:"value"`

	// Namespace name of the custom property
	NamespaceName *string `mandatory:"false" json:"namespaceName"`
}

func (m CustomPropertySetUsage) String() string {
	return common.PointerString(m)
}
