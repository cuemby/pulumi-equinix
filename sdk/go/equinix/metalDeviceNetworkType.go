// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// See the Network Types Guide for examples of this resource and to
// learn about the recommended `MetalPort` alternative.
//
// ## Import
//
// This resource can also be imported using existing device ID
//
// ```sh
//  $ pulumi import equinix:index/metalDeviceNetworkType:MetalDeviceNetworkType equinix_metal_device_network_type {existing device_id}
// ```
type MetalDeviceNetworkType struct {
	pulumi.CustomResourceState

	// The ID of the device on which the network type should be set.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual`
	// and `layer2-bonded`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMetalDeviceNetworkType registers a new resource with the given unique name, arguments, and options.
func NewMetalDeviceNetworkType(ctx *pulumi.Context,
	name string, args *MetalDeviceNetworkTypeArgs, opts ...pulumi.ResourceOption) (*MetalDeviceNetworkType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MetalDeviceNetworkType
	err := ctx.RegisterResource("equinix:index/metalDeviceNetworkType:MetalDeviceNetworkType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalDeviceNetworkType gets an existing MetalDeviceNetworkType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalDeviceNetworkType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalDeviceNetworkTypeState, opts ...pulumi.ResourceOption) (*MetalDeviceNetworkType, error) {
	var resource MetalDeviceNetworkType
	err := ctx.ReadResource("equinix:index/metalDeviceNetworkType:MetalDeviceNetworkType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalDeviceNetworkType resources.
type metalDeviceNetworkTypeState struct {
	// The ID of the device on which the network type should be set.
	DeviceId *string `pulumi:"deviceId"`
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual`
	// and `layer2-bonded`.
	Type *string `pulumi:"type"`
}

type MetalDeviceNetworkTypeState struct {
	// The ID of the device on which the network type should be set.
	DeviceId pulumi.StringPtrInput
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual`
	// and `layer2-bonded`.
	Type pulumi.StringPtrInput
}

func (MetalDeviceNetworkTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalDeviceNetworkTypeState)(nil)).Elem()
}

type metalDeviceNetworkTypeArgs struct {
	// The ID of the device on which the network type should be set.
	DeviceId string `pulumi:"deviceId"`
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual`
	// and `layer2-bonded`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a MetalDeviceNetworkType resource.
type MetalDeviceNetworkTypeArgs struct {
	// The ID of the device on which the network type should be set.
	DeviceId pulumi.StringInput
	// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual`
	// and `layer2-bonded`.
	Type pulumi.StringInput
}

func (MetalDeviceNetworkTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalDeviceNetworkTypeArgs)(nil)).Elem()
}

type MetalDeviceNetworkTypeInput interface {
	pulumi.Input

	ToMetalDeviceNetworkTypeOutput() MetalDeviceNetworkTypeOutput
	ToMetalDeviceNetworkTypeOutputWithContext(ctx context.Context) MetalDeviceNetworkTypeOutput
}

func (*MetalDeviceNetworkType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalDeviceNetworkType)(nil)).Elem()
}

func (i *MetalDeviceNetworkType) ToMetalDeviceNetworkTypeOutput() MetalDeviceNetworkTypeOutput {
	return i.ToMetalDeviceNetworkTypeOutputWithContext(context.Background())
}

func (i *MetalDeviceNetworkType) ToMetalDeviceNetworkTypeOutputWithContext(ctx context.Context) MetalDeviceNetworkTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalDeviceNetworkTypeOutput)
}

// MetalDeviceNetworkTypeArrayInput is an input type that accepts MetalDeviceNetworkTypeArray and MetalDeviceNetworkTypeArrayOutput values.
// You can construct a concrete instance of `MetalDeviceNetworkTypeArrayInput` via:
//
//          MetalDeviceNetworkTypeArray{ MetalDeviceNetworkTypeArgs{...} }
type MetalDeviceNetworkTypeArrayInput interface {
	pulumi.Input

	ToMetalDeviceNetworkTypeArrayOutput() MetalDeviceNetworkTypeArrayOutput
	ToMetalDeviceNetworkTypeArrayOutputWithContext(context.Context) MetalDeviceNetworkTypeArrayOutput
}

type MetalDeviceNetworkTypeArray []MetalDeviceNetworkTypeInput

func (MetalDeviceNetworkTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalDeviceNetworkType)(nil)).Elem()
}

func (i MetalDeviceNetworkTypeArray) ToMetalDeviceNetworkTypeArrayOutput() MetalDeviceNetworkTypeArrayOutput {
	return i.ToMetalDeviceNetworkTypeArrayOutputWithContext(context.Background())
}

func (i MetalDeviceNetworkTypeArray) ToMetalDeviceNetworkTypeArrayOutputWithContext(ctx context.Context) MetalDeviceNetworkTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalDeviceNetworkTypeArrayOutput)
}

// MetalDeviceNetworkTypeMapInput is an input type that accepts MetalDeviceNetworkTypeMap and MetalDeviceNetworkTypeMapOutput values.
// You can construct a concrete instance of `MetalDeviceNetworkTypeMapInput` via:
//
//          MetalDeviceNetworkTypeMap{ "key": MetalDeviceNetworkTypeArgs{...} }
type MetalDeviceNetworkTypeMapInput interface {
	pulumi.Input

	ToMetalDeviceNetworkTypeMapOutput() MetalDeviceNetworkTypeMapOutput
	ToMetalDeviceNetworkTypeMapOutputWithContext(context.Context) MetalDeviceNetworkTypeMapOutput
}

type MetalDeviceNetworkTypeMap map[string]MetalDeviceNetworkTypeInput

func (MetalDeviceNetworkTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalDeviceNetworkType)(nil)).Elem()
}

func (i MetalDeviceNetworkTypeMap) ToMetalDeviceNetworkTypeMapOutput() MetalDeviceNetworkTypeMapOutput {
	return i.ToMetalDeviceNetworkTypeMapOutputWithContext(context.Background())
}

func (i MetalDeviceNetworkTypeMap) ToMetalDeviceNetworkTypeMapOutputWithContext(ctx context.Context) MetalDeviceNetworkTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalDeviceNetworkTypeMapOutput)
}

type MetalDeviceNetworkTypeOutput struct{ *pulumi.OutputState }

func (MetalDeviceNetworkTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalDeviceNetworkType)(nil)).Elem()
}

func (o MetalDeviceNetworkTypeOutput) ToMetalDeviceNetworkTypeOutput() MetalDeviceNetworkTypeOutput {
	return o
}

func (o MetalDeviceNetworkTypeOutput) ToMetalDeviceNetworkTypeOutputWithContext(ctx context.Context) MetalDeviceNetworkTypeOutput {
	return o
}

// The ID of the device on which the network type should be set.
func (o MetalDeviceNetworkTypeOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalDeviceNetworkType) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// Network type to set. Must be one of `layer3`, `hybrid`, `layer2-individual`
// and `layer2-bonded`.
func (o MetalDeviceNetworkTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalDeviceNetworkType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type MetalDeviceNetworkTypeArrayOutput struct{ *pulumi.OutputState }

func (MetalDeviceNetworkTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalDeviceNetworkType)(nil)).Elem()
}

func (o MetalDeviceNetworkTypeArrayOutput) ToMetalDeviceNetworkTypeArrayOutput() MetalDeviceNetworkTypeArrayOutput {
	return o
}

func (o MetalDeviceNetworkTypeArrayOutput) ToMetalDeviceNetworkTypeArrayOutputWithContext(ctx context.Context) MetalDeviceNetworkTypeArrayOutput {
	return o
}

func (o MetalDeviceNetworkTypeArrayOutput) Index(i pulumi.IntInput) MetalDeviceNetworkTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalDeviceNetworkType {
		return vs[0].([]*MetalDeviceNetworkType)[vs[1].(int)]
	}).(MetalDeviceNetworkTypeOutput)
}

type MetalDeviceNetworkTypeMapOutput struct{ *pulumi.OutputState }

func (MetalDeviceNetworkTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalDeviceNetworkType)(nil)).Elem()
}

func (o MetalDeviceNetworkTypeMapOutput) ToMetalDeviceNetworkTypeMapOutput() MetalDeviceNetworkTypeMapOutput {
	return o
}

func (o MetalDeviceNetworkTypeMapOutput) ToMetalDeviceNetworkTypeMapOutputWithContext(ctx context.Context) MetalDeviceNetworkTypeMapOutput {
	return o
}

func (o MetalDeviceNetworkTypeMapOutput) MapIndex(k pulumi.StringInput) MetalDeviceNetworkTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalDeviceNetworkType {
		return vs[0].(map[string]*MetalDeviceNetworkType)[vs[1].(string)]
	}).(MetalDeviceNetworkTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalDeviceNetworkTypeInput)(nil)).Elem(), &MetalDeviceNetworkType{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalDeviceNetworkTypeArrayInput)(nil)).Elem(), MetalDeviceNetworkTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalDeviceNetworkTypeMapInput)(nil)).Elem(), MetalDeviceNetworkTypeMap{})
	pulumi.RegisterOutputType(MetalDeviceNetworkTypeOutput{})
	pulumi.RegisterOutputType(MetalDeviceNetworkTypeArrayOutput{})
	pulumi.RegisterOutputType(MetalDeviceNetworkTypeMapOutput{})
}
