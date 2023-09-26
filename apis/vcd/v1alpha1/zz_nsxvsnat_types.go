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

type NsxvSnatObservation struct {

	// NAT rule description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Edge gateway name in which NAT Rule is located
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Whether the rule should be enabled. Default 'true'
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether logging should be enabled for this rule. Default 'false'
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// Org or external network name
	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	// Network type. One of 'ext', 'org'
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Original address or address range. This is the the source address for SNAT rules
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// Optional. Allows to set custom rule tag
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Read only. Possible values 'user', 'internal_high'
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// Translated address or address range
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type NsxvSnatParameters struct {

	// NAT rule description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Edge gateway name in which NAT Rule is located
	// +kubebuilder:validation:Optional
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// Whether the rule should be enabled. Default 'true'
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether logging should be enabled for this rule. Default 'false'
	// +kubebuilder:validation:Optional
	LoggingEnabled *bool `json:"loggingEnabled,omitempty" tf:"logging_enabled,omitempty"`

	// Org or external network name
	// +kubebuilder:validation:Optional
	NetworkName *string `json:"networkName,omitempty" tf:"network_name,omitempty"`

	// Network type. One of 'ext', 'org'
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Original address or address range. This is the the source address for SNAT rules
	// +kubebuilder:validation:Optional
	OriginalAddress *string `json:"originalAddress,omitempty" tf:"original_address,omitempty"`

	// Optional. Allows to set custom rule tag
	// +kubebuilder:validation:Optional
	RuleTag *float64 `json:"ruleTag,omitempty" tf:"rule_tag,omitempty"`

	// Read only. Possible values 'user', 'internal_high'
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`

	// Translated address or address range
	// +kubebuilder:validation:Optional
	TranslatedAddress *string `json:"translatedAddress,omitempty" tf:"translated_address,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// NsxvSnatSpec defines the desired state of NsxvSnat
type NsxvSnatSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxvSnatParameters `json:"forProvider"`
}

// NsxvSnatStatus defines the observed state of NsxvSnat.
type NsxvSnatStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxvSnatObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxvSnat is the Schema for the NsxvSnats API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxvSnat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.edgeGateway)",message="edgeGateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.networkName)",message="networkName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.networkType)",message="networkType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.originalAddress)",message="originalAddress is a required parameter"
	Spec   NsxvSnatSpec   `json:"spec"`
	Status NsxvSnatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxvSnatList contains a list of NsxvSnats
type NsxvSnatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxvSnat `json:"items"`
}

// Repository type metadata.
var (
	NsxvSnat_Kind             = "NsxvSnat"
	NsxvSnat_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxvSnat_Kind}.String()
	NsxvSnat_KindAPIVersion   = NsxvSnat_Kind + "." + CRDGroupVersion.String()
	NsxvSnat_GroupVersionKind = CRDGroupVersion.WithKind(NsxvSnat_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxvSnat{}, &NsxvSnatList{})
}