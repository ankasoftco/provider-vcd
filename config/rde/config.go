package rde

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_rde", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "Rde"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_rde_interface", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "RdeInterface"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_rde_interface_behavior", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "RdeInterfaceBehavior"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_rde_type", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "RdeType"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_rde_type_behavior", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "RdeTypeBehavior"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_rde_type_behavior_acl", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "RdeTypeBehaviorAcl"
		r.Version = "v1alpha1"
	})
}