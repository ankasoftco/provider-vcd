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

type DHCPPoolObservation struct {

	// The default DHCP lease time to use
	DefaultLeaseTime *float64 `json:"defaultLeaseTime,omitempty" tf:"default_lease_time,omitempty"`

	// The final address in the IP Range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// The maximum DHCP lease time to use
	MaxLeaseTime *float64 `json:"maxLeaseTime,omitempty" tf:"max_lease_time,omitempty"`

	// The first address in the IP Range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type DHCPPoolParameters struct {

	// The default DHCP lease time to use
	// +kubebuilder:validation:Optional
	DefaultLeaseTime *float64 `json:"defaultLeaseTime,omitempty" tf:"default_lease_time,omitempty"`

	// The final address in the IP Range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// The maximum DHCP lease time to use
	// +kubebuilder:validation:Optional
	MaxLeaseTime *float64 `json:"maxLeaseTime,omitempty" tf:"max_lease_time,omitempty"`

	// The first address in the IP Range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

type NetworkIsolatedMetadataEntryObservation struct {

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

type NetworkIsolatedMetadataEntryParameters struct {

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

type NetworkIsolatedObservation struct {

	// A range of IPs to issue to virtual machines that don't have a static IP
	DHCPPool []DHCPPoolObservation `json:"dhcpPool,omitempty" tf:"dhcp_pool,omitempty"`

	// A FQDN for the virtual machines on this network
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Optional description for the network
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// First DNS server to use
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Second DNS server to use
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// The gateway for this network
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Network Hyper Reference
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key value map of metadata to assign to this network. Key and value can be any string
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given Network
	MetadataEntry []NetworkIsolatedMetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A unique name for this network
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The netmask for the new network
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Defines if this network is shared between multiple VDCs in the Org
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// A range of IPs permitted to be used as static IPs for virtual machines
	StaticIPPool []NetworkIsolatedStaticIPPoolObservation `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NetworkIsolatedParameters struct {

	// A range of IPs to issue to virtual machines that don't have a static IP
	// +kubebuilder:validation:Optional
	DHCPPool []DHCPPoolParameters `json:"dhcpPool,omitempty" tf:"dhcp_pool,omitempty"`

	// A FQDN for the virtual machines on this network
	// +kubebuilder:validation:Optional
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	// Optional description for the network
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// First DNS server to use
	// +kubebuilder:validation:Optional
	Dns1 *string `json:"dns1,omitempty" tf:"dns1,omitempty"`

	// Second DNS server to use
	// +kubebuilder:validation:Optional
	Dns2 *string `json:"dns2,omitempty" tf:"dns2,omitempty"`

	// The gateway for this network
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Key value map of metadata to assign to this network. Key and value can be any string
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given Network
	// +kubebuilder:validation:Optional
	MetadataEntry []NetworkIsolatedMetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A unique name for this network
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The netmask for the new network
	// +kubebuilder:validation:Optional
	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Defines if this network is shared between multiple VDCs in the Org
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// A range of IPs permitted to be used as static IPs for virtual machines
	// +kubebuilder:validation:Optional
	StaticIPPool []NetworkIsolatedStaticIPPoolParameters `json:"staticIpPool,omitempty" tf:"static_ip_pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NetworkIsolatedStaticIPPoolObservation struct {

	// The final address in the IP Range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// The first address in the IP Range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type NetworkIsolatedStaticIPPoolParameters struct {

	// The final address in the IP Range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// The first address in the IP Range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

// NetworkIsolatedSpec defines the desired state of NetworkIsolated
type NetworkIsolatedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkIsolatedParameters `json:"forProvider"`
}

// NetworkIsolatedStatus defines the observed state of NetworkIsolated.
type NetworkIsolatedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkIsolatedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkIsolated is the Schema for the NetworkIsolateds API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NetworkIsolated struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   NetworkIsolatedSpec   `json:"spec"`
	Status NetworkIsolatedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkIsolatedList contains a list of NetworkIsolateds
type NetworkIsolatedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkIsolated `json:"items"`
}

// Repository type metadata.
var (
	NetworkIsolated_Kind             = "NetworkIsolated"
	NetworkIsolated_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkIsolated_Kind}.String()
	NetworkIsolated_KindAPIVersion   = NetworkIsolated_Kind + "." + CRDGroupVersion.String()
	NetworkIsolated_GroupVersionKind = CRDGroupVersion.WithKind(NetworkIsolated_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkIsolated{}, &NetworkIsolatedList{})
}
