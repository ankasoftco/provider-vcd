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

type NextHopObservation struct {

	// Admin distance of next hop
	AdminDistance *float64 `json:"adminDistance,omitempty" tf:"admin_distance,omitempty"`

	// IP Address of next hop
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	Scope []ScopeObservation `json:"scope,omitempty" tf:"scope,omitempty"`
}

type NextHopParameters struct {

	// Admin distance of next hop
	// +kubebuilder:validation:Required
	AdminDistance *float64 `json:"adminDistance" tf:"admin_distance,omitempty"`

	// IP Address of next hop
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`
}

type NsxtEdgeGatewayStaticRouteObservation struct {

	// Description of Static Route
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Edge gateway ID for Static Route configuration
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of Static Route
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network CIDR (e.g. 192.168.1.1/24) for Static Route
	NetworkCidr *string `json:"networkCidr,omitempty" tf:"network_cidr,omitempty"`

	// A set of next hops to use within the static route
	NextHop []NextHopObservation `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`
}

type NsxtEdgeGatewayStaticRouteParameters struct {

	// Description of Static Route
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Edge gateway ID for Static Route configuration
	// +kubebuilder:validation:Optional
	EdgeGatewayID *string `json:"edgeGatewayId,omitempty" tf:"edge_gateway_id,omitempty"`

	// Name of Static Route
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network CIDR (e.g. 192.168.1.1/24) for Static Route
	// +kubebuilder:validation:Optional
	NetworkCidr *string `json:"networkCidr,omitempty" tf:"network_cidr,omitempty"`

	// A set of next hops to use within the static route
	// +kubebuilder:validation:Optional
	NextHop []NextHopParameters `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`
}

type ScopeObservation struct {

	// ID of Scope element
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of Scope element
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Scope type - One of 'NETWORK', 'SYSTEM_OWNED'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScopeParameters struct {

	// ID of Scope element
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Scope type - One of 'NETWORK', 'SYSTEM_OWNED'
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// NsxtEdgeGatewayStaticRouteSpec defines the desired state of NsxtEdgeGatewayStaticRoute
type NsxtEdgeGatewayStaticRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsxtEdgeGatewayStaticRouteParameters `json:"forProvider"`
}

// NsxtEdgeGatewayStaticRouteStatus defines the observed state of NsxtEdgeGatewayStaticRoute.
type NsxtEdgeGatewayStaticRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsxtEdgeGatewayStaticRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtEdgeGatewayStaticRoute is the Schema for the NsxtEdgeGatewayStaticRoutes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type NsxtEdgeGatewayStaticRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.edgeGatewayId)",message="edgeGatewayId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.networkCidr)",message="networkCidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.nextHop)",message="nextHop is a required parameter"
	Spec   NsxtEdgeGatewayStaticRouteSpec   `json:"spec"`
	Status NsxtEdgeGatewayStaticRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsxtEdgeGatewayStaticRouteList contains a list of NsxtEdgeGatewayStaticRoutes
type NsxtEdgeGatewayStaticRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsxtEdgeGatewayStaticRoute `json:"items"`
}

// Repository type metadata.
var (
	NsxtEdgeGatewayStaticRoute_Kind             = "NsxtEdgeGatewayStaticRoute"
	NsxtEdgeGatewayStaticRoute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsxtEdgeGatewayStaticRoute_Kind}.String()
	NsxtEdgeGatewayStaticRoute_KindAPIVersion   = NsxtEdgeGatewayStaticRoute_Kind + "." + CRDGroupVersion.String()
	NsxtEdgeGatewayStaticRoute_GroupVersionKind = CRDGroupVersion.WithKind(NsxtEdgeGatewayStaticRoute_Kind)
)

func init() {
	SchemeBuilder.Register(&NsxtEdgeGatewayStaticRoute{}, &NsxtEdgeGatewayStaticRouteList{})
}
