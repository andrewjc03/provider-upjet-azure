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

type AADAdminInitParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	Login *string `json:"login,omitempty" tf:"login"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type AADAdminObservation struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AADAdminParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login"`

	// The object id of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type AzureDevopsRepoInitParameters struct {

	// Specifies the Azure DevOps account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the collaboration branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// The last commit ID.
	LastCommitID *string `json:"lastCommitId,omitempty" tf:"last_commit_id,omitempty"`

	// Specifies the name of the Azure DevOps project.
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`

	// the ID of the tenant for the Azure DevOps account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureDevopsRepoObservation struct {

	// Specifies the Azure DevOps account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the collaboration branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// The last commit ID.
	LastCommitID *string `json:"lastCommitId,omitempty" tf:"last_commit_id,omitempty"`

	// Specifies the name of the Azure DevOps project.
	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`

	// the ID of the tenant for the Azure DevOps account.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type AzureDevopsRepoParameters struct {

	// Specifies the Azure DevOps account name.
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// Specifies the collaboration branch of the repository to get code from.
	// +kubebuilder:validation:Optional
	BranchName *string `json:"branchName" tf:"branch_name,omitempty"`

	// The last commit ID.
	// +kubebuilder:validation:Optional
	LastCommitID *string `json:"lastCommitId,omitempty" tf:"last_commit_id,omitempty"`

	// Specifies the name of the Azure DevOps project.
	// +kubebuilder:validation:Optional
	ProjectName *string `json:"projectName" tf:"project_name,omitempty"`

	// Specifies the name of the git repository.
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	// +kubebuilder:validation:Optional
	RootFolder *string `json:"rootFolder" tf:"root_folder,omitempty"`

	// the ID of the tenant for the Azure DevOps account.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type CustomerManagedKeyInitParameters struct {

	// An identifier for the key. Name needs to match the name of the key used with the azurerm_synapse_workspace_key resource. Defaults to "cmk" if not specified.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The Azure Key Vault Key Versionless ID to be used as the Customer Managed Key (CMK) for double encryption (e.g. https://example-keyvault.vault.azure.net/type/cmk/).
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_id",true)
	KeyVersionlessID *string `json:"keyVersionlessId,omitempty" tf:"key_versionless_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVersionlessId.
	// +kubebuilder:validation:Optional
	KeyVersionlessIDRef *v1.Reference `json:"keyVersionlessIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVersionlessId.
	// +kubebuilder:validation:Optional
	KeyVersionlessIDSelector *v1.Selector `json:"keyVersionlessIdSelector,omitempty" tf:"-"`
}

type CustomerManagedKeyObservation struct {

	// An identifier for the key. Name needs to match the name of the key used with the azurerm_synapse_workspace_key resource. Defaults to "cmk" if not specified.
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The Azure Key Vault Key Versionless ID to be used as the Customer Managed Key (CMK) for double encryption (e.g. https://example-keyvault.vault.azure.net/type/cmk/).
	KeyVersionlessID *string `json:"keyVersionlessId,omitempty" tf:"key_versionless_id,omitempty"`
}

type CustomerManagedKeyParameters struct {

	// An identifier for the key. Name needs to match the name of the key used with the azurerm_synapse_workspace_key resource. Defaults to "cmk" if not specified.
	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name,omitempty"`

	// The Azure Key Vault Key Versionless ID to be used as the Customer Managed Key (CMK) for double encryption (e.g. https://example-keyvault.vault.azure.net/type/cmk/).
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_id",true)
	// +kubebuilder:validation:Optional
	KeyVersionlessID *string `json:"keyVersionlessId,omitempty" tf:"key_versionless_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVersionlessId.
	// +kubebuilder:validation:Optional
	KeyVersionlessIDRef *v1.Reference `json:"keyVersionlessIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVersionlessId.
	// +kubebuilder:validation:Optional
	KeyVersionlessIDSelector *v1.Selector `json:"keyVersionlessIdSelector,omitempty" tf:"-"`
}

type GithubRepoInitParameters struct {

	// Specifies the GitHub account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the collaboration branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com.
	GitURL *string `json:"gitUrl,omitempty" tf:"git_url,omitempty"`

	// The last commit ID.
	LastCommitID *string `json:"lastCommitId,omitempty" tf:"last_commit_id,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`
}

type GithubRepoObservation struct {

	// Specifies the GitHub account name.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Specifies the collaboration branch of the repository to get code from.
	BranchName *string `json:"branchName,omitempty" tf:"branch_name,omitempty"`

	// Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com.
	GitURL *string `json:"gitUrl,omitempty" tf:"git_url,omitempty"`

	// The last commit ID.
	LastCommitID *string `json:"lastCommitId,omitempty" tf:"last_commit_id,omitempty"`

	// Specifies the name of the git repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	RootFolder *string `json:"rootFolder,omitempty" tf:"root_folder,omitempty"`
}

type GithubRepoParameters struct {

	// Specifies the GitHub account name.
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// Specifies the collaboration branch of the repository to get code from.
	// +kubebuilder:validation:Optional
	BranchName *string `json:"branchName" tf:"branch_name,omitempty"`

	// Specifies the GitHub Enterprise host name. For example: https://github.mydomain.com.
	// +kubebuilder:validation:Optional
	GitURL *string `json:"gitUrl,omitempty" tf:"git_url,omitempty"`

	// The last commit ID.
	// +kubebuilder:validation:Optional
	LastCommitID *string `json:"lastCommitId,omitempty" tf:"last_commit_id,omitempty"`

	// Specifies the name of the git repository.
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName" tf:"repository_name,omitempty"`

	// Specifies the root folder within the repository. Set to / for the top level.
	// +kubebuilder:validation:Optional
	RootFolder *string `json:"rootFolder" tf:"root_folder,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Synapse Workspace.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be associated with this Synapse Workspace. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Synapse Workspace.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this Synapse Workspace.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this Synapse Workspace.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be associated with this Synapse Workspace. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Synapse Workspace.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be associated with this Synapse Workspace. Possible values are SystemAssigned, UserAssigned and SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type SQLAADAdminInitParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace SQL.
	Login *string `json:"login,omitempty" tf:"login"`

	// The object id of the Azure AD Administrator of this Synapse Workspace SQL.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace SQL.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type SQLAADAdminObservation struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace SQL.
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// The object id of the Azure AD Administrator of this Synapse Workspace SQL.
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace SQL.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type SQLAADAdminParameters struct {

	// The login name of the Azure AD Administrator of this Synapse Workspace SQL.
	// +kubebuilder:validation:Optional
	Login *string `json:"login,omitempty" tf:"login"`

	// The object id of the Azure AD Administrator of this Synapse Workspace SQL.
	// +kubebuilder:validation:Optional
	ObjectID *string `json:"objectId,omitempty" tf:"object_id"`

	// The tenant id of the Azure AD Administrator of this Synapse Workspace SQL.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type WorkspaceInitParameters struct {

	// An aad_admin block as defined below. Conflicts with customer_managed_key.
	AADAdmin []AADAdminInitParameters `json:"aadAdmin,omitempty" tf:"aad_admin,omitempty"`

	// An azure_devops_repo block as defined below.
	AzureDevopsRepo []AzureDevopsRepoInitParameters `json:"azureDevopsRepo,omitempty" tf:"azure_devops_repo,omitempty"`

	// Subnet ID used for computes in workspace Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	ComputeSubnetID *string `json:"computeSubnetId,omitempty" tf:"compute_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate computeSubnetId.
	// +kubebuilder:validation:Optional
	ComputeSubnetIDRef *v1.Reference `json:"computeSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate computeSubnetId.
	// +kubebuilder:validation:Optional
	ComputeSubnetIDSelector *v1.Selector `json:"computeSubnetIdSelector,omitempty" tf:"-"`

	// A customer_managed_key block as defined below. Conflicts with aad_admin.
	CustomerManagedKey []CustomerManagedKeyInitParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Is data exfiltration protection enabled in this workspace? If set to true, managed_virtual_network_enabled must also be set to true. Changing this forces a new resource to be created.
	DataExfiltrationProtectionEnabled *bool `json:"dataExfiltrationProtectionEnabled,omitempty" tf:"data_exfiltration_protection_enabled,omitempty"`

	// A github_repo block as defined below.
	GithubRepo []GithubRepoInitParameters `json:"githubRepo,omitempty" tf:"github_repo,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Allowed AAD Tenant Ids For Linking.
	LinkingAllowedForAADTenantIds []*string `json:"linkingAllowedForAadTenantIds,omitempty" tf:"linking_allowed_for_aad_tenant_ids,omitempty"`

	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Workspace managed resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameRef *v1.Reference `json:"managedResourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameSelector *v1.Selector `json:"managedResourceGroupNameSelector,omitempty" tf:"-"`

	// Is Virtual Network enabled for all computes in this workspace? Changing this forces a new resource to be created.
	ManagedVirtualNetworkEnabled *bool `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled,omitempty"`

	// Whether public network access is allowed for the Cognitive Account. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The ID of purview account.
	PurviewID *string `json:"purviewId,omitempty" tf:"purview_id,omitempty"`

	// An sql_aad_admin block as defined below.
	SQLAADAdmin []SQLAADAdminInitParameters `json:"sqlAadAdmin,omitempty" tf:"sql_aad_admin,omitempty"`

	// Specifies The login name of the SQL administrator. Changing this forces a new resource to be created. If this is not provided aad_admin or customer_managed_key must be provided.
	SQLAdministratorLogin *string `json:"sqlAdministratorLogin,omitempty" tf:"sql_administrator_login,omitempty"`

	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	SQLIdentityControlEnabled *bool `json:"sqlIdentityControlEnabled,omitempty" tf:"sql_identity_control_enabled,omitempty"`

	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.DataLakeGen2FileSystem
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	StorageDataLakeGen2FileSystemID *string `json:"storageDataLakeGen2FilesystemId,omitempty" tf:"storage_data_lake_gen2_filesystem_id,omitempty"`

	// Reference to a DataLakeGen2FileSystem in storage to populate storageDataLakeGen2FilesystemId.
	// +kubebuilder:validation:Optional
	StorageDataLakeGen2FileSystemIDRef *v1.Reference `json:"storageDataLakeGen2FilesystemIdRef,omitempty" tf:"-"`

	// Selector for a DataLakeGen2FileSystem in storage to populate storageDataLakeGen2FilesystemId.
	// +kubebuilder:validation:Optional
	StorageDataLakeGen2FileSystemIDSelector *v1.Selector `json:"storageDataLakeGen2FilesystemIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Synapse Workspace.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WorkspaceObservation struct {

	// An aad_admin block as defined below. Conflicts with customer_managed_key.
	AADAdmin []AADAdminObservation `json:"aadAdmin,omitempty" tf:"aad_admin,omitempty"`

	// An azure_devops_repo block as defined below.
	AzureDevopsRepo []AzureDevopsRepoObservation `json:"azureDevopsRepo,omitempty" tf:"azure_devops_repo,omitempty"`

	// Subnet ID used for computes in workspace Changing this forces a new resource to be created.
	ComputeSubnetID *string `json:"computeSubnetId,omitempty" tf:"compute_subnet_id,omitempty"`

	// A list of Connectivity endpoints for this Synapse Workspace.
	// +mapType=granular
	ConnectivityEndpoints map[string]*string `json:"connectivityEndpoints,omitempty" tf:"connectivity_endpoints,omitempty"`

	// A customer_managed_key block as defined below. Conflicts with aad_admin.
	CustomerManagedKey []CustomerManagedKeyObservation `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Is data exfiltration protection enabled in this workspace? If set to true, managed_virtual_network_enabled must also be set to true. Changing this forces a new resource to be created.
	DataExfiltrationProtectionEnabled *bool `json:"dataExfiltrationProtectionEnabled,omitempty" tf:"data_exfiltration_protection_enabled,omitempty"`

	// A github_repo block as defined below.
	GithubRepo []GithubRepoObservation `json:"githubRepo,omitempty" tf:"github_repo,omitempty"`

	// The ID of the synapse Workspace.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Allowed AAD Tenant Ids For Linking.
	LinkingAllowedForAADTenantIds []*string `json:"linkingAllowedForAadTenantIds,omitempty" tf:"linking_allowed_for_aad_tenant_ids,omitempty"`

	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Workspace managed resource group. Changing this forces a new resource to be created.
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`

	// Is Virtual Network enabled for all computes in this workspace? Changing this forces a new resource to be created.
	ManagedVirtualNetworkEnabled *bool `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled,omitempty"`

	// Whether public network access is allowed for the Cognitive Account. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The ID of purview account.
	PurviewID *string `json:"purviewId,omitempty" tf:"purview_id,omitempty"`

	// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// An sql_aad_admin block as defined below.
	SQLAADAdmin []SQLAADAdminObservation `json:"sqlAadAdmin,omitempty" tf:"sql_aad_admin,omitempty"`

	// Specifies The login name of the SQL administrator. Changing this forces a new resource to be created. If this is not provided aad_admin or customer_managed_key must be provided.
	SQLAdministratorLogin *string `json:"sqlAdministratorLogin,omitempty" tf:"sql_administrator_login,omitempty"`

	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	SQLIdentityControlEnabled *bool `json:"sqlIdentityControlEnabled,omitempty" tf:"sql_identity_control_enabled,omitempty"`

	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	StorageDataLakeGen2FileSystemID *string `json:"storageDataLakeGen2FilesystemId,omitempty" tf:"storage_data_lake_gen2_filesystem_id,omitempty"`

	// A mapping of tags which should be assigned to the Synapse Workspace.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WorkspaceParameters struct {

	// An aad_admin block as defined below. Conflicts with customer_managed_key.
	// +kubebuilder:validation:Optional
	AADAdmin []AADAdminParameters `json:"aadAdmin,omitempty" tf:"aad_admin,omitempty"`

	// An azure_devops_repo block as defined below.
	// +kubebuilder:validation:Optional
	AzureDevopsRepo []AzureDevopsRepoParameters `json:"azureDevopsRepo,omitempty" tf:"azure_devops_repo,omitempty"`

	// Subnet ID used for computes in workspace Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ComputeSubnetID *string `json:"computeSubnetId,omitempty" tf:"compute_subnet_id,omitempty"`

	// Reference to a Subnet in network to populate computeSubnetId.
	// +kubebuilder:validation:Optional
	ComputeSubnetIDRef *v1.Reference `json:"computeSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate computeSubnetId.
	// +kubebuilder:validation:Optional
	ComputeSubnetIDSelector *v1.Selector `json:"computeSubnetIdSelector,omitempty" tf:"-"`

	// A customer_managed_key block as defined below. Conflicts with aad_admin.
	// +kubebuilder:validation:Optional
	CustomerManagedKey []CustomerManagedKeyParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Is data exfiltration protection enabled in this workspace? If set to true, managed_virtual_network_enabled must also be set to true. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataExfiltrationProtectionEnabled *bool `json:"dataExfiltrationProtectionEnabled,omitempty" tf:"data_exfiltration_protection_enabled,omitempty"`

	// A github_repo block as defined below.
	// +kubebuilder:validation:Optional
	GithubRepo []GithubRepoParameters `json:"githubRepo,omitempty" tf:"github_repo,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Allowed AAD Tenant Ids For Linking.
	// +kubebuilder:validation:Optional
	LinkingAllowedForAADTenantIds []*string `json:"linkingAllowedForAadTenantIds,omitempty" tf:"linking_allowed_for_aad_tenant_ids,omitempty"`

	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Workspace managed resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ManagedResourceGroupName *string `json:"managedResourceGroupName,omitempty" tf:"managed_resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameRef *v1.Reference `json:"managedResourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate managedResourceGroupName.
	// +kubebuilder:validation:Optional
	ManagedResourceGroupNameSelector *v1.Selector `json:"managedResourceGroupNameSelector,omitempty" tf:"-"`

	// Is Virtual Network enabled for all computes in this workspace? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ManagedVirtualNetworkEnabled *bool `json:"managedVirtualNetworkEnabled,omitempty" tf:"managed_virtual_network_enabled,omitempty"`

	// Whether public network access is allowed for the Cognitive Account. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The ID of purview account.
	// +kubebuilder:validation:Optional
	PurviewID *string `json:"purviewId,omitempty" tf:"purview_id,omitempty"`

	// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// An sql_aad_admin block as defined below.
	// +kubebuilder:validation:Optional
	SQLAADAdmin []SQLAADAdminParameters `json:"sqlAadAdmin,omitempty" tf:"sql_aad_admin,omitempty"`

	// Specifies The login name of the SQL administrator. Changing this forces a new resource to be created. If this is not provided aad_admin or customer_managed_key must be provided.
	// +kubebuilder:validation:Optional
	SQLAdministratorLogin *string `json:"sqlAdministratorLogin,omitempty" tf:"sql_administrator_login,omitempty"`

	// The Password associated with the sql_administrator_login for the SQL administrator. If this is not provided aad_admin or customer_managed_key must be provided.
	// +kubebuilder:validation:Optional
	SQLAdministratorLoginPasswordSecretRef *v1.SecretKeySelector `json:"sqlAdministratorLoginPasswordSecretRef,omitempty" tf:"-"`

	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	// +kubebuilder:validation:Optional
	SQLIdentityControlEnabled *bool `json:"sqlIdentityControlEnabled,omitempty" tf:"sql_identity_control_enabled,omitempty"`

	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.DataLakeGen2FileSystem
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageDataLakeGen2FileSystemID *string `json:"storageDataLakeGen2FilesystemId,omitempty" tf:"storage_data_lake_gen2_filesystem_id,omitempty"`

	// Reference to a DataLakeGen2FileSystem in storage to populate storageDataLakeGen2FilesystemId.
	// +kubebuilder:validation:Optional
	StorageDataLakeGen2FileSystemIDRef *v1.Reference `json:"storageDataLakeGen2FilesystemIdRef,omitempty" tf:"-"`

	// Selector for a DataLakeGen2FileSystem in storage to populate storageDataLakeGen2FilesystemId.
	// +kubebuilder:validation:Optional
	StorageDataLakeGen2FileSystemIDSelector *v1.Selector `json:"storageDataLakeGen2FilesystemIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Synapse Workspace.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
type WorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceParameters `json:"forProvider"`
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
	InitProvider WorkspaceInitParameters `json:"initProvider,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace.
type WorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Workspace is the Schema for the Workspaces API. Manages a Synapse Workspace.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   WorkspaceSpec   `json:"spec"`
	Status WorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceList contains a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

// Repository type metadata.
var (
	Workspace_Kind             = "Workspace"
	Workspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workspace_Kind}.String()
	Workspace_KindAPIVersion   = Workspace_Kind + "." + CRDGroupVersion.String()
	Workspace_GroupVersionKind = CRDGroupVersion.WithKind(Workspace_Kind)
)

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
