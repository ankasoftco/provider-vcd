package org

import "github.com/upbound/upjet/pkg/config"

const shortGroup string = "vcd"
const version string = "v1alpha1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_org", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Org"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_org_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrgGroup"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_org_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrgLdap"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_org_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrgSaml"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_org_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrgUser"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_org_vdc", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrgVdc"
		r.Version = version
	})

	p.AddResourceConfigurator("vcd_org_vdc_access_control", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "OrgVdcAccessControl"
		r.Version = version
	})
}
