package org

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vcd_org", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "Org"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_org_group", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "OrgGroup"
		r.Version = "v1alpha1"
	})
	
	p.AddResourceConfigurator("vcd_org_ldap", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "OrgLdap"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_org_saml", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "OrgSaml"
		r.Version = "v1alpha1"
	})
	
	p.AddResourceConfigurator("vcd_org_user", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "OrgUser"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_org_vdc", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "OrgVdc"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vcd_org_vdc_access_control", func(r *config.Resource) {
		r.ShortGroup = "vcd"
		r.Kind = "OrgVdcAccessControl"
		r.Version = "v1alpha1"
	})
}