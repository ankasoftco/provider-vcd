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
	networkdirect "github.com/ankasoftco/provider-vcd/internal/controller/vcd/networkdirect"
	networkisolated "github.com/ankasoftco/provider-vcd/internal/controller/vcd/networkisolated"
	networkisolatedv2 "github.com/ankasoftco/provider-vcd/internal/controller/vcd/networkisolatedv2"
	networkrouted "github.com/ankasoftco/provider-vcd/internal/controller/vcd/networkrouted"
	networkroutedv2 "github.com/ankasoftco/provider-vcd/internal/controller/vcd/networkroutedv2"
	nsxtalbcloud "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtalbcloud"
	nsxtalbcontroller "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtalbcontroller"
	nsxtalbedgegatewayserviceenginegroup "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtalbedgegatewayserviceenginegroup"
	nsxtalbpool "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtalbpool"
	nsxtalbserviceenginegroup "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtalbserviceenginegroup"
	nsxtalbsettings "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtalbsettings"
	nsxtalbvirtualservice "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtalbvirtualservice"
	nsxtappportprofile "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtappportprofile"
	nsxtdistributedfirewall "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtdistributedfirewall"
	nsxtdistributedfirewallrule "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtdistributedfirewallrule"
	nsxtdynamicsecuritygroup "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtdynamicsecuritygroup"
	nsxtedgegateway "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtedgegateway"
	nsxtedgegatewaybgpconfiguration "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtedgegatewaybgpconfiguration"
	nsxtedgegatewaybgpipprefixlist "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtedgegatewaybgpipprefixlist"
	nsxtedgegatewaybgpneighbor "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtedgegatewaybgpneighbor"
	nsxtedgegatewaydhcpv6 "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtedgegatewaydhcpv6"
	nsxtedgegatewayratelimit "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtedgegatewayratelimit"
	nsxtedgegatewaystaticroute "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtedgegatewaystaticroute"
	nsxtfirewall "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtfirewall"
	nsxtipsecvpntunnel "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtipsecvpntunnel"
	nsxtipset "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtipset"
	nsxtnatrule "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtnatrule"
	nsxtnetworkdhcp "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtnetworkdhcp"
	nsxtnetworkdhcpbinding "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtnetworkdhcpbinding"
	nsxtnetworkimported "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtnetworkimported"
	nsxtrouteadvertisement "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtrouteadvertisement"
	nsxtsecuritygroup "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxtsecuritygroup"
	nsxvdhcprelay "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxvdhcprelay"
	nsxvdistributedfirewall "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxvdistributedfirewall"
	nsxvdnat "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxvdnat"
	nsxvfirewallrule "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxvfirewallrule"
	nsxvipset "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxvipset"
	nsxvsnat "github.com/ankasoftco/provider-vcd/internal/controller/vcd/nsxvsnat"
	org "github.com/ankasoftco/provider-vcd/internal/controller/vcd/org"
	orggroup "github.com/ankasoftco/provider-vcd/internal/controller/vcd/orggroup"
	orgldap "github.com/ankasoftco/provider-vcd/internal/controller/vcd/orgldap"
	orgsaml "github.com/ankasoftco/provider-vcd/internal/controller/vcd/orgsaml"
	orguser "github.com/ankasoftco/provider-vcd/internal/controller/vcd/orguser"
	orgvdc "github.com/ankasoftco/provider-vcd/internal/controller/vcd/orgvdc"
	orgvdcaccesscontrol "github.com/ankasoftco/provider-vcd/internal/controller/vcd/orgvdcaccesscontrol"
	providervdc "github.com/ankasoftco/provider-vcd/internal/controller/vcd/providervdc"
	rde "github.com/ankasoftco/provider-vcd/internal/controller/vcd/rde"
	rdeinterface "github.com/ankasoftco/provider-vcd/internal/controller/vcd/rdeinterface"
	rdeinterfacebehavior "github.com/ankasoftco/provider-vcd/internal/controller/vcd/rdeinterfacebehavior"
	rdetype "github.com/ankasoftco/provider-vcd/internal/controller/vcd/rdetype"
	rdetypebehavior "github.com/ankasoftco/provider-vcd/internal/controller/vcd/rdetypebehavior"
	rdetypebehavioracl "github.com/ankasoftco/provider-vcd/internal/controller/vcd/rdetypebehavioracl"
	rightsbundle "github.com/ankasoftco/provider-vcd/internal/controller/vcd/rightsbundle"
	role "github.com/ankasoftco/provider-vcd/internal/controller/vcd/role"
	securitytag "github.com/ankasoftco/provider-vcd/internal/controller/vcd/securitytag"
	serviceaccount "github.com/ankasoftco/provider-vcd/internal/controller/vcd/serviceaccount"
	subscribedcatalog "github.com/ankasoftco/provider-vcd/internal/controller/vcd/subscribedcatalog"
	uiplugin "github.com/ankasoftco/provider-vcd/internal/controller/vcd/uiplugin"
	vapp "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vapp"
	vappaccesscontrol "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vappaccesscontrol"
	vappfirewallrules "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vappfirewallrules"
	vappnatrules "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vappnatrules"
	vappnetwork "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vappnetwork"
	vapporgnetwork "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vapporgnetwork"
	vappstaticrouting "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vappstaticrouting"
	vappvm "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vappvm"
	vdcgroup "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vdcgroup"
	vm "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vm"
	vmaffinityrule "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vmaffinityrule"
	vminternaldisk "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vminternaldisk"
	vmplacementpolicy "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vmplacementpolicy"
	vmsizingpolicy "github.com/ankasoftco/provider-vcd/internal/controller/vcd/vmsizingpolicy"
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
		networkdirect.Setup,
		networkisolated.Setup,
		networkisolatedv2.Setup,
		networkrouted.Setup,
		networkroutedv2.Setup,
		nsxtalbcloud.Setup,
		nsxtalbcontroller.Setup,
		nsxtalbedgegatewayserviceenginegroup.Setup,
		nsxtalbpool.Setup,
		nsxtalbserviceenginegroup.Setup,
		nsxtalbsettings.Setup,
		nsxtalbvirtualservice.Setup,
		nsxtappportprofile.Setup,
		nsxtdistributedfirewall.Setup,
		nsxtdistributedfirewallrule.Setup,
		nsxtdynamicsecuritygroup.Setup,
		nsxtedgegateway.Setup,
		nsxtedgegatewaybgpconfiguration.Setup,
		nsxtedgegatewaybgpipprefixlist.Setup,
		nsxtedgegatewaybgpneighbor.Setup,
		nsxtedgegatewaydhcpv6.Setup,
		nsxtedgegatewayratelimit.Setup,
		nsxtedgegatewaystaticroute.Setup,
		nsxtfirewall.Setup,
		nsxtipsecvpntunnel.Setup,
		nsxtipset.Setup,
		nsxtnatrule.Setup,
		nsxtnetworkdhcp.Setup,
		nsxtnetworkdhcpbinding.Setup,
		nsxtnetworkimported.Setup,
		nsxtrouteadvertisement.Setup,
		nsxtsecuritygroup.Setup,
		nsxvdhcprelay.Setup,
		nsxvdistributedfirewall.Setup,
		nsxvdnat.Setup,
		nsxvfirewallrule.Setup,
		nsxvipset.Setup,
		nsxvsnat.Setup,
		org.Setup,
		orggroup.Setup,
		orgldap.Setup,
		orgsaml.Setup,
		orguser.Setup,
		orgvdc.Setup,
		orgvdcaccesscontrol.Setup,
		providervdc.Setup,
		rde.Setup,
		rdeinterface.Setup,
		rdeinterfacebehavior.Setup,
		rdetype.Setup,
		rdetypebehavior.Setup,
		rdetypebehavioracl.Setup,
		rightsbundle.Setup,
		role.Setup,
		securitytag.Setup,
		serviceaccount.Setup,
		subscribedcatalog.Setup,
		uiplugin.Setup,
		vapp.Setup,
		vappaccesscontrol.Setup,
		vappfirewallrules.Setup,
		vappnatrules.Setup,
		vappnetwork.Setup,
		vapporgnetwork.Setup,
		vappstaticrouting.Setup,
		vappvm.Setup,
		vdcgroup.Setup,
		vm.Setup,
		vmaffinityrule.Setup,
		vminternaldisk.Setup,
		vmplacementpolicy.Setup,
		vmsizingpolicy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
