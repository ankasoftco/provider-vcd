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
	globalrole "github.com/ankasoftco/provider-vcd/internal/controller/vcd/globalrole"
	independentdisk "github.com/ankasoftco/provider-vcd/internal/controller/vcd/independentdisk"
	insertedmedia "github.com/ankasoftco/provider-vcd/internal/controller/vcd/insertedmedia"
	ipspace "github.com/ankasoftco/provider-vcd/internal/controller/vcd/ipspace"
	ipspacecustomquota "github.com/ankasoftco/provider-vcd/internal/controller/vcd/ipspacecustomquota"
	ipspaceipallocation "github.com/ankasoftco/provider-vcd/internal/controller/vcd/ipspaceipallocation"
	ipspaceuplink "github.com/ankasoftco/provider-vcd/internal/controller/vcd/ipspaceuplink"
	lbappprofile "github.com/ankasoftco/provider-vcd/internal/controller/vcd/lbappprofile"
	lbapprule "github.com/ankasoftco/provider-vcd/internal/controller/vcd/lbapprule"
	lbserverpool "github.com/ankasoftco/provider-vcd/internal/controller/vcd/lbserverpool"
	lbservicemonitor "github.com/ankasoftco/provider-vcd/internal/controller/vcd/lbservicemonitor"
	lbvirtualserver "github.com/ankasoftco/provider-vcd/internal/controller/vcd/lbvirtualserver"
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
		globalrole.Setup,
		independentdisk.Setup,
		insertedmedia.Setup,
		ipspace.Setup,
		ipspacecustomquota.Setup,
		ipspaceipallocation.Setup,
		ipspaceuplink.Setup,
		lbappprofile.Setup,
		lbapprule.Setup,
		lbserverpool.Setup,
		lbservicemonitor.Setup,
		lbvirtualserver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
