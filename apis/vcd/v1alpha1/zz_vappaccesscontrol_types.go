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

type VAppAccessControlObservation struct {

	// Access level when the vApp is shared with everyone (one of ReadOnly, Change, FullControl). Required when shared_with_everyone is set
	EveryoneAccessLevel *string `json:"everyoneAccessLevel,omitempty" tf:"everyone_access_level,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	SharedWith []VAppAccessControlSharedWithObservation `json:"sharedWith,omitempty" tf:"shared_with,omitempty"`

	// Whether the vApp is shared with everyone
	SharedWithEveryone *bool `json:"sharedWithEveryone,omitempty" tf:"shared_with_everyone,omitempty"`

	// vApp identifier
	VappID *string `json:"vappId,omitempty" tf:"vapp_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppAccessControlParameters struct {

	// Access level when the vApp is shared with everyone (one of ReadOnly, Change, FullControl). Required when shared_with_everyone is set
	// +kubebuilder:validation:Optional
	EveryoneAccessLevel *string `json:"everyoneAccessLevel,omitempty" tf:"everyone_access_level,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// +kubebuilder:validation:Optional
	SharedWith []VAppAccessControlSharedWithParameters `json:"sharedWith,omitempty" tf:"shared_with,omitempty"`

	// Whether the vApp is shared with everyone
	// +kubebuilder:validation:Optional
	SharedWithEveryone *bool `json:"sharedWithEveryone,omitempty" tf:"shared_with_everyone,omitempty"`

	// vApp identifier
	// +kubebuilder:validation:Optional
	VappID *string `json:"vappId,omitempty" tf:"vapp_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VAppAccessControlSharedWithObservation struct {

	// The access level for the user or group to which we are sharing. (One of ReadOnly, Change, FullControl)
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// ID of the group to which we are sharing. Required if user_id is not set
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Name of the subject (group or user) with which we are sharing
	SubjectName *string `json:"subjectName,omitempty" tf:"subject_name,omitempty"`

	// ID of the user to which we are sharing. Required if group_id is not set
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type VAppAccessControlSharedWithParameters struct {

	// The access level for the user or group to which we are sharing. (One of ReadOnly, Change, FullControl)
	// +kubebuilder:validation:Required
	AccessLevel *string `json:"accessLevel" tf:"access_level,omitempty"`

	// ID of the group to which we are sharing. Required if user_id is not set
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// ID of the user to which we are sharing. Required if group_id is not set
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// VAppAccessControlSpec defines the desired state of VAppAccessControl
type VAppAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VAppAccessControlParameters `json:"forProvider"`
}

// VAppAccessControlStatus defines the observed state of VAppAccessControl.
type VAppAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VAppAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VAppAccessControl is the Schema for the VAppAccessControls API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type VAppAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.sharedWithEveryone)",message="sharedWithEveryone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vappId)",message="vappId is a required parameter"
	Spec   VAppAccessControlSpec   `json:"spec"`
	Status VAppAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VAppAccessControlList contains a list of VAppAccessControls
type VAppAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VAppAccessControl `json:"items"`
}

// Repository type metadata.
var (
	VAppAccessControl_Kind             = "VAppAccessControl"
	VAppAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VAppAccessControl_Kind}.String()
	VAppAccessControl_KindAPIVersion   = VAppAccessControl_Kind + "." + CRDGroupVersion.String()
	VAppAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(VAppAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&VAppAccessControl{}, &VAppAccessControlList{})
}