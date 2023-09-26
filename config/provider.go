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
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
