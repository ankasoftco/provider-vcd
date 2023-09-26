package externalNetwork

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_external_network", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "ExternalNetwork"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_external_network_v2", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "ExternalNetworkV2"
		r.Version = "v1alpha1"
	})
}