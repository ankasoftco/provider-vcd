/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"vcd_api_token": 									config.IdentifierFromProvider,
	"vcd_catalog": 										config.IdentifierFromProvider,
	"vcd_catalog_acces_control": 						config.IdentifierFromProvider,
	"vcd_catalog_item": 								config.IdentifierFromProvider,
	"vcd_catalog_media": 								config.IdentifierFromProvider,
	"vcd_catalog_vapp_template": 						config.IdentifierFromProvider,
	"vcd_certificate_library": 							config.IdentifierFromProvider,
	"vcd_cloned_vapp": 									config.IdentifierFromProvider,
	"vcd_edgegateway": 									config.IdentifierFromProvider,
	"vcd_edgegateway_settings": 						config.IdentifierFromProvider,
	"vcd_edgegateway_vpn": 								config.IdentifierFromProvider,
	"vcd_external_network": 							config.IdentifierFromProvider,
	"vcd_external_network_v2": 							config.IdentifierFromProvider,
	"vcd_global_role":		 							config.IdentifierFromProvider,
	"vcd_independent_disk":		 						config.IdentifierFromProvider,
	"vcd_inserted_media":		 						config.IdentifierFromProvider,
	"vcd_ip_space":		 								config.IdentifierFromProvider,
	"vcd_ip_space_custom_quota":						config.IdentifierFromProvider,
	"vcd_ip_space_ip_allocation":						config.IdentifierFromProvider,
	"vcd_ip_space_uplink":		 						config.IdentifierFromProvider,
	"vcd_lb_app_profile":		 						config.IdentifierFromProvider,
	"vcd_lb_app_rule":		 							config.IdentifierFromProvider,
	"vcd_lb_server_pool":		 						config.IdentifierFromProvider,
	"vcd_lb_service_monitor":						 	config.IdentifierFromProvider,
	"vcd_lb_virtual_server":							config.IdentifierFromProvider,
	"vcd_network_direct":								config.IdentifierFromProvider,
	"vcd_network_isolated":								config.IdentifierFromProvider,
	"vcd_network_isolated_v2":							config.IdentifierFromProvider,
	"vcd_network_routed":								config.IdentifierFromProvider,
	"vcd_network_routed_v2":							config.IdentifierFromProvider,
	"vcd_nsxt_alb_cloud":								config.IdentifierFromProvider,
	"vcd_nsxt_alb_controller":							config.IdentifierFromProvider,
	"vcd_nsxt_alb_edgegateway_service_engine_group":	config.IdentifierFromProvider,
	"vcd_nsxt_alb_pool":								config.IdentifierFromProvider,
	"vcd_nsxt_alb_service_engine_group":				config.IdentifierFromProvider,
	"vcd_nsxt_alb_settings":							config.IdentifierFromProvider,
	"vcd_nsxt_alb_virtual_service":						config.IdentifierFromProvider,
	"vcd_nsxt_app_port_profile":						config.IdentifierFromProvider,
	"vcd_nsxt_distributed_firewall":					config.IdentifierFromProvider,
	"vcd_nsxt_distributed_firewall_rule":				config.IdentifierFromProvider,
	"vcd_nsxt_dynamic_security_group":					config.IdentifierFromProvider,
	"vcd_nsxt_edgegateway":								config.IdentifierFromProvider,
	"vcd_nsxt_edgegateway_bgp_configuration":			config.IdentifierFromProvider,
	"vcd_nsxt_edgegateway_bgp_ip_prefix_list":			config.IdentifierFromProvider,
	"vcd_nsxt_edgegateway_bgp_neighbor":				config.IdentifierFromProvider,
	"vcd_nsxt_edgegateway_dhcpv6":						config.IdentifierFromProvider,
	"vcd_nsxt_edgegateway_rate_limiting":				config.IdentifierFromProvider,
	"vcd_nsxt_edgegateway_static_route":				config.IdentifierFromProvider,
	"vcd_nsxt_firewall":								config.IdentifierFromProvider,
	"vcd_nsxt_ip_set":									config.IdentifierFromProvider,
	"vcd_nsxt_ipsec_vpn_tunnel":						config.IdentifierFromProvider,
	"vcd_nsxt_nat_rule":								config.IdentifierFromProvider,
	"vcd_nsxt_network_dhcp":							config.IdentifierFromProvider,
	"vcd_nsxt_network_dhcp_binding":					config.IdentifierFromProvider,
	"vcd_nsxt_network_imported":						config.IdentifierFromProvider,
	"vcd_nsxt_route_advertisement":						config.IdentifierFromProvider,
	"vcd_nsxt_security_group":							config.IdentifierFromProvider,
	"vcd_nsxv_dhcp_relay":								config.IdentifierFromProvider,
	"vcd_nsxv_distributed_firewall":					config.IdentifierFromProvider,
	"vcd_nsxv_dnat":									config.IdentifierFromProvider,
	"vcd_nsxv_firewall_rule":							config.IdentifierFromProvider,
	"vcd_nsxv_snat":									config.IdentifierFromProvider,
	"vcd_nsxv_ip_set":									config.IdentifierFromProvider,
	"vcd_org":											config.IdentifierFromProvider,
	"vcd_org_group":									config.IdentifierFromProvider,
	"vcd_org_ldap":										config.IdentifierFromProvider,
	"vcd_org_saml":										config.IdentifierFromProvider,
	"vcd_org_user":										config.IdentifierFromProvider,
	"vcd_org_vdc":										config.IdentifierFromProvider,
	"vcd_org_vdc_access_control":						config.IdentifierFromProvider,
	"vcd_provider_vdc":									config.IdentifierFromProvider,
	"vcd_rde":											config.IdentifierFromProvider,
	"vcd_rde_interface":								config.IdentifierFromProvider,
	"vcd_rde_interface_behavior":						config.IdentifierFromProvider,
	"vcd_rde_type":										config.IdentifierFromProvider,
	"vcd_rde_type_behavior":							config.IdentifierFromProvider,
	"vcd_rde_type_behavior_acl":						config.IdentifierFromProvider,
	"vcd_rights_bundle":								config.IdentifierFromProvider,
	"vcd_role":											config.IdentifierFromProvider,
	"vcd_security_tag":									config.IdentifierFromProvider,
	"vcd_service_account":								config.IdentifierFromProvider,
	"vcd_subscribed_catalog":							config.IdentifierFromProvider,
	"vcd_ui_plugin":									config.IdentifierFromProvider,
	"vcd_vapp":											config.IdentifierFromProvider,
	"vcd_vapp_access_control":							config.IdentifierFromProvider,
	"vcd_vapp_firewall_rules":							config.IdentifierFromProvider,
	"vcd_vapp_nat_rules":								config.IdentifierFromProvider,
	"vcd_vapp_network":									config.IdentifierFromProvider,
	"vcd_vapp_org_network":								config.IdentifierFromProvider,
	"vcd_vapp_static_routing":							config.IdentifierFromProvider,
	"vcd_vapp_vm":										config.IdentifierFromProvider,
	"vcd_vdc_group":									config.IdentifierFromProvider,
	"vcd_vm":											config.IdentifierFromProvider,
	"vcd_vm_affinity_rule":								config.IdentifierFromProvider,
	"vcd_vm_internal_disk":								config.IdentifierFromProvider,
	"vcd_vm_placement_policy":							config.IdentifierFromProvider,
	"vcd_vm_sizing_policy":								config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
