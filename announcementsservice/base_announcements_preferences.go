// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v54/common"
)

// BaseAnnouncementsPreferences The object that contains the announcement email preferences configured for the tenancy (root compartment).
type BaseAnnouncementsPreferences interface {

	// The OCID of the compartment for which the email preferences apply. Because announcements are
	// specific to a tenancy, specify the tenancy by providing the root compartment OCID.
	GetCompartmentId() *string

	// The ID of the preferences.
	GetId() *string

	// A Boolean value to indicate whether the specified compartment chooses to not to receive informational announcements by email.
	// (Manage preferences for receiving announcements by email by specifying the `preferenceType` attribute instead.)
	GetIsUnsubscribed() *bool

	// When the preferences were set initially.
	GetTimeCreated() *common.SDKTime

	// When the preferences were last updated.
	GetTimeUpdated() *common.SDKTime

	// The string representing the user's preference regarding receiving announcements by email.
	GetPreferenceType() BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum
}

type baseannouncementspreferences struct {
	JsonData       []byte
	CompartmentId  *string                                                     `mandatory:"false" json:"compartmentId"`
	Id             *string                                                     `mandatory:"false" json:"id"`
	IsUnsubscribed *bool                                                       `mandatory:"false" json:"isUnsubscribed"`
	TimeCreated    *common.SDKTime                                             `mandatory:"false" json:"timeCreated"`
	TimeUpdated    *common.SDKTime                                             `mandatory:"false" json:"timeUpdated"`
	PreferenceType BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum `mandatory:"false" json:"preferenceType,omitempty"`
	Type           string                                                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *baseannouncementspreferences) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbaseannouncementspreferences baseannouncementspreferences
	s := struct {
		Model Unmarshalerbaseannouncementspreferences
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.Id = s.Model.Id
	m.IsUnsubscribed = s.Model.IsUnsubscribed
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.PreferenceType = s.Model.PreferenceType
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *baseannouncementspreferences) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "AnnouncementsPreferencesSummary":
		mm := AnnouncementsPreferencesSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AnnouncementsPreferences":
		mm := AnnouncementsPreferences{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCompartmentId returns CompartmentId
func (m baseannouncementspreferences) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetId returns Id
func (m baseannouncementspreferences) GetId() *string {
	return m.Id
}

//GetIsUnsubscribed returns IsUnsubscribed
func (m baseannouncementspreferences) GetIsUnsubscribed() *bool {
	return m.IsUnsubscribed
}

//GetTimeCreated returns TimeCreated
func (m baseannouncementspreferences) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m baseannouncementspreferences) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetPreferenceType returns PreferenceType
func (m baseannouncementspreferences) GetPreferenceType() BaseCreateAnnouncementsPreferencesDetailsPreferenceTypeEnum {
	return m.PreferenceType
}

func (m baseannouncementspreferences) String() string {
	return common.PointerString(m)
}
