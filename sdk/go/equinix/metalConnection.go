// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to request the creation an Interconnection asset to connect with other parties using [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/).
//
// > Equinix Metal connection with serviceTokenType `aSide` is not generally available and may not be enabled yet for your organization.
//
// ## Example Usage
// ### Shared Connection with zSide token - Non-redundant Connection from your own Equinix Fabric Port to Equinix Metal
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
// 		exampleMetalConnection, err := equinix.NewMetalConnection(ctx, "exampleMetalConnection", &equinix.MetalConnectionArgs{
// 			ProjectId:        pulumi.Any(local.Project_id),
// 			Type:             pulumi.String("shared"),
// 			Redundancy:       pulumi.String("primary"),
// 			Metro:            pulumi.String("FR"),
// 			Speed:            pulumi.String("200Mbps"),
// 			ServiceTokenType: pulumi.String("z_side"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleGetECXPort, err := equinix.GetECXPort(ctx, &GetECXPortArgs{
// 			Name: "CX-FR5-NL-Dot1q-BO-1G-PRI",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.NewECXL2Connection(ctx, "exampleECXL2Connection", &equinix.ECXL2ConnectionArgs{
// 			ZsideServiceToken: exampleMetalConnection.ServiceTokens.ApplyT(func(serviceTokens []MetalConnectionServiceToken) (string, error) {
// 				return serviceTokens[0].Id, nil
// 			}).(pulumi.StringOutput),
// 			Speed:     pulumi.Int(200),
// 			SpeedUnit: pulumi.String("MB"),
// 			Notifications: pulumi.StringArray{
// 				pulumi.String("example@equinix.com"),
// 			},
// 			SellerMetroCode: pulumi.String("FR"),
// 			PortUuid:        pulumi.String(exampleGetECXPort.Id),
// 			VlanStag:        pulumi.Int(1020),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MetalConnection struct {
	pulumi.CustomResourceState

	// Description for the connection resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Facility where the connection will be created.
	Facility pulumi.StringOutput `pulumi:"facility"`
	// Metro where the connection will be created.
	Metro pulumi.StringOutput `pulumi:"metro"`
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Name of the connection resource
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization where the connection is scoped to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of
	// port is described in documentation of the
	// MetalConnection datasource.
	Ports MetalConnectionPortArrayOutput `pulumi:"ports"`
	// ID of the project where the connection is scoped to, must be set for.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Connection redundancy - redundant or primary.
	Redundancy pulumi.StringOutput `pulumi:"redundancy"`
	// Only used with shared connection. Type of service token to use for the connection, aSide or z_side.
	ServiceTokenType pulumi.StringPtrOutput `pulumi:"serviceTokenType"`
	// List of connection service tokens with attributes. Scehma of serviceToken is described in documentation of the MetalConnection datasource.
	ServiceTokens MetalConnectionServiceTokenArrayOutput `pulumi:"serviceTokens"`
	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	Speed pulumi.StringOutput `pulumi:"speed"`
	// Status of the connection resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// String list of tags.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	//
	// Deprecated: token is deprecated. Use service_tokens instead
	Token pulumi.StringOutput `pulumi:"token"`
	// Connection type - dedicated or shared.
	Type pulumi.StringOutput `pulumi:"type"`
	// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
	Vlans pulumi.IntArrayOutput `pulumi:"vlans"`
}

// NewMetalConnection registers a new resource with the given unique name, arguments, and options.
func NewMetalConnection(ctx *pulumi.Context,
	name string, args *MetalConnectionArgs, opts ...pulumi.ResourceOption) (*MetalConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Redundancy == nil {
		return nil, errors.New("invalid value for required argument 'Redundancy'")
	}
	if args.Speed == nil {
		return nil, errors.New("invalid value for required argument 'Speed'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MetalConnection
	err := ctx.RegisterResource("equinix:index/metalConnection:MetalConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalConnection gets an existing MetalConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalConnectionState, opts ...pulumi.ResourceOption) (*MetalConnection, error) {
	var resource MetalConnection
	err := ctx.ReadResource("equinix:index/metalConnection:MetalConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalConnection resources.
type metalConnectionState struct {
	// Description for the connection resource.
	Description *string `pulumi:"description"`
	// Facility where the connection will be created.
	Facility *string `pulumi:"facility"`
	// Metro where the connection will be created.
	Metro *string `pulumi:"metro"`
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
	Mode *string `pulumi:"mode"`
	// Name of the connection resource
	Name *string `pulumi:"name"`
	// ID of the organization where the connection is scoped to.
	OrganizationId *string `pulumi:"organizationId"`
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of
	// port is described in documentation of the
	// MetalConnection datasource.
	Ports []MetalConnectionPort `pulumi:"ports"`
	// ID of the project where the connection is scoped to, must be set for.
	ProjectId *string `pulumi:"projectId"`
	// Connection redundancy - redundant or primary.
	Redundancy *string `pulumi:"redundancy"`
	// Only used with shared connection. Type of service token to use for the connection, aSide or z_side.
	ServiceTokenType *string `pulumi:"serviceTokenType"`
	// List of connection service tokens with attributes. Scehma of serviceToken is described in documentation of the MetalConnection datasource.
	ServiceTokens []MetalConnectionServiceToken `pulumi:"serviceTokens"`
	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	Speed *string `pulumi:"speed"`
	// Status of the connection resource.
	Status *string `pulumi:"status"`
	// String list of tags.
	Tags []string `pulumi:"tags"`
	// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	//
	// Deprecated: token is deprecated. Use service_tokens instead
	Token *string `pulumi:"token"`
	// Connection type - dedicated or shared.
	Type *string `pulumi:"type"`
	// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
	Vlans []int `pulumi:"vlans"`
}

type MetalConnectionState struct {
	// Description for the connection resource.
	Description pulumi.StringPtrInput
	// Facility where the connection will be created.
	Facility pulumi.StringPtrInput
	// Metro where the connection will be created.
	Metro pulumi.StringPtrInput
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
	Mode pulumi.StringPtrInput
	// Name of the connection resource
	Name pulumi.StringPtrInput
	// ID of the organization where the connection is scoped to.
	OrganizationId pulumi.StringPtrInput
	// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of
	// port is described in documentation of the
	// MetalConnection datasource.
	Ports MetalConnectionPortArrayInput
	// ID of the project where the connection is scoped to, must be set for.
	ProjectId pulumi.StringPtrInput
	// Connection redundancy - redundant or primary.
	Redundancy pulumi.StringPtrInput
	// Only used with shared connection. Type of service token to use for the connection, aSide or z_side.
	ServiceTokenType pulumi.StringPtrInput
	// List of connection service tokens with attributes. Scehma of serviceToken is described in documentation of the MetalConnection datasource.
	ServiceTokens MetalConnectionServiceTokenArrayInput
	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	Speed pulumi.StringPtrInput
	// Status of the connection resource.
	Status pulumi.StringPtrInput
	// String list of tags.
	Tags pulumi.StringArrayInput
	// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
	//
	// Deprecated: token is deprecated. Use service_tokens instead
	Token pulumi.StringPtrInput
	// Connection type - dedicated or shared.
	Type pulumi.StringPtrInput
	// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
	Vlans pulumi.IntArrayInput
}

func (MetalConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalConnectionState)(nil)).Elem()
}

type metalConnectionArgs struct {
	// Description for the connection resource.
	Description *string `pulumi:"description"`
	// Facility where the connection will be created.
	Facility *string `pulumi:"facility"`
	// Metro where the connection will be created.
	Metro *string `pulumi:"metro"`
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
	Mode *string `pulumi:"mode"`
	// Name of the connection resource
	Name *string `pulumi:"name"`
	// ID of the organization where the connection is scoped to.
	OrganizationId *string `pulumi:"organizationId"`
	// ID of the project where the connection is scoped to, must be set for.
	ProjectId *string `pulumi:"projectId"`
	// Connection redundancy - redundant or primary.
	Redundancy string `pulumi:"redundancy"`
	// Only used with shared connection. Type of service token to use for the connection, aSide or z_side.
	ServiceTokenType *string `pulumi:"serviceTokenType"`
	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	Speed string `pulumi:"speed"`
	// String list of tags.
	Tags []string `pulumi:"tags"`
	// Connection type - dedicated or shared.
	Type string `pulumi:"type"`
	// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
	Vlans []int `pulumi:"vlans"`
}

// The set of arguments for constructing a MetalConnection resource.
type MetalConnectionArgs struct {
	// Description for the connection resource.
	Description pulumi.StringPtrInput
	// Facility where the connection will be created.
	Facility pulumi.StringPtrInput
	// Metro where the connection will be created.
	Metro pulumi.StringPtrInput
	// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
	Mode pulumi.StringPtrInput
	// Name of the connection resource
	Name pulumi.StringPtrInput
	// ID of the organization where the connection is scoped to.
	OrganizationId pulumi.StringPtrInput
	// ID of the project where the connection is scoped to, must be set for.
	ProjectId pulumi.StringPtrInput
	// Connection redundancy - redundant or primary.
	Redundancy pulumi.StringInput
	// Only used with shared connection. Type of service token to use for the connection, aSide or z_side.
	ServiceTokenType pulumi.StringPtrInput
	// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
	Speed pulumi.StringInput
	// String list of tags.
	Tags pulumi.StringArrayInput
	// Connection type - dedicated or shared.
	Type pulumi.StringInput
	// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
	Vlans pulumi.IntArrayInput
}

func (MetalConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalConnectionArgs)(nil)).Elem()
}

type MetalConnectionInput interface {
	pulumi.Input

	ToMetalConnectionOutput() MetalConnectionOutput
	ToMetalConnectionOutputWithContext(ctx context.Context) MetalConnectionOutput
}

func (*MetalConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalConnection)(nil)).Elem()
}

func (i *MetalConnection) ToMetalConnectionOutput() MetalConnectionOutput {
	return i.ToMetalConnectionOutputWithContext(context.Background())
}

func (i *MetalConnection) ToMetalConnectionOutputWithContext(ctx context.Context) MetalConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalConnectionOutput)
}

// MetalConnectionArrayInput is an input type that accepts MetalConnectionArray and MetalConnectionArrayOutput values.
// You can construct a concrete instance of `MetalConnectionArrayInput` via:
//
//          MetalConnectionArray{ MetalConnectionArgs{...} }
type MetalConnectionArrayInput interface {
	pulumi.Input

	ToMetalConnectionArrayOutput() MetalConnectionArrayOutput
	ToMetalConnectionArrayOutputWithContext(context.Context) MetalConnectionArrayOutput
}

type MetalConnectionArray []MetalConnectionInput

func (MetalConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalConnection)(nil)).Elem()
}

func (i MetalConnectionArray) ToMetalConnectionArrayOutput() MetalConnectionArrayOutput {
	return i.ToMetalConnectionArrayOutputWithContext(context.Background())
}

func (i MetalConnectionArray) ToMetalConnectionArrayOutputWithContext(ctx context.Context) MetalConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalConnectionArrayOutput)
}

// MetalConnectionMapInput is an input type that accepts MetalConnectionMap and MetalConnectionMapOutput values.
// You can construct a concrete instance of `MetalConnectionMapInput` via:
//
//          MetalConnectionMap{ "key": MetalConnectionArgs{...} }
type MetalConnectionMapInput interface {
	pulumi.Input

	ToMetalConnectionMapOutput() MetalConnectionMapOutput
	ToMetalConnectionMapOutputWithContext(context.Context) MetalConnectionMapOutput
}

type MetalConnectionMap map[string]MetalConnectionInput

func (MetalConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalConnection)(nil)).Elem()
}

func (i MetalConnectionMap) ToMetalConnectionMapOutput() MetalConnectionMapOutput {
	return i.ToMetalConnectionMapOutputWithContext(context.Background())
}

func (i MetalConnectionMap) ToMetalConnectionMapOutputWithContext(ctx context.Context) MetalConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalConnectionMapOutput)
}

type MetalConnectionOutput struct{ *pulumi.OutputState }

func (MetalConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalConnection)(nil)).Elem()
}

func (o MetalConnectionOutput) ToMetalConnectionOutput() MetalConnectionOutput {
	return o
}

func (o MetalConnectionOutput) ToMetalConnectionOutputWithContext(ctx context.Context) MetalConnectionOutput {
	return o
}

// Description for the connection resource.
func (o MetalConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Facility where the connection will be created.
func (o MetalConnectionOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Facility }).(pulumi.StringOutput)
}

// Metro where the connection will be created.
func (o MetalConnectionOutput) Metro() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Metro }).(pulumi.StringOutput)
}

// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
func (o MetalConnectionOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// Name of the connection resource
func (o MetalConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization where the connection is scoped to.
func (o MetalConnectionOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of
// port is described in documentation of the
// MetalConnection datasource.
func (o MetalConnectionOutput) Ports() MetalConnectionPortArrayOutput {
	return o.ApplyT(func(v *MetalConnection) MetalConnectionPortArrayOutput { return v.Ports }).(MetalConnectionPortArrayOutput)
}

// ID of the project where the connection is scoped to, must be set for.
func (o MetalConnectionOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Connection redundancy - redundant or primary.
func (o MetalConnectionOutput) Redundancy() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Redundancy }).(pulumi.StringOutput)
}

// Only used with shared connection. Type of service token to use for the connection, aSide or z_side.
func (o MetalConnectionOutput) ServiceTokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringPtrOutput { return v.ServiceTokenType }).(pulumi.StringPtrOutput)
}

// List of connection service tokens with attributes. Scehma of serviceToken is described in documentation of the MetalConnection datasource.
func (o MetalConnectionOutput) ServiceTokens() MetalConnectionServiceTokenArrayOutput {
	return o.ApplyT(func(v *MetalConnection) MetalConnectionServiceTokenArrayOutput { return v.ServiceTokens }).(MetalConnectionServiceTokenArrayOutput)
}

// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
func (o MetalConnectionOutput) Speed() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Speed }).(pulumi.StringOutput)
}

// Status of the connection resource.
func (o MetalConnectionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// String list of tags.
func (o MetalConnectionOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
//
// Deprecated: token is deprecated. Use service_tokens instead
func (o MetalConnectionOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// Connection type - dedicated or shared.
func (o MetalConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
func (o MetalConnectionOutput) Vlans() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *MetalConnection) pulumi.IntArrayOutput { return v.Vlans }).(pulumi.IntArrayOutput)
}

type MetalConnectionArrayOutput struct{ *pulumi.OutputState }

func (MetalConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalConnection)(nil)).Elem()
}

func (o MetalConnectionArrayOutput) ToMetalConnectionArrayOutput() MetalConnectionArrayOutput {
	return o
}

func (o MetalConnectionArrayOutput) ToMetalConnectionArrayOutputWithContext(ctx context.Context) MetalConnectionArrayOutput {
	return o
}

func (o MetalConnectionArrayOutput) Index(i pulumi.IntInput) MetalConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalConnection {
		return vs[0].([]*MetalConnection)[vs[1].(int)]
	}).(MetalConnectionOutput)
}

type MetalConnectionMapOutput struct{ *pulumi.OutputState }

func (MetalConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalConnection)(nil)).Elem()
}

func (o MetalConnectionMapOutput) ToMetalConnectionMapOutput() MetalConnectionMapOutput {
	return o
}

func (o MetalConnectionMapOutput) ToMetalConnectionMapOutputWithContext(ctx context.Context) MetalConnectionMapOutput {
	return o
}

func (o MetalConnectionMapOutput) MapIndex(k pulumi.StringInput) MetalConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalConnection {
		return vs[0].(map[string]*MetalConnection)[vs[1].(string)]
	}).(MetalConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalConnectionInput)(nil)).Elem(), &MetalConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalConnectionArrayInput)(nil)).Elem(), MetalConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalConnectionMapInput)(nil)).Elem(), MetalConnectionMap{})
	pulumi.RegisterOutputType(MetalConnectionOutput{})
	pulumi.RegisterOutputType(MetalConnectionArrayOutput{})
	pulumi.RegisterOutputType(MetalConnectionMapOutput{})
}
