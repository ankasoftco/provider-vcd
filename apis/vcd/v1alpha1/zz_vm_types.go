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

type VmCustomizationObservation struct {

	// Allow local administrator password
	AllowLocalAdminPassword *bool `json:"allowLocalAdminPassword,omitempty" tf:"allow_local_admin_password,omitempty"`

	// Auto generate password
	AutoGeneratePassword *bool `json:"autoGeneratePassword,omitempty" tf:"auto_generate_password,omitempty"`

	// 'true' value will change SID. Applicable only for Windows VMs
	ChangeSid *bool `json:"changeSid,omitempty" tf:"change_sid,omitempty"`

	// 'true' value will enable guest customization. It may occur on first boot or when 'force' is used
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// 'true' value will cause the VM to reboot on every 'apply' operation
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// Script to run on initial boot or with customization.force=true set
	Initscript *string `json:"initscript,omitempty" tf:"initscript,omitempty"`

	// Enable this VM to join a domain
	JoinDomain *bool `json:"joinDomain,omitempty" tf:"join_domain,omitempty"`

	// Account organizational unit for domain name join
	JoinDomainAccountOu *string `json:"joinDomainAccountOu,omitempty" tf:"join_domain_account_ou,omitempty"`

	// Custom domain name for join
	JoinDomainName *string `json:"joinDomainName,omitempty" tf:"join_domain_name,omitempty"`

	// Username for custom domain name join
	JoinDomainUser *string `json:"joinDomainUser,omitempty" tf:"join_domain_user,omitempty"`

	// Use organization's domain for joining
	JoinOrgDomain *bool `json:"joinOrgDomain,omitempty" tf:"join_org_domain,omitempty"`

	// Require Administrator to change password on first login
	MustChangePasswordOnFirstLogin *bool `json:"mustChangePasswordOnFirstLogin,omitempty" tf:"must_change_password_on_first_login,omitempty"`

	// Number of times to log on automatically. '0' - disabled.
	NumberOfAutoLogons *float64 `json:"numberOfAutoLogons,omitempty" tf:"number_of_auto_logons,omitempty"`
}

type VmCustomizationParameters struct {

	// Manually specify admin password
	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// Allow local administrator password
	// +kubebuilder:validation:Optional
	AllowLocalAdminPassword *bool `json:"allowLocalAdminPassword,omitempty" tf:"allow_local_admin_password,omitempty"`

	// Auto generate password
	// +kubebuilder:validation:Optional
	AutoGeneratePassword *bool `json:"autoGeneratePassword,omitempty" tf:"auto_generate_password,omitempty"`

	// 'true' value will change SID. Applicable only for Windows VMs
	// +kubebuilder:validation:Optional
	ChangeSid *bool `json:"changeSid,omitempty" tf:"change_sid,omitempty"`

	// 'true' value will enable guest customization. It may occur on first boot or when 'force' is used
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// 'true' value will cause the VM to reboot on every 'apply' operation
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// Script to run on initial boot or with customization.force=true set
	// +kubebuilder:validation:Optional
	Initscript *string `json:"initscript,omitempty" tf:"initscript,omitempty"`

	// Enable this VM to join a domain
	// +kubebuilder:validation:Optional
	JoinDomain *bool `json:"joinDomain,omitempty" tf:"join_domain,omitempty"`

	// Account organizational unit for domain name join
	// +kubebuilder:validation:Optional
	JoinDomainAccountOu *string `json:"joinDomainAccountOu,omitempty" tf:"join_domain_account_ou,omitempty"`

	// Custom domain name for join
	// +kubebuilder:validation:Optional
	JoinDomainName *string `json:"joinDomainName,omitempty" tf:"join_domain_name,omitempty"`

	// Password for custom domain name join
	// +kubebuilder:validation:Optional
	JoinDomainPasswordSecretRef *v1.SecretKeySelector `json:"joinDomainPasswordSecretRef,omitempty" tf:"-"`

	// Username for custom domain name join
	// +kubebuilder:validation:Optional
	JoinDomainUser *string `json:"joinDomainUser,omitempty" tf:"join_domain_user,omitempty"`

	// Use organization's domain for joining
	// +kubebuilder:validation:Optional
	JoinOrgDomain *bool `json:"joinOrgDomain,omitempty" tf:"join_org_domain,omitempty"`

	// Require Administrator to change password on first login
	// +kubebuilder:validation:Optional
	MustChangePasswordOnFirstLogin *bool `json:"mustChangePasswordOnFirstLogin,omitempty" tf:"must_change_password_on_first_login,omitempty"`

	// Number of times to log on automatically. '0' - disabled.
	// +kubebuilder:validation:Optional
	NumberOfAutoLogons *float64 `json:"numberOfAutoLogons,omitempty" tf:"number_of_auto_logons,omitempty"`
}

type VmDiskObservation struct {

	// Bus number on which to place the disk controller
	BusNumber *string `json:"busNumber,omitempty" tf:"bus_number,omitempty"`

	// Independent disk name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the disk in MB.
	SizeInMb *float64 `json:"sizeInMb,omitempty" tf:"size_in_mb,omitempty"`

	// Unit number (slot) on the bus specified by BusNumber
	UnitNumber *string `json:"unitNumber,omitempty" tf:"unit_number,omitempty"`
}

type VmDiskParameters struct {

	// Bus number on which to place the disk controller
	// +kubebuilder:validation:Required
	BusNumber *string `json:"busNumber" tf:"bus_number,omitempty"`

	// Independent disk name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Unit number (slot) on the bus specified by BusNumber
	// +kubebuilder:validation:Required
	UnitNumber *string `json:"unitNumber" tf:"unit_number,omitempty"`
}

type VmInternalDiskObservation struct {
	BusNumber *float64 `json:"busNumber,omitempty" tf:"bus_number,omitempty"`

	BusType *string `json:"busType,omitempty" tf:"bus_type,omitempty"`

	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	SizeInMb *float64 `json:"sizeInMb,omitempty" tf:"size_in_mb,omitempty"`

	StorageProfile *string `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`

	ThinProvisioned *bool `json:"thinProvisioned,omitempty" tf:"thin_provisioned,omitempty"`

	UnitNumber *float64 `json:"unitNumber,omitempty" tf:"unit_number,omitempty"`
}

type VmInternalDiskParameters struct {
}

type VmMetadataEntryObservation struct {

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

type VmMetadataEntryParameters struct {

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

type VmNetworkObservation struct {

	// Network card adapter type. (e.g. 'E1000', 'E1000E', 'SRIOVETHERNETCARD', 'VMXNET3', 'PCNet32')
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	// It defines if NIC is connected or not.
	Connected *bool `json:"connected,omitempty" tf:"connected,omitempty"`

	// IP of the VM. Settings depend on `ip_allocation_mode`. Omitted or empty for DHCP, POOL, NONE. Required for MANUAL
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IP address allocation mode. One of POOL, DHCP, MANUAL, NONE
	IPAllocationMode *string `json:"ipAllocationMode,omitempty" tf:"ip_allocation_mode,omitempty"`

	// Set to true if network interface should be primary. First network card in the list will be primary by default
	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	// Mac address of network interface
	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// Name of the network this VM should connect to. Always required except for `type` `NONE`
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network type to use: 'vapp', 'org' or 'none'. Use 'vapp' for vApp network, 'org' to attach Org VDC network. 'none' for empty NIC.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VmNetworkParameters struct {

	// Network card adapter type. (e.g. 'E1000', 'E1000E', 'SRIOVETHERNETCARD', 'VMXNET3', 'PCNet32')
	// +kubebuilder:validation:Optional
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	// It defines if NIC is connected or not.
	// +kubebuilder:validation:Optional
	Connected *bool `json:"connected,omitempty" tf:"connected,omitempty"`

	// IP of the VM. Settings depend on `ip_allocation_mode`. Omitted or empty for DHCP, POOL, NONE. Required for MANUAL
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// IP address allocation mode. One of POOL, DHCP, MANUAL, NONE
	// +kubebuilder:validation:Optional
	IPAllocationMode *string `json:"ipAllocationMode,omitempty" tf:"ip_allocation_mode,omitempty"`

	// Set to true if network interface should be primary. First network card in the list will be primary by default
	// +kubebuilder:validation:Optional
	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`

	// Mac address of network interface
	// +kubebuilder:validation:Optional
	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// Name of the network this VM should connect to. Always required except for `type` `NONE`
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Network type to use: 'vapp', 'org' or 'none'. Use 'vapp' for vApp network, 'org' to attach Org VDC network. 'none' for empty NIC.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VmObservation struct {

	// Automatically accept EULA if OVA has it
	AcceptAllEulas *bool `json:"acceptAllEulas,omitempty" tf:"accept_all_eulas,omitempty"`

	// Media name to add as boot image.
	BootImage *string `json:"bootImage,omitempty" tf:"boot_image,omitempty"`

	// The URN of the media to use as boot image.
	BootImageID *string `json:"bootImageId,omitempty" tf:"boot_image_id,omitempty"`

	// The number of cores per socket
	CPUCores *float64 `json:"cpuCores,omitempty" tf:"cpu_cores,omitempty"`

	// True if the virtual machine supports addition of virtual CPUs while powered on.
	CPUHotAddEnabled *bool `json:"cpuHotAddEnabled,omitempty" tf:"cpu_hot_add_enabled,omitempty"`

	// The limit for how much of CPU can be consumed on the underlying virtualization infrastructure. This is only valid when the resource allocation is not unlimited.
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload
	CPUPriority *string `json:"cpuPriority,omitempty" tf:"cpu_priority,omitempty"`

	// The amount of MHz reservation on the underlying virtualization infrastructure
	CPUReservation *float64 `json:"cpuReservation,omitempty" tf:"cpu_reservation,omitempty"`

	// Custom priority for the resource. This is a read-only, unless the `cpu_priority` is CUSTOM
	CPUShares *float64 `json:"cpuShares,omitempty" tf:"cpu_shares,omitempty"`

	// The catalog name in which to find the given vApp Template or media for boot_image
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Computer name to assign to this virtual machine
	ComputerName *string `json:"computerName,omitempty" tf:"computer_name,omitempty"`

	// The number of virtual CPUs to allocate to the VM
	Cpus *float64 `json:"cpus,omitempty" tf:"cpus,omitempty"`

	// Guest customization block
	Customization []VmCustomizationObservation `json:"customization,omitempty" tf:"customization,omitempty"`

	// The VM description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Disk []VmDiskObservation `json:"disk,omitempty" tf:"disk,omitempty"`

	// Expose hardware-assisted CPU virtualization to guest OS.
	ExposeHardwareVirtualization *bool `json:"exposeHardwareVirtualization,omitempty" tf:"expose_hardware_virtualization,omitempty"`

	// Key/value settings for guest properties
	GuestProperties map[string]*string `json:"guestProperties,omitempty" tf:"guest_properties,omitempty"`

	// Virtual Hardware Version (e.g.`vmx-14`, `vmx-13`, `vmx-12`, etc.)
	HardwareVersion *string `json:"hardwareVersion,omitempty" tf:"hardware_version,omitempty"`

	// VM Hyper Reference
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A block will show internal disk details
	InternalDisk []VmInternalDiskObservation `json:"internalDisk,omitempty" tf:"internal_disk,omitempty"`

	// The amount of RAM (in MB) to allocate to the VM
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// True if the virtual machine supports addition of memory while powered on.
	MemoryHotAddEnabled *bool `json:"memoryHotAddEnabled,omitempty" tf:"memory_hot_add_enabled,omitempty"`

	// The limit for how much of memory can be consumed on the underlying virtualization infrastructure. This is only valid when the resource allocation is not unlimited.
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload
	MemoryPriority *string `json:"memoryPriority,omitempty" tf:"memory_priority,omitempty"`

	// The amount of RAM (in MB) reservation on the underlying virtualization infrastructure
	MemoryReservation *float64 `json:"memoryReservation,omitempty" tf:"memory_reservation,omitempty"`

	// Custom priority for the resource. This is a read-only, unless the `memory_priority` is CUSTOM
	MemoryShares *float64 `json:"memoryShares,omitempty" tf:"memory_shares,omitempty"`

	// Key value map of metadata to assign to this VM
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given VM
	MetadataEntry []VmMetadataEntryObservation `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A name for the VM, unique within the vApp
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block to define network interface. Multiple can be used.
	Network []VmNetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	// Optional number of seconds to try and wait for DHCP IP (valid for 'network' block only)
	NetworkDHCPWaitSeconds *float64 `json:"networkDhcpWaitSeconds,omitempty" tf:"network_dhcp_wait_seconds,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Operating System type. Possible values can be found in documentation.
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// A block to match internal_disk interface in template. Multiple can be used. Disk will be matched by bus_type, bus_number and unit_number.
	OverrideTemplateDisk []VmOverrideTemplateDiskObservation `json:"overrideTemplateDisk,omitempty" tf:"override_template_disk,omitempty"`

	// VM placement policy ID. Has to be assigned to Org VDC.
	PlacementPolicyID *string `json:"placementPolicyId,omitempty" tf:"placement_policy_id,omitempty"`

	// A boolean value stating if this VM should be powered on
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// True if the update of resource should fail when virtual machine power off needed.
	PreventUpdatePowerOff *bool `json:"preventUpdatePowerOff,omitempty" tf:"prevent_update_power_off,omitempty"`

	// Security tags to assign to this VM
	SecurityTags []*string `json:"securityTags,omitempty" tf:"security_tags,omitempty"`

	// VM sizing policy ID. Has to be assigned to Org VDC.
	SizingPolicyID *string `json:"sizingPolicyId,omitempty" tf:"sizing_policy_id,omitempty"`

	// Shows the status code of the VM
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// Shows the status of the VM
	StatusText *string `json:"statusText,omitempty" tf:"status_text,omitempty"`

	// Storage profile to override the default one
	StorageProfile *string `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`

	// The name of the vApp Template to use
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// The name of the VM in vApp Template to use. In cases when vApp template has more than one VM
	VMNameInTemplate *string `json:"vmNameInTemplate,omitempty" tf:"vm_name_in_template,omitempty"`

	// Type of VM: either 'vcd_vapp_vm' or 'vcd_vm'
	VMType *string `json:"vmType,omitempty" tf:"vm_type,omitempty"`

	// The vApp this VM belongs to - Required, unless it is a standalone VM
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The URN of the vApp Template to use
	VappTemplateID *string `json:"vappTemplateId,omitempty" tf:"vapp_template_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

type VmOverrideTemplateDiskObservation struct {

	// The number of the SCSI or IDE controller itself.
	BusNumber *float64 `json:"busNumber,omitempty" tf:"bus_number,omitempty"`

	// The type of disk controller. Possible values: ide, parallel( LSI Logic Parallel SCSI), sas(LSI Logic SAS (SCSI)), paravirtual(Paravirtual (SCSI)), sata, nvme
	BusType *string `json:"busType,omitempty" tf:"bus_type,omitempty"`

	// Specifies the IOPS for the disk. Default is 0.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The size of the disk in MB.
	SizeInMb *float64 `json:"sizeInMb,omitempty" tf:"size_in_mb,omitempty"`

	// Storage profile to override the VM default one
	StorageProfile *string `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`

	// The device number on the SCSI or IDE controller of the disk.
	UnitNumber *float64 `json:"unitNumber,omitempty" tf:"unit_number,omitempty"`
}

type VmOverrideTemplateDiskParameters struct {

	// The number of the SCSI or IDE controller itself.
	// +kubebuilder:validation:Required
	BusNumber *float64 `json:"busNumber" tf:"bus_number,omitempty"`

	// The type of disk controller. Possible values: ide, parallel( LSI Logic Parallel SCSI), sas(LSI Logic SAS (SCSI)), paravirtual(Paravirtual (SCSI)), sata, nvme
	// +kubebuilder:validation:Required
	BusType *string `json:"busType" tf:"bus_type,omitempty"`

	// Specifies the IOPS for the disk. Default is 0.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// The size of the disk in MB.
	// +kubebuilder:validation:Required
	SizeInMb *float64 `json:"sizeInMb" tf:"size_in_mb,omitempty"`

	// Storage profile to override the VM default one
	// +kubebuilder:validation:Optional
	StorageProfile *string `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`

	// The device number on the SCSI or IDE controller of the disk.
	// +kubebuilder:validation:Required
	UnitNumber *float64 `json:"unitNumber" tf:"unit_number,omitempty"`
}

type VmParameters struct {

	// Automatically accept EULA if OVA has it
	// +kubebuilder:validation:Optional
	AcceptAllEulas *bool `json:"acceptAllEulas,omitempty" tf:"accept_all_eulas,omitempty"`

	// Media name to add as boot image.
	// +kubebuilder:validation:Optional
	BootImage *string `json:"bootImage,omitempty" tf:"boot_image,omitempty"`

	// The URN of the media to use as boot image.
	// +kubebuilder:validation:Optional
	BootImageID *string `json:"bootImageId,omitempty" tf:"boot_image_id,omitempty"`

	// The number of cores per socket
	// +kubebuilder:validation:Optional
	CPUCores *float64 `json:"cpuCores,omitempty" tf:"cpu_cores,omitempty"`

	// True if the virtual machine supports addition of virtual CPUs while powered on.
	// +kubebuilder:validation:Optional
	CPUHotAddEnabled *bool `json:"cpuHotAddEnabled,omitempty" tf:"cpu_hot_add_enabled,omitempty"`

	// The limit for how much of CPU can be consumed on the underlying virtualization infrastructure. This is only valid when the resource allocation is not unlimited.
	// +kubebuilder:validation:Optional
	CPULimit *float64 `json:"cpuLimit,omitempty" tf:"cpu_limit,omitempty"`

	// Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload
	// +kubebuilder:validation:Optional
	CPUPriority *string `json:"cpuPriority,omitempty" tf:"cpu_priority,omitempty"`

	// The amount of MHz reservation on the underlying virtualization infrastructure
	// +kubebuilder:validation:Optional
	CPUReservation *float64 `json:"cpuReservation,omitempty" tf:"cpu_reservation,omitempty"`

	// Custom priority for the resource. This is a read-only, unless the `cpu_priority` is CUSTOM
	// +kubebuilder:validation:Optional
	CPUShares *float64 `json:"cpuShares,omitempty" tf:"cpu_shares,omitempty"`

	// The catalog name in which to find the given vApp Template or media for boot_image
	// +kubebuilder:validation:Optional
	CatalogName *string `json:"catalogName,omitempty" tf:"catalog_name,omitempty"`

	// Computer name to assign to this virtual machine
	// +kubebuilder:validation:Optional
	ComputerName *string `json:"computerName,omitempty" tf:"computer_name,omitempty"`

	// The number of virtual CPUs to allocate to the VM
	// +kubebuilder:validation:Optional
	Cpus *float64 `json:"cpus,omitempty" tf:"cpus,omitempty"`

	// Guest customization block
	// +kubebuilder:validation:Optional
	Customization []VmCustomizationParameters `json:"customization,omitempty" tf:"customization,omitempty"`

	// The VM description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Disk []VmDiskParameters `json:"disk,omitempty" tf:"disk,omitempty"`

	// Expose hardware-assisted CPU virtualization to guest OS.
	// +kubebuilder:validation:Optional
	ExposeHardwareVirtualization *bool `json:"exposeHardwareVirtualization,omitempty" tf:"expose_hardware_virtualization,omitempty"`

	// Key/value settings for guest properties
	// +kubebuilder:validation:Optional
	GuestProperties map[string]*string `json:"guestProperties,omitempty" tf:"guest_properties,omitempty"`

	// Virtual Hardware Version (e.g.`vmx-14`, `vmx-13`, `vmx-12`, etc.)
	// +kubebuilder:validation:Optional
	HardwareVersion *string `json:"hardwareVersion,omitempty" tf:"hardware_version,omitempty"`

	// VM Hyper Reference
	// +kubebuilder:validation:Optional
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	// The amount of RAM (in MB) to allocate to the VM
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// True if the virtual machine supports addition of memory while powered on.
	// +kubebuilder:validation:Optional
	MemoryHotAddEnabled *bool `json:"memoryHotAddEnabled,omitempty" tf:"memory_hot_add_enabled,omitempty"`

	// The limit for how much of memory can be consumed on the underlying virtualization infrastructure. This is only valid when the resource allocation is not unlimited.
	// +kubebuilder:validation:Optional
	MemoryLimit *float64 `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`

	// Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload
	// +kubebuilder:validation:Optional
	MemoryPriority *string `json:"memoryPriority,omitempty" tf:"memory_priority,omitempty"`

	// The amount of RAM (in MB) reservation on the underlying virtualization infrastructure
	// +kubebuilder:validation:Optional
	MemoryReservation *float64 `json:"memoryReservation,omitempty" tf:"memory_reservation,omitempty"`

	// Custom priority for the resource. This is a read-only, unless the `memory_priority` is CUSTOM
	// +kubebuilder:validation:Optional
	MemoryShares *float64 `json:"memoryShares,omitempty" tf:"memory_shares,omitempty"`

	// Key value map of metadata to assign to this VM
	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Metadata entries for the given VM
	// +kubebuilder:validation:Optional
	MetadataEntry []VmMetadataEntryParameters `json:"metadataEntry,omitempty" tf:"metadata_entry,omitempty"`

	// A name for the VM, unique within the vApp
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A block to define network interface. Multiple can be used.
	// +kubebuilder:validation:Optional
	Network []VmNetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// Optional number of seconds to try and wait for DHCP IP (valid for 'network' block only)
	// +kubebuilder:validation:Optional
	NetworkDHCPWaitSeconds *float64 `json:"networkDhcpWaitSeconds,omitempty" tf:"network_dhcp_wait_seconds,omitempty"`

	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	// +kubebuilder:validation:Optional
	Org *string `json:"org,omitempty" tf:"org,omitempty"`

	// Operating System type. Possible values can be found in documentation.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// A block to match internal_disk interface in template. Multiple can be used. Disk will be matched by bus_type, bus_number and unit_number.
	// +kubebuilder:validation:Optional
	OverrideTemplateDisk []VmOverrideTemplateDiskParameters `json:"overrideTemplateDisk,omitempty" tf:"override_template_disk,omitempty"`

	// VM placement policy ID. Has to be assigned to Org VDC.
	// +kubebuilder:validation:Optional
	PlacementPolicyID *string `json:"placementPolicyId,omitempty" tf:"placement_policy_id,omitempty"`

	// A boolean value stating if this VM should be powered on
	// +kubebuilder:validation:Optional
	PowerOn *bool `json:"powerOn,omitempty" tf:"power_on,omitempty"`

	// True if the update of resource should fail when virtual machine power off needed.
	// +kubebuilder:validation:Optional
	PreventUpdatePowerOff *bool `json:"preventUpdatePowerOff,omitempty" tf:"prevent_update_power_off,omitempty"`

	// Security tags to assign to this VM
	// +kubebuilder:validation:Optional
	SecurityTags []*string `json:"securityTags,omitempty" tf:"security_tags,omitempty"`

	// VM sizing policy ID. Has to be assigned to Org VDC.
	// +kubebuilder:validation:Optional
	SizingPolicyID *string `json:"sizingPolicyId,omitempty" tf:"sizing_policy_id,omitempty"`

	// Storage profile to override the default one
	// +kubebuilder:validation:Optional
	StorageProfile *string `json:"storageProfile,omitempty" tf:"storage_profile,omitempty"`

	// The name of the vApp Template to use
	// +kubebuilder:validation:Optional
	TemplateName *string `json:"templateName,omitempty" tf:"template_name,omitempty"`

	// The name of the VM in vApp Template to use. In cases when vApp template has more than one VM
	// +kubebuilder:validation:Optional
	VMNameInTemplate *string `json:"vmNameInTemplate,omitempty" tf:"vm_name_in_template,omitempty"`

	// The vApp this VM belongs to - Required, unless it is a standalone VM
	// +kubebuilder:validation:Optional
	VappName *string `json:"vappName,omitempty" tf:"vapp_name,omitempty"`

	// The URN of the vApp Template to use
	// +kubebuilder:validation:Optional
	VappTemplateID *string `json:"vappTemplateId,omitempty" tf:"vapp_template_id,omitempty"`

	// The name of VDC to use, optional if defined at provider level
	// +kubebuilder:validation:Optional
	Vdc *string `json:"vdc,omitempty" tf:"vdc,omitempty"`
}

// VmSpec defines the desired state of Vm
type VmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VmParameters `json:"forProvider"`
}

// VmStatus defines the observed state of Vm.
type VmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vm is the Schema for the Vms API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vcd}
type Vm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   VmSpec   `json:"spec"`
	Status VmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VmList contains a list of Vms
type VmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vm `json:"items"`
}

// Repository type metadata.
var (
	Vm_Kind             = "Vm"
	Vm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vm_Kind}.String()
	Vm_KindAPIVersion   = Vm_Kind + "." + CRDGroupVersion.String()
	Vm_GroupVersionKind = CRDGroupVersion.WithKind(Vm_Kind)
)

func init() {
	SchemeBuilder.Register(&Vm{}, &VmList{})
}
