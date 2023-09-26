package network

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_network_direct", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NetworkDirect"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_network_isolated", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NetworkIsolated"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_network_isolated_v2", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NetworkIsolatedV2"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_network_routed", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NetworkRouted"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_network_routed_v2", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NetworkRoutedV2"
		r.Version = "v1alpha1"
	})
}