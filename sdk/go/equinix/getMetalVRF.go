// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a VRF resource.
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
// 		_, err := equinix.GetMetalVRF(ctx, &GetMetalVRFArgs{
// 			VrfId: "48630899-9ff2-4ce6-a93f-50ff4ebcdf6e",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMetalVRF(ctx *pulumi.Context, args *LookupMetalVRFArgs, opts ...pulumi.InvokeOption) (*LookupMetalVRFResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMetalVRFResult
	err := ctx.Invoke("equinix:index/getMetalVRF:GetMetalVRF", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalVRF.
type LookupMetalVRFArgs struct {
	// ID of the VRF resource
	VrfId string `pulumi:"vrfId"`
}

// A collection of values returned by GetMetalVRF.
type LookupMetalVRFResult struct {
	// Description of the VRF.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// All IPv4 and IPv6 Ranges that will be available to BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must not overlap other ranges within the VRF.
	IpRanges []string `pulumi:"ipRanges"`
	// The 4-byte ASN set on the VRF.
	LocalAsn int `pulumi:"localAsn"`
	// Metro ID or Code where the VRF will be deployed.
	Metro string `pulumi:"metro"`
	// User-supplied name of the VRF, unique to the project
	Name string `pulumi:"name"`
	// Project ID where the VRF will be deployed.
	ProjectId string `pulumi:"projectId"`
	VrfId     string `pulumi:"vrfId"`
}

func LookupMetalVRFOutput(ctx *pulumi.Context, args LookupMetalVRFOutputArgs, opts ...pulumi.InvokeOption) LookupMetalVRFResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetalVRFResult, error) {
			args := v.(LookupMetalVRFArgs)
			r, err := LookupMetalVRF(ctx, &args, opts...)
			var s LookupMetalVRFResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetalVRFResultOutput)
}

// A collection of arguments for invoking GetMetalVRF.
type LookupMetalVRFOutputArgs struct {
	// ID of the VRF resource
	VrfId pulumi.StringInput `pulumi:"vrfId"`
}

func (LookupMetalVRFOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalVRFArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalVRF.
type LookupMetalVRFResultOutput struct{ *pulumi.OutputState }

func (LookupMetalVRFResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalVRFResult)(nil)).Elem()
}

func (o LookupMetalVRFResultOutput) ToLookupMetalVRFResultOutput() LookupMetalVRFResultOutput {
	return o
}

func (o LookupMetalVRFResultOutput) ToLookupMetalVRFResultOutputWithContext(ctx context.Context) LookupMetalVRFResultOutput {
	return o
}

// Description of the VRF.
func (o LookupMetalVRFResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMetalVRFResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) string { return v.Id }).(pulumi.StringOutput)
}

// All IPv4 and IPv6 Ranges that will be available to BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must not overlap other ranges within the VRF.
func (o LookupMetalVRFResultOutput) IpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) []string { return v.IpRanges }).(pulumi.StringArrayOutput)
}

// The 4-byte ASN set on the VRF.
func (o LookupMetalVRFResultOutput) LocalAsn() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) int { return v.LocalAsn }).(pulumi.IntOutput)
}

// Metro ID or Code where the VRF will be deployed.
func (o LookupMetalVRFResultOutput) Metro() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) string { return v.Metro }).(pulumi.StringOutput)
}

// User-supplied name of the VRF, unique to the project
func (o LookupMetalVRFResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project ID where the VRF will be deployed.
func (o LookupMetalVRFResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupMetalVRFResultOutput) VrfId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVRFResult) string { return v.VrfId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetalVRFResultOutput{})
}
