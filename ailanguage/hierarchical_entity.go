// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
//     pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
//     This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// HierarchicalEntity Hierarchical entity object
type HierarchicalEntity struct {

	// The number of Unicode code points preceding this entity in the submitted text.
	Offset *int `mandatory:"false" json:"offset"`

	// Length of entity text
	Length *int `mandatory:"false" json:"length"`

	// Entity text like name of person, location, and so on.
	Text *string `mandatory:"false" json:"text"`

	// Type of entity text like PER, LOC.
	Type *string `mandatory:"false" json:"type"`

	// Sub-type of entity text like GPE for LOCATION type
	SubType *string `mandatory:"false" json:"subType"`

	// Score or confidence for detected entity.
	Score *float64 `mandatory:"false" json:"score"`
}

func (m HierarchicalEntity) String() string {
	return common.PointerString(m)
}
