// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// UpdateDomainGovernanceDetails The parameters for updating a domain govenance entity.
type UpdateDomainGovernanceDetails struct {

	// The email to notify the user, and that the ONS subscription will be created with. The ONS subscription for the previous email will also be deleted.
	SubscriptionEmail *string `mandatory:"false" json:"subscriptionEmail"`

	// Indicates whether governance is enabled for this domain.
	IsGovernanceEnabled *bool `mandatory:"false" json:"isGovernanceEnabled"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateDomainGovernanceDetails) String() string {
	return common.PointerString(m)
}
