// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Equinix Metal facility datasource.
func GetMetalFacility(ctx *pulumi.Context, args *GetMetalFacilityArgs, opts ...pulumi.InvokeOption) (*GetMetalFacilityResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMetalFacilityResult
	err := ctx.Invoke("equinix:index/getMetalFacility:GetMetalFacility", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalFacility.
type GetMetalFacilityArgs struct {
	// One or more device plans for which the facility must have capacity.
	Capacities []GetMetalFacilityCapacity `pulumi:"capacities"`
	// The facility code to search for facilities.
	Code string `pulumi:"code"`
	// Set of feature strings that the facility must have. Some
	// possible values are `baremetal`, `ibx`, `storage`, `globalIpv4`, `backendTransfer`, `layer2`.
	FeaturesRequireds []string `pulumi:"featuresRequireds"`
}

// A collection of values returned by GetMetalFacility.
type GetMetalFacilityResult struct {
	Capacities []GetMetalFacilityCapacity `pulumi:"capacities"`
	Code       string                     `pulumi:"code"`
	// The features of the facility.
	Features          []string `pulumi:"features"`
	FeaturesRequireds []string `pulumi:"featuresRequireds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The metro code the facility is part of.
	Metro string `pulumi:"metro"`
	// The name of the facility.
	Name string `pulumi:"name"`
}

func GetMetalFacilityOutput(ctx *pulumi.Context, args GetMetalFacilityOutputArgs, opts ...pulumi.InvokeOption) GetMetalFacilityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMetalFacilityResult, error) {
			args := v.(GetMetalFacilityArgs)
			r, err := GetMetalFacility(ctx, &args, opts...)
			var s GetMetalFacilityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMetalFacilityResultOutput)
}

// A collection of arguments for invoking GetMetalFacility.
type GetMetalFacilityOutputArgs struct {
	// One or more device plans for which the facility must have capacity.
	Capacities GetMetalFacilityCapacityArrayInput `pulumi:"capacities"`
	// The facility code to search for facilities.
	Code pulumi.StringInput `pulumi:"code"`
	// Set of feature strings that the facility must have. Some
	// possible values are `baremetal`, `ibx`, `storage`, `globalIpv4`, `backendTransfer`, `layer2`.
	FeaturesRequireds pulumi.StringArrayInput `pulumi:"featuresRequireds"`
}

func (GetMetalFacilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMetalFacilityArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalFacility.
type GetMetalFacilityResultOutput struct{ *pulumi.OutputState }

func (GetMetalFacilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMetalFacilityResult)(nil)).Elem()
}

func (o GetMetalFacilityResultOutput) ToGetMetalFacilityResultOutput() GetMetalFacilityResultOutput {
	return o
}

func (o GetMetalFacilityResultOutput) ToGetMetalFacilityResultOutputWithContext(ctx context.Context) GetMetalFacilityResultOutput {
	return o
}

func (o GetMetalFacilityResultOutput) Capacities() GetMetalFacilityCapacityArrayOutput {
	return o.ApplyT(func(v GetMetalFacilityResult) []GetMetalFacilityCapacity { return v.Capacities }).(GetMetalFacilityCapacityArrayOutput)
}

func (o GetMetalFacilityResultOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalFacilityResult) string { return v.Code }).(pulumi.StringOutput)
}

// The features of the facility.
func (o GetMetalFacilityResultOutput) Features() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMetalFacilityResult) []string { return v.Features }).(pulumi.StringArrayOutput)
}

func (o GetMetalFacilityResultOutput) FeaturesRequireds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMetalFacilityResult) []string { return v.FeaturesRequireds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMetalFacilityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalFacilityResult) string { return v.Id }).(pulumi.StringOutput)
}

// The metro code the facility is part of.
func (o GetMetalFacilityResultOutput) Metro() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalFacilityResult) string { return v.Metro }).(pulumi.StringOutput)
}

// The name of the facility.
func (o GetMetalFacilityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalFacilityResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMetalFacilityResultOutput{})
}
