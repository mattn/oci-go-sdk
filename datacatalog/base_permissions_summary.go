// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v29/common"
)

// BasePermissionsSummary Permissions object sent as part of the response.
type BasePermissionsSummary struct {

	// An array of permissions.
	UserPermissions []string `mandatory:"false" json:"userPermissions"`
}

func (m BasePermissionsSummary) String() string {
	return common.PointerString(m)
}
