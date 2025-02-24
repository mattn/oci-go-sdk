// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Billing Center Gateway API
//
// This site describes all the Rest endpoints of Billing Center Gateway.
//

package ospgateway

import (
	"github.com/oracle/oci-go-sdk/v54/common"
)

// InvoiceLineSummary Product items of the invoice
type InvoiceLineSummary struct {

	// Product of the item
	Product *string `mandatory:"true" json:"product"`

	// Product of the item
	OrderNo *string `mandatory:"false" json:"orderNo"`

	// Part number
	PartNumber *string `mandatory:"false" json:"partNumber"`

	// Start date
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// End date
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Quantity of the ordered product
	Quantity *float32 `mandatory:"false" json:"quantity"`

	// Unit price of the ordered product
	NetUnitPrice *float32 `mandatory:"false" json:"netUnitPrice"`

	// Total price of the ordered product (Net unit price x quantity)
	TotalPrice *float32 `mandatory:"false" json:"totalPrice"`

	Currency *Currency `mandatory:"false" json:"currency"`
}

func (m InvoiceLineSummary) String() string {
	return common.PointerString(m)
}
