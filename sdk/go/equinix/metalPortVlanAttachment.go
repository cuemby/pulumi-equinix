// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to attach device ports to VLANs.
//
// Device and VLAN must be in the same facility.
//
// If you need this resource to add the port back to bond on removal, set `forceBond = true`.
//
// To learn more about Layer 2 networking in Equinix Metal, refer to
//
// * <https://metal.equinix.com/developers/docs/networking/layer2/>
// * <https://metal.equinix.com/developers/docs/networking/layer2-configs/>
//
// ## Example Usage
// ### Hybrid network type
//
// ```go
// package main
//
// import (
// 	"github.com/cuemby/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testMetalVlan, err := equinix.NewMetalVlan(ctx, "testMetalVlan", &equinix.MetalVlanArgs{
// 			Description: pulumi.String("VLAN in New Jersey"),
// 			Facility:    pulumi.String("ny5"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testMetalDevice, err := equinix.NewMetalDevice(ctx, "testMetalDevice", &equinix.MetalDeviceArgs{
// 			Hostname: pulumi.String("test"),
// 			Plan:     pulumi.String("c3.small.x86"),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("ny5"),
// 			},
// 			OperatingSystem: pulumi.String("ubuntu_20_04"),
// 			BillingCycle:    pulumi.String("hourly"),
// 			ProjectId:       pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testMetalDeviceNetworkType, err := equinix.NewMetalDeviceNetworkType(ctx, "testMetalDeviceNetworkType", &equinix.MetalDeviceNetworkTypeArgs{
// 			DeviceId: testMetalDevice.ID(),
// 			Type:     pulumi.String("hybrid"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.NewMetalPortVlanAttachment(ctx, "testMetalPortVlanAttachment", &equinix.MetalPortVlanAttachmentArgs{
// 			DeviceId: testMetalDeviceNetworkType.ID(),
// 			PortName: pulumi.String("eth1"),
// 			VlanVnid: testMetalVlan.Vxlan,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Layer 2 network
//
// ```go
// package main
//
// import (
// 	"github.com/cuemby/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testMetalDevice, err := equinix.NewMetalDevice(ctx, "testMetalDevice", &equinix.MetalDeviceArgs{
// 			Hostname: pulumi.String("test"),
// 			Plan:     pulumi.String("c3.small.x86"),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("ny5"),
// 			},
// 			OperatingSystem: pulumi.String("ubuntu_20_04"),
// 			BillingCycle:    pulumi.String("hourly"),
// 			ProjectId:       pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testMetalDeviceNetworkType, err := equinix.NewMetalDeviceNetworkType(ctx, "testMetalDeviceNetworkType", &equinix.MetalDeviceNetworkTypeArgs{
// 			DeviceId: testMetalDevice.ID(),
// 			Type:     pulumi.String("layer2-individual"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		test1MetalVlan, err := equinix.NewMetalVlan(ctx, "test1MetalVlan", &equinix.MetalVlanArgs{
// 			Description: pulumi.String("VLAN in New Jersey"),
// 			Facility:    pulumi.String("ny5"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		test2MetalVlan, err := equinix.NewMetalVlan(ctx, "test2MetalVlan", &equinix.MetalVlanArgs{
// 			Description: pulumi.String("VLAN in New Jersey"),
// 			Facility:    pulumi.String("ny5"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.NewMetalPortVlanAttachment(ctx, "test1MetalPortVlanAttachment", &equinix.MetalPortVlanAttachmentArgs{
// 			DeviceId: testMetalDeviceNetworkType.ID(),
// 			VlanVnid: test1MetalVlan.Vxlan,
// 			PortName: pulumi.String("eth1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.NewMetalPortVlanAttachment(ctx, "test2MetalPortVlanAttachment", &equinix.MetalPortVlanAttachmentArgs{
// 			DeviceId: testMetalDeviceNetworkType.ID(),
// 			VlanVnid: test2MetalVlan.Vxlan,
// 			PortName: pulumi.String("eth1"),
// 			Native:   pulumi.Bool(true),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			pulumi.Resource("equinix_metal_port_vlan_attachment.test1"),
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Attribute Referece
//
// In addition to all arguments above, the following attributes are exported:
//
// * `id` - UUID of device port used in the assignment.
// * `vlanId` - UUID of VLAN API resource.
// * `portId` - UUID of device port.
type MetalPortVlanAttachment struct {
	pulumi.CustomResourceState

	// ID of device to be assigned to the VLAN.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// Add port back to the bond when this resource is removed. Default is
	// `false`.
	ForceBond pulumi.BoolPtrOutput `pulumi:"forceBond"`
	// Mark this VLAN a native VLAN on the port. This can be used only if this
	// assignment assigns second or further VLAN to the port. To ensure that this attachment is not first
	// on a port, you can use `dependsOn` pointing to another `MetalPortVlanAttachment`, just
	// like in the layer2-individual example above.
	Native pulumi.BoolPtrOutput `pulumi:"native"`
	// UUID of device port
	PortId pulumi.StringOutput `pulumi:"portId"`
	// Name of network port to be assigned to the VLAN.
	PortName pulumi.StringOutput `pulumi:"portName"`
	// UUID of VLAN API resource
	VlanId pulumi.StringOutput `pulumi:"vlanId"`
	// VXLAN Network Identifier.
	VlanVnid pulumi.IntOutput `pulumi:"vlanVnid"`
}

// NewMetalPortVlanAttachment registers a new resource with the given unique name, arguments, and options.
func NewMetalPortVlanAttachment(ctx *pulumi.Context,
	name string, args *MetalPortVlanAttachmentArgs, opts ...pulumi.ResourceOption) (*MetalPortVlanAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.PortName == nil {
		return nil, errors.New("invalid value for required argument 'PortName'")
	}
	if args.VlanVnid == nil {
		return nil, errors.New("invalid value for required argument 'VlanVnid'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MetalPortVlanAttachment
	err := ctx.RegisterResource("equinix:index/metalPortVlanAttachment:MetalPortVlanAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalPortVlanAttachment gets an existing MetalPortVlanAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalPortVlanAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalPortVlanAttachmentState, opts ...pulumi.ResourceOption) (*MetalPortVlanAttachment, error) {
	var resource MetalPortVlanAttachment
	err := ctx.ReadResource("equinix:index/metalPortVlanAttachment:MetalPortVlanAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalPortVlanAttachment resources.
type metalPortVlanAttachmentState struct {
	// ID of device to be assigned to the VLAN.
	DeviceId *string `pulumi:"deviceId"`
	// Add port back to the bond when this resource is removed. Default is
	// `false`.
	ForceBond *bool `pulumi:"forceBond"`
	// Mark this VLAN a native VLAN on the port. This can be used only if this
	// assignment assigns second or further VLAN to the port. To ensure that this attachment is not first
	// on a port, you can use `dependsOn` pointing to another `MetalPortVlanAttachment`, just
	// like in the layer2-individual example above.
	Native *bool `pulumi:"native"`
	// UUID of device port
	PortId *string `pulumi:"portId"`
	// Name of network port to be assigned to the VLAN.
	PortName *string `pulumi:"portName"`
	// UUID of VLAN API resource
	VlanId *string `pulumi:"vlanId"`
	// VXLAN Network Identifier.
	VlanVnid *int `pulumi:"vlanVnid"`
}

type MetalPortVlanAttachmentState struct {
	// ID of device to be assigned to the VLAN.
	DeviceId pulumi.StringPtrInput
	// Add port back to the bond when this resource is removed. Default is
	// `false`.
	ForceBond pulumi.BoolPtrInput
	// Mark this VLAN a native VLAN on the port. This can be used only if this
	// assignment assigns second or further VLAN to the port. To ensure that this attachment is not first
	// on a port, you can use `dependsOn` pointing to another `MetalPortVlanAttachment`, just
	// like in the layer2-individual example above.
	Native pulumi.BoolPtrInput
	// UUID of device port
	PortId pulumi.StringPtrInput
	// Name of network port to be assigned to the VLAN.
	PortName pulumi.StringPtrInput
	// UUID of VLAN API resource
	VlanId pulumi.StringPtrInput
	// VXLAN Network Identifier.
	VlanVnid pulumi.IntPtrInput
}

func (MetalPortVlanAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalPortVlanAttachmentState)(nil)).Elem()
}

type metalPortVlanAttachmentArgs struct {
	// ID of device to be assigned to the VLAN.
	DeviceId string `pulumi:"deviceId"`
	// Add port back to the bond when this resource is removed. Default is
	// `false`.
	ForceBond *bool `pulumi:"forceBond"`
	// Mark this VLAN a native VLAN on the port. This can be used only if this
	// assignment assigns second or further VLAN to the port. To ensure that this attachment is not first
	// on a port, you can use `dependsOn` pointing to another `MetalPortVlanAttachment`, just
	// like in the layer2-individual example above.
	Native *bool `pulumi:"native"`
	// Name of network port to be assigned to the VLAN.
	PortName string `pulumi:"portName"`
	// VXLAN Network Identifier.
	VlanVnid int `pulumi:"vlanVnid"`
}

// The set of arguments for constructing a MetalPortVlanAttachment resource.
type MetalPortVlanAttachmentArgs struct {
	// ID of device to be assigned to the VLAN.
	DeviceId pulumi.StringInput
	// Add port back to the bond when this resource is removed. Default is
	// `false`.
	ForceBond pulumi.BoolPtrInput
	// Mark this VLAN a native VLAN on the port. This can be used only if this
	// assignment assigns second or further VLAN to the port. To ensure that this attachment is not first
	// on a port, you can use `dependsOn` pointing to another `MetalPortVlanAttachment`, just
	// like in the layer2-individual example above.
	Native pulumi.BoolPtrInput
	// Name of network port to be assigned to the VLAN.
	PortName pulumi.StringInput
	// VXLAN Network Identifier.
	VlanVnid pulumi.IntInput
}

func (MetalPortVlanAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalPortVlanAttachmentArgs)(nil)).Elem()
}

type MetalPortVlanAttachmentInput interface {
	pulumi.Input

	ToMetalPortVlanAttachmentOutput() MetalPortVlanAttachmentOutput
	ToMetalPortVlanAttachmentOutputWithContext(ctx context.Context) MetalPortVlanAttachmentOutput
}

func (*MetalPortVlanAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalPortVlanAttachment)(nil)).Elem()
}

func (i *MetalPortVlanAttachment) ToMetalPortVlanAttachmentOutput() MetalPortVlanAttachmentOutput {
	return i.ToMetalPortVlanAttachmentOutputWithContext(context.Background())
}

func (i *MetalPortVlanAttachment) ToMetalPortVlanAttachmentOutputWithContext(ctx context.Context) MetalPortVlanAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalPortVlanAttachmentOutput)
}

// MetalPortVlanAttachmentArrayInput is an input type that accepts MetalPortVlanAttachmentArray and MetalPortVlanAttachmentArrayOutput values.
// You can construct a concrete instance of `MetalPortVlanAttachmentArrayInput` via:
//
//          MetalPortVlanAttachmentArray{ MetalPortVlanAttachmentArgs{...} }
type MetalPortVlanAttachmentArrayInput interface {
	pulumi.Input

	ToMetalPortVlanAttachmentArrayOutput() MetalPortVlanAttachmentArrayOutput
	ToMetalPortVlanAttachmentArrayOutputWithContext(context.Context) MetalPortVlanAttachmentArrayOutput
}

type MetalPortVlanAttachmentArray []MetalPortVlanAttachmentInput

func (MetalPortVlanAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalPortVlanAttachment)(nil)).Elem()
}

func (i MetalPortVlanAttachmentArray) ToMetalPortVlanAttachmentArrayOutput() MetalPortVlanAttachmentArrayOutput {
	return i.ToMetalPortVlanAttachmentArrayOutputWithContext(context.Background())
}

func (i MetalPortVlanAttachmentArray) ToMetalPortVlanAttachmentArrayOutputWithContext(ctx context.Context) MetalPortVlanAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalPortVlanAttachmentArrayOutput)
}

// MetalPortVlanAttachmentMapInput is an input type that accepts MetalPortVlanAttachmentMap and MetalPortVlanAttachmentMapOutput values.
// You can construct a concrete instance of `MetalPortVlanAttachmentMapInput` via:
//
//          MetalPortVlanAttachmentMap{ "key": MetalPortVlanAttachmentArgs{...} }
type MetalPortVlanAttachmentMapInput interface {
	pulumi.Input

	ToMetalPortVlanAttachmentMapOutput() MetalPortVlanAttachmentMapOutput
	ToMetalPortVlanAttachmentMapOutputWithContext(context.Context) MetalPortVlanAttachmentMapOutput
}

type MetalPortVlanAttachmentMap map[string]MetalPortVlanAttachmentInput

func (MetalPortVlanAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalPortVlanAttachment)(nil)).Elem()
}

func (i MetalPortVlanAttachmentMap) ToMetalPortVlanAttachmentMapOutput() MetalPortVlanAttachmentMapOutput {
	return i.ToMetalPortVlanAttachmentMapOutputWithContext(context.Background())
}

func (i MetalPortVlanAttachmentMap) ToMetalPortVlanAttachmentMapOutputWithContext(ctx context.Context) MetalPortVlanAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalPortVlanAttachmentMapOutput)
}

type MetalPortVlanAttachmentOutput struct{ *pulumi.OutputState }

func (MetalPortVlanAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalPortVlanAttachment)(nil)).Elem()
}

func (o MetalPortVlanAttachmentOutput) ToMetalPortVlanAttachmentOutput() MetalPortVlanAttachmentOutput {
	return o
}

func (o MetalPortVlanAttachmentOutput) ToMetalPortVlanAttachmentOutputWithContext(ctx context.Context) MetalPortVlanAttachmentOutput {
	return o
}

// ID of device to be assigned to the VLAN.
func (o MetalPortVlanAttachmentOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalPortVlanAttachment) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// Add port back to the bond when this resource is removed. Default is
// `false`.
func (o MetalPortVlanAttachmentOutput) ForceBond() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MetalPortVlanAttachment) pulumi.BoolPtrOutput { return v.ForceBond }).(pulumi.BoolPtrOutput)
}

// Mark this VLAN a native VLAN on the port. This can be used only if this
// assignment assigns second or further VLAN to the port. To ensure that this attachment is not first
// on a port, you can use `dependsOn` pointing to another `MetalPortVlanAttachment`, just
// like in the layer2-individual example above.
func (o MetalPortVlanAttachmentOutput) Native() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MetalPortVlanAttachment) pulumi.BoolPtrOutput { return v.Native }).(pulumi.BoolPtrOutput)
}

// UUID of device port
func (o MetalPortVlanAttachmentOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalPortVlanAttachment) pulumi.StringOutput { return v.PortId }).(pulumi.StringOutput)
}

// Name of network port to be assigned to the VLAN.
func (o MetalPortVlanAttachmentOutput) PortName() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalPortVlanAttachment) pulumi.StringOutput { return v.PortName }).(pulumi.StringOutput)
}

// UUID of VLAN API resource
func (o MetalPortVlanAttachmentOutput) VlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalPortVlanAttachment) pulumi.StringOutput { return v.VlanId }).(pulumi.StringOutput)
}

// VXLAN Network Identifier.
func (o MetalPortVlanAttachmentOutput) VlanVnid() pulumi.IntOutput {
	return o.ApplyT(func(v *MetalPortVlanAttachment) pulumi.IntOutput { return v.VlanVnid }).(pulumi.IntOutput)
}

type MetalPortVlanAttachmentArrayOutput struct{ *pulumi.OutputState }

func (MetalPortVlanAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalPortVlanAttachment)(nil)).Elem()
}

func (o MetalPortVlanAttachmentArrayOutput) ToMetalPortVlanAttachmentArrayOutput() MetalPortVlanAttachmentArrayOutput {
	return o
}

func (o MetalPortVlanAttachmentArrayOutput) ToMetalPortVlanAttachmentArrayOutputWithContext(ctx context.Context) MetalPortVlanAttachmentArrayOutput {
	return o
}

func (o MetalPortVlanAttachmentArrayOutput) Index(i pulumi.IntInput) MetalPortVlanAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalPortVlanAttachment {
		return vs[0].([]*MetalPortVlanAttachment)[vs[1].(int)]
	}).(MetalPortVlanAttachmentOutput)
}

type MetalPortVlanAttachmentMapOutput struct{ *pulumi.OutputState }

func (MetalPortVlanAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalPortVlanAttachment)(nil)).Elem()
}

func (o MetalPortVlanAttachmentMapOutput) ToMetalPortVlanAttachmentMapOutput() MetalPortVlanAttachmentMapOutput {
	return o
}

func (o MetalPortVlanAttachmentMapOutput) ToMetalPortVlanAttachmentMapOutputWithContext(ctx context.Context) MetalPortVlanAttachmentMapOutput {
	return o
}

func (o MetalPortVlanAttachmentMapOutput) MapIndex(k pulumi.StringInput) MetalPortVlanAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalPortVlanAttachment {
		return vs[0].(map[string]*MetalPortVlanAttachment)[vs[1].(string)]
	}).(MetalPortVlanAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalPortVlanAttachmentInput)(nil)).Elem(), &MetalPortVlanAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalPortVlanAttachmentArrayInput)(nil)).Elem(), MetalPortVlanAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalPortVlanAttachmentMapInput)(nil)).Elem(), MetalPortVlanAttachmentMap{})
	pulumi.RegisterOutputType(MetalPortVlanAttachmentOutput{})
	pulumi.RegisterOutputType(MetalPortVlanAttachmentArrayOutput{})
	pulumi.RegisterOutputType(MetalPortVlanAttachmentMapOutput{})
}
