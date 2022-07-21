// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to find IP address blocks in Equinix Metal. You can use IP address or a block
// ID for lookup.
//
// > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
func LookupMetalReservedIPBlock(ctx *pulumi.Context, args *LookupMetalReservedIPBlockArgs, opts ...pulumi.InvokeOption) (*LookupMetalReservedIPBlockResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMetalReservedIPBlockResult
	err := ctx.Invoke("equinix:index/getMetalReservedIPBlock:GetMetalReservedIPBlock", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalReservedIPBlock.
type LookupMetalReservedIPBlockArgs struct {
	// UUID of the IP address block to look up.
	Id *string `pulumi:"id"`
	// Block containing this IP address will be returned.
	IpAddress *string `pulumi:"ipAddress"`
	// UUID of the project where the searched block should be.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by GetMetalReservedIPBlock.
type LookupMetalReservedIPBlockResult struct {
	Address       string  `pulumi:"address"`
	AddressFamily int     `pulumi:"addressFamily"`
	Cidr          int     `pulumi:"cidr"`
	CidrNotation  string  `pulumi:"cidrNotation"`
	Facility      string  `pulumi:"facility"`
	Gateway       string  `pulumi:"gateway"`
	Global        bool    `pulumi:"global"`
	Id            string  `pulumi:"id"`
	IpAddress     *string `pulumi:"ipAddress"`
	Manageable    bool    `pulumi:"manageable"`
	Management    bool    `pulumi:"management"`
	Metro         string  `pulumi:"metro"`
	Netmask       string  `pulumi:"netmask"`
	Network       string  `pulumi:"network"`
	ProjectId     string  `pulumi:"projectId"`
	Public        bool    `pulumi:"public"`
	Quantity      int     `pulumi:"quantity"`
	// One of `globalIpv4`, `publicIpv4`, `privateIpv4`, `publicIpv6`,or `vrf`
	Type  string `pulumi:"type"`
	VrfId int    `pulumi:"vrfId"`
}

func LookupMetalReservedIPBlockOutput(ctx *pulumi.Context, args LookupMetalReservedIPBlockOutputArgs, opts ...pulumi.InvokeOption) LookupMetalReservedIPBlockResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetalReservedIPBlockResult, error) {
			args := v.(LookupMetalReservedIPBlockArgs)
			r, err := LookupMetalReservedIPBlock(ctx, &args, opts...)
			var s LookupMetalReservedIPBlockResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetalReservedIPBlockResultOutput)
}

// A collection of arguments for invoking GetMetalReservedIPBlock.
type LookupMetalReservedIPBlockOutputArgs struct {
	// UUID of the IP address block to look up.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Block containing this IP address will be returned.
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
	// UUID of the project where the searched block should be.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupMetalReservedIPBlockOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalReservedIPBlockArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalReservedIPBlock.
type LookupMetalReservedIPBlockResultOutput struct{ *pulumi.OutputState }

func (LookupMetalReservedIPBlockResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalReservedIPBlockResult)(nil)).Elem()
}

func (o LookupMetalReservedIPBlockResultOutput) ToLookupMetalReservedIPBlockResultOutput() LookupMetalReservedIPBlockResultOutput {
	return o
}

func (o LookupMetalReservedIPBlockResultOutput) ToLookupMetalReservedIPBlockResultOutputWithContext(ctx context.Context) LookupMetalReservedIPBlockResultOutput {
	return o
}

func (o LookupMetalReservedIPBlockResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) AddressFamily() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) int { return v.AddressFamily }).(pulumi.IntOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Cidr() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) int { return v.Cidr }).(pulumi.IntOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) CidrNotation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.CidrNotation }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Facility }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Gateway }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Global() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) bool { return v.Global }).(pulumi.BoolOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Manageable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) bool { return v.Manageable }).(pulumi.BoolOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Management() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) bool { return v.Management }).(pulumi.BoolOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Metro() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Metro }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Netmask }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) bool { return v.Public }).(pulumi.BoolOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) Quantity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) int { return v.Quantity }).(pulumi.IntOutput)
}

// One of `globalIpv4`, `publicIpv4`, `privateIpv4`, `publicIpv6`,or `vrf`
func (o LookupMetalReservedIPBlockResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMetalReservedIPBlockResultOutput) VrfId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalReservedIPBlockResult) int { return v.VrfId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetalReservedIPBlockResultOutput{})
}
