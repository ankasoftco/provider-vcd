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

type RdeTypeObservation struct {

	// The description of the Runtime Defined Entity Type
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An external entity's ID that this definition may apply to
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// To be used when creating a new version of a Runtime Defined Entity Type. Specifies the version of the type that will be the template for the authorization configuration of the new version.The Type ACLs and the access requirements of the Type Behaviors of the new version will be copied from those of the inherited version.If not set, then the new type version will not inherit another version and will have the default authorization settings, just like the first version of a new type
	InheritedVersion *string `json:"inheritedVersion,omitempty" tf:"inherited_version,omitempty"`

	// Set of Defined Interface URNs that this Runtime Defined Entity Type is referenced by
	InterfaceIds []*string `json:"interfaceIds,omitempty" tf:"interface_ids,omitempty"`

	// The name of the Runtime Defined Entity Type
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A unique namespace associated with the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique
	Nss *string `json:"nss,omitempty" tf:"nss,omitempty"`

	// True if the Runtime Defined Entity Type cannot be modified
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`

	// The JSON-Schema valid definition of the Runtime Defined Entity Type
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// URL that should point to a JSON-Schema valid definition file of the Runtime Defined Entity Type
	SchemaURL *string `json:"schemaUrl,omitempty" tf:"schema_url,omitempty"`

	// The vendor name for the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`

	// The version of the Runtime Defined Entity Type. The version string must follow semantic versioning rules. Combination of `vendor`, `nss` and `version` must be unique
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RdeTypeParameters struct {

	// The description of the Runtime Defined Entity Type
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An external entity's ID that this definition may apply to
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// To be used when creating a new version of a Runtime Defined Entity Type. Specifies the version of the type that will be the template for the authorization configuration of the new version.The Type ACLs and the access requirements of the Type Behaviors of the new version will be copied from those of the inherited version.If not set, then the new type version will not inherit another version and will have the default authorization settings, just like the first version of a new type
	// +kubebuilder:validation:Optional
	InheritedVersion *string `json:"inheritedVersion,omitempty" tf:"inherited_version,omitempty"`

	// Set of Defined Interface URNs that this Runtime Defined Entity Type is referenced by
	// +kubebuilder:validation:Optional
	InterfaceIds []*string `json:"interfaceIds,omitempty" tf:"interface_ids,omitempty"`

	// The name of the Runtime Defined Entity Type
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A unique namespace associated with the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique
	// +kubebuilder:validation:Optional
	Nss *string `json:"nss,omitempty" tf:"nss,omitempty"`

	// The JSON-Schema valid definition of the Runtime Defined Entity Type
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`

	// URL that should point to a JSON-Schema valid definition file of the Runtime Defined Entity Type
	// +kubebuilder:validation:Optional
	SchemaURL *string `json:"schemaUrl,omitempty" tf:"schema_url,omitempty"`

	// The vendor name for the Runtime Defined Entity Type. Combination of `vendor`, `nss` and `version` must be unique
	// +kubebuilder:validation:Optional
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`

	// The version of the Runtime Defined Entity Type. The version string must follow semantic versioning rules. Combination of `vendor`, `nss` and `version` must be unique
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// RdeTypeSpec defines the desired state of RdeType
type RdeTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RdeTypeParameters `json:"forProvider"`
}

// RdeTypeStatus defines the observed state of RdeType.
type RdeTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RdeTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RdeType is the Schema for the RdeTypes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type RdeType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.nss)",message="nss is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vendor)",message="vendor is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.version)",message="version is a required parameter"
	Spec   RdeTypeSpec   `json:"spec"`
	Status RdeTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdeTypeList contains a list of RdeTypes
type RdeTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdeType `json:"items"`
}

// Repository type metadata.
var (
	RdeType_Kind             = "RdeType"
	RdeType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RdeType_Kind}.String()
	RdeType_KindAPIVersion   = RdeType_Kind + "." + CRDGroupVersion.String()
	RdeType_GroupVersionKind = CRDGroupVersion.WithKind(RdeType_Kind)
)

func init() {
	SchemeBuilder.Register(&RdeType{}, &RdeTypeList{})
}
