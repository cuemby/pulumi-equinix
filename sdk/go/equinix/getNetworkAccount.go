// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get number and identifier of Equinix Network Edge
// billing account in a given metro location.
//
// Billing account reference is required to create Network Edge virtual device
// in corresponding metro location.
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
// 		dc, err := equinix.GetNetworkAccount(ctx, &GetNetworkAccountArgs{
// 			MetroCode: "DC",
// 			Status:    pulumi.StringRef("Active"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("number", dc.Number)
// 		return nil
// 	})
// }
// ```
func GetNetworkAccount(ctx *pulumi.Context, args *GetNetworkAccountArgs, opts ...pulumi.InvokeOption) (*GetNetworkAccountResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetNetworkAccountResult
	err := ctx.Invoke("equinix:index/getNetworkAccount:GetNetworkAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNetworkAccount.
type GetNetworkAccountArgs struct {
	// Account location metro code.
	MetroCode string `pulumi:"metroCode"`
	// Account name for filtering.
	Name *string `pulumi:"name"`
	// Account status for filtering. Possible values are: `Active`, `Processing`,
	// `Submitted`, `Staged`.
	Status *string `pulumi:"status"`
}

// A collection of values returned by GetNetworkAccount.
type GetNetworkAccountResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	MetroCode string `pulumi:"metroCode"`
	Name      string `pulumi:"name"`
	// Account unique number.
	Number string `pulumi:"number"`
	Status string `pulumi:"status"`
	// Account unique identifier.
	UcmId string `pulumi:"ucmId"`
}

func GetNetworkAccountOutput(ctx *pulumi.Context, args GetNetworkAccountOutputArgs, opts ...pulumi.InvokeOption) GetNetworkAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNetworkAccountResult, error) {
			args := v.(GetNetworkAccountArgs)
			r, err := GetNetworkAccount(ctx, &args, opts...)
			var s GetNetworkAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNetworkAccountResultOutput)
}

// A collection of arguments for invoking GetNetworkAccount.
type GetNetworkAccountOutputArgs struct {
	// Account location metro code.
	MetroCode pulumi.StringInput `pulumi:"metroCode"`
	// Account name for filtering.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Account status for filtering. Possible values are: `Active`, `Processing`,
	// `Submitted`, `Staged`.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetNetworkAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkAccountArgs)(nil)).Elem()
}

// A collection of values returned by GetNetworkAccount.
type GetNetworkAccountResultOutput struct{ *pulumi.OutputState }

func (GetNetworkAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkAccountResult)(nil)).Elem()
}

func (o GetNetworkAccountResultOutput) ToGetNetworkAccountResultOutput() GetNetworkAccountResultOutput {
	return o
}

func (o GetNetworkAccountResultOutput) ToGetNetworkAccountResultOutputWithContext(ctx context.Context) GetNetworkAccountResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNetworkAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNetworkAccountResultOutput) MetroCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkAccountResult) string { return v.MetroCode }).(pulumi.StringOutput)
}

func (o GetNetworkAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Account unique number.
func (o GetNetworkAccountResultOutput) Number() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkAccountResult) string { return v.Number }).(pulumi.StringOutput)
}

func (o GetNetworkAccountResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkAccountResult) string { return v.Status }).(pulumi.StringOutput)
}

// Account unique identifier.
func (o GetNetworkAccountResultOutput) UcmId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkAccountResult) string { return v.UcmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworkAccountResultOutput{})
}
