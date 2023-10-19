package nsxv

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxv_dhcp_relay", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxvDhcpRelay"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxv_distributed_firewall", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxvDistributedFirewall"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxv_dnat", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxvDnat"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxv_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxvFirewallRule"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxv_ip_set", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxvIpSet"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxv_snat", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxvSnat"
		r.Version = version
	})
}
