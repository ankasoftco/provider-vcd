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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
