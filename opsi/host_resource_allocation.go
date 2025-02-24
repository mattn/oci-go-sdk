// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// HostResourceAllocation Resource Allocation metric for the host
type HostResourceAllocation struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Total number of CPUs available
	TotalCpus *int `mandatory:"false" json:"totalCpus"`

	// Total amount of usable physical memory in gibabytes
	TotalMemoryInGB *float64 `mandatory:"false" json:"totalMemoryInGB"`
}

//GetTimeCollected returns TimeCollected
func (m HostResourceAllocation) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostResourceAllocation) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HostResourceAllocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostResourceAllocation HostResourceAllocation
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostResourceAllocation
	}{
		"HOST_RESOURCE_ALLOCATION",
		(MarshalTypeHostResourceAllocation)(m),
	}

	return json.Marshal(&s)
}
