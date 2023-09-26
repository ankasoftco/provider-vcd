/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"vcd_api_token": 					config.IdentifierFromProvider,
	"vcd_catalog": 						config.IdentifierFromProvider,
	"vcd_catalog_acces_control": 		config.IdentifierFromProvider,
	"vcd_catalog_item": 				config.IdentifierFromProvider,
	"vcd_catalog_media": 				config.IdentifierFromProvider,
	"vcd_catalog_vapp_template": 		config.IdentifierFromProvider,
	"vcd_certificate_library": 			config.IdentifierFromProvider,
	"vcd_cloned_vapp": 					config.IdentifierFromProvider,
	"vcd_edgegateway": 					config.IdentifierFromProvider,
	"vcd_edgegateway_settings": 		config.IdentifierFromProvider,
	"vcd_edgegateway_vpn": 				config.IdentifierFromProvider,
	"vcd_external_network": 			config.IdentifierFromProvider,
	"vcd_external_network_v2": 			config.IdentifierFromProvider,
	"vcd_global_role":		 			config.IdentifierFromProvider,
	"vcd_independent_disk":		 		config.IdentifierFromProvider,
	"vcd_inserted_media":		 		config.IdentifierFromProvider,
	"vcd_ip_space":		 				config.IdentifierFromProvider,
	"vcd_ip_space_custom_quota":		config.IdentifierFromProvider,
	"vcd_ip_space_ip_allocation":		config.IdentifierFromProvider,
	"vcd_ip_space_uplink":		 		config.IdentifierFromProvider,
	"vcd_lb_app_profile":		 		config.IdentifierFromProvider,
	"vcd_lb_app_rule":		 			config.IdentifierFromProvider,
	"vcd_lb_server_pool":		 		config.IdentifierFromProvider,
	"vcd_lb_service_monitor":		 	config.IdentifierFromProvider,
	"vcd_lb_virtual_server":			config.IdentifierFromProvider,
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
