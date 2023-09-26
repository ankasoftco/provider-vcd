package independentDisk

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_independent_disk", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "IndependentDisk"
		r.Version = "v1alpha1"
	})
}