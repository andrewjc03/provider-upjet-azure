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

type TableEntityInitParameters struct {

	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	// +mapType=granular
	Entity map[string]*string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey *string `json:"rowKey,omitempty" tf:"row_key,omitempty"`

	// Specifies the storage account in which to create the storage table entity. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// The name of the storage table in which to create the storage table entity. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Table
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table in storage to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table in storage to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`
}

type TableEntityObservation struct {

	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	// +mapType=granular
	Entity map[string]*string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The ID of the Entity within the Table in the Storage Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	RowKey *string `json:"rowKey,omitempty" tf:"row_key,omitempty"`

	// Specifies the storage account in which to create the storage table entity. Changing this forces a new resource to be created.
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// The name of the storage table in which to create the storage table entity. Changing this forces a new resource to be created.
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type TableEntityParameters struct {

	// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Entity map[string]*string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
	// +kubebuilder:validation:Optional
	PartitionKey *string `json:"partitionKey,omitempty" tf:"partition_key,omitempty"`

	// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
	// +kubebuilder:validation:Optional
	RowKey *string `json:"rowKey,omitempty" tf:"row_key,omitempty"`

	// Specifies the storage account in which to create the storage table entity. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// The name of the storage table in which to create the storage table entity. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Table
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table in storage to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table in storage to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`
}

// TableEntitySpec defines the desired state of TableEntity
type TableEntitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableEntityParameters `json:"forProvider"`
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
	InitProvider TableEntityInitParameters `json:"initProvider,omitempty"`
}

// TableEntityStatus defines the observed state of TableEntity.
type TableEntityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableEntityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TableEntity is the Schema for the TableEntitys API. Manages an Entity within a Table in an Azure Storage Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TableEntity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entity) || (has(self.initProvider) && has(self.initProvider.entity))",message="spec.forProvider.entity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.partitionKey) || (has(self.initProvider) && has(self.initProvider.partitionKey))",message="spec.forProvider.partitionKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rowKey) || (has(self.initProvider) && has(self.initProvider.rowKey))",message="spec.forProvider.rowKey is a required parameter"
	Spec   TableEntitySpec   `json:"spec"`
	Status TableEntityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableEntityList contains a list of TableEntitys
type TableEntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableEntity `json:"items"`
}

// Repository type metadata.
var (
	TableEntity_Kind             = "TableEntity"
	TableEntity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableEntity_Kind}.String()
	TableEntity_KindAPIVersion   = TableEntity_Kind + "." + CRDGroupVersion.String()
	TableEntity_GroupVersionKind = CRDGroupVersion.WithKind(TableEntity_Kind)
)

func init() {
	SchemeBuilder.Register(&TableEntity{}, &TableEntityList{})
}
