package serviceAccount

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_service_account", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "ServiceAccount"
		r.Version = "v1alpha1"
	})
}