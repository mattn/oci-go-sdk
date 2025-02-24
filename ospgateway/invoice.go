// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Billing Center Gateway API
//
// This site describes all the Rest endpoints of Billing Center Gateway.
//

package ospgateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// Invoice Invoice details
type Invoice struct {

	// Invoice identifier which is generated on the on-premise sie. Pls note this is not an OCID
	InvoiceId *string `mandatory:"true" json:"invoiceId"`

	// Invoice external reference
	InvoiceNumber *string `mandatory:"false" json:"invoiceNumber"`

	// Transaction identifier
	InternalInvoiceId *string `mandatory:"false" json:"internalInvoiceId"`

	// Is credit card payment eligible
	IsCreditCardPayable *bool `mandatory:"false" json:"isCreditCardPayable"`

	// Date of invoice
	TimeInvoice *common.SDKTime `mandatory:"false" json:"timeInvoice"`

	// Tax of invoice amount
	Tax *float32 `mandatory:"false" json:"tax"`

	// Total amount of invoice
	InvoiceAmount *float32 `mandatory:"false" json:"invoiceAmount"`

	// Balance of invoice
	InvoiceAmountDue *float32 `mandatory:"false" json:"invoiceAmountDue"`

	// Invoice amount credit
	InvoiceAmountCredited *float32 `mandatory:"false" json:"invoiceAmountCredited"`

	// Invoice amount adjust
	InvoiceAmountAdjusted *float32 `mandatory:"false" json:"invoiceAmountAdjusted"`

	// Invoice amount applied
	InvoiceAmountApplied *float32 `mandatory:"false" json:"invoiceAmountApplied"`

	Currency *Currency `mandatory:"false" json:"currency"`

	// Type of invoice
	InvoiceType InvoiceInvoiceTypeEnum `mandatory:"false" json:"invoiceType,omitempty"`

	// Due date of invoice
	TimeInvoiceDue *common.SDKTime `mandatory:"false" json:"timeInvoiceDue"`

	// Invoice reference number
	InvoiceRefNumber *string `mandatory:"false" json:"invoiceRefNumber"`

	// Invoice PO number
	InvoicePoNumber *string `mandatory:"false" json:"invoicePoNumber"`

	// Invoice status
	InvoiceStatus InvoiceInvoiceStatusEnum `mandatory:"false" json:"invoiceStatus,omitempty"`

	// Preferred Email on the invoice
	PreferredEmail *string `mandatory:"false" json:"preferredEmail"`

	// Is emailing pdf allowed
	IsPdfEmailAvailable *bool `mandatory:"false" json:"isPdfEmailAvailable"`

	// Is pdf download access allowed
	IsDisplayDownloadPdf *bool `mandatory:"false" json:"isDisplayDownloadPdf"`

	// Whether invoice can be payed
	IsPayable *bool `mandatory:"false" json:"isPayable"`

	// Payment terms
	PaymentTerms *string `mandatory:"false" json:"paymentTerms"`

	LastPaymentDetail PaymentDetail `mandatory:"false" json:"lastPaymentDetail"`

	BillToAddress *Address `mandatory:"false" json:"billToAddress"`

	// List of subscription identifiers
	SubscriptionIds []string `mandatory:"false" json:"subscriptionIds"`
}

func (m Invoice) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Invoice) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		InvoiceNumber         *string                  `json:"invoiceNumber"`
		InternalInvoiceId     *string                  `json:"internalInvoiceId"`
		IsCreditCardPayable   *bool                    `json:"isCreditCardPayable"`
		TimeInvoice           *common.SDKTime          `json:"timeInvoice"`
		Tax                   *float32                 `json:"tax"`
		InvoiceAmount         *float32                 `json:"invoiceAmount"`
		InvoiceAmountDue      *float32                 `json:"invoiceAmountDue"`
		InvoiceAmountCredited *float32                 `json:"invoiceAmountCredited"`
		InvoiceAmountAdjusted *float32                 `json:"invoiceAmountAdjusted"`
		InvoiceAmountApplied  *float32                 `json:"invoiceAmountApplied"`
		Currency              *Currency                `json:"currency"`
		InvoiceType           InvoiceInvoiceTypeEnum   `json:"invoiceType"`
		TimeInvoiceDue        *common.SDKTime          `json:"timeInvoiceDue"`
		InvoiceRefNumber      *string                  `json:"invoiceRefNumber"`
		InvoicePoNumber       *string                  `json:"invoicePoNumber"`
		InvoiceStatus         InvoiceInvoiceStatusEnum `json:"invoiceStatus"`
		PreferredEmail        *string                  `json:"preferredEmail"`
		IsPdfEmailAvailable   *bool                    `json:"isPdfEmailAvailable"`
		IsDisplayDownloadPdf  *bool                    `json:"isDisplayDownloadPdf"`
		IsPayable             *bool                    `json:"isPayable"`
		PaymentTerms          *string                  `json:"paymentTerms"`
		LastPaymentDetail     paymentdetail            `json:"lastPaymentDetail"`
		BillToAddress         *Address                 `json:"billToAddress"`
		SubscriptionIds       []string                 `json:"subscriptionIds"`
		InvoiceId             *string                  `json:"invoiceId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.InvoiceNumber = model.InvoiceNumber

	m.InternalInvoiceId = model.InternalInvoiceId

	m.IsCreditCardPayable = model.IsCreditCardPayable

	m.TimeInvoice = model.TimeInvoice

	m.Tax = model.Tax

	m.InvoiceAmount = model.InvoiceAmount

	m.InvoiceAmountDue = model.InvoiceAmountDue

	m.InvoiceAmountCredited = model.InvoiceAmountCredited

	m.InvoiceAmountAdjusted = model.InvoiceAmountAdjusted

	m.InvoiceAmountApplied = model.InvoiceAmountApplied

	m.Currency = model.Currency

	m.InvoiceType = model.InvoiceType

	m.TimeInvoiceDue = model.TimeInvoiceDue

	m.InvoiceRefNumber = model.InvoiceRefNumber

	m.InvoicePoNumber = model.InvoicePoNumber

	m.InvoiceStatus = model.InvoiceStatus

	m.PreferredEmail = model.PreferredEmail

	m.IsPdfEmailAvailable = model.IsPdfEmailAvailable

	m.IsDisplayDownloadPdf = model.IsDisplayDownloadPdf

	m.IsPayable = model.IsPayable

	m.PaymentTerms = model.PaymentTerms

	nn, e = model.LastPaymentDetail.UnmarshalPolymorphicJSON(model.LastPaymentDetail.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LastPaymentDetail = nn.(PaymentDetail)
	} else {
		m.LastPaymentDetail = nil
	}

	m.BillToAddress = model.BillToAddress

	m.SubscriptionIds = make([]string, len(model.SubscriptionIds))
	for i, n := range model.SubscriptionIds {
		m.SubscriptionIds[i] = n
	}

	m.InvoiceId = model.InvoiceId

	return
}

// InvoiceInvoiceTypeEnum Enum with underlying type: string
type InvoiceInvoiceTypeEnum string

// Set of constants representing the allowable values for InvoiceInvoiceTypeEnum
const (
	InvoiceInvoiceTypeHardware     InvoiceInvoiceTypeEnum = "HARDWARE"
	InvoiceInvoiceTypeSubscription InvoiceInvoiceTypeEnum = "SUBSCRIPTION"
	InvoiceInvoiceTypeSupport      InvoiceInvoiceTypeEnum = "SUPPORT"
	InvoiceInvoiceTypeLicense      InvoiceInvoiceTypeEnum = "LICENSE"
	InvoiceInvoiceTypeEducation    InvoiceInvoiceTypeEnum = "EDUCATION"
	InvoiceInvoiceTypeConsulting   InvoiceInvoiceTypeEnum = "CONSULTING"
	InvoiceInvoiceTypeService      InvoiceInvoiceTypeEnum = "SERVICE"
	InvoiceInvoiceTypeUsage        InvoiceInvoiceTypeEnum = "USAGE"
)

var mappingInvoiceInvoiceType = map[string]InvoiceInvoiceTypeEnum{
	"HARDWARE":     InvoiceInvoiceTypeHardware,
	"SUBSCRIPTION": InvoiceInvoiceTypeSubscription,
	"SUPPORT":      InvoiceInvoiceTypeSupport,
	"LICENSE":      InvoiceInvoiceTypeLicense,
	"EDUCATION":    InvoiceInvoiceTypeEducation,
	"CONSULTING":   InvoiceInvoiceTypeConsulting,
	"SERVICE":      InvoiceInvoiceTypeService,
	"USAGE":        InvoiceInvoiceTypeUsage,
}

// GetInvoiceInvoiceTypeEnumValues Enumerates the set of values for InvoiceInvoiceTypeEnum
func GetInvoiceInvoiceTypeEnumValues() []InvoiceInvoiceTypeEnum {
	values := make([]InvoiceInvoiceTypeEnum, 0)
	for _, v := range mappingInvoiceInvoiceType {
		values = append(values, v)
	}
	return values
}

// InvoiceInvoiceStatusEnum Enum with underlying type: string
type InvoiceInvoiceStatusEnum string

// Set of constants representing the allowable values for InvoiceInvoiceStatusEnum
const (
	InvoiceInvoiceStatusOpen             InvoiceInvoiceStatusEnum = "OPEN"
	InvoiceInvoiceStatusPastDue          InvoiceInvoiceStatusEnum = "PAST_DUE"
	InvoiceInvoiceStatusPaymentSubmitted InvoiceInvoiceStatusEnum = "PAYMENT_SUBMITTED"
	InvoiceInvoiceStatusClosed           InvoiceInvoiceStatusEnum = "CLOSED"
)

var mappingInvoiceInvoiceStatus = map[string]InvoiceInvoiceStatusEnum{
	"OPEN":              InvoiceInvoiceStatusOpen,
	"PAST_DUE":          InvoiceInvoiceStatusPastDue,
	"PAYMENT_SUBMITTED": InvoiceInvoiceStatusPaymentSubmitted,
	"CLOSED":            InvoiceInvoiceStatusClosed,
}

// GetInvoiceInvoiceStatusEnumValues Enumerates the set of values for InvoiceInvoiceStatusEnum
func GetInvoiceInvoiceStatusEnumValues() []InvoiceInvoiceStatusEnum {
	values := make([]InvoiceInvoiceStatusEnum, 0)
	for _, v := range mappingInvoiceInvoiceStatus {
		values = append(values, v)
	}
	return values
}
