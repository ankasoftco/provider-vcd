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

type InsertedMediaObservation struct {

	// catalog name where to find media file
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// When ejecting answers automatically to question yes
	EjectForce *bool `json:"ejectForce,omitempty" tf:"eject_force,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// media name to use
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// VM in vApp in which media will be inserted or ejected
	VMName *string `json:"vmName,omitempty" tf:"vm_name,omitempty"`

	// vApp to use
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type InsertedMediaParameters struct {

	// catalog name where to find media file
	// +kubebuilder:validation:Optional
	Catalog *string `json:"catalog,omitempty" tf:"catalog,omitempty"`

	// When ejecting answers automatically to question yes
	// +kubebuilder:validation:Optional
	EjectForce *bool `json:"ejectForce,omitempty" tf:"eject_force,omitempty"`

	// media name to use
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// VM in vApp in which media will be inserted or ejected
	// +kubebuilder:validation:Optional
	VMName *string `json:"vmName,omitempty" tf:"vm_name,omitempty"`

	// vApp to use
	// +kubebuilder:validation:Optional
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// InsertedMediaSpec defines the desired state of InsertedMedia
type InsertedMediaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InsertedMediaParameters `json:"forProvider"`
}

// InsertedMediaStatus defines the observed state of InsertedMedia.
type InsertedMediaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InsertedMediaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InsertedMedia is the Schema for the InsertedMedias API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type InsertedMedia struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.catalog)",message="catalog is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vappName)",message="vappName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vmName)",message="vmName is a required parameter"
	Spec   InsertedMediaSpec   `json:"spec"`
	Status InsertedMediaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InsertedMediaList contains a list of InsertedMedias
type InsertedMediaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InsertedMedia `json:"items"`
}

// Repository type metadata.
var (
	InsertedMedia_Kind             = "InsertedMedia"
	InsertedMedia_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InsertedMedia_Kind}.String()
	InsertedMedia_KindAPIVersion   = InsertedMedia_Kind + "." + CRDGroupVersion.String()
	InsertedMedia_GroupVersionKind = CRDGroupVersion.WithKind(InsertedMedia_Kind)
)

func init() {
	SchemeBuilder.Register(&InsertedMedia{}, &InsertedMediaList{})
}
