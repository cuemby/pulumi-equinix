// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this datasource to retrieve list of BGP neighbors of a device in the Equinix Metal host.
//
// To have any BGP neighbors listed, the device must be in BGP-enabled project
// and have a BGP session assigned.
//
// To learn more about using BGP in Equinix Metal, see the
// MetalBGPSession resource documentation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		test, err := equinix.GetMetalDeviceBGPNeighbors(ctx, &GetMetalDeviceBGPNeighborsArgs{
// 			DeviceId: "4c641195-25e5-4c3c-b2b7-4cd7a42c7b40",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("bgpNeighborsListing", test.BgpNeighbors)
// 		return nil
// 	})
// }
// ```
func GetMetalDeviceBGPNeighbors(ctx *pulumi.Context, args *GetMetalDeviceBGPNeighborsArgs, opts ...pulumi.InvokeOption) (*GetMetalDeviceBGPNeighborsResult, error) {
	var rv GetMetalDeviceBGPNeighborsResult
	err := ctx.Invoke("equinix:index/getMetalDeviceBGPNeighbors:GetMetalDeviceBGPNeighbors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalDeviceBGPNeighbors.
type GetMetalDeviceBGPNeighborsArgs struct {
	// UUID of BGP-enabled device whose neighbors to list.
	DeviceId string `pulumi:"deviceId"`
}

// A collection of values returned by GetMetalDeviceBGPNeighbors.
type GetMetalDeviceBGPNeighborsResult struct {
	// array of BGP neighbor records with attributes:
	BgpNeighbors []GetMetalDeviceBGPNeighborsBgpNeighbor `pulumi:"bgpNeighbors"`
	DeviceId     string                                  `pulumi:"deviceId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetMetalDeviceBGPNeighborsOutput(ctx *pulumi.Context, args GetMetalDeviceBGPNeighborsOutputArgs, opts ...pulumi.InvokeOption) GetMetalDeviceBGPNeighborsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMetalDeviceBGPNeighborsResult, error) {
			args := v.(GetMetalDeviceBGPNeighborsArgs)
			r, err := GetMetalDeviceBGPNeighbors(ctx, &args, opts...)
			var s GetMetalDeviceBGPNeighborsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMetalDeviceBGPNeighborsResultOutput)
}

// A collection of arguments for invoking GetMetalDeviceBGPNeighbors.
type GetMetalDeviceBGPNeighborsOutputArgs struct {
	// UUID of BGP-enabled device whose neighbors to list.
	DeviceId pulumi.StringInput `pulumi:"deviceId"`
}

func (GetMetalDeviceBGPNeighborsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMetalDeviceBGPNeighborsArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalDeviceBGPNeighbors.
type GetMetalDeviceBGPNeighborsResultOutput struct{ *pulumi.OutputState }

func (GetMetalDeviceBGPNeighborsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMetalDeviceBGPNeighborsResult)(nil)).Elem()
}

func (o GetMetalDeviceBGPNeighborsResultOutput) ToGetMetalDeviceBGPNeighborsResultOutput() GetMetalDeviceBGPNeighborsResultOutput {
	return o
}

func (o GetMetalDeviceBGPNeighborsResultOutput) ToGetMetalDeviceBGPNeighborsResultOutputWithContext(ctx context.Context) GetMetalDeviceBGPNeighborsResultOutput {
	return o
}

// array of BGP neighbor records with attributes:
func (o GetMetalDeviceBGPNeighborsResultOutput) BgpNeighbors() GetMetalDeviceBGPNeighborsBgpNeighborArrayOutput {
	return o.ApplyT(func(v GetMetalDeviceBGPNeighborsResult) []GetMetalDeviceBGPNeighborsBgpNeighbor {
		return v.BgpNeighbors
	}).(GetMetalDeviceBGPNeighborsBgpNeighborArrayOutput)
}

func (o GetMetalDeviceBGPNeighborsResultOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalDeviceBGPNeighborsResult) string { return v.DeviceId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMetalDeviceBGPNeighborsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalDeviceBGPNeighborsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMetalDeviceBGPNeighborsResultOutput{})
}
