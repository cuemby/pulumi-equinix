// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a [connection resource](https://metal.equinix.com/developers/docs/networking/fabric/)
//
// > Equinix Metal connection with serviceTokenType `aSide` is not generally available and may not be enabled yet for your organization.
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
// 		_, err := equinix.GetMetalConnection(ctx, &GetMetalConnectionArgs{
// 			ConnectionId: "4347e805-eb46-4699-9eb9-5c116e6a017d",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMetalConnection(ctx *pulumi.Context, args *LookupMetalConnectionArgs, opts ...pulumi.InvokeOption) (*LookupMetalConnectionResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMetalConnectionResult
	err := ctx.Invoke("equinix:index/getMetalConnection:GetMetalConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalConnection.
type LookupMetalConnectionArgs struct {
	// ID of the connection resource.
	ConnectionId string `pulumi:"connectionId"`
}

// A collection of values returned by GetMetalConnection.
type LookupMetalConnectionResult struct {
	ConnectionId string `pulumi:"connectionId"`
	// Description of the connection resource.
	Description string `pulumi:"description"`
	// Slug of a facility to which the connection belongs.
	Facility string `pulumi:"facility"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Slug of a metro to which the connection belongs.
	Metro string `pulumi:"metro"`
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel.
	Mode string `pulumi:"mode"`
	// Port name.
	Name string `pulumi:"name"`
	// ID of the organization where the connection is scoped to.
	OrganizationId string `pulumi:"organizationId"`
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
	Ports []GetMetalConnectionPort `pulumi:"ports"`
	// ID of project to which the connection belongs.
	ProjectId string `pulumi:"projectId"`
	// Connection redundancy, reduntant or primary.
	Redundancy string `pulumi:"redundancy"`
	// Type of service token, aSide or z_side. One available in shared connection.
	ServiceTokenType string `pulumi:"serviceTokenType"`
	// List of connection service tokens with attributes
	ServiceTokens []GetMetalConnectionServiceToken `pulumi:"serviceTokens"`
	// Port speed in bits per second.
	Speed string `pulumi:"speed"`
	// Port status.
	Status string `pulumi:"status"`
	// String list of tags.
	Tags []string `pulumi:"tags"`
	// (Deprecated) Token to configure the connection in the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard).
	Token string `pulumi:"token"`
	// Token type, `aSide` or `zSide`.
	Type string `pulumi:"type"`
	// Attached VLANs. Only available in shared connection. One vlan for Primary/Single connection and two vlans for Redundant connection.
	Vlans []int `pulumi:"vlans"`
}

func LookupMetalConnectionOutput(ctx *pulumi.Context, args LookupMetalConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupMetalConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetalConnectionResult, error) {
			args := v.(LookupMetalConnectionArgs)
			r, err := LookupMetalConnection(ctx, &args, opts...)
			var s LookupMetalConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetalConnectionResultOutput)
}

// A collection of arguments for invoking GetMetalConnection.
type LookupMetalConnectionOutputArgs struct {
	// ID of the connection resource.
	ConnectionId pulumi.StringInput `pulumi:"connectionId"`
}

func (LookupMetalConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalConnectionArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalConnection.
type LookupMetalConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupMetalConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalConnectionResult)(nil)).Elem()
}

func (o LookupMetalConnectionResultOutput) ToLookupMetalConnectionResultOutput() LookupMetalConnectionResultOutput {
	return o
}

func (o LookupMetalConnectionResultOutput) ToLookupMetalConnectionResultOutputWithContext(ctx context.Context) LookupMetalConnectionResultOutput {
	return o
}

func (o LookupMetalConnectionResultOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.ConnectionId }).(pulumi.StringOutput)
}

// Description of the connection resource.
func (o LookupMetalConnectionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Description }).(pulumi.StringOutput)
}

// Slug of a facility to which the connection belongs.
func (o LookupMetalConnectionResultOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Facility }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMetalConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Slug of a metro to which the connection belongs.
func (o LookupMetalConnectionResultOutput) Metro() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Metro }).(pulumi.StringOutput)
}

// Mode for connections in IBX facilities with the dedicated type - standard or tunnel.
func (o LookupMetalConnectionResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Mode }).(pulumi.StringOutput)
}

// Port name.
func (o LookupMetalConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization where the connection is scoped to.
func (o LookupMetalConnectionResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`)
func (o LookupMetalConnectionResultOutput) Ports() GetMetalConnectionPortArrayOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) []GetMetalConnectionPort { return v.Ports }).(GetMetalConnectionPortArrayOutput)
}

// ID of project to which the connection belongs.
func (o LookupMetalConnectionResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Connection redundancy, reduntant or primary.
func (o LookupMetalConnectionResultOutput) Redundancy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Redundancy }).(pulumi.StringOutput)
}

// Type of service token, aSide or z_side. One available in shared connection.
func (o LookupMetalConnectionResultOutput) ServiceTokenType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.ServiceTokenType }).(pulumi.StringOutput)
}

// List of connection service tokens with attributes
func (o LookupMetalConnectionResultOutput) ServiceTokens() GetMetalConnectionServiceTokenArrayOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) []GetMetalConnectionServiceToken { return v.ServiceTokens }).(GetMetalConnectionServiceTokenArrayOutput)
}

// Port speed in bits per second.
func (o LookupMetalConnectionResultOutput) Speed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Speed }).(pulumi.StringOutput)
}

// Port status.
func (o LookupMetalConnectionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Status }).(pulumi.StringOutput)
}

// String list of tags.
func (o LookupMetalConnectionResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// (Deprecated) Token to configure the connection in the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard).
func (o LookupMetalConnectionResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Token }).(pulumi.StringOutput)
}

// Token type, `aSide` or `zSide`.
func (o LookupMetalConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Attached VLANs. Only available in shared connection. One vlan for Primary/Single connection and two vlans for Redundant connection.
func (o LookupMetalConnectionResultOutput) Vlans() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupMetalConnectionResult) []int { return v.Vlans }).(pulumi.IntArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetalConnectionResultOutput{})
}
