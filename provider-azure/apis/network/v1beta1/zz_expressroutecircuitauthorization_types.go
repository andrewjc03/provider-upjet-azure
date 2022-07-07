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

type ExpressRouteCircuitAuthorizationObservation struct {
	AuthorizationUseStatus *string `json:"authorizationUseStatus,omitempty" tf:"authorization_use_status,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ExpressRouteCircuitAuthorizationParameters struct {

	// +crossplane:generate:reference:type=ExpressRouteCircuit
	// +kubebuilder:validation:Optional
	ExpressRouteCircuitName *string `json:"expressRouteCircuitName,omitempty" tf:"express_route_circuit_name,omitempty"`

	// +kubebuilder:validation:Optional
	ExpressRouteCircuitNameRef *v1.Reference `json:"expressRouteCircuitNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ExpressRouteCircuitNameSelector *v1.Selector `json:"expressRouteCircuitNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// ExpressRouteCircuitAuthorizationSpec defines the desired state of ExpressRouteCircuitAuthorization
type ExpressRouteCircuitAuthorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitAuthorizationParameters `json:"forProvider"`
}

// ExpressRouteCircuitAuthorizationStatus defines the observed state of ExpressRouteCircuitAuthorization.
type ExpressRouteCircuitAuthorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitAuthorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitAuthorization is the Schema for the ExpressRouteCircuitAuthorizations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuitAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitAuthorizationSpec   `json:"spec"`
	Status            ExpressRouteCircuitAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitAuthorizationList contains a list of ExpressRouteCircuitAuthorizations
type ExpressRouteCircuitAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuitAuthorization `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuitAuthorization_Kind             = "ExpressRouteCircuitAuthorization"
	ExpressRouteCircuitAuthorization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteCircuitAuthorization_Kind}.String()
	ExpressRouteCircuitAuthorization_KindAPIVersion   = ExpressRouteCircuitAuthorization_Kind + "." + CRDGroupVersion.String()
	ExpressRouteCircuitAuthorization_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteCircuitAuthorization_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuitAuthorization{}, &ExpressRouteCircuitAuthorizationList{})
}
