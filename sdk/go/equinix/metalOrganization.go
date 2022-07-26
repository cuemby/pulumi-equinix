// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage organization resource in Equinix Metal.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/cuemby/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := equinix.NewMetalOrganization(ctx, "tfOrganization1", &equinix.MetalOrganizationArgs{
// 			Description: pulumi.String("quux"),
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
// This resource can be imported using an existing organization ID
//
// ```sh
//  $ pulumi import equinix:index/metalOrganization:MetalOrganization equinix_metal_organization {existing_organization_id}
// ```
type MetalOrganization struct {
	pulumi.CustomResourceState

	// Postal address.
	Address MetalOrganizationAddressOutput `pulumi:"address"`
	// The timestamp for when the organization was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// Description string.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Logo URL.
	Logo pulumi.StringPtrOutput `pulumi:"logo"`
	// The name of the Organization.
	Name pulumi.StringOutput `pulumi:"name"`
	// Twitter handle.
	Twitter pulumi.StringPtrOutput `pulumi:"twitter"`
	// The timestamp for the last time the organization was updated.
	Updated pulumi.StringOutput `pulumi:"updated"`
	// Website link.
	Website pulumi.StringPtrOutput `pulumi:"website"`
}

// NewMetalOrganization registers a new resource with the given unique name, arguments, and options.
func NewMetalOrganization(ctx *pulumi.Context,
	name string, args *MetalOrganizationArgs, opts ...pulumi.ResourceOption) (*MetalOrganization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MetalOrganization
	err := ctx.RegisterResource("equinix:index/metalOrganization:MetalOrganization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalOrganization gets an existing MetalOrganization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalOrganizationState, opts ...pulumi.ResourceOption) (*MetalOrganization, error) {
	var resource MetalOrganization
	err := ctx.ReadResource("equinix:index/metalOrganization:MetalOrganization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalOrganization resources.
type metalOrganizationState struct {
	// Postal address.
	Address *MetalOrganizationAddress `pulumi:"address"`
	// The timestamp for when the organization was created.
	Created *string `pulumi:"created"`
	// Description string.
	Description *string `pulumi:"description"`
	// Logo URL.
	Logo *string `pulumi:"logo"`
	// The name of the Organization.
	Name *string `pulumi:"name"`
	// Twitter handle.
	Twitter *string `pulumi:"twitter"`
	// The timestamp for the last time the organization was updated.
	Updated *string `pulumi:"updated"`
	// Website link.
	Website *string `pulumi:"website"`
}

type MetalOrganizationState struct {
	// Postal address.
	Address MetalOrganizationAddressPtrInput
	// The timestamp for when the organization was created.
	Created pulumi.StringPtrInput
	// Description string.
	Description pulumi.StringPtrInput
	// Logo URL.
	Logo pulumi.StringPtrInput
	// The name of the Organization.
	Name pulumi.StringPtrInput
	// Twitter handle.
	Twitter pulumi.StringPtrInput
	// The timestamp for the last time the organization was updated.
	Updated pulumi.StringPtrInput
	// Website link.
	Website pulumi.StringPtrInput
}

func (MetalOrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalOrganizationState)(nil)).Elem()
}

type metalOrganizationArgs struct {
	// Postal address.
	Address MetalOrganizationAddress `pulumi:"address"`
	// Description string.
	Description *string `pulumi:"description"`
	// Logo URL.
	Logo *string `pulumi:"logo"`
	// The name of the Organization.
	Name *string `pulumi:"name"`
	// Twitter handle.
	Twitter *string `pulumi:"twitter"`
	// Website link.
	Website *string `pulumi:"website"`
}

// The set of arguments for constructing a MetalOrganization resource.
type MetalOrganizationArgs struct {
	// Postal address.
	Address MetalOrganizationAddressInput
	// Description string.
	Description pulumi.StringPtrInput
	// Logo URL.
	Logo pulumi.StringPtrInput
	// The name of the Organization.
	Name pulumi.StringPtrInput
	// Twitter handle.
	Twitter pulumi.StringPtrInput
	// Website link.
	Website pulumi.StringPtrInput
}

func (MetalOrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalOrganizationArgs)(nil)).Elem()
}

type MetalOrganizationInput interface {
	pulumi.Input

	ToMetalOrganizationOutput() MetalOrganizationOutput
	ToMetalOrganizationOutputWithContext(ctx context.Context) MetalOrganizationOutput
}

func (*MetalOrganization) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalOrganization)(nil)).Elem()
}

func (i *MetalOrganization) ToMetalOrganizationOutput() MetalOrganizationOutput {
	return i.ToMetalOrganizationOutputWithContext(context.Background())
}

func (i *MetalOrganization) ToMetalOrganizationOutputWithContext(ctx context.Context) MetalOrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalOrganizationOutput)
}

// MetalOrganizationArrayInput is an input type that accepts MetalOrganizationArray and MetalOrganizationArrayOutput values.
// You can construct a concrete instance of `MetalOrganizationArrayInput` via:
//
//          MetalOrganizationArray{ MetalOrganizationArgs{...} }
type MetalOrganizationArrayInput interface {
	pulumi.Input

	ToMetalOrganizationArrayOutput() MetalOrganizationArrayOutput
	ToMetalOrganizationArrayOutputWithContext(context.Context) MetalOrganizationArrayOutput
}

type MetalOrganizationArray []MetalOrganizationInput

func (MetalOrganizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalOrganization)(nil)).Elem()
}

func (i MetalOrganizationArray) ToMetalOrganizationArrayOutput() MetalOrganizationArrayOutput {
	return i.ToMetalOrganizationArrayOutputWithContext(context.Background())
}

func (i MetalOrganizationArray) ToMetalOrganizationArrayOutputWithContext(ctx context.Context) MetalOrganizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalOrganizationArrayOutput)
}

// MetalOrganizationMapInput is an input type that accepts MetalOrganizationMap and MetalOrganizationMapOutput values.
// You can construct a concrete instance of `MetalOrganizationMapInput` via:
//
//          MetalOrganizationMap{ "key": MetalOrganizationArgs{...} }
type MetalOrganizationMapInput interface {
	pulumi.Input

	ToMetalOrganizationMapOutput() MetalOrganizationMapOutput
	ToMetalOrganizationMapOutputWithContext(context.Context) MetalOrganizationMapOutput
}

type MetalOrganizationMap map[string]MetalOrganizationInput

func (MetalOrganizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalOrganization)(nil)).Elem()
}

func (i MetalOrganizationMap) ToMetalOrganizationMapOutput() MetalOrganizationMapOutput {
	return i.ToMetalOrganizationMapOutputWithContext(context.Background())
}

func (i MetalOrganizationMap) ToMetalOrganizationMapOutputWithContext(ctx context.Context) MetalOrganizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalOrganizationMapOutput)
}

type MetalOrganizationOutput struct{ *pulumi.OutputState }

func (MetalOrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalOrganization)(nil)).Elem()
}

func (o MetalOrganizationOutput) ToMetalOrganizationOutput() MetalOrganizationOutput {
	return o
}

func (o MetalOrganizationOutput) ToMetalOrganizationOutputWithContext(ctx context.Context) MetalOrganizationOutput {
	return o
}

// Postal address.
func (o MetalOrganizationOutput) Address() MetalOrganizationAddressOutput {
	return o.ApplyT(func(v *MetalOrganization) MetalOrganizationAddressOutput { return v.Address }).(MetalOrganizationAddressOutput)
}

// The timestamp for when the organization was created.
func (o MetalOrganizationOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalOrganization) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Description string.
func (o MetalOrganizationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalOrganization) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Logo URL.
func (o MetalOrganizationOutput) Logo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalOrganization) pulumi.StringPtrOutput { return v.Logo }).(pulumi.StringPtrOutput)
}

// The name of the Organization.
func (o MetalOrganizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalOrganization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Twitter handle.
func (o MetalOrganizationOutput) Twitter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalOrganization) pulumi.StringPtrOutput { return v.Twitter }).(pulumi.StringPtrOutput)
}

// The timestamp for the last time the organization was updated.
func (o MetalOrganizationOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalOrganization) pulumi.StringOutput { return v.Updated }).(pulumi.StringOutput)
}

// Website link.
func (o MetalOrganizationOutput) Website() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetalOrganization) pulumi.StringPtrOutput { return v.Website }).(pulumi.StringPtrOutput)
}

type MetalOrganizationArrayOutput struct{ *pulumi.OutputState }

func (MetalOrganizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalOrganization)(nil)).Elem()
}

func (o MetalOrganizationArrayOutput) ToMetalOrganizationArrayOutput() MetalOrganizationArrayOutput {
	return o
}

func (o MetalOrganizationArrayOutput) ToMetalOrganizationArrayOutputWithContext(ctx context.Context) MetalOrganizationArrayOutput {
	return o
}

func (o MetalOrganizationArrayOutput) Index(i pulumi.IntInput) MetalOrganizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalOrganization {
		return vs[0].([]*MetalOrganization)[vs[1].(int)]
	}).(MetalOrganizationOutput)
}

type MetalOrganizationMapOutput struct{ *pulumi.OutputState }

func (MetalOrganizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalOrganization)(nil)).Elem()
}

func (o MetalOrganizationMapOutput) ToMetalOrganizationMapOutput() MetalOrganizationMapOutput {
	return o
}

func (o MetalOrganizationMapOutput) ToMetalOrganizationMapOutputWithContext(ctx context.Context) MetalOrganizationMapOutput {
	return o
}

func (o MetalOrganizationMapOutput) MapIndex(k pulumi.StringInput) MetalOrganizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalOrganization {
		return vs[0].(map[string]*MetalOrganization)[vs[1].(string)]
	}).(MetalOrganizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalOrganizationInput)(nil)).Elem(), &MetalOrganization{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalOrganizationArrayInput)(nil)).Elem(), MetalOrganizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalOrganizationMapInput)(nil)).Elem(), MetalOrganizationMap{})
	pulumi.RegisterOutputType(MetalOrganizationOutput{})
	pulumi.RegisterOutputType(MetalOrganizationArrayOutput{})
	pulumi.RegisterOutputType(MetalOrganizationMapOutput{})
}
