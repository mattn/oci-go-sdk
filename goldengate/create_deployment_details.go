// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v38/common"
)

// CreateDeploymentDetails The information about a new Deployment.
type CreateDeploymentDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Oracle license model that applies to a Deployment.
	LicenseModel LicenseModelEnum `mandatory:"true" json:"licenseModel"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The Minimum number of OCPUs to be made available for this Deployment.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// Indicates if auto scaling is enabled for the Deployment's CPU core count.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// The deployment type.
	DeploymentType DeploymentTypeEnum `mandatory:"true" json:"deploymentType"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup being referenced.
	DeploymentBackupId *string `mandatory:"false" json:"deploymentBackupId"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// An array of Network Security Group (https://docs.cloud.oracle.com/Content/Network/Concepts/networksecuritygroups.htm) OCIDs used to define network access for a deployment.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// True if this object is publicly available.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	OggData *CreateOggDeploymentDetails `mandatory:"false" json:"oggData"`
}

func (m CreateDeploymentDetails) String() string {
	return common.PointerString(m)
}
