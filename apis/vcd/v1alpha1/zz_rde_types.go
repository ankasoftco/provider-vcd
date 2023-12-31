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

type RdeObservation struct {

	// A computed representation of the actual Runtime Defined Entity JSON retrieved from VCD. Useful to see the actual entity contents if it is being changed by a third party in VCD
	ComputedEntity *string `json:"computedEntity,omitempty" tf:"computed_entity,omitempty"`

	// If true, `computed_entity` is equal to either `input_entity` or the contents of `input_entity_url`
	EntityInSync *bool `json:"entityInSync,omitempty" tf:"entity_in_sync,omitempty"`

	// An external entity's ID that this Runtime Defined Entity may have a relation to
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A JSON representation of the Runtime Defined Entity that is defined by the user and is used to initialize/override its contents
	InputEntity *string `json:"inputEntity,omitempty" tf:"input_entity,omitempty"`

	// URL that should point to a JSON representation of the Runtime Defined Entity and is used to initialize/override its contents
	InputEntityURL *string `json:"inputEntityUrl,omitempty" tf:"input_entity_url,omitempty"`

	// The name of the Runtime Defined Entity. It can be non-unique
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization that will own this Runtime Defined Entity, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The organization of the Runtime Defined Entity
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// The ID of the user that owns the Runtime Defined Entity
	OwnerUserID *string `json:"ownerUserId,omitempty" tf:"owner_user_id,omitempty"`

	// The Runtime Defined Entity Type ID
	RdeTypeID *string `json:"rdeTypeId,omitempty" tf:"rde_type_id,omitempty"`

	// If `true`, the Runtime Defined Entity will be resolved by this provider. If `false`, it won't beresolved and must be done either by an external component action or by an update. The Runtime Defined Entity can't bedeleted until the entity is resolved.
	Resolve *bool `json:"resolve,omitempty" tf:"resolve,omitempty"`

	// If `true`, the Runtime Defined Entity will be resolved before it gets deleted, to ensure forced deletion.Destroy will fail if it is not resolved.
	ResolveOnRemoval *bool `json:"resolveOnRemoval,omitempty" tf:"resolve_on_removal,omitempty"`

	// Specifies whether the entity is correctly resolved or not. When created it will be in PRE_CREATED state. If the entity is correctly validated against its RDE Type schema, the state will be RESOLVED,otherwise it will be RESOLUTION_ERROR. If an entity resolution ends in a RESOLUTION_ERROR state, it will require to be updated to a correct JSON to be usable
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type RdeParameters struct {

	// An external entity's ID that this Runtime Defined Entity may have a relation to
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// A JSON representation of the Runtime Defined Entity that is defined by the user and is used to initialize/override its contents
	// +kubebuilder:validation:Optional
	InputEntity *string `json:"inputEntity,omitempty" tf:"input_entity,omitempty"`

	// URL that should point to a JSON representation of the Runtime Defined Entity and is used to initialize/override its contents
	// +kubebuilder:validation:Optional
	InputEntityURL *string `json:"inputEntityUrl,omitempty" tf:"input_entity_url,omitempty"`

	// The name of the Runtime Defined Entity. It can be non-unique
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization that will own this Runtime Defined Entity, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// The Runtime Defined Entity Type ID
	// +kubebuilder:validation:Optional
	RdeTypeID *string `json:"rdeTypeId,omitempty" tf:"rde_type_id,omitempty"`

	// If `true`, the Runtime Defined Entity will be resolved by this provider. If `false`, it won't beresolved and must be done either by an external component action or by an update. The Runtime Defined Entity can't bedeleted until the entity is resolved.
	// +kubebuilder:validation:Optional
	Resolve *bool `json:"resolve,omitempty" tf:"resolve,omitempty"`

	// If `true`, the Runtime Defined Entity will be resolved before it gets deleted, to ensure forced deletion.Destroy will fail if it is not resolved.
	// +kubebuilder:validation:Optional
	ResolveOnRemoval *bool `json:"resolveOnRemoval,omitempty" tf:"resolve_on_removal,omitempty"`
}

// RdeSpec defines the desired state of Rde
type RdeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RdeParameters `json:"forProvider"`
}

// RdeStatus defines the observed state of Rde.
type RdeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RdeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rde is the Schema for the Rdes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type Rde struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.rdeTypeId)",message="rdeTypeId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.resolve)",message="resolve is a required parameter"
	Spec   RdeSpec   `json:"spec"`
	Status RdeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdeList contains a list of Rdes
type RdeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rde `json:"items"`
}

// Repository type metadata.
var (
	Rde_Kind             = "Rde"
	Rde_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rde_Kind}.String()
	Rde_KindAPIVersion   = Rde_Kind + "." + CRDGroupVersion.String()
	Rde_GroupVersionKind = CRDGroupVersion.WithKind(Rde_Kind)
)

func init() {
	SchemeBuilder.Register(&Rde{}, &RdeList{})
}
