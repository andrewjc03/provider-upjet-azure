// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IOTHubDPSSharedAccessPolicyInitParameters struct {

	// Adds EnrollmentRead permission to this Shared Access Account. It allows read access to enrollment data.
	EnrollmentRead *bool `json:"enrollmentRead,omitempty" tf:"enrollment_read,omitempty"`

	// Adds EnrollmentWrite permission to this Shared Access Account. It allows write access to enrollment data.
	EnrollmentWrite *bool `json:"enrollmentWrite,omitempty" tf:"enrollment_write,omitempty"`

	// Adds RegistrationStatusRead permission to this Shared Access Account. It allows read access to device registrations.
	RegistrationRead *bool `json:"registrationRead,omitempty" tf:"registration_read,omitempty"`

	// Adds RegistrationStatusWrite permission to this Shared Access Account. It allows write access to device registrations.
	RegistrationWrite *bool `json:"registrationWrite,omitempty" tf:"registration_write,omitempty"`

	// Adds ServiceConfig permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	ServiceConfig *bool `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`
}

type IOTHubDPSSharedAccessPolicyObservation struct {

	// Adds EnrollmentRead permission to this Shared Access Account. It allows read access to enrollment data.
	EnrollmentRead *bool `json:"enrollmentRead,omitempty" tf:"enrollment_read,omitempty"`

	// Adds EnrollmentWrite permission to this Shared Access Account. It allows write access to enrollment data.
	EnrollmentWrite *bool `json:"enrollmentWrite,omitempty" tf:"enrollment_write,omitempty"`

	// The ID of the IoTHub Device Provisioning Service Shared Access Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IOTHubDPSName *string `json:"iothubDpsName,omitempty" tf:"iothub_dps_name,omitempty"`

	// Adds RegistrationStatusRead permission to this Shared Access Account. It allows read access to device registrations.
	RegistrationRead *bool `json:"registrationRead,omitempty" tf:"registration_read,omitempty"`

	// Adds RegistrationStatusWrite permission to this Shared Access Account. It allows write access to device registrations.
	RegistrationWrite *bool `json:"registrationWrite,omitempty" tf:"registration_write,omitempty"`

	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Adds ServiceConfig permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	ServiceConfig *bool `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`
}

type IOTHubDPSSharedAccessPolicyParameters struct {

	// Adds EnrollmentRead permission to this Shared Access Account. It allows read access to enrollment data.
	// +kubebuilder:validation:Optional
	EnrollmentRead *bool `json:"enrollmentRead,omitempty" tf:"enrollment_read,omitempty"`

	// Adds EnrollmentWrite permission to this Shared Access Account. It allows write access to enrollment data.
	// +kubebuilder:validation:Optional
	EnrollmentWrite *bool `json:"enrollmentWrite,omitempty" tf:"enrollment_write,omitempty"`

	// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHubDPS
	// +kubebuilder:validation:Optional
	IOTHubDPSName *string `json:"iothubDpsName,omitempty" tf:"iothub_dps_name,omitempty"`

	// Reference to a IOTHubDPS in devices to populate iothubDpsName.
	// +kubebuilder:validation:Optional
	IOTHubDPSNameRef *v1.Reference `json:"iothubDpsNameRef,omitempty" tf:"-"`

	// Selector for a IOTHubDPS in devices to populate iothubDpsName.
	// +kubebuilder:validation:Optional
	IOTHubDPSNameSelector *v1.Selector `json:"iothubDpsNameSelector,omitempty" tf:"-"`

	// Adds RegistrationStatusRead permission to this Shared Access Account. It allows read access to device registrations.
	// +kubebuilder:validation:Optional
	RegistrationRead *bool `json:"registrationRead,omitempty" tf:"registration_read,omitempty"`

	// Adds RegistrationStatusWrite permission to this Shared Access Account. It allows write access to device registrations.
	// +kubebuilder:validation:Optional
	RegistrationWrite *bool `json:"registrationWrite,omitempty" tf:"registration_write,omitempty"`

	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Adds ServiceConfig permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	// +kubebuilder:validation:Optional
	ServiceConfig *bool `json:"serviceConfig,omitempty" tf:"service_config,omitempty"`
}

// IOTHubDPSSharedAccessPolicySpec defines the desired state of IOTHubDPSSharedAccessPolicy
type IOTHubDPSSharedAccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubDPSSharedAccessPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IOTHubDPSSharedAccessPolicyInitParameters `json:"initProvider,omitempty"`
}

// IOTHubDPSSharedAccessPolicyStatus defines the observed state of IOTHubDPSSharedAccessPolicy.
type IOTHubDPSSharedAccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubDPSSharedAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IOTHubDPSSharedAccessPolicy is the Schema for the IOTHubDPSSharedAccessPolicys API. Manages an IotHub Device Provisioning Service Shared Access Policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubDPSSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubDPSSharedAccessPolicySpec   `json:"spec"`
	Status            IOTHubDPSSharedAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubDPSSharedAccessPolicyList contains a list of IOTHubDPSSharedAccessPolicys
type IOTHubDPSSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubDPSSharedAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	IOTHubDPSSharedAccessPolicy_Kind             = "IOTHubDPSSharedAccessPolicy"
	IOTHubDPSSharedAccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubDPSSharedAccessPolicy_Kind}.String()
	IOTHubDPSSharedAccessPolicy_KindAPIVersion   = IOTHubDPSSharedAccessPolicy_Kind + "." + CRDGroupVersion.String()
	IOTHubDPSSharedAccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubDPSSharedAccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubDPSSharedAccessPolicy{}, &IOTHubDPSSharedAccessPolicyList{})
}
