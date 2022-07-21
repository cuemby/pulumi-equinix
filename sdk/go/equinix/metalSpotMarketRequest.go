// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Equinix Metal Spot Market Request resource to allow you to
// manage spot market requests on your account. For more detail on Spot Market,
// see [this article in Equinix Metal documentation](https://metal.equinix.com/developers/docs/deploy/spot-market/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/cuemby/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := equinix.NewMetalSpotMarketRequest(ctx, "req", &equinix.MetalSpotMarketRequestArgs{
// 			ProjectId:   pulumi.Any(local.Project_id),
// 			MaxBidPrice: pulumi.Float64(0.03),
// 			Facilities: pulumi.StringArray{
// 				pulumi.String("ny5"),
// 			},
// 			DevicesMin: pulumi.Int(1),
// 			DevicesMax: pulumi.Int(1),
// 			InstanceParameters: &MetalSpotMarketRequestInstanceParametersArgs{
// 				Hostname:        pulumi.String("testspot"),
// 				BillingCycle:    pulumi.String("hourly"),
// 				OperatingSystem: pulumi.String("ubuntu_20_04"),
// 				Plan:            pulumi.String("c3.small.x86"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// This resource can be imported using an existing spot market request ID
//
// ```sh
//  $ pulumi import equinix:index/metalSpotMarketRequest:MetalSpotMarketRequest equinix_metal_spot_market_request {existing_spot_market_request_id}
// ```
type MetalSpotMarketRequest struct {
	pulumi.CustomResourceState

	// Maximum number devices to be created.
	DevicesMax pulumi.IntOutput `pulumi:"devicesMax"`
	// Miniumum number devices to be created.
	DevicesMin pulumi.IntOutput `pulumi:"devicesMin"`
	// Facility IDs where devices should be created.
	Facilities pulumi.StringArrayOutput `pulumi:"facilities"`
	// Key/Value pairs of parameters for devices provisioned from
	// this request. Valid keys are: `billingCycle`, `plan`, `operatingSystem`, `hostname`,
	// `termintationTime`, `alwaysPxe`, `description`, `features`, `locked`, `projectSshKeys`,
	// `userSshKeys`, `userdata`, `customdata`, `ipxeScriptUrl`, `tags`. You can find each parameter
	// description in MetalDevice docs.
	InstanceParameters MetalSpotMarketRequestInstanceParametersOutput `pulumi:"instanceParameters"`
	// Maximum price user is willing to pay per hour per device.
	MaxBidPrice pulumi.Float64Output `pulumi:"maxBidPrice"`
	// Metro where devices should be created.
	Metro pulumi.StringPtrOutput `pulumi:"metro"`
	// Project ID.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// On resource creation wait until all desired devices are active.
	// On resource destruction wait until devices are removed.
	WaitForDevices pulumi.BoolPtrOutput `pulumi:"waitForDevices"`
}

// NewMetalSpotMarketRequest registers a new resource with the given unique name, arguments, and options.
func NewMetalSpotMarketRequest(ctx *pulumi.Context,
	name string, args *MetalSpotMarketRequestArgs, opts ...pulumi.ResourceOption) (*MetalSpotMarketRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DevicesMax == nil {
		return nil, errors.New("invalid value for required argument 'DevicesMax'")
	}
	if args.DevicesMin == nil {
		return nil, errors.New("invalid value for required argument 'DevicesMin'")
	}
	if args.InstanceParameters == nil {
		return nil, errors.New("invalid value for required argument 'InstanceParameters'")
	}
	if args.MaxBidPrice == nil {
		return nil, errors.New("invalid value for required argument 'MaxBidPrice'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MetalSpotMarketRequest
	err := ctx.RegisterResource("equinix:index/metalSpotMarketRequest:MetalSpotMarketRequest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalSpotMarketRequest gets an existing MetalSpotMarketRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalSpotMarketRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalSpotMarketRequestState, opts ...pulumi.ResourceOption) (*MetalSpotMarketRequest, error) {
	var resource MetalSpotMarketRequest
	err := ctx.ReadResource("equinix:index/metalSpotMarketRequest:MetalSpotMarketRequest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalSpotMarketRequest resources.
type metalSpotMarketRequestState struct {
	// Maximum number devices to be created.
	DevicesMax *int `pulumi:"devicesMax"`
	// Miniumum number devices to be created.
	DevicesMin *int `pulumi:"devicesMin"`
	// Facility IDs where devices should be created.
	Facilities []string `pulumi:"facilities"`
	// Key/Value pairs of parameters for devices provisioned from
	// this request. Valid keys are: `billingCycle`, `plan`, `operatingSystem`, `hostname`,
	// `termintationTime`, `alwaysPxe`, `description`, `features`, `locked`, `projectSshKeys`,
	// `userSshKeys`, `userdata`, `customdata`, `ipxeScriptUrl`, `tags`. You can find each parameter
	// description in MetalDevice docs.
	InstanceParameters *MetalSpotMarketRequestInstanceParameters `pulumi:"instanceParameters"`
	// Maximum price user is willing to pay per hour per device.
	MaxBidPrice *float64 `pulumi:"maxBidPrice"`
	// Metro where devices should be created.
	Metro *string `pulumi:"metro"`
	// Project ID.
	ProjectId *string `pulumi:"projectId"`
	// On resource creation wait until all desired devices are active.
	// On resource destruction wait until devices are removed.
	WaitForDevices *bool `pulumi:"waitForDevices"`
}

type MetalSpotMarketRequestState struct {
	// Maximum number devices to be created.
	DevicesMax pulumi.IntPtrInput
	// Miniumum number devices to be created.
	DevicesMin pulumi.IntPtrInput
	// Facility IDs where devices should be created.
	Facilities pulumi.StringArrayInput
	// Key/Value pairs of parameters for devices provisioned from
	// this request. Valid keys are: `billingCycle`, `plan`, `operatingSystem`, `hostname`,
	// `termintationTime`, `alwaysPxe`, `description`, `features`, `locked`, `projectSshKeys`,
	// `userSshKeys`, `userdata`, `customdata`, `ipxeScriptUrl`, `tags`. You can find each parameter
	// description in MetalDevice docs.
	InstanceParameters MetalSpotMarketRequestInstanceParametersPtrInput
	// Maximum price user is willing to pay per hour per device.
	MaxBidPrice pulumi.Float64PtrInput
	// Metro where devices should be created.
	Metro pulumi.StringPtrInput
	// Project ID.
	ProjectId pulumi.StringPtrInput
	// On resource creation wait until all desired devices are active.
	// On resource destruction wait until devices are removed.
	WaitForDevices pulumi.BoolPtrInput
}

func (MetalSpotMarketRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalSpotMarketRequestState)(nil)).Elem()
}

type metalSpotMarketRequestArgs struct {
	// Maximum number devices to be created.
	DevicesMax int `pulumi:"devicesMax"`
	// Miniumum number devices to be created.
	DevicesMin int `pulumi:"devicesMin"`
	// Facility IDs where devices should be created.
	Facilities []string `pulumi:"facilities"`
	// Key/Value pairs of parameters for devices provisioned from
	// this request. Valid keys are: `billingCycle`, `plan`, `operatingSystem`, `hostname`,
	// `termintationTime`, `alwaysPxe`, `description`, `features`, `locked`, `projectSshKeys`,
	// `userSshKeys`, `userdata`, `customdata`, `ipxeScriptUrl`, `tags`. You can find each parameter
	// description in MetalDevice docs.
	InstanceParameters MetalSpotMarketRequestInstanceParameters `pulumi:"instanceParameters"`
	// Maximum price user is willing to pay per hour per device.
	MaxBidPrice float64 `pulumi:"maxBidPrice"`
	// Metro where devices should be created.
	Metro *string `pulumi:"metro"`
	// Project ID.
	ProjectId string `pulumi:"projectId"`
	// On resource creation wait until all desired devices are active.
	// On resource destruction wait until devices are removed.
	WaitForDevices *bool `pulumi:"waitForDevices"`
}

// The set of arguments for constructing a MetalSpotMarketRequest resource.
type MetalSpotMarketRequestArgs struct {
	// Maximum number devices to be created.
	DevicesMax pulumi.IntInput
	// Miniumum number devices to be created.
	DevicesMin pulumi.IntInput
	// Facility IDs where devices should be created.
	Facilities pulumi.StringArrayInput
	// Key/Value pairs of parameters for devices provisioned from
	// this request. Valid keys are: `billingCycle`, `plan`, `operatingSystem`, `hostname`,
	// `termintationTime`, `alwaysPxe`, `description`, `features`, `locked`, `projectSshKeys`,
	// `userSshKeys`, `userdata`, `customdata`, `ipxeScriptUrl`, `tags`. You can find each parameter
	// description in MetalDevice docs.
	InstanceParameters MetalSpotMarketRequestInstanceParametersInput
	// Maximum price user is willing to pay per hour per device.
	MaxBidPrice pulumi.Float64Input
	// Metro where devices should be created.
	Metro pulumi.StringPtrInput
	// Project ID.
	ProjectId pulumi.StringInput
	// On resource creation wait until all desired devices are active.
	// On resource destruction wait until devices are removed.
	WaitForDevices pulumi.BoolPtrInput
}

func (MetalSpotMarketRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalSpotMarketRequestArgs)(nil)).Elem()
}

type MetalSpotMarketRequestInput interface {
	pulumi.Input

	ToMetalSpotMarketRequestOutput() MetalSpotMarketRequestOutput
	ToMetalSpotMarketRequestOutputWithContext(ctx context.Context) MetalSpotMarketRequestOutput
}

func (*MetalSpotMarketRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalSpotMarketRequest)(nil)).Elem()
}

func (i *MetalSpotMarketRequest) ToMetalSpotMarketRequestOutput() MetalSpotMarketRequestOutput {
	return i.ToMetalSpotMarketRequestOutputWithContext(context.Background())
}

func (i *MetalSpotMarketRequest) ToMetalSpotMarketRequestOutputWithContext(ctx context.Context) MetalSpotMarketRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalSpotMarketRequestOutput)
}

// MetalSpotMarketRequestArrayInput is an input type that accepts MetalSpotMarketRequestArray and MetalSpotMarketRequestArrayOutput values.
// You can construct a concrete instance of `MetalSpotMarketRequestArrayInput` via:
//
//          MetalSpotMarketRequestArray{ MetalSpotMarketRequestArgs{...} }
type MetalSpotMarketRequestArrayInput interface {
	pulumi.Input

	ToMetalSpotMarketRequestArrayOutput() MetalSpotMarketRequestArrayOutput
	ToMetalSpotMarketRequestArrayOutputWithContext(context.Context) MetalSpotMarketRequestArrayOutput
}

type MetalSpotMarketRequestArray []MetalSpotMarketRequestInput

func (MetalSpotMarketRequestArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalSpotMarketRequest)(nil)).Elem()
}

func (i MetalSpotMarketRequestArray) ToMetalSpotMarketRequestArrayOutput() MetalSpotMarketRequestArrayOutput {
	return i.ToMetalSpotMarketRequestArrayOutputWithContext(context.Background())
}

func (i MetalSpotMarketRequestArray) ToMetalSpotMarketRequestArrayOutputWithContext(ctx context.Context) MetalSpotMarketRequestArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalSpotMarketRequestArrayOutput)
}

// MetalSpotMarketRequestMapInput is an input type that accepts MetalSpotMarketRequestMap and MetalSpotMarketRequestMapOutput values.
// You can construct a concrete instance of `MetalSpotMarketRequestMapInput` via:
//
//          MetalSpotMarketRequestMap{ "key": MetalSpotMarketRequestArgs{...} }
type MetalSpotMarketRequestMapInput interface {
	pulumi.Input

	ToMetalSpotMarketRequestMapOutput() MetalSpotMarketRequestMapOutput
	ToMetalSpotMarketRequestMapOutputWithContext(context.Context) MetalSpotMarketRequestMapOutput
}

type MetalSpotMarketRequestMap map[string]MetalSpotMarketRequestInput

func (MetalSpotMarketRequestMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalSpotMarketRequest)(nil)).Elem()
}

func (i MetalSpotMarketRequestMap) ToMetalSpotMarketRequestMapOutput() MetalSpotMarketRequestMapOutput {
	return i.ToMetalSpotMarketRequestMapOutputWithContext(context.Background())
}

func (i MetalSpotMarketRequestMap) ToMetalSpotMarketRequestMapOutputWithContext(ctx context.Context) MetalSpotMarketRequestMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalSpotMarketRequestMapOutput)
}

type MetalSpotMarketRequestOutput struct{ *pulumi.OutputState }

func (MetalSpotMarketRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalSpotMarketRequest)(nil)).Elem()
}

func (o MetalSpotMarketRequestOutput) ToMetalSpotMarketRequestOutput() MetalSpotMarketRequestOutput {
	return o
}

func (o MetalSpotMarketRequestOutput) ToMetalSpotMarketRequestOutputWithContext(ctx context.Context) MetalSpotMarketRequestOutput {
	return o
}

// Maximum number devices to be created.
func (o MetalSpotMarketRequestOutput) DevicesMax() pulumi.IntOutput {
	return o.ApplyT(func(v *MetalSpotMarketRequest) pulumi.IntOutput { return v.DevicesMax }).(pulumi.IntOutput)
}

// Miniumum number devices to be created.
func (o MetalSpotMarketRequestOutput) DevicesMin() pulumi.IntOutput {
	return o.ApplyT(func(v *MetalSpotMarketRequest) pulumi.IntOutput { return v.DevicesMin }).(pulumi.IntOutput)
}

// Facility IDs where devices should be created.
func (o MetalSpotMarketRequestOutput) Facilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetalSpotMarketRequest) pulumi.StringArrayOutput { return v.Facilities }).(pulumi.StringArrayOutput)
}

// Key/Value pairs of parameters for devices provisioned from
// this request. Valid keys are: `billingCycle`, `plan`, `operatingSystem`, `hostname`,
// `termintationTime`, `alwaysPxe`, `description`, `features`, `locked`, `projectSshKeys`,
// `userSshKeys`, `userdata`, `customdata`, `ipxeScriptUrl`, `tags`. You can find each parameter
// description in MetalDevice docs.
func (o MetalSpotMarketRequestOutput) InstanceParameters() MetalSpotMarketRequestInstanceParametersOutput {
	return o.ApplyT(func(v *MetalSpotMarketRequest) MetalSpotMarketRequestInstanceParametersOutput {
		return v.InstanceParameters
	}).(MetalSpotMarketRequestInstanceParametersOutput)
}

// Maximum price user is willing to pay per hour per device.
func (o MetalSpotMarketRequestOutput) MaxBidPrice() pulumi.Float64Output {
	return o.ApplyT(func(v *MetalSpotMarketRequest) pulumi.Float64Output { return v.MaxBidPrice }).(pulumi.Float64Output)
}

// Metro where devices should be created.
func (o MetalSpotMarketRequestOutput) Metro() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalSpotMarketRequest) pulumi.StringPtrOutput { return v.Metro }).(pulumi.StringPtrOutput)
}

// Project ID.
func (o MetalSpotMarketRequestOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalSpotMarketRequest) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// On resource creation wait until all desired devices are active.
// On resource destruction wait until devices are removed.
func (o MetalSpotMarketRequestOutput) WaitForDevices() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MetalSpotMarketRequest) pulumi.BoolPtrOutput { return v.WaitForDevices }).(pulumi.BoolPtrOutput)
}

type MetalSpotMarketRequestArrayOutput struct{ *pulumi.OutputState }

func (MetalSpotMarketRequestArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalSpotMarketRequest)(nil)).Elem()
}

func (o MetalSpotMarketRequestArrayOutput) ToMetalSpotMarketRequestArrayOutput() MetalSpotMarketRequestArrayOutput {
	return o
}

func (o MetalSpotMarketRequestArrayOutput) ToMetalSpotMarketRequestArrayOutputWithContext(ctx context.Context) MetalSpotMarketRequestArrayOutput {
	return o
}

func (o MetalSpotMarketRequestArrayOutput) Index(i pulumi.IntInput) MetalSpotMarketRequestOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalSpotMarketRequest {
		return vs[0].([]*MetalSpotMarketRequest)[vs[1].(int)]
	}).(MetalSpotMarketRequestOutput)
}

type MetalSpotMarketRequestMapOutput struct{ *pulumi.OutputState }

func (MetalSpotMarketRequestMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalSpotMarketRequest)(nil)).Elem()
}

func (o MetalSpotMarketRequestMapOutput) ToMetalSpotMarketRequestMapOutput() MetalSpotMarketRequestMapOutput {
	return o
}

func (o MetalSpotMarketRequestMapOutput) ToMetalSpotMarketRequestMapOutputWithContext(ctx context.Context) MetalSpotMarketRequestMapOutput {
	return o
}

func (o MetalSpotMarketRequestMapOutput) MapIndex(k pulumi.StringInput) MetalSpotMarketRequestOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalSpotMarketRequest {
		return vs[0].(map[string]*MetalSpotMarketRequest)[vs[1].(string)]
	}).(MetalSpotMarketRequestOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalSpotMarketRequestInput)(nil)).Elem(), &MetalSpotMarketRequest{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalSpotMarketRequestArrayInput)(nil)).Elem(), MetalSpotMarketRequestArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalSpotMarketRequestMapInput)(nil)).Elem(), MetalSpotMarketRequestMap{})
	pulumi.RegisterOutputType(MetalSpotMarketRequestOutput{})
	pulumi.RegisterOutputType(MetalSpotMarketRequestArrayOutput{})
	pulumi.RegisterOutputType(MetalSpotMarketRequestMapOutput{})
}
