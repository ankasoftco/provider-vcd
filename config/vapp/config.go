package vapp

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_vapp", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VApp"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vapp_access_control", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VAppAccessControl"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vapp_firewall_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VAppFirewallRules"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vapp_nat_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VAppNatRules"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vapp_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VAppNetwork"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vapp_org_network", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VAppOrgNetwork"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vapp_static_routing", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VAppStaticRouting"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_vapp_vm", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VAppVm"
		r.Version = version
	})
}
