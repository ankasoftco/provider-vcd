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

type ClonedvAppObservation struct {

	// If true, it will delete the source (vApp or template) after creating the new vApp
	DeleteSource *bool `json:"deleteSource,omitempty" tf:"delete_source,omitempty"`

	// Optional description of the vApp
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// vApp Hyper Reference
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A name for the vApp, unique withing the VDC
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// The identifier of the source to use for the creation of this vApp
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// The type of the source to use for the creation of this vApp (one of 'vapp' or 'template')
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// Shows the status code of the vApp
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// Shows the status of the vApp
	StatusText *string `json:"statusText,omitempty" tf:"status_text,omitempty"`

	// List of VMs contained in this vApp (in alphabetic order)
	VMList []*string `json:"vmList,omitempty" tf:"vm_list,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type ClonedvAppParameters struct {

	// If true, it will delete the source (vApp or template) after creating the new vApp
	// +kubebuilder:validation:Optional
	DeleteSource *bool `json:"deleteSource,omitempty" tf:"delete_source,omitempty"`

	// Optional description of the vApp
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A name for the vApp, unique withing the VDC
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A boolean value stating if this vApp should be powered on
	// +kubebuilder:validation:Optional
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// The identifier of the source to use for the creation of this vApp
	// +kubebuilder:validation:Optional
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// The type of the source to use for the creation of this vApp (one of 'vapp' or 'template')
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// ClonedvAppSpec defines the desired state of ClonedvApp
type ClonedvAppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClonedvAppParameters `json:"forProvider"`
}

// ClonedvAppStatus defines the observed state of ClonedvApp.
type ClonedvAppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClonedvAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClonedvApp is the Schema for the ClonedvApps API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type ClonedvApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.sourceId)",message="sourceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.sourceType)",message="sourceType is a required parameter"
	Spec   ClonedvAppSpec   `json:"spec"`
	Status ClonedvAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClonedvAppList contains a list of ClonedvApps
type ClonedvAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClonedvApp `json:"items"`
}

// Repository type metadata.
var (
	ClonedvApp_Kind             = "ClonedvApp"
	ClonedvApp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClonedvApp_Kind}.String()
	ClonedvApp_KindAPIVersion   = ClonedvApp_Kind + "." + CRDGroupVersion.String()
	ClonedvApp_GroupVersionKind = CRDGroupVersion.WithKind(ClonedvApp_Kind)
)

func init() {
	SchemeBuilder.Register(&ClonedvApp{}, &ClonedvAppList{})
}
