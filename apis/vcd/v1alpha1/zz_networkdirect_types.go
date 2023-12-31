/*
Copyright 2023 ANKASOFT
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkDirectMetadataEntryObservation struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry. Required if the metadata entry is not empty
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry. Required if the metadata entry is not empty
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NetworkDirectMetadataEntryParameters struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL
	// +kubebuilder:validation:Optional
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	// +kubebuilder:validation:Optional
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type NetworkDirectObservation struct {

	// Optional description for the network
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the external network
	ExternalNetwork *string `json:"externalNetwork,omitempty" tf:"external_network,omitempty"`

	// DNS suffix of the external network
	ExternalNetworkDNSSuffix *string `json:"externalNetworkDnsSuffix,omitempty" tf:"external_network_dns_suffix,omitempty"`

	// Main DNS of the external network
	ExternalNetworkDns1 *string `json:"externalNetworkDns1,omitempty" tf:"external_network_dns1,omitempty"`

	// Secondary DNS of the external network
	ExternalNetworkDns2 *string `json:"externalNetworkDns2,omitempty" tf:"external_network_dns2,omitempty"`

	// Gateway of the external network
	ExternalNetworkGateway *string `json:"externalNetworkGateway,omitempty" tf:"external_network_gateway,omitempty"`

	// Net mask of the external network
	ExternalNetworkNetmask *string `json:"externalNetworkNetmask,omitempty" tf:"external_network_netmask,omitempty"`

	// Network Hypertext Reference
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key value map of metadata to assign to this network. Key and value can be any string
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given Network
	MetadataEntry []NetworkDirectMetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A unique name for this network
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Defines if this network is shared between multiple VDCs in the Org
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NetworkDirectParameters struct {

	// Optional description for the network
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the external network
	// +kubebuilder:validation:Optional
	ExternalNetwork *string `json:"externalNetwork,omitempty" tf:"external_network,omitempty"`

	// Key value map of metadata to assign to this network. Key and value can be any string
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given Network
	// +kubebuilder:validation:Optional
	MetadataEntry []NetworkDirectMetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A unique name for this network
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Defines if this network is shared between multiple VDCs in the Org
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NetworkDirectSpec defines the desired state of NetworkDirect
type NetworkDirectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkDirectParameters `json:"forProvider"`
}

// NetworkDirectStatus defines the observed state of NetworkDirect.
type NetworkDirectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkDirectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkDirect is the Schema for the NetworkDirects API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NetworkDirect struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.externalNetwork)",message="externalNetwork is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   NetworkDirectSpec   `json:"spec"`
	Status NetworkDirectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkDirectList contains a list of NetworkDirects
type NetworkDirectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkDirect `json:"items"`
}

// Repository type metadata.
var (
	NetworkDirect_Kind             = "NetworkDirect"
	NetworkDirect_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkDirect_Kind}.String()
	NetworkDirect_KindAPIVersion   = NetworkDirect_Kind + "." + CRDGroupVersion.String()
	NetworkDirect_GroupVersionKind = CRDGroupVersion.WithKind(NetworkDirect_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkDirect{}, &NetworkDirectList{})
}
