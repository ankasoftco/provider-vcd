package nsxv

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxv_dhcp_relay", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxvDhcpRelay"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxv_distributed_firewall", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxvDistributedFirewall"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxv_dnat", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxvDnat"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxv_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxvFirewallRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxv_ip_set", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxvIpSet"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxv_snat", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxvSnat"
		r.Version = "v1alpha1"
	})
}