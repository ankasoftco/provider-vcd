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

type VAppOrgNetworkObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are accessed in this vApp
	IsFenced *bool `json:"isFenced,omitempty" tf:"is_fenced,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Organization network name to which vApp network is connected to
	OrgNetworkName *string `json:"orgNetworkName,omitempty" tf:"org_network_name,omitempty"`

	// Specifies whether the vApp should be rebooted when the vApp network is removed. Default is false.
	RebootVappOnRemoval *bool `json:"rebootVappOnRemoval,omitempty" tf:"reboot_vapp_on_removal,omitempty"`

	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	RetainIPMacEnabled *bool `json:"retainIpMacEnabled,omitempty" tf:"retain_ip_mac_enabled,omitempty"`

	// vApp network name
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppOrgNetworkParameters struct {

	// Fencing allows identical virtual machines in different vApp networks connect to organization VDC networks that are accessed in this vApp
	// +kubebuilder:validation:Optional
	IsFenced *bool `json:"isFenced,omitempty" tf:"is_fenced,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Organization network name to which vApp network is connected to
	// +kubebuilder:validation:Optional
	OrgNetworkName *string `json:"orgNetworkName,omitempty" tf:"org_network_name,omitempty"`

	// Specifies whether the vApp should be rebooted when the vApp network is removed. Default is false.
	// +kubebuilder:validation:Optional
	RebootVappOnRemoval *bool `json:"rebootVappOnRemoval,omitempty" tf:"reboot_vapp_on_removal,omitempty"`

	// Specifies whether the network resources such as IP/MAC of router will be retained across deployments. Default is false.
	// +kubebuilder:validation:Optional
	RetainIPMacEnabled *bool `json:"retainIpMacEnabled,omitempty" tf:"retain_ip_mac_enabled,omitempty"`

	// vApp network name
	// +kubebuilder:validation:Optional
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// VAppOrgNetworkSpec defines the desired state of VAppOrgNetwork
type VAppOrgNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VAppOrgNetworkParameters `json:"forProvider"`
}

// VAppOrgNetworkStatus defines the observed state of VAppOrgNetwork.
type VAppOrgNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VAppOrgNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VAppOrgNetwork is the Schema for the VAppOrgNetworks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type VAppOrgNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.orgNetworkName)",message="orgNetworkName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vappName)",message="vappName is a required parameter"
	Spec   VAppOrgNetworkSpec   `json:"spec"`
	Status VAppOrgNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VAppOrgNetworkList contains a list of VAppOrgNetworks
type VAppOrgNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VAppOrgNetwork `json:"items"`
}

// Repository type metadata.
var (
	VAppOrgNetwork_Kind             = "VAppOrgNetwork"
	VAppOrgNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VAppOrgNetwork_Kind}.String()
	VAppOrgNetwork_KindAPIVersion   = VAppOrgNetwork_Kind + "." + CRDGroupVersion.String()
	VAppOrgNetwork_GroupVersionKind = CRDGroupVersion.WithKind(VAppOrgNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&VAppOrgNetwork{}, &VAppOrgNetworkList{})
}
