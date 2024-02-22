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

type MonitorInitParameters struct {

	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	CompanyName *string `json:"companyName,omitempty" tf:"company_name,omitempty"`

	// Whether the resource monitoring is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	EnterpriseAppID *string `json:"enterpriseAppId,omitempty" tf:"enterprise_app_id,omitempty"`

	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A plan block as defined below. Changing this forces a new resource to be created.
	Plan []PlanInitParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// A mapping of tags which should be assigned to the logz Monitor.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A user block as defined below. Changing this forces a new resource to be created.
	User []UserInitParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type MonitorObservation struct {

	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	CompanyName *string `json:"companyName,omitempty" tf:"company_name,omitempty"`

	// Whether the resource monitoring is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	EnterpriseAppID *string `json:"enterpriseAppId,omitempty" tf:"enterprise_app_id,omitempty"`

	// The ID of the logz Monitor.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID associated with the logz organization of this logz Monitor.
	LogzOrganizationID *string `json:"logzOrganizationId,omitempty" tf:"logz_organization_id,omitempty"`

	// A plan block as defined below. Changing this forces a new resource to be created.
	Plan []PlanObservation `json:"plan,omitempty" tf:"plan,omitempty"`

	// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The single sign on url associated with the logz organization of this logz Monitor.
	SingleSignOnURL *string `json:"singleSignOnUrl,omitempty" tf:"single_sign_on_url,omitempty"`

	// A mapping of tags which should be assigned to the logz Monitor.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A user block as defined below. Changing this forces a new resource to be created.
	User []UserObservation `json:"user,omitempty" tf:"user,omitempty"`
}

type MonitorParameters struct {

	// Name of the Logz organization. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	CompanyName *string `json:"companyName,omitempty" tf:"company_name,omitempty"`

	// Whether the resource monitoring is enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Enterprise App. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	EnterpriseAppID *string `json:"enterpriseAppId,omitempty" tf:"enterprise_app_id,omitempty"`

	// The Azure Region where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A plan block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Plan []PlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// The name of the Resource Group where the logz Monitor should exist. Changing this forces a new logz Monitor to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the logz Monitor.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A user block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	User []UserParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type PlanInitParameters struct {

	// Different billing cycles. Possible values are MONTHLY or WEEKLY. Changing this forces a new logz Monitor to be created.
	BillingCycle *string `json:"billingCycle,omitempty" tf:"billing_cycle,omitempty"`

	// Date when plan was applied. Changing this forces a new logz Monitor to be created.
	EffectiveDate *string `json:"effectiveDate,omitempty" tf:"effective_date,omitempty"`

	// Plan id as published by Logz. The only possible value is 100gb14days. Defaults to 100gb14days. Changing this forces a new logz Monitor to be created.
	PlanID *string `json:"planId,omitempty" tf:"plan_id,omitempty"`

	// Different usage types. Possible values are PAYG or COMMITTED. Changing this forces a new logz Monitor to be created.
	UsageType *string `json:"usageType,omitempty" tf:"usage_type,omitempty"`
}

type PlanObservation struct {

	// Different billing cycles. Possible values are MONTHLY or WEEKLY. Changing this forces a new logz Monitor to be created.
	BillingCycle *string `json:"billingCycle,omitempty" tf:"billing_cycle,omitempty"`

	// Date when plan was applied. Changing this forces a new logz Monitor to be created.
	EffectiveDate *string `json:"effectiveDate,omitempty" tf:"effective_date,omitempty"`

	// Plan id as published by Logz. The only possible value is 100gb14days. Defaults to 100gb14days. Changing this forces a new logz Monitor to be created.
	PlanID *string `json:"planId,omitempty" tf:"plan_id,omitempty"`

	// Different usage types. Possible values are PAYG or COMMITTED. Changing this forces a new logz Monitor to be created.
	UsageType *string `json:"usageType,omitempty" tf:"usage_type,omitempty"`
}

type PlanParameters struct {

	// Different billing cycles. Possible values are MONTHLY or WEEKLY. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	BillingCycle *string `json:"billingCycle" tf:"billing_cycle,omitempty"`

	// Date when plan was applied. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	EffectiveDate *string `json:"effectiveDate" tf:"effective_date,omitempty"`

	// Plan id as published by Logz. The only possible value is 100gb14days. Defaults to 100gb14days. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	PlanID *string `json:"planId,omitempty" tf:"plan_id,omitempty"`

	// Different usage types. Possible values are PAYG or COMMITTED. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	UsageType *string `json:"usageType" tf:"usage_type,omitempty"`
}

type UserInitParameters struct {

	// Email of the user used by Logz for contacting them if needed. Changing this forces a new logz Monitor to be created.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// First Name of the user. Changing this forces a new logz Monitor to be created.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// Last Name of the user. Changing this forces a new logz Monitor to be created.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// Phone number of the user used by Logz for contacting them if needed. Changing this forces a new logz Monitor to be created.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type UserObservation struct {

	// Email of the user used by Logz for contacting them if needed. Changing this forces a new logz Monitor to be created.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// First Name of the user. Changing this forces a new logz Monitor to be created.
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// Last Name of the user. Changing this forces a new logz Monitor to be created.
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// Phone number of the user used by Logz for contacting them if needed. Changing this forces a new logz Monitor to be created.
	PhoneNumber *string `json:"phoneNumber,omitempty" tf:"phone_number,omitempty"`
}

type UserParameters struct {

	// Email of the user used by Logz for contacting them if needed. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	Email *string `json:"email" tf:"email,omitempty"`

	// First Name of the user. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName" tf:"first_name,omitempty"`

	// Last Name of the user. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName" tf:"last_name,omitempty"`

	// Phone number of the user used by Logz for contacting them if needed. Changing this forces a new logz Monitor to be created.
	// +kubebuilder:validation:Optional
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number,omitempty"`
}

// MonitorSpec defines the desired state of Monitor
type MonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorParameters `json:"forProvider"`
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
	InitProvider MonitorInitParameters `json:"initProvider,omitempty"`
}

// MonitorStatus defines the observed state of Monitor.
type MonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Monitor is the Schema for the Monitors API. Manages a logz Monitor.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Monitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.plan) || (has(self.initProvider) && has(self.initProvider.plan))",message="spec.forProvider.plan is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.user) || (has(self.initProvider) && has(self.initProvider.user))",message="spec.forProvider.user is a required parameter"
	Spec   MonitorSpec   `json:"spec"`
	Status MonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorList contains a list of Monitors
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Monitor `json:"items"`
}

// Repository type metadata.
var (
	Monitor_Kind             = "Monitor"
	Monitor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Monitor_Kind}.String()
	Monitor_KindAPIVersion   = Monitor_Kind + "." + CRDGroupVersion.String()
	Monitor_GroupVersionKind = CRDGroupVersion.WithKind(Monitor_Kind)
)

func init() {
	SchemeBuilder.Register(&Monitor{}, &MonitorList{})
}
