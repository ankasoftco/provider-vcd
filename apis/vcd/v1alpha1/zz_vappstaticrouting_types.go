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

type VAppStaticRoutingObservation struct {

	// Enable or disable static Routing. Default is `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// vApp network identifier
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	Rule []VAppStaticRoutingRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// vApp identifier
	VappID *string `json:"vappId,omitempty" tf:"vapp_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppStaticRoutingParameters struct {

	// Enable or disable static Routing. Default is `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// vApp network identifier
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// +kubebuilder:validation:Optional
	Rule []VAppStaticRoutingRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// vApp identifier
	// +kubebuilder:validation:Optional
	VappID *string `json:"vappId,omitempty" tf:"vapp_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppStaticRoutingRuleObservation struct {

	// Name for the static route.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// network specification in CIDR.
	NetworkCidr *string `json:"networkCidr,omitempty" tf:"network_cidr,omitempty"`

	// IP Address of Next Hop router/gateway.
	NextHopIP *string `json:"nextHopIp,omitempty" tf:"next_hop_ip,omitempty"`
}

type VAppStaticRoutingRuleParameters struct {

	// Name for the static route.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// network specification in CIDR.
	// +kubebuilder:validation:Required
	NetworkCidr *string `json:"networkCidr" tf:"network_cidr,omitempty"`

	// IP Address of Next Hop router/gateway.
	// +kubebuilder:validation:Required
	NextHopIP *string `json:"nextHopIp" tf:"next_hop_ip,omitempty"`
}

// VAppStaticRoutingSpec defines the desired state of VAppStaticRouting
type VAppStaticRoutingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VAppStaticRoutingParameters `json:"forProvider"`
}

// VAppStaticRoutingStatus defines the observed state of VAppStaticRouting.
type VAppStaticRoutingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VAppStaticRoutingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VAppStaticRouting is the Schema for the VAppStaticRoutings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type VAppStaticRouting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.networkId)",message="networkId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vappId)",message="vappId is a required parameter"
	Spec   VAppStaticRoutingSpec   `json:"spec"`
	Status VAppStaticRoutingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VAppStaticRoutingList contains a list of VAppStaticRoutings
type VAppStaticRoutingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VAppStaticRouting `json:"items"`
}

// Repository type metadata.
var (
	VAppStaticRouting_Kind             = "VAppStaticRouting"
	VAppStaticRouting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VAppStaticRouting_Kind}.String()
	VAppStaticRouting_KindAPIVersion   = VAppStaticRouting_Kind + "." + CRDGroupVersion.String()
	VAppStaticRouting_GroupVersionKind = CRDGroupVersion.WithKind(VAppStaticRouting_Kind)
)

func init() {
	SchemeBuilder.Register(&VAppStaticRouting{}, &VAppStaticRoutingList{})
}
