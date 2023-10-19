package lb

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_lb_app_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbAppProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_lb_app_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbAppRule"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_lb_server_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbServerPool"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_lb_service_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbServiceMonitor"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_lb_virtual_server", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LbVirtualServer"
		r.Version = version
	})
}
