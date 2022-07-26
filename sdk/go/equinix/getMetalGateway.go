// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
//
// > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
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
// 		_, err := equinix.NewMetalVlan(ctx, "testMetalVlan", &equinix.MetalVlanArgs{
// 			Description: pulumi.String("test VLAN in SV"),
// 			Metro:       pulumi.String("sv"),
// 			ProjectId:   pulumi.Any(local.Project_id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.GetMetalGateway(ctx, &GetMetalGatewayArgs{
// 			GatewayId: local.Gateway_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMetalGateway(ctx *pulumi.Context, args *LookupMetalGatewayArgs, opts ...pulumi.InvokeOption) (*LookupMetalGatewayResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMetalGatewayResult
	err := ctx.Invoke("equinix:index/getMetalGateway:GetMetalGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalGateway.
type LookupMetalGatewayArgs struct {
	// UUID of the metal gateway resource to retrieve.
	GatewayId string `pulumi:"gatewayId"`
}

// A collection of values returned by GetMetalGateway.
type LookupMetalGatewayResult struct {
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// UUID of IP reservation block bound to the gateway.
	IpReservationId string `pulumi:"ipReservationId"`
	// Size of the private IPv4 subnet bound to this metal gateway. One of
	// `8`, `16`, `32`, `64`, `128`.
	PrivateIpv4SubnetSize int `pulumi:"privateIpv4SubnetSize"`
	// UUID of the project where the gateway is scoped to.
	ProjectId string `pulumi:"projectId"`
	// Status of the gateway resource.
	State string `pulumi:"state"`
	// UUID of the VLAN where the gateway is scoped to.
	VlanId string `pulumi:"vlanId"`
	// UUID of the VRF associated with the IP Reservation.
	VrfId string `pulumi:"vrfId"`
}

func LookupMetalGatewayOutput(ctx *pulumi.Context, args LookupMetalGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupMetalGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetalGatewayResult, error) {
			args := v.(LookupMetalGatewayArgs)
			r, err := LookupMetalGateway(ctx, &args, opts...)
			var s LookupMetalGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetalGatewayResultOutput)
}

// A collection of arguments for invoking GetMetalGateway.
type LookupMetalGatewayOutputArgs struct {
	// UUID of the metal gateway resource to retrieve.
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
}

func (LookupMetalGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalGatewayArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalGateway.
type LookupMetalGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupMetalGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalGatewayResult)(nil)).Elem()
}

func (o LookupMetalGatewayResultOutput) ToLookupMetalGatewayResultOutput() LookupMetalGatewayResultOutput {
	return o
}

func (o LookupMetalGatewayResultOutput) ToLookupMetalGatewayResultOutputWithContext(ctx context.Context) LookupMetalGatewayResultOutput {
	return o
}

func (o LookupMetalGatewayResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMetalGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// UUID of IP reservation block bound to the gateway.
func (o LookupMetalGatewayResultOutput) IpReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) string { return v.IpReservationId }).(pulumi.StringOutput)
}

// Size of the private IPv4 subnet bound to this metal gateway. One of
// `8`, `16`, `32`, `64`, `128`.
func (o LookupMetalGatewayResultOutput) PrivateIpv4SubnetSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) int { return v.PrivateIpv4SubnetSize }).(pulumi.IntOutput)
}

// UUID of the project where the gateway is scoped to.
func (o LookupMetalGatewayResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Status of the gateway resource.
func (o LookupMetalGatewayResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) string { return v.State }).(pulumi.StringOutput)
}

// UUID of the VLAN where the gateway is scoped to.
func (o LookupMetalGatewayResultOutput) VlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) string { return v.VlanId }).(pulumi.StringOutput)
}

// UUID of the VRF associated with the IP Reservation.
func (o LookupMetalGatewayResultOutput) VrfId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalGatewayResult) string { return v.VrfId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetalGatewayResultOutput{})
}
