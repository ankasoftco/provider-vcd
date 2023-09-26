package vApp

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_vapp", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VApp"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vapp_access_control", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VAppAccessControl"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vapp_firewall_rules", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VAppFirewallRules"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vapp_nat_rules", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VAppNatRules"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vapp_network", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VAppNetwork"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vapp_org_network", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VAppOrgNetwork"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vapp_static_routing", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VAppStaticRouting"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_vapp_vm", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "VAppVm"
		r.Version = "v1alpha1"
	})
}