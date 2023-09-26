package nsxt

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_nsxt_alb_cloud", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAlbCloud"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_controller", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAlbController"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_edgegateway_service_engine_group", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAlbEdgegatewayServiceEngineGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_pool", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAlbPool"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_service_engine_group", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAlbServiceEngineGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_settings", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAlbSettings"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_alb_virtual_service", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAlbVirtualService"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_app_port_profile", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtAppPortProfile"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_distributed_firewall", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtDistributedFirewall"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_distributed_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtDistributedFirewallRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_dynamic_security_group", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtDynamicSecurityGroup"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgeGateway"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_configuration", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgegatewayBgpConfiguration"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_ip_prefix_list", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgeGatewayBgpIpPrefixList"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_bgp_neighbor", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgeGatewayBgpNeighbor"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_dhcp_forwarding", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgeGatewayDhcpForwarding"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_dhcpv6", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgeGatewayDhcpV6"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_rate_limiting", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgeGatewayRateLimit"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_edgegateway_static_route", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtEdgeGatewayStaticRoute"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_firewall", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtFirewall"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_ip_set", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtIpSet"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_ipsec_vpn_tunnel", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtIpsecVpnTunnel"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_nat_rule", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtNatRule"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_network_dhcp", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtNetworkDhcp"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_network_dhcp_binding", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtNetworkDhcpBinding"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_network_imported", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtNetworkImported"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_route_advertisement", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtRouteAdvertisement"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_nsxt_security_group", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "NsxtSecurityGroup"
		r.Version = "v1alpha1"
	})
}