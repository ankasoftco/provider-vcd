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

type LbServiceMonitorObservation struct {

	// Edge gateway name in which the LB Service Monitor is located
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// String that the monitor expects to match in the status line of the http or https response (for example, HTTP/1.1)
	Expected *string `json:"expected,omitempty" tf:"expected,omitempty"`

	// Advanced monitor parameters as key=value pairs
	Extension map[string]*string `json:"extension,omitempty" tf:"extension,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Interval in seconds at which a server is to be monitored (defaults to 10)
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Number of times the specified monitoring Method must fail sequentially before the server is declared down  (defaults to 3)
	MaxRetries *float64 `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`

	// Method to be used to detect server status. One of OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Unique LB Service Monitor name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// String to be matched in the response content
	Receive *string `json:"receive,omitempty" tf:"receive,omitempty"`

	// Data to be sent
	Send *string `json:"send,omitempty" tf:"send,omitempty"`

	// Maximum time in seconds within which a response from the server must be received  (defaults to 15)
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Way in which you want to send the health check request to the server. One of http, https, tcp, icmp, or udp
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// URL to be used in the server status request
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type LbServiceMonitorParameters struct {

	// Edge gateway name in which the LB Service Monitor is located
	// +kubebuilder:validation:Optional
	EdgeGateway *string `json:"edgeGateway,omitempty" tf:"edge_gateway,omitempty"`

	// String that the monitor expects to match in the status line of the http or https response (for example, HTTP/1.1)
	// +kubebuilder:validation:Optional
	Expected *string `json:"expected,omitempty" tf:"expected,omitempty"`

	// Advanced monitor parameters as key=value pairs
	// +kubebuilder:validation:Optional
	Extension map[string]*string `json:"extension,omitempty" tf:"extension,omitempty"`

	// Interval in seconds at which a server is to be monitored (defaults to 10)
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// Number of times the specified monitoring Method must fail sequentially before the server is declared down  (defaults to 3)
	// +kubebuilder:validation:Optional
	MaxRetries *float64 `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`

	// Method to be used to detect server status. One of OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Unique LB Service Monitor name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// String to be matched in the response content
	// +kubebuilder:validation:Optional
	Receive *string `json:"receive,omitempty" tf:"receive,omitempty"`

	// Data to be sent
	// +kubebuilder:validation:Optional
	Send *string `json:"send,omitempty" tf:"send,omitempty"`

	// Maximum time in seconds within which a response from the server must be received  (defaults to 15)
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Way in which you want to send the health check request to the server. One of http, https, tcp, icmp, or udp
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// URL to be used in the server status request
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// LbServiceMonitorSpec defines the desired state of LbServiceMonitor
type LbServiceMonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LbServiceMonitorParameters `json:"forProvider"`
}

// LbServiceMonitorStatus defines the observed state of LbServiceMonitor.
type LbServiceMonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LbServiceMonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbServiceMonitor is the Schema for the LbServiceMonitors API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type LbServiceMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.edgeGateway)",message="edgeGateway is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.type)",message="type is a required parameter"
	Spec   LbServiceMonitorSpec   `json:"spec"`
	Status LbServiceMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbServiceMonitorList contains a list of LbServiceMonitors
type LbServiceMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbServiceMonitor `json:"items"`
}

// Repository type metadata.
var (
	LbServiceMonitor_Kind             = "LbServiceMonitor"
	LbServiceMonitor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LbServiceMonitor_Kind}.String()
	LbServiceMonitor_KindAPIVersion   = LbServiceMonitor_Kind + "." + CRDGroupVersion.String()
	LbServiceMonitor_GroupVersionKind = CRDGroupVersion.WithKind(LbServiceMonitor_Kind)
)

func init() {
	SchemeBuilder.Register(&LbServiceMonitor{}, &LbServiceMonitorList{})
}
