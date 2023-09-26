package catalog

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_catalog", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "Catalog"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_catalog_access_control", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "CatalogAccessControl"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_catalog_item", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "CatalogItem"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_catalog_media", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "CatalogMedia"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_catalog_vapp_template", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "CatalogvAppTemplate"
		r.Version = "v1alpha1"
	})
}