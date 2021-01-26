// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v34/common"
)

// LogAnalyticsEntityTypeSummary Summary of a log analytics entity type.
type LogAnalyticsEntityTypeSummary struct {

	// Log analytics entity type name.
	Name *string `mandatory:"true" json:"name"`

	// Internal name for the log analytics entity type.
	InternalName *string `mandatory:"true" json:"internalName"`

	// Log analytics entity type category. Category will be used for grouping and filtering.
	Category *string `mandatory:"true" json:"category"`

	// Log analytics entity type group. This can be CLOUD (OCI) or NON_CLOUD otherwise.
	CloudType EntityCloudTypeEnum `mandatory:"true" json:"cloudType"`

	// The current lifecycle state of the log analytics entity type.
	LifecycleState EntityLifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Time the log analytics entity type was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time the log analytics entity type was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m LogAnalyticsEntityTypeSummary) String() string {
	return common.PointerString(m)
}
