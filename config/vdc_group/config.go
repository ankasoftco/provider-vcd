package vdcGroup

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_vdc_group", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VdcGroup"
		r.Version = "v1alpha1"
	})
}