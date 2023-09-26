/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	apiToken "github.com/ankasoftco/provider-vcd/config/api_token"
	catalog "github.com/ankasoftco/provider-vcd/config/catalog"
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
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
