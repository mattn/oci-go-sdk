// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// ContextualData The representation of ContextualData
type ContextualData struct {

	// The unique client identifier
	ClientId *string `mandatory:"true" json:"clientId"`

	// The schema name
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The schema version
	SchemaVersion *string `mandatory:"true" json:"schemaVersion"`

	// The context data payload
	Payload *string `mandatory:"true" json:"payload"`
}

func (m ContextualData) String() string {
	return common.PointerString(m)
}
