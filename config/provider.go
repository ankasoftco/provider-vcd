/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	apiToken "github.com/ankasoftco/provider-vcd/config/api_token"
	catalog "github.com/ankasoftco/provider-vcd/config/catalog"
	certificateLibrary "github.com/ankasoftco/provider-vcd/config/certificate_library"
	clonedvApp "github.com/ankasoftco/provider-vcd/config/cloned_vapp"
	edgegateway "github.com/ankasoftco/provider-vcd/config/edgegateway"
	externalNetwork "github.com/ankasoftco/provider-vcd/config/external_network"
	globalRole "github.com/ankasoftco/provider-vcd/config/global_role"
	independentDisk "github.com/ankasoftco/provider-vcd/config/independent_disk"
	insertedMedia "github.com/ankasoftco/provider-vcd/config/inserted_media"
	ipSpace "github.com/ankasoftco/provider-vcd/config/ip_space"
	lb "github.com/ankasoftco/provider-vcd/config/lb"
	network "github.com/ankasoftco/provider-vcd/config/network"
	nsxt "github.com/ankasoftco/provider-vcd/config/nsxt"
	nsxv "github.com/ankasoftco/provider-vcd/config/nsxv"
	org "github.com/ankasoftco/provider-vcd/config/org"
	providerVdc "github.com/ankasoftco/provider-vcd/config/provider_vdc"
	rde "github.com/ankasoftco/provider-vcd/config/rde"
	rightsBundle "github.com/ankasoftco/provider-vcd/config/rights_bundle"
	role "github.com/ankasoftco/provider-vcd/config/role"
	securityTag "github.com/ankasoftco/provider-vcd/config/security_tag"
	serviceAccount "github.com/ankasoftco/provider-vcd/config/service_account"
	subscribedCatalog "github.com/ankasoftco/provider-vcd/config/subscribed_catalog"
	uiPlugin "github.com/ankasoftco/provider-vcd/config/ui_plugin"
	vApp "github.com/ankasoftco/provider-vcd/config/vapp"
	vdcGroup "github.com/ankasoftco/provider-vcd/config/vdc_group"
	vm "github.com/ankasoftco/provider-vcd/config/vm"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "vcd"
	modulePath     = "github.com/ankasoftco/provider-vcd"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		apiToken.Configure,
		catalog.Configure,
		certificateLibrary.Configure,
		clonedvApp.Configure,
		edgegateway.Configure,
		externalNetwork.Configure,
		globalRole.Configure,
		independentDisk.Configure,
		insertedMedia.Configure,
		ipSpace.Configure,
		lb.Configure,
		network.Configure,
		nsxt.Configure,
		nsxv.Configure,
		org.Configure,
		providerVdc.Configure,
		rde.Configure,
		rightsBundle.Configure,
		role.Configure,
		securityTag.Configure,
		serviceAccount.Configure,
		subscribedCatalog.Configure,
		uiPlugin.Configure,
		vApp.Configure,
		vdcGroup.Configure,
		vm.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
