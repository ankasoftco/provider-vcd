package securityTag

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_security_tag", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "SecurityTag"
		r.Version = "v1alpha1"
	})
}