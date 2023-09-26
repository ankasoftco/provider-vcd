package vm

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_vm", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "Vm"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vm_affinity_rule", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VmAffinityRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vm_internal_disk", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VmInternalDisk"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vm_placement_policy", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VmPlacementPolicy"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vm_sizing_policy", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VmSizingPolicy"
		r.Version = "v1alpha1"
	})
}