package lb

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_lb_app_profile", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "LbAppProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_lb_app_rule", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "LbAppRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_lb_server_pool", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "LbServerPool"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_lb_service_monitor", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "LbServiceMonitor"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_lb_virtual_server", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "LbVirtualServer"
		r.Version = "v1alpha1"
	})
}