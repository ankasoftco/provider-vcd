package providerVdc

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_provider_vdc", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "ProviderVdc"
		r.Version = "v1alpha1"
	})
}