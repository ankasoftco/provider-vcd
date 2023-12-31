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

type NsxtNetworkDhcpObservation struct {

	// The DNS server IPs to be assigned by this DHCP service. 2 values maximum.
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Lease time in seconds
	LeaseTime *float64 `json:"leaseTime,omitempty" tf:"lease_time,omitempty"`

	// IP Address of DHCP server in network. Only applicable when mode=NETWORK
	ListenerIPAddress *string `json:"listenerIpAddress,omitempty" tf:"listener_ip_address,omitempty"`

	// DHCP mode. One of 'EDGE' (default), 'NETWORK', 'RELAY'
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Parent Org VDC network ID
	OrgNetworkID *string `json:"orgNetworkId,omitempty" tf:"org_network_id,omitempty"`

	// IP ranges used for DHCP pool allocation in the network
	Pool []PoolObservation `json:"pool,omitempty" tf:"pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtNetworkDhcpParameters struct {

	// The DNS server IPs to be assigned by this DHCP service. 2 values maximum.
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Lease time in seconds
	// +kubebuilder:validation:Optional
	LeaseTime *float64 `json:"leaseTime,omitempty" tf:"lease_time,omitempty"`

	// IP Address of DHCP server in network. Only applicable when mode=NETWORK
	// +kubebuilder:validation:Optional
	ListenerIPAddress *string `json:"listenerIpAddress,omitempty" tf:"listener_ip_address,omitempty"`

	// DHCP mode. One of 'EDGE' (default), 'NETWORK', 'RELAY'
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Parent Org VDC network ID
	// +kubebuilder:validation:Optional
	OrgNetworkID *string `json:"orgNetworkId,omitempty" tf:"org_network_id,omitempty"`

	// IP ranges used for DHCP pool allocation in the network
	// +kubebuilder:validation:Optional
	Pool []PoolParameters `json:"pool,omitempty" tf:"pool,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type PoolObservation struct {

	// End address of DHCP pool IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// Start address of DHCP pool IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type PoolParameters struct {

	// End address of DHCP pool IP range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// Start address of DHCP pool IP range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

// NsxtNetworkDhcpSpec defines the desired state of NsxtNetworkDhcp
type NsxtNetworkDhcpSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtNetworkDhcpParameters `json:"forProvider"`
}

// NsxtNetworkDhcpStatus defines the observed state of NsxtNetworkDhcp.
type NsxtNetworkDhcpStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtNetworkDhcpObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtNetworkDhcp is the Schema for the NsxtNetworkDhcps API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtNetworkDhcp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.orgNetworkId)",message="orgNetworkId is a required parameter"
	Spec   NsxtNetworkDhcpSpec   `json:"spec"`
	Status NsxtNetworkDhcpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtNetworkDhcpList contains a list of NsxtNetworkDhcps
type NsxtNetworkDhcpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtNetworkDhcp `json:"items"`
}

// Repository type metadata.
var (
	NsxtNetworkDhcp_Kind             = "NsxtNetworkDhcp"
	NsxtNetworkDhcp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtNetworkDhcp_Kind}.String()
	NsxtNetworkDhcp_KindAPIVersion   = NsxtNetworkDhcp_Kind + "." + CRDGroupVersion.String()
	NsxtNetworkDhcp_GroupVersionKind = CRDGroupVersion.WithKind(NsxtNetworkDhcp_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtNetworkDhcp{}, &NsxtNetworkDhcpList{})
}
