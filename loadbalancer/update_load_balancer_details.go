// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// UpdateLoadBalancerDetails. Configuration details to update a load balancer.
type UpdateLoadBalancerDetails struct {

	// The user-friendly display name for the load balancer. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My load balancer`
	DisplayName *string `mandatory:"true" json:"displayName,omitempty"`
}

func (model UpdateLoadBalancerDetails) String() string {
	return common.PointerString(model)
}