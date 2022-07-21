// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a virtual circuit resource from
// [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)
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
// 		exampleConnection, err := equinix.GetMetalConnection(ctx, &GetMetalConnectionArgs{
// 			ConnectionId: "4347e805-eb46-4699-9eb9-5c116e6a017d",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.GetMetalVirtualCircuit(ctx, &GetMetalVirtualCircuitArgs{
// 			VirtualCircuitId: exampleConnection.Ports[1].VirtualCircuitIds[0],
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMetalVirtualCircuit(ctx *pulumi.Context, args *LookupMetalVirtualCircuitArgs, opts ...pulumi.InvokeOption) (*LookupMetalVirtualCircuitResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMetalVirtualCircuitResult
	err := ctx.Invoke("equinix:index/getMetalVirtualCircuit:GetMetalVirtualCircuit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalVirtualCircuit.
type LookupMetalVirtualCircuitArgs struct {
	// ID of the virtual circuit resource
	VirtualCircuitId string `pulumi:"virtualCircuitId"`
}

// A collection of values returned by GetMetalVirtualCircuit.
type LookupMetalVirtualCircuitResult struct {
	// UUID of Connection where the VC is scoped to.
	ConnectionId string `pulumi:"connectionId"`
	// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
	CustomerIp string `pulumi:"customerIp"`
	// Description for the Virtual Circuit resource.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The password that can be set for the VRF BGP peer
	Md5 string `pulumi:"md5"`
	// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
	MetalIp string `pulumi:"metalIp"`
	// Name of the virtual circuit resource.
	Name    string `pulumi:"name"`
	NniVlan int    `pulumi:"nniVlan"`
	NniVnid int    `pulumi:"nniVnid"`
	// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the localAsn of the VRF.
	PeerAsn int `pulumi:"peerAsn"`
	// UUID of the Connection Port where the VC is scoped to.
	PortId string `pulumi:"portId"`
	// ID of project to which the VC belongs.
	// * `vnid`, `nniVlan`, `nniNvid` - VLAN parameters, see the
	//   [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).
	ProjectId string `pulumi:"projectId"`
	// Speed of the Virtual Circuit resource.
	Speed string `pulumi:"speed"`
	// Status of the virtal circuit.
	Status string `pulumi:"status"`
	// A subnet from one of the IP
	// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
	// * For a /31 block, it will only have two IP addresses, which will be used for
	//   the metalIp and customer_ip.
	// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
	Subnet string `pulumi:"subnet"`
	// Tags for the Virtual Circuit resource.
	Tags             []string `pulumi:"tags"`
	VirtualCircuitId string   `pulumi:"virtualCircuitId"`
	VlanId           string   `pulumi:"vlanId"`
	Vnid             int      `pulumi:"vnid"`
	// UUID of the VLAN to associate.
	VrfId string `pulumi:"vrfId"`
}

func LookupMetalVirtualCircuitOutput(ctx *pulumi.Context, args LookupMetalVirtualCircuitOutputArgs, opts ...pulumi.InvokeOption) LookupMetalVirtualCircuitResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetalVirtualCircuitResult, error) {
			args := v.(LookupMetalVirtualCircuitArgs)
			r, err := LookupMetalVirtualCircuit(ctx, &args, opts...)
			var s LookupMetalVirtualCircuitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetalVirtualCircuitResultOutput)
}

// A collection of arguments for invoking GetMetalVirtualCircuit.
type LookupMetalVirtualCircuitOutputArgs struct {
	// ID of the virtual circuit resource
	VirtualCircuitId pulumi.StringInput `pulumi:"virtualCircuitId"`
}

func (LookupMetalVirtualCircuitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalVirtualCircuitArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalVirtualCircuit.
type LookupMetalVirtualCircuitResultOutput struct{ *pulumi.OutputState }

func (LookupMetalVirtualCircuitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetalVirtualCircuitResult)(nil)).Elem()
}

func (o LookupMetalVirtualCircuitResultOutput) ToLookupMetalVirtualCircuitResultOutput() LookupMetalVirtualCircuitResultOutput {
	return o
}

func (o LookupMetalVirtualCircuitResultOutput) ToLookupMetalVirtualCircuitResultOutputWithContext(ctx context.Context) LookupMetalVirtualCircuitResultOutput {
	return o
}

// UUID of Connection where the VC is scoped to.
func (o LookupMetalVirtualCircuitResultOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.ConnectionId }).(pulumi.StringOutput)
}

// The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
func (o LookupMetalVirtualCircuitResultOutput) CustomerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.CustomerIp }).(pulumi.StringOutput)
}

// Description for the Virtual Circuit resource.
func (o LookupMetalVirtualCircuitResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMetalVirtualCircuitResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.Id }).(pulumi.StringOutput)
}

// The password that can be set for the VRF BGP peer
func (o LookupMetalVirtualCircuitResultOutput) Md5() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.Md5 }).(pulumi.StringOutput)
}

// The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
func (o LookupMetalVirtualCircuitResultOutput) MetalIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.MetalIp }).(pulumi.StringOutput)
}

// Name of the virtual circuit resource.
func (o LookupMetalVirtualCircuitResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMetalVirtualCircuitResultOutput) NniVlan() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) int { return v.NniVlan }).(pulumi.IntOutput)
}

func (o LookupMetalVirtualCircuitResultOutput) NniVnid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) int { return v.NniVnid }).(pulumi.IntOutput)
}

// The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the localAsn of the VRF.
func (o LookupMetalVirtualCircuitResultOutput) PeerAsn() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) int { return v.PeerAsn }).(pulumi.IntOutput)
}

// UUID of the Connection Port where the VC is scoped to.
func (o LookupMetalVirtualCircuitResultOutput) PortId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.PortId }).(pulumi.StringOutput)
}

// ID of project to which the VC belongs.
// * `vnid`, `nniVlan`, `nniNvid` - VLAN parameters, see the
//   [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).
func (o LookupMetalVirtualCircuitResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Speed of the Virtual Circuit resource.
func (o LookupMetalVirtualCircuitResultOutput) Speed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.Speed }).(pulumi.StringOutput)
}

// Status of the virtal circuit.
func (o LookupMetalVirtualCircuitResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.Status }).(pulumi.StringOutput)
}

// A subnet from one of the IP
// blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
// * For a /31 block, it will only have two IP addresses, which will be used for
//   the metalIp and customer_ip.
// * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
func (o LookupMetalVirtualCircuitResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.Subnet }).(pulumi.StringOutput)
}

// Tags for the Virtual Circuit resource.
func (o LookupMetalVirtualCircuitResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupMetalVirtualCircuitResultOutput) VirtualCircuitId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.VirtualCircuitId }).(pulumi.StringOutput)
}

func (o LookupMetalVirtualCircuitResultOutput) VlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.VlanId }).(pulumi.StringOutput)
}

func (o LookupMetalVirtualCircuitResultOutput) Vnid() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) int { return v.Vnid }).(pulumi.IntOutput)
}

// UUID of the VLAN to associate.
func (o LookupMetalVirtualCircuitResultOutput) VrfId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetalVirtualCircuitResult) string { return v.VrfId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetalVirtualCircuitResultOutput{})
}
