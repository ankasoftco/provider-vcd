package clonedvapp

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_cloned_vapp", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "ClonedvApp"
		r.Version = "v1alpha1"
	})
}
