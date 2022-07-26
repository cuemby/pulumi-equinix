// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create Metal User API Key resources in Equinix Metal. Each API key contains a
// token which can be used for authentication in Equinix Metal HTTP API (in HTTP request header
// `X-Auth-Token`).
//
// Read-only keys only allow to list and view existing resources, read-write keys can also be used to
// create resources.
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
// 		_, err := equinix.NewMetalUserAPIKey(ctx, "test", &equinix.MetalUserAPIKeyArgs{
// 			Description: pulumi.String("Read-only user key"),
// 			ReadOnly:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MetalUserAPIKey struct {
	pulumi.CustomResourceState

	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringOutput `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients.
	Token pulumi.StringOutput `pulumi:"token"`
	// UUID of the owner of the API key.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewMetalUserAPIKey registers a new resource with the given unique name, arguments, and options.
func NewMetalUserAPIKey(ctx *pulumi.Context,
	name string, args *MetalUserAPIKeyArgs, opts ...pulumi.ResourceOption) (*MetalUserAPIKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ReadOnly == nil {
		return nil, errors.New("invalid value for required argument 'ReadOnly'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MetalUserAPIKey
	err := ctx.RegisterResource("equinix:index/metalUserAPIKey:MetalUserAPIKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalUserAPIKey gets an existing MetalUserAPIKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalUserAPIKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalUserAPIKeyState, opts ...pulumi.ResourceOption) (*MetalUserAPIKey, error) {
	var resource MetalUserAPIKey
	err := ctx.ReadResource("equinix:index/metalUserAPIKey:MetalUserAPIKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalUserAPIKey resources.
type metalUserAPIKeyState struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description *string `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly *bool `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients.
	Token *string `pulumi:"token"`
	// UUID of the owner of the API key.
	UserId *string `pulumi:"userId"`
}

type MetalUserAPIKeyState struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringPtrInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolPtrInput
	// API token which can be used in Equinix Metal API clients.
	Token pulumi.StringPtrInput
	// UUID of the owner of the API key.
	UserId pulumi.StringPtrInput
}

func (MetalUserAPIKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalUserAPIKeyState)(nil)).Elem()
}

type metalUserAPIKeyArgs struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description string `pulumi:"description"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly bool `pulumi:"readOnly"`
}

// The set of arguments for constructing a MetalUserAPIKey resource.
type MetalUserAPIKeyArgs struct {
	// Description string for the User API Key resource.
	// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolInput
}

func (MetalUserAPIKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalUserAPIKeyArgs)(nil)).Elem()
}

type MetalUserAPIKeyInput interface {
	pulumi.Input

	ToMetalUserAPIKeyOutput() MetalUserAPIKeyOutput
	ToMetalUserAPIKeyOutputWithContext(ctx context.Context) MetalUserAPIKeyOutput
}

func (*MetalUserAPIKey) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalUserAPIKey)(nil)).Elem()
}

func (i *MetalUserAPIKey) ToMetalUserAPIKeyOutput() MetalUserAPIKeyOutput {
	return i.ToMetalUserAPIKeyOutputWithContext(context.Background())
}

func (i *MetalUserAPIKey) ToMetalUserAPIKeyOutputWithContext(ctx context.Context) MetalUserAPIKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalUserAPIKeyOutput)
}

// MetalUserAPIKeyArrayInput is an input type that accepts MetalUserAPIKeyArray and MetalUserAPIKeyArrayOutput values.
// You can construct a concrete instance of `MetalUserAPIKeyArrayInput` via:
//
//          MetalUserAPIKeyArray{ MetalUserAPIKeyArgs{...} }
type MetalUserAPIKeyArrayInput interface {
	pulumi.Input

	ToMetalUserAPIKeyArrayOutput() MetalUserAPIKeyArrayOutput
	ToMetalUserAPIKeyArrayOutputWithContext(context.Context) MetalUserAPIKeyArrayOutput
}

type MetalUserAPIKeyArray []MetalUserAPIKeyInput

func (MetalUserAPIKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalUserAPIKey)(nil)).Elem()
}

func (i MetalUserAPIKeyArray) ToMetalUserAPIKeyArrayOutput() MetalUserAPIKeyArrayOutput {
	return i.ToMetalUserAPIKeyArrayOutputWithContext(context.Background())
}

func (i MetalUserAPIKeyArray) ToMetalUserAPIKeyArrayOutputWithContext(ctx context.Context) MetalUserAPIKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalUserAPIKeyArrayOutput)
}

// MetalUserAPIKeyMapInput is an input type that accepts MetalUserAPIKeyMap and MetalUserAPIKeyMapOutput values.
// You can construct a concrete instance of `MetalUserAPIKeyMapInput` via:
//
//          MetalUserAPIKeyMap{ "key": MetalUserAPIKeyArgs{...} }
type MetalUserAPIKeyMapInput interface {
	pulumi.Input

	ToMetalUserAPIKeyMapOutput() MetalUserAPIKeyMapOutput
	ToMetalUserAPIKeyMapOutputWithContext(context.Context) MetalUserAPIKeyMapOutput
}

type MetalUserAPIKeyMap map[string]MetalUserAPIKeyInput

func (MetalUserAPIKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalUserAPIKey)(nil)).Elem()
}

func (i MetalUserAPIKeyMap) ToMetalUserAPIKeyMapOutput() MetalUserAPIKeyMapOutput {
	return i.ToMetalUserAPIKeyMapOutputWithContext(context.Background())
}

func (i MetalUserAPIKeyMap) ToMetalUserAPIKeyMapOutputWithContext(ctx context.Context) MetalUserAPIKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalUserAPIKeyMapOutput)
}

type MetalUserAPIKeyOutput struct{ *pulumi.OutputState }

func (MetalUserAPIKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalUserAPIKey)(nil)).Elem()
}

func (o MetalUserAPIKeyOutput) ToMetalUserAPIKeyOutput() MetalUserAPIKeyOutput {
	return o
}

func (o MetalUserAPIKeyOutput) ToMetalUserAPIKeyOutputWithContext(ctx context.Context) MetalUserAPIKeyOutput {
	return o
}

// Description string for the User API Key resource.
// * `read-only` - (Required) Flag indicating whether the API key shoud be read-only.
func (o MetalUserAPIKeyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalUserAPIKey) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Flag indicating whether the API key shoud be read-only
func (o MetalUserAPIKeyOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *MetalUserAPIKey) pulumi.BoolOutput { return v.ReadOnly }).(pulumi.BoolOutput)
}

// API token which can be used in Equinix Metal API clients.
func (o MetalUserAPIKeyOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalUserAPIKey) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// UUID of the owner of the API key.
func (o MetalUserAPIKeyOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalUserAPIKey) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type MetalUserAPIKeyArrayOutput struct{ *pulumi.OutputState }

func (MetalUserAPIKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalUserAPIKey)(nil)).Elem()
}

func (o MetalUserAPIKeyArrayOutput) ToMetalUserAPIKeyArrayOutput() MetalUserAPIKeyArrayOutput {
	return o
}

func (o MetalUserAPIKeyArrayOutput) ToMetalUserAPIKeyArrayOutputWithContext(ctx context.Context) MetalUserAPIKeyArrayOutput {
	return o
}

func (o MetalUserAPIKeyArrayOutput) Index(i pulumi.IntInput) MetalUserAPIKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalUserAPIKey {
		return vs[0].([]*MetalUserAPIKey)[vs[1].(int)]
	}).(MetalUserAPIKeyOutput)
}

type MetalUserAPIKeyMapOutput struct{ *pulumi.OutputState }

func (MetalUserAPIKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalUserAPIKey)(nil)).Elem()
}

func (o MetalUserAPIKeyMapOutput) ToMetalUserAPIKeyMapOutput() MetalUserAPIKeyMapOutput {
	return o
}

func (o MetalUserAPIKeyMapOutput) ToMetalUserAPIKeyMapOutputWithContext(ctx context.Context) MetalUserAPIKeyMapOutput {
	return o
}

func (o MetalUserAPIKeyMapOutput) MapIndex(k pulumi.StringInput) MetalUserAPIKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalUserAPIKey {
		return vs[0].(map[string]*MetalUserAPIKey)[vs[1].(string)]
	}).(MetalUserAPIKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalUserAPIKeyInput)(nil)).Elem(), &MetalUserAPIKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalUserAPIKeyArrayInput)(nil)).Elem(), MetalUserAPIKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalUserAPIKeyMapInput)(nil)).Elem(), MetalUserAPIKeyMap{})
	pulumi.RegisterOutputType(MetalUserAPIKeyOutput{})
	pulumi.RegisterOutputType(MetalUserAPIKeyArrayOutput{})
	pulumi.RegisterOutputType(MetalUserAPIKeyMapOutput{})
}
