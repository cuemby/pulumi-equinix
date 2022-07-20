// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "equinix:index/eCXL2Connection:ECXL2Connection":
		r = &ECXL2Connection{}
	case "equinix:index/eCXL2ConnectionAccepter:ECXL2ConnectionAccepter":
		r = &ECXL2ConnectionAccepter{}
	case "equinix:index/eCXL2Serviceprofile:ECXL2Serviceprofile":
		r = &ECXL2Serviceprofile{}
	case "equinix:index/metalBGPSession:MetalBGPSession":
		r = &MetalBGPSession{}
	case "equinix:index/metalConnection:MetalConnection":
		r = &MetalConnection{}
	case "equinix:index/metalDevice:MetalDevice":
		r = &MetalDevice{}
	case "equinix:index/metalDeviceNetworkType:MetalDeviceNetworkType":
		r = &MetalDeviceNetworkType{}
	case "equinix:index/metalGateway:MetalGateway":
		r = &MetalGateway{}
	case "equinix:index/metalIPAttachment:MetalIPAttachment":
		r = &MetalIPAttachment{}
	case "equinix:index/metalOrganization:MetalOrganization":
		r = &MetalOrganization{}
	case "equinix:index/metalPort:MetalPort":
		r = &MetalPort{}
	case "equinix:index/metalPortVlanAttachment:MetalPortVlanAttachment":
		r = &MetalPortVlanAttachment{}
	case "equinix:index/metalProject:MetalProject":
		r = &MetalProject{}
	case "equinix:index/metalProjectAPIKey:MetalProjectAPIKey":
		r = &MetalProjectAPIKey{}
	case "equinix:index/metalProjectSSHKey:MetalProjectSSHKey":
		r = &MetalProjectSSHKey{}
	case "equinix:index/metalReservedIPBlock:MetalReservedIPBlock":
		r = &MetalReservedIPBlock{}
	case "equinix:index/metalSSHKey:MetalSSHKey":
		r = &MetalSSHKey{}
	case "equinix:index/metalSpotMarketRequest:MetalSpotMarketRequest":
		r = &MetalSpotMarketRequest{}
	case "equinix:index/metalUserAPIKey:MetalUserAPIKey":
		r = &MetalUserAPIKey{}
	case "equinix:index/metalVRF:MetalVRF":
		r = &MetalVRF{}
	case "equinix:index/metalVirtualCircuit:MetalVirtualCircuit":
		r = &MetalVirtualCircuit{}
	case "equinix:index/metalVlan:MetalVlan":
		r = &MetalVlan{}
	case "equinix:index/networkACLTemplate:NetworkACLTemplate":
		r = &NetworkACLTemplate{}
	case "equinix:index/networkBGP:NetworkBGP":
		r = &NetworkBGP{}
	case "equinix:index/networkDevice:NetworkDevice":
		r = &NetworkDevice{}
	case "equinix:index/networkDeviceLink:NetworkDeviceLink":
		r = &NetworkDeviceLink{}
	case "equinix:index/networkSSHKey:NetworkSSHKey":
		r = &NetworkSSHKey{}
	case "equinix:index/networkSSHUser:NetworkSSHUser":
		r = &NetworkSSHUser{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:equinix" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"equinix",
		"index/eCXL2Connection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/eCXL2ConnectionAccepter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/eCXL2Serviceprofile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalBGPSession",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalDevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalDeviceNetworkType",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalIPAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalOrganization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalPort",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalPortVlanAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalProject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalProjectAPIKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalProjectSSHKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalReservedIPBlock",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalSSHKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalSpotMarketRequest",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalUserAPIKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalVRF",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalVirtualCircuit",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/metalVlan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/networkACLTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/networkBGP",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/networkDevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/networkDeviceLink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/networkSSHKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"equinix",
		"index/networkSSHUser",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"equinix",
		&pkg{version},
	)
}
