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

type RdeInterfaceBehaviorObservation struct {

	// A description specifying the contract of the Behavior
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Execution map of the Behavior
	Execution map[string]*string `json:"execution,omitempty" tf:"execution,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Behavior
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the RDE Interface that owns the Behavior
	RdeInterfaceID *string `json:"rdeInterfaceId,omitempty" tf:"rde_interface_id,omitempty"`

	// The Behavior invocation reference to be used for polymorphic behavior invocations
	Ref *string `json:"ref,omitempty" tf:"ref,omitempty"`
}

type RdeInterfaceBehaviorParameters struct {

	// A description specifying the contract of the Behavior
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Execution map of the Behavior
	// +kubebuilder:validation:Optional
	Execution map[string]*string `json:"execution,omitempty" tf:"execution,omitempty"`

	// Name of the Behavior
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the RDE Interface that owns the Behavior
	// +kubebuilder:validation:Optional
	RdeInterfaceID *string `json:"rdeInterfaceId,omitempty" tf:"rde_interface_id,omitempty"`
}

// RdeInterfaceBehaviorSpec defines the desired state of RdeInterfaceBehavior
type RdeInterfaceBehaviorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RdeInterfaceBehaviorParameters `json:"forProvider"`
}

// RdeInterfaceBehaviorStatus defines the observed state of RdeInterfaceBehavior.
type RdeInterfaceBehaviorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RdeInterfaceBehaviorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RdeInterfaceBehavior is the Schema for the RdeInterfaceBehaviors API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type RdeInterfaceBehavior struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.execution)",message="execution is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.rdeInterfaceId)",message="rdeInterfaceId is a required parameter"
	Spec   RdeInterfaceBehaviorSpec   `json:"spec"`
	Status RdeInterfaceBehaviorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdeInterfaceBehaviorList contains a list of RdeInterfaceBehaviors
type RdeInterfaceBehaviorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdeInterfaceBehavior `json:"items"`
}

// Repository type metadata.
var (
	RdeInterfaceBehavior_Kind             = "RdeInterfaceBehavior"
	RdeInterfaceBehavior_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RdeInterfaceBehavior_Kind}.String()
	RdeInterfaceBehavior_KindAPIVersion   = RdeInterfaceBehavior_Kind + "." + CRDGroupVersion.String()
	RdeInterfaceBehavior_GroupVersionKind = CRDGroupVersion.WithKind(RdeInterfaceBehavior_Kind)
)

func init() {
	SchemeBuilder.Register(&RdeInterfaceBehavior{}, &RdeInterfaceBehaviorList{})
}
