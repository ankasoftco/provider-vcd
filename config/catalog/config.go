package catalog

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_catalog", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Catalog"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_catalog_access_control", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CatalogAccessControl"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_catalog_item", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CatalogItem"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_catalog_media", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CatalogMedia"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_catalog_vapp_template", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CatalogvAppTemplate"
		r.Version = version
	})
}
