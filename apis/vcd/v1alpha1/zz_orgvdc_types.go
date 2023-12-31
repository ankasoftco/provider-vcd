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

type CPUObservation struct {

	// Capacity that is committed to be available. Value in MB or MHz. Used with AllocationPool (Allocation pool) and ReservationPool (Reservation pool).
	Allocated *float64 `json:"allocated,omitempty" tf:"allocated,omitempty"`

	// Capacity limit relative to the value specified for Allocation. It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp (Pay as you go).
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	Reserved *float64 `json:"reserved,omitempty" tf:"reserved,omitempty"`

	Used *float64 `json:"used,omitempty" tf:"used,omitempty"`
}

type CPUParameters struct {

	// Capacity that is committed to be available. Value in MB or MHz. Used with AllocationPool (Allocation pool) and ReservationPool (Reservation pool).
	// +kubebuilder:validation:Optional
	Allocated *float64 `json:"allocated,omitempty" tf:"allocated,omitempty"`

	// Capacity limit relative to the value specified for Allocation. It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp (Pay as you go).
	// +kubebuilder:validation:Optional
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`
}

type ComputeCapacityObservation struct {
	CPU []CPUObservation `json:"cpu,omitempty" tf:"cpu,omitempty"`

	Memory []MemoryObservation `json:"memory,omitempty" tf:"memory,omitempty"`
}

type ComputeCapacityParameters struct {

	// +kubebuilder:validation:Required
	CPU []CPUParameters `json:"cpu" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Required
	Memory []MemoryParameters `json:"memory" tf:"memory,omitempty"`
}

type MemoryObservation struct {

	// Capacity that is committed to be available. Value in MB or MHz. Used with AllocationPool (Allocation pool) and ReservationPool (Reservation pool).
	Allocated *float64 `json:"allocated,omitempty" tf:"allocated,omitempty"`

	// Capacity limit relative to the value specified for Allocation. It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp (Pay as you go).
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	Reserved *float64 `json:"reserved,omitempty" tf:"reserved,omitempty"`

	Used *float64 `json:"used,omitempty" tf:"used,omitempty"`
}

type MemoryParameters struct {

	// Capacity that is committed to be available. Value in MB or MHz. Used with AllocationPool (Allocation pool) and ReservationPool (Reservation pool).
	// +kubebuilder:validation:Optional
	Allocated *float64 `json:"allocated,omitempty" tf:"allocated,omitempty"`

	// Capacity limit relative to the value specified for Allocation. It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp (Pay as you go).
	// +kubebuilder:validation:Optional
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`
}

type OrgVdcMetadataEntryObservation struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry. Required if the metadata entry is not empty
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry. Required if the metadata entry is not empty
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OrgVdcMetadataEntryParameters struct {

	// Domain for this metadata entry. true if it belongs to SYSTEM, false if it belongs to GENERAL
	// +kubebuilder:validation:Optional
	IsSystem *bool `json:"isSystem,omitempty" tf:"is_system,omitempty"`

	// Key of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
	// +kubebuilder:validation:Optional
	UserAccess *string `json:"userAccess,omitempty" tf:"user_access,omitempty"`

	// Value of this metadata entry. Required if the metadata entry is not empty
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OrgVdcObservation struct {

	// The allocation model used by this VDC; must be one of {AllocationVApp, AllocationPool, ReservationPool, Flex}
	AllocationModel *string `json:"allocationModel,omitempty" tf:"allocation_model,omitempty"`

	// Set to false to disallow creation of the VDC if the AllocationModel is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
	AllowOverCommit *bool `json:"allowOverCommit,omitempty" tf:"allow_over_commit,omitempty"`

	// Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If the element is empty, vCD sets a value
	CPUGuaranteed *float64 `json:"cpuGuaranteed,omitempty" tf:"cpu_guaranteed,omitempty"`

	// Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when AllocationModel is AllocationVApp or AllocationPool, and may not be less than 256 MHz. Defaults to 1000 MHz if the element is empty or missing.
	CPUSpeed *float64 `json:"cpuSpeed,omitempty" tf:"cpu_speed,omitempty"`

	// The compute capacity allocated to this VDC.
	ComputeCapacity []ComputeCapacityObservation `json:"computeCapacity,omitempty" tf:"compute_capacity,omitempty"`

	// ID of default Compute policy for this VDC, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
	DefaultComputePolicyID *string `json:"defaultComputePolicyId,omitempty" tf:"default_compute_policy_id,omitempty"`

	// ID of default VM Compute policy, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
	DefaultVMSizingPolicyID *string `json:"defaultVmSizingPolicyId,omitempty" tf:"default_vm_sizing_policy_id,omitempty"`

	// When destroying use delete_force=True to remove a VDC and any objects it contains, regardless of their state.
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// When destroying use delete_recursive=True to remove the VDC and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of NSX-T Edge Cluster (provider vApp networking services and DHCP capability for Isolated networks)
	EdgeClusterID *string `json:"edgeClusterId,omitempty" tf:"edge_cluster_id,omitempty"`

	// Set to true to indicate if the Flex VDC is to be elastic.
	Elasticity *bool `json:"elasticity,omitempty" tf:"elasticity,omitempty"`

	// Request for fast provisioning. Request will be honored only if the underlying datas tore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.
	EnableFastProvisioning *bool `json:"enableFastProvisioning,omitempty" tf:"enable_fast_provisioning,omitempty"`

	// Set to true to enable distributed firewall - Only applies to NSX-V VDCs
	EnableNsxvDistributedFirewall *bool `json:"enableNsxvDistributedFirewall,omitempty" tf:"enable_nsxv_distributed_firewall,omitempty"`

	// Boolean to request thin provisioning. Request will be honored only if the underlying datastore supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
	EnableThinProvisioning *bool `json:"enableThinProvisioning,omitempty" tf:"enable_thin_provisioning,omitempty"`

	// True if discovery of vCenter VMs is enabled for resource pools backing this VDC. If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.
	EnableVMDiscovery *bool `json:"enableVmDiscovery,omitempty" tf:"enable_vm_discovery,omitempty"`

	// True if this VDC is enabled for use by the organization VDCs. Default is true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set to true to indicate if the Flex VDC is to include memory overhead into its accounting for admission control.
	IncludeVMMemoryOverhead *bool `json:"includeVmMemoryOverhead,omitempty" tf:"include_vm_memory_overhead,omitempty"`

	// Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. When Allocation model is AllocationPool minimum value is 0.2. If the element is empty, vCD sets a value.
	MemoryGuaranteed *float64 `json:"memoryGuaranteed,omitempty" tf:"memory_guaranteed,omitempty"`

	// Key and value pairs for Org VDC metadata
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given VDC
	MetadataEntry []OrgVdcMetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
	NetworkPoolName *string `json:"networkPoolName,omitempty" tf:"network_pool_name,omitempty"`

	// Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be deployed.
	NetworkQuota *float64 `json:"networkQuota,omitempty" tf:"network_quota,omitempty"`

	// Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
	NicQuota *float64 `json:"nicQuota,omitempty" tf:"nic_quota,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A reference to the Provider VDC from which this organization VDC is provisioned.
	ProviderVdcName *string `json:"providerVdcName,omitempty" tf:"provider_vdc_name,omitempty"`

	// Storage profiles supported by this VDC.
	StorageProfile []StorageProfileObservation `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`

	// Set of VM Placement Policy IDs
	VMPlacementPolicyIds []*string `json:"vmPlacementPolicyIds,omitempty" tf:"vm_placement_policy_ids,omitempty"`

	// The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.
	VMQuota *float64 `json:"vmQuota,omitempty" tf:"vm_quota,omitempty"`

	// Set of VM Sizing Policy IDs
	VMSizingPolicyIds []*string `json:"vmSizingPolicyIds,omitempty" tf:"vm_sizing_policy_ids,omitempty"`
}

type OrgVdcParameters struct {

	// The allocation model used by this VDC; must be one of {AllocationVApp, AllocationPool, ReservationPool, Flex}
	// +kubebuilder:validation:Optional
	AllocationModel *string `json:"allocationModel,omitempty" tf:"allocation_model,omitempty"`

	// Set to false to disallow creation of the VDC if the AllocationModel is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
	// +kubebuilder:validation:Optional
	AllowOverCommit *bool `json:"allowOverCommit,omitempty" tf:"allow_over_commit,omitempty"`

	// Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. If the element is empty, vCD sets a value
	// +kubebuilder:validation:Optional
	CPUGuaranteed *float64 `json:"cpuGuaranteed,omitempty" tf:"cpu_guaranteed,omitempty"`

	// Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when AllocationModel is AllocationVApp or AllocationPool, and may not be less than 256 MHz. Defaults to 1000 MHz if the element is empty or missing.
	// +kubebuilder:validation:Optional
	CPUSpeed *float64 `json:"cpuSpeed,omitempty" tf:"cpu_speed,omitempty"`

	// The compute capacity allocated to this VDC.
	// +kubebuilder:validation:Optional
	ComputeCapacity []ComputeCapacityParameters `json:"computeCapacity,omitempty" tf:"compute_capacity,omitempty"`

	// ID of default Compute policy for this VDC, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
	// +kubebuilder:validation:Optional
	DefaultComputePolicyID *string `json:"defaultComputePolicyId,omitempty" tf:"default_compute_policy_id,omitempty"`

	// ID of default VM Compute policy, which can be a VM Sizing Policy, VM Placement Policy or vGPU Policy
	// +kubebuilder:validation:Optional
	DefaultVMSizingPolicyID *string `json:"defaultVmSizingPolicyId,omitempty" tf:"default_vm_sizing_policy_id,omitempty"`

	// When destroying use delete_force=True to remove a VDC and any objects it contains, regardless of their state.
	// +kubebuilder:validation:Optional
	DeleteForce *bool `json:"deleteForce,omitempty" tf:"delete_force,omitempty"`

	// When destroying use delete_recursive=True to remove the VDC and any objects it contains that are in a state that normally allows removal.
	// +kubebuilder:validation:Optional
	DeleteRecursive *bool `json:"deleteRecursive,omitempty" tf:"delete_recursive,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of NSX-T Edge Cluster (provider vApp networking services and DHCP capability for Isolated networks)
	// +kubebuilder:validation:Optional
	EdgeClusterID *string `json:"edgeClusterId,omitempty" tf:"edge_cluster_id,omitempty"`

	// Set to true to indicate if the Flex VDC is to be elastic.
	// +kubebuilder:validation:Optional
	Elasticity *bool `json:"elasticity,omitempty" tf:"elasticity,omitempty"`

	// Request for fast provisioning. Request will be honored only if the underlying datas tore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.
	// +kubebuilder:validation:Optional
	EnableFastProvisioning *bool `json:"enableFastProvisioning,omitempty" tf:"enable_fast_provisioning,omitempty"`

	// Set to true to enable distributed firewall - Only applies to NSX-V VDCs
	// +kubebuilder:validation:Optional
	EnableNsxvDistributedFirewall *bool `json:"enableNsxvDistributedFirewall,omitempty" tf:"enable_nsxv_distributed_firewall,omitempty"`

	// Boolean to request thin provisioning. Request will be honored only if the underlying datastore supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
	// +kubebuilder:validation:Optional
	EnableThinProvisioning *bool `json:"enableThinProvisioning,omitempty" tf:"enable_thin_provisioning,omitempty"`

	// True if discovery of vCenter VMs is enabled for resource pools backing this VDC. If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.
	// +kubebuilder:validation:Optional
	EnableVMDiscovery *bool `json:"enableVmDiscovery,omitempty" tf:"enable_vm_discovery,omitempty"`

	// True if this VDC is enabled for use by the organization VDCs. Default is true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Set to true to indicate if the Flex VDC is to include memory overhead into its accounting for admission control.
	// +kubebuilder:validation:Optional
	IncludeVMMemoryOverhead *bool `json:"includeVmMemoryOverhead,omitempty" tf:"include_vm_memory_overhead,omitempty"`

	// Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when AllocationModel is AllocationVApp or AllocationPool. When Allocation model is AllocationPool minimum value is 0.2. If the element is empty, vCD sets a value.
	// +kubebuilder:validation:Optional
	MemoryGuaranteed *float64 `json:"memoryGuaranteed,omitempty" tf:"memory_guaranteed,omitempty"`

	// Key and value pairs for Org VDC metadata
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given VDC
	// +kubebuilder:validation:Optional
	MetadataEntry []OrgVdcMetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
	// +kubebuilder:validation:Optional
	NetworkPoolName *string `json:"networkPoolName,omitempty" tf:"network_pool_name,omitempty"`

	// Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be deployed.
	// +kubebuilder:validation:Optional
	NetworkQuota *float64 `json:"networkQuota,omitempty" tf:"network_quota,omitempty"`

	// Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
	// +kubebuilder:validation:Optional
	NicQuota *float64 `json:"nicQuota,omitempty" tf:"nic_quota,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// A reference to the Provider VDC from which this organization VDC is provisioned.
	// +kubebuilder:validation:Optional
	ProviderVdcName *string `json:"providerVdcName,omitempty" tf:"provider_vdc_name,omitempty"`

	// Storage profiles supported by this VDC.
	// +kubebuilder:validation:Optional
	StorageProfile []StorageProfileParameters `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`

	// Set of VM Placement Policy IDs
	// +kubebuilder:validation:Optional
	VMPlacementPolicyIds []*string `json:"vmPlacementPolicyIds,omitempty" tf:"vm_placement_policy_ids,omitempty"`

	// The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.
	// +kubebuilder:validation:Optional
	VMQuota *float64 `json:"vmQuota,omitempty" tf:"vm_quota,omitempty"`

	// Set of VM Sizing Policy IDs
	// +kubebuilder:validation:Optional
	VMSizingPolicyIds []*string `json:"vmSizingPolicyIds,omitempty" tf:"vm_sizing_policy_ids,omitempty"`
}

type StorageProfileObservation struct {

	// True if this is default storage profile for this VDC. The default storage profile is used when an object that can specify a storage profile is created with no storage profile specified.
	Default *bool `json:"default,omitempty" tf:"default,omitempty"`

	// True if this storage profile is enabled for use in the VDC.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Maximum number of MB allocated for this storage profile. A value of 0 specifies unlimited MB.
	Limit *float64 `json:"limit,omitempty" tf:"limit,omitempty"`

	// Name of Provider VDC storage profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Storage used in MB
	StorageUsedInMb *float64 `json:"storageUsedInMb,omitempty" tf:"storage_used_in_mb,omitempty"`
}

type StorageProfileParameters struct {

	// True if this is default storage profile for this VDC. The default storage profile is used when an object that can specify a storage profile is created with no storage profile specified.
	// +kubebuilder:validation:Required
	Default *bool `json:"default" tf:"default,omitempty"`

	// True if this storage profile is enabled for use in the VDC.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Maximum number of MB allocated for this storage profile. A value of 0 specifies unlimited MB.
	// +kubebuilder:validation:Required
	Limit *float64 `json:"limit" tf:"limit,omitempty"`

	// Name of Provider VDC storage profile.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// OrgVdcSpec defines the desired state of OrgVdc
type OrgVdcSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrgVdcParameters `json:"forProvider"`
}

// OrgVdcStatus defines the observed state of OrgVdc.
type OrgVdcStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrgVdcObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrgVdc is the Schema for the OrgVdcs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type OrgVdc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.allocationModel)",message="allocationModel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.computeCapacity)",message="computeCapacity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.deleteForce)",message="deleteForce is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.deleteRecursive)",message="deleteRecursive is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.providerVdcName)",message="providerVdcName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.storageProfile)",message="storageProfile is a required parameter"
	Spec   OrgVdcSpec   `json:"spec"`
	Status OrgVdcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrgVdcList contains a list of OrgVdcs
type OrgVdcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrgVdc `json:"items"`
}

// Repository type metadata.
var (
	OrgVdc_Kind             = "OrgVdc"
	OrgVdc_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrgVdc_Kind}.String()
	OrgVdc_KindAPIVersion   = OrgVdc_Kind + "." + CRDGroupVersion.String()
	OrgVdc_GroupVersionKind = CRDGroupVersion.WithKind(OrgVdc_Kind)
)

func init() {
	SchemeBuilder.Register(&OrgVdc{}, &OrgVdcList{})
}
