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

type IPPrefixObservation struct {

	// Floating IP quota
	DefaultQuota *string `json:"defaultQuota,omitempty" tf:"default_quota,omitempty"`

	// One or more prefixes
	Prefix []PrefixObservation `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type IPPrefixParameters struct {

	// Floating IP quota
	// +kubebuilder:validation:Optional
	DefaultQuota *string `json:"defaultQuota,omitempty" tf:"default_quota,omitempty"`

	// One or more prefixes
	// +kubebuilder:validation:Required
	Prefix []PrefixParameters `json:"prefix" tf:"prefix,omitempty"`
}

type IPRangeObservation struct {

	// End address of the IP range
	EndAddress *string `json:"endAddress,omitempty" tf:"end_address,omitempty"`

	// ID of IP Range
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Start address of the IP range
	StartAddress *string `json:"startAddress,omitempty" tf:"start_address,omitempty"`
}

type IPRangeParameters struct {

	// End address of the IP range
	// +kubebuilder:validation:Required
	EndAddress *string `json:"endAddress" tf:"end_address,omitempty"`

	// Start address of the IP range
	// +kubebuilder:validation:Required
	StartAddress *string `json:"startAddress" tf:"start_address,omitempty"`
}

type IpSpaceObservation struct {

	// Description of IP space
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// External scope in CIDR format
	ExternalScope *string `json:"externalScope,omitempty" tf:"external_scope,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more IP prefixes within internal scope
	IPPrefix []IPPrefixObservation `json:"ipPrefix,omitempty" tf:"ip_prefix,omitempty"`

	// One or more IP ranges for floating IP allocation
	IPRange []IPRangeObservation `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// IP ranges quota. '-1' - unlimited, '0' - no quota
	IPRangeQuota *string `json:"ipRangeQuota,omitempty" tf:"ip_range_quota,omitempty"`

	// A set of internal scope IPs in CIDR format
	InternalScope []*string `json:"internalScope,omitempty" tf:"internal_scope,omitempty"`

	// Name of IP space
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Org ID for 'SHARED' IP spaces
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Flag whether route advertisement should be enabled
	RouteAdvertisementEnabled *bool `json:"routeAdvertisementEnabled,omitempty" tf:"route_advertisement_enabled,omitempty"`

	// Type of IP space
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IpSpaceParameters struct {

	// Description of IP space
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// External scope in CIDR format
	// +kubebuilder:validation:Optional
	ExternalScope *string `json:"externalScope,omitempty" tf:"external_scope,omitempty"`

	// One or more IP prefixes within internal scope
	// +kubebuilder:validation:Optional
	IPPrefix []IPPrefixParameters `json:"ipPrefix,omitempty" tf:"ip_prefix,omitempty"`

	// One or more IP ranges for floating IP allocation
	// +kubebuilder:validation:Optional
	IPRange []IPRangeParameters `json:"ipRange,omitempty" tf:"ip_range,omitempty"`

	// IP ranges quota. '-1' - unlimited, '0' - no quota
	// +kubebuilder:validation:Optional
	IPRangeQuota *string `json:"ipRangeQuota,omitempty" tf:"ip_range_quota,omitempty"`

	// A set of internal scope IPs in CIDR format
	// +kubebuilder:validation:Optional
	InternalScope []*string `json:"internalScope,omitempty" tf:"internal_scope,omitempty"`

	// Name of IP space
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Org ID for 'SHARED' IP spaces
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Flag whether route advertisement should be enabled
	// +kubebuilder:validation:Optional
	RouteAdvertisementEnabled *bool `json:"routeAdvertisementEnabled,omitempty" tf:"route_advertisement_enabled,omitempty"`

	// Type of IP space
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PrefixObservation struct {

	// First IP
	FirstIP *string `json:"firstIp,omitempty" tf:"first_ip,omitempty"`

	// ID of IP Prefix
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Number of prefixes to define
	PrefixCount *string `json:"prefixCount,omitempty" tf:"prefix_count,omitempty"`

	// Prefix length
	PrefixLength *string `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`
}

type PrefixParameters struct {

	// First IP
	// +kubebuilder:validation:Required
	FirstIP *string `json:"firstIp" tf:"first_ip,omitempty"`

	// Number of prefixes to define
	// +kubebuilder:validation:Required
	PrefixCount *string `json:"prefixCount" tf:"prefix_count,omitempty"`

	// Prefix length
	// +kubebuilder:validation:Required
	PrefixLength *string `json:"prefixLength" tf:"prefix_length,omitempty"`
}

// IpSpaceSpec defines the desired state of IpSpace
type IpSpaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpSpaceParameters `json:"forProvider"`
}

// IpSpaceStatus defines the observed state of IpSpace.
type IpSpaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpSpaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IpSpace is the Schema for the IpSpaces API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type IpSpace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.internalScope)",message="internalScope is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.type)",message="type is a required parameter"
	Spec   IpSpaceSpec   `json:"spec"`
	Status IpSpaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpSpaceList contains a list of IpSpaces
type IpSpaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IpSpace `json:"items"`
}

// Repository type metadata.
var (
	IpSpace_Kind             = "IpSpace"
	IpSpace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IpSpace_Kind}.String()
	IpSpace_KindAPIVersion   = IpSpace_Kind + "." + CRDGroupVersion.String()
	IpSpace_GroupVersionKind = CRDGroupVersion.WithKind(IpSpace_Kind)
)

func init() {
	SchemeBuilder.Register(&IpSpace{}, &IpSpaceList{})
}
