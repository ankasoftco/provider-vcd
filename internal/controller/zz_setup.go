/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	providerconfig "github.com/ankasoftco/provider-vcd/internal/controller/providerconfig"
	apitoken "github.com/ankasoftco/provider-vcd/internal/controller/vcd/apitoken"
	catalog "github.com/ankasoftco/provider-vcd/internal/controller/vcd/catalog"
	catalogitem "github.com/ankasoftco/provider-vcd/internal/controller/vcd/catalogitem"
	catalogmedia "github.com/ankasoftco/provider-vcd/internal/controller/vcd/catalogmedia"
	catalogvapptemplate "github.com/ankasoftco/provider-vcd/internal/controller/vcd/catalogvapptemplate"
	clonedvapp "github.com/ankasoftco/provider-vcd/internal/controller/vcd/clonedvapp"
	edgegateway "github.com/ankasoftco/provider-vcd/internal/controller/vcd/edgegateway"
	edgegatewaysettings "github.com/ankasoftco/provider-vcd/internal/controller/vcd/edgegatewaysettings"
	edgegatewayvpn "github.com/ankasoftco/provider-vcd/internal/controller/vcd/edgegatewayvpn"
	externalnetwork "github.com/ankasoftco/provider-vcd/internal/controller/vcd/externalnetwork"
	externalnetworkv2 "github.com/ankasoftco/provider-vcd/internal/controller/vcd/externalnetworkv2"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		apitoken.Setup,
		catalog.Setup,
		catalogitem.Setup,
		catalogmedia.Setup,
		catalogvapptemplate.Setup,
		clonedvapp.Setup,
		edgegateway.Setup,
		edgegatewaysettings.Setup,
		edgegatewayvpn.Setup,
		externalnetwork.Setup,
		externalnetworkv2.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
