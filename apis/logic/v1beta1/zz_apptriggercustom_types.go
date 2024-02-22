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

type AppTriggerCustomInitParameters struct {

	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`
}

type AppTriggerCustomObservation struct {

	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// The ID of the Trigger within the Logic App Workflow.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`
}

type AppTriggerCustomParameters struct {

	// Specifies the JSON Blob defining the Body of this Custom Trigger.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Specifies the ID of the Logic App Workflow. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logic/v1beta1.AppWorkflow
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogicAppID *string `json:"logicAppId,omitempty" tf:"logic_app_id,omitempty"`

	// Reference to a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDRef *v1.Reference `json:"logicAppIdRef,omitempty" tf:"-"`

	// Selector for a AppWorkflow in logic to populate logicAppId.
	// +kubebuilder:validation:Optional
	LogicAppIDSelector *v1.Selector `json:"logicAppIdSelector,omitempty" tf:"-"`
}

// AppTriggerCustomSpec defines the desired state of AppTriggerCustom
type AppTriggerCustomSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppTriggerCustomParameters `json:"forProvider"`
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
	InitProvider AppTriggerCustomInitParameters `json:"initProvider,omitempty"`
}

// AppTriggerCustomStatus defines the observed state of AppTriggerCustom.
type AppTriggerCustomStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppTriggerCustomObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AppTriggerCustom is the Schema for the AppTriggerCustoms API. Manages a Custom Trigger within a Logic App Workflow
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppTriggerCustom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	Spec   AppTriggerCustomSpec   `json:"spec"`
	Status AppTriggerCustomStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppTriggerCustomList contains a list of AppTriggerCustoms
type AppTriggerCustomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppTriggerCustom `json:"items"`
}

// Repository type metadata.
var (
	AppTriggerCustom_Kind             = "AppTriggerCustom"
	AppTriggerCustom_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppTriggerCustom_Kind}.String()
	AppTriggerCustom_KindAPIVersion   = AppTriggerCustom_Kind + "." + CRDGroupVersion.String()
	AppTriggerCustom_GroupVersionKind = CRDGroupVersion.WithKind(AppTriggerCustom_Kind)
)

func init() {
	SchemeBuilder.Register(&AppTriggerCustom{}, &AppTriggerCustomList{})
}
