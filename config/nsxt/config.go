package nsxt

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxt_alb_cloud", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAlbCloud"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_controller", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAlbController"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_edgegateway_service_engine_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAlbEdgegatewayServiceEngineGroup"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAlbPool"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_service_engine_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAlbServiceEngineGroup"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAlbSettings"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_virtual_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAlbVirtualService"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_app_port_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtAppPortProfile"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_distributed_firewall", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtDistributedFirewall"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_distributed_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtDistributedFirewallRule"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_dynamic_security_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtDynamicSecurityGroup"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgeGateway"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_configuration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgegatewayBgpConfiguration"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_ip_prefix_list", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgeGatewayBgpIpPrefixList"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_neighbor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgeGatewayBgpNeighbor"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_dhcp_forwarding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgeGatewayDhcpForwarding"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_dhcpv6", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgeGatewayDhcpV6"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_rate_limiting", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgeGatewayRateLimit"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_static_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtEdgeGatewayStaticRoute"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_firewall", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtFirewall"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_ip_set", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtIpSet"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_ipsec_vpn_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtIpsecVpnTunnel"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_nat_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtNatRule"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_network_dhcp", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtNetworkDhcp"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_network_dhcp_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtNetworkDhcpBinding"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_network_imported", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtNetworkImported"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_nsxt_route_advertisement", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtRouteAdvertisement"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_security_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "NsxtSecurityGroup"
		r.Version = version
	})
}
