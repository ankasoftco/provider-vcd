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

type RdeTypeBehaviorObservation struct {

	// The description of the contract of the overridden Behavior
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Execution map of the Behavior that overrides the original
	Execution map[string]*string `json:"execution,omitempty" tf:"execution,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the overridden Behavior
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the original RDE Interface Behavior to override
	RdeInterfaceBehaviorID *string `json:"rdeInterfaceBehaviorId,omitempty" tf:"rde_interface_behavior_id,omitempty"`

	// The ID of the RDE Type that owns the Behavior override
	RdeTypeID *string `json:"rdeTypeId,omitempty" tf:"rde_type_id,omitempty"`

	// The Behavior invocation reference to be used for polymorphic behavior invocations
	Ref *string `json:"ref,omitempty" tf:"ref,omitempty"`
}

type RdeTypeBehaviorParameters struct {

	// The description of the contract of the overridden Behavior
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Execution map of the Behavior that overrides the original
	// +kubebuilder:validation:Optional
	Execution map[string]*string `json:"execution,omitempty" tf:"execution,omitempty"`

	// The ID of the original RDE Interface Behavior to override
	// +kubebuilder:validation:Optional
	RdeInterfaceBehaviorID *string `json:"rdeInterfaceBehaviorId,omitempty" tf:"rde_interface_behavior_id,omitempty"`

	// The ID of the RDE Type that owns the Behavior override
	// +kubebuilder:validation:Optional
	RdeTypeID *string `json:"rdeTypeId,omitempty" tf:"rde_type_id,omitempty"`
}

// RdeTypeBehaviorSpec defines the desired state of RdeTypeBehavior
type RdeTypeBehaviorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RdeTypeBehaviorParameters `json:"forProvider"`
}

// RdeTypeBehaviorStatus defines the observed state of RdeTypeBehavior.
type RdeTypeBehaviorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RdeTypeBehaviorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RdeTypeBehavior is the Schema for the RdeTypeBehaviors API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type RdeTypeBehavior struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.rdeInterfaceBehaviorId)",message="rdeInterfaceBehaviorId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.rdeTypeId)",message="rdeTypeId is a required parameter"
	Spec   RdeTypeBehaviorSpec   `json:"spec"`
	Status RdeTypeBehaviorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdeTypeBehaviorList contains a list of RdeTypeBehaviors
type RdeTypeBehaviorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdeTypeBehavior `json:"items"`
}

// Repository type metadata.
var (
	RdeTypeBehavior_Kind             = "RdeTypeBehavior"
	RdeTypeBehavior_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RdeTypeBehavior_Kind}.String()
	RdeTypeBehavior_KindAPIVersion   = RdeTypeBehavior_Kind + "." + CRDGroupVersion.String()
	RdeTypeBehavior_GroupVersionKind = CRDGroupVersion.WithKind(RdeTypeBehavior_Kind)
)

func init() {
	SchemeBuilder.Register(&RdeTypeBehavior{}, &RdeTypeBehaviorList{})
}
