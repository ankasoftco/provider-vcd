package globalrole

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_global_role", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "GlobalRole"
		r.Version = "v1alpha1"
	})
}
