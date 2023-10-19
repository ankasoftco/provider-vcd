package edgegateway

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_edgegateway", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "Edgegateway"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_edgegateway_settings", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "EdgegatewaySettings"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_edgegateway_vpn", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "EdgegatewayVPN"
		r.Version = "v1alpha1"
	})
}
