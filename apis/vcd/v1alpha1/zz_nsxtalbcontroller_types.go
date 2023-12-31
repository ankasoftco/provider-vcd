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

type NsxtAlbControllerObservation struct {

	// NSX-T ALB Controller description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// NSX-T ALB License type. One of 'BASIC', 'ENTERPRISE'. Must not be used from VCD 10.4.0 onwards
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// NSX-T ALB Controller name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// NSX-T ALB Controller URL
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// NSX-T ALB Controller Username
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// NSX-T ALB Controller version
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NsxtAlbControllerParameters struct {

	// NSX-T ALB Controller description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// NSX-T ALB License type. One of 'BASIC', 'ENTERPRISE'. Must not be used from VCD 10.4.0 onwards
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// NSX-T ALB Controller name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// NSX-T ALB Controller Password
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// NSX-T ALB Controller URL
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// NSX-T ALB Controller Username
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// NsxtAlbControllerSpec defines the desired state of NsxtAlbController
type NsxtAlbControllerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtAlbControllerParameters `json:"forProvider"`
}

// NsxtAlbControllerStatus defines the observed state of NsxtAlbController.
type NsxtAlbControllerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtAlbControllerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtAlbController is the Schema for the NsxtAlbControllers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtAlbController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.passwordSecretRef)",message="passwordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.url)",message="url is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.username)",message="username is a required parameter"
	Spec   NsxtAlbControllerSpec   `json:"spec"`
	Status NsxtAlbControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtAlbControllerList contains a list of NsxtAlbControllers
type NsxtAlbControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtAlbController `json:"items"`
}

// Repository type metadata.
var (
	NsxtAlbController_Kind             = "NsxtAlbController"
	NsxtAlbController_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtAlbController_Kind}.String()
	NsxtAlbController_KindAPIVersion   = NsxtAlbController_Kind + "." + CRDGroupVersion.String()
	NsxtAlbController_GroupVersionKind = CRDGroupVersion.WithKind(NsxtAlbController_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtAlbController{}, &NsxtAlbControllerList{})
}
