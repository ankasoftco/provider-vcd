package vm

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_vm", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Vm"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vm_affinity_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VmAffinityRule"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vm_internal_disk", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VmInternalDisk"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vm_placement_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VmPlacementPolicy"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vm_sizing_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VmSizingPolicy"
		r.Version = version
	})
}
