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

type NsxtSecurityGroupMemberVmsObservation struct {
	VMID *string `json:"vmId,omitempty" tf:"vm_id,omitempty"`

	VMName *string `json:"vmName,omitempty" tf:"vm_name,omitempty"`

	VappID *string `json:"vappId,omitempty" tf:"vapp_id,omitempty"`

	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`
}

type NsxtSecurityGroupMemberVmsParameters struct {
}

type NsxtSecurityGroupObservation struct {

	// Security Group description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Edge Gateway ID in which security group is located
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of Org VDC network IDs attached to this security group
	MemberOrgNetworkIds []*string `json:"memberOrgNetworkIds,omitempty" tf:"member_org_network_ids,omitempty"`

	// Set of VM IDs
	MemberVms []NsxtSecurityGroupMemberVmsObservation `json:"memberVms,omitempty" tf:"member_vms,omitempty"`

	// Security Group name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// ID of VDC or VDC Group
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxtSecurityGroupParameters struct {

	// Security Group description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Edge Gateway ID in which security group is located
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Set of Org VDC network IDs attached to this security group
	// +kubebuilder:validation:Optional
	MemberOrgNetworkIds []*string `json:"memberOrgNetworkIds,omitempty" tf:"member_org_network_ids,omitempty"`

	// Security Group name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxtSecurityGroupSpec defines the desired state of NsxtSecurityGroup
type NsxtSecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtSecurityGroupParameters `json:"forProvider"`
}

// NsxtSecurityGroupStatus defines the observed state of NsxtSecurityGroup.
type NsxtSecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtSecurityGroup is the Schema for the NsxtSecurityGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.edgeGatewayId)",message="edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   NsxtSecurityGroupSpec   `json:"spec"`
	Status NsxtSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtSecurityGroupList contains a list of NsxtSecurityGroups
type NsxtSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	NsxtSecurityGroup_Kind             = "NsxtSecurityGroup"
	NsxtSecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtSecurityGroup_Kind}.String()
	NsxtSecurityGroup_KindAPIVersion   = NsxtSecurityGroup_Kind + "." + CRDGroupVersion.String()
	NsxtSecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(NsxtSecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtSecurityGroup{}, &NsxtSecurityGroupList{})
}
