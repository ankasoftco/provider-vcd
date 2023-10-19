package apitoken

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_api_token", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "ApiToken"
		r.Version = "v1alpha1"
	})
}
