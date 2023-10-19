package network

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_network_direct", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NetworkDirect"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_network_isolated", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NetworkIsolated"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_network_isolated_v2", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NetworkIsolatedV2"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_network_routed", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NetworkRouted"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_network_routed_v2", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NetworkRoutedV2"
		r.Version = version
	})
}
