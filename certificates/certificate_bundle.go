// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Retrieval API
//
// API for retrieving certificates.
//

package certificates

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// CertificateBundle The contents of the certificate, properties of the certificate (and certificate version), and user-provided contextual metadata for the certificate.
type CertificateBundle interface {

	// The OCID of the certificate.
	GetCertificateId() *string

	// The name of the certificate.
	GetCertificateName() *string

	// The version number of the certificate.
	GetVersionNumber() *int64

	// A unique certificate identifier used in certificate revocation tracking, formatted as octets.
	// Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF`
	GetSerialNumber() *string

	// An optional property indicating when the certificate version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	GetValidity() *Validity

	// A list of rotation states for the certificate bundle.
	GetStages() []VersionStageEnum

	// The certificate in PEM format.
	GetCertificatePem() *string

	// The certificate chain (in PEM format) for the certificate bundle.
	GetCertChainPem() *string

	// The name of the certificate version.
	GetVersionName() *string

	GetRevocationStatus() *RevocationStatus
}

type certificatebundle struct {
	JsonData              []byte
	CertificateId         *string            `mandatory:"true" json:"certificateId"`
	CertificateName       *string            `mandatory:"true" json:"certificateName"`
	VersionNumber         *int64             `mandatory:"true" json:"versionNumber"`
	SerialNumber          *string            `mandatory:"true" json:"serialNumber"`
	TimeCreated           *common.SDKTime    `mandatory:"true" json:"timeCreated"`
	Validity              *Validity          `mandatory:"true" json:"validity"`
	Stages                []VersionStageEnum `mandatory:"true" json:"stages"`
	CertificatePem        *string            `mandatory:"false" json:"certificatePem"`
	CertChainPem          *string            `mandatory:"false" json:"certChainPem"`
	VersionName           *string            `mandatory:"false" json:"versionName"`
	RevocationStatus      *RevocationStatus  `mandatory:"false" json:"revocationStatus"`
	CertificateBundleType string             `json:"certificateBundleType"`
}

// UnmarshalJSON unmarshals json
func (m *certificatebundle) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercertificatebundle certificatebundle
	s := struct {
		Model Unmarshalercertificatebundle
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CertificateId = s.Model.CertificateId
	m.CertificateName = s.Model.CertificateName
	m.VersionNumber = s.Model.VersionNumber
	m.SerialNumber = s.Model.SerialNumber
	m.TimeCreated = s.Model.TimeCreated
	m.Validity = s.Model.Validity
	m.Stages = s.Model.Stages
	m.CertificatePem = s.Model.CertificatePem
	m.CertChainPem = s.Model.CertChainPem
	m.VersionName = s.Model.VersionName
	m.RevocationStatus = s.Model.RevocationStatus
	m.CertificateBundleType = s.Model.CertificateBundleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *certificatebundle) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CertificateBundleType {
	case "CERTIFICATE_CONTENT_PUBLIC_ONLY":
		mm := CertificateBundlePublicOnly{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CERTIFICATE_CONTENT_WITH_PRIVATE_KEY":
		mm := CertificateBundleWithPrivateKey{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCertificateId returns CertificateId
func (m certificatebundle) GetCertificateId() *string {
	return m.CertificateId
}

//GetCertificateName returns CertificateName
func (m certificatebundle) GetCertificateName() *string {
	return m.CertificateName
}

//GetVersionNumber returns VersionNumber
func (m certificatebundle) GetVersionNumber() *int64 {
	return m.VersionNumber
}

//GetSerialNumber returns SerialNumber
func (m certificatebundle) GetSerialNumber() *string {
	return m.SerialNumber
}

//GetTimeCreated returns TimeCreated
func (m certificatebundle) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetValidity returns Validity
func (m certificatebundle) GetValidity() *Validity {
	return m.Validity
}

//GetStages returns Stages
func (m certificatebundle) GetStages() []VersionStageEnum {
	return m.Stages
}

//GetCertificatePem returns CertificatePem
func (m certificatebundle) GetCertificatePem() *string {
	return m.CertificatePem
}

//GetCertChainPem returns CertChainPem
func (m certificatebundle) GetCertChainPem() *string {
	return m.CertChainPem
}

//GetVersionName returns VersionName
func (m certificatebundle) GetVersionName() *string {
	return m.VersionName
}

//GetRevocationStatus returns RevocationStatus
func (m certificatebundle) GetRevocationStatus() *RevocationStatus {
	return m.RevocationStatus
}

func (m certificatebundle) String() string {
	return common.PointerString(m)
}

// CertificateBundleCertificateBundleTypeEnum Enum with underlying type: string
type CertificateBundleCertificateBundleTypeEnum string

// Set of constants representing the allowable values for CertificateBundleCertificateBundleTypeEnum
const (
	CertificateBundleCertificateBundleTypePublicOnly     CertificateBundleCertificateBundleTypeEnum = "CERTIFICATE_CONTENT_PUBLIC_ONLY"
	CertificateBundleCertificateBundleTypeWithPrivateKey CertificateBundleCertificateBundleTypeEnum = "CERTIFICATE_CONTENT_WITH_PRIVATE_KEY"
)

var mappingCertificateBundleCertificateBundleType = map[string]CertificateBundleCertificateBundleTypeEnum{
	"CERTIFICATE_CONTENT_PUBLIC_ONLY":      CertificateBundleCertificateBundleTypePublicOnly,
	"CERTIFICATE_CONTENT_WITH_PRIVATE_KEY": CertificateBundleCertificateBundleTypeWithPrivateKey,
}

// GetCertificateBundleCertificateBundleTypeEnumValues Enumerates the set of values for CertificateBundleCertificateBundleTypeEnum
func GetCertificateBundleCertificateBundleTypeEnumValues() []CertificateBundleCertificateBundleTypeEnum {
	values := make([]CertificateBundleCertificateBundleTypeEnum, 0)
	for _, v := range mappingCertificateBundleCertificateBundleType {
		values = append(values, v)
	}
	return values
}
