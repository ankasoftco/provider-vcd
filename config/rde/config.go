package rde

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_rde", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Rde"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_rde_interface", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RdeInterface"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_rde_interface_behavior", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RdeInterfaceBehavior"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_rde_type", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RdeType"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_rde_type_behavior", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RdeTypeBehavior"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_rde_type_behavior_acl", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RdeTypeBehaviorAcl"
		r.Version = version
	})
}
