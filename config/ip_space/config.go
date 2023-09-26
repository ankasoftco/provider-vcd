package ipSpace

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_ip_space", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpace"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_ip_space_custom_quota", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpaceCustomQuota"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_ip_space_ip_allocation", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpaceIpAllocation"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_ip_space_uplink", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IpSpaceUplink"
		r.Version = "v1alpha1"
	})
}