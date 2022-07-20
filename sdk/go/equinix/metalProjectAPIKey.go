// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create Metal Project API Key resources in Equinix Metal. Project API keys can
// be used to create and read resources in a single project. Each API key contains a token which can
// be used for authentication in Equinix Metal HTTP API (in HTTP request header `X-Auth-Token`).
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
// 	"github.com/pulumi/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := equinix.NewMetalProjectAPIKey(ctx, "test", &equinix.MetalProjectAPIKeyArgs{
// 			ProjectId:   pulumi.Any(local.Existing_project_id),
// 			Description: pulumi.String("Read-only key scoped to a projct"),
// 			ReadOnly:    pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MetalProjectAPIKey struct {
	pulumi.CustomResourceState

	// Description string for the Project API Key resource.
	// * `read-only` - (Optional) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringOutput `pulumi:"description"`
	// UUID of the project where the API key is scoped to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolOutput `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewMetalProjectAPIKey registers a new resource with the given unique name, arguments, and options.
func NewMetalProjectAPIKey(ctx *pulumi.Context,
	name string, args *MetalProjectAPIKeyArgs, opts ...pulumi.ResourceOption) (*MetalProjectAPIKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ReadOnly == nil {
		return nil, errors.New("invalid value for required argument 'ReadOnly'")
	}
	var resource MetalProjectAPIKey
	err := ctx.RegisterResource("equinix:index/metalProjectAPIKey:MetalProjectAPIKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalProjectAPIKey gets an existing MetalProjectAPIKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalProjectAPIKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalProjectAPIKeyState, opts ...pulumi.ResourceOption) (*MetalProjectAPIKey, error) {
	var resource MetalProjectAPIKey
	err := ctx.ReadResource("equinix:index/metalProjectAPIKey:MetalProjectAPIKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalProjectAPIKey resources.
type metalProjectAPIKeyState struct {
	// Description string for the Project API Key resource.
	// * `read-only` - (Optional) Flag indicating whether the API key shoud be read-only.
	Description *string `pulumi:"description"`
	// UUID of the project where the API key is scoped to.
	ProjectId *string `pulumi:"projectId"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly *bool `pulumi:"readOnly"`
	// API token which can be used in Equinix Metal API clients
	Token *string `pulumi:"token"`
}

type MetalProjectAPIKeyState struct {
	// Description string for the Project API Key resource.
	// * `read-only` - (Optional) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringPtrInput
	// UUID of the project where the API key is scoped to.
	ProjectId pulumi.StringPtrInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolPtrInput
	// API token which can be used in Equinix Metal API clients
	Token pulumi.StringPtrInput
}

func (MetalProjectAPIKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalProjectAPIKeyState)(nil)).Elem()
}

type metalProjectAPIKeyArgs struct {
	// Description string for the Project API Key resource.
	// * `read-only` - (Optional) Flag indicating whether the API key shoud be read-only.
	Description string `pulumi:"description"`
	// UUID of the project where the API key is scoped to.
	ProjectId string `pulumi:"projectId"`
	// Flag indicating whether the API key shoud be read-only
	ReadOnly bool `pulumi:"readOnly"`
}

// The set of arguments for constructing a MetalProjectAPIKey resource.
type MetalProjectAPIKeyArgs struct {
	// Description string for the Project API Key resource.
	// * `read-only` - (Optional) Flag indicating whether the API key shoud be read-only.
	Description pulumi.StringInput
	// UUID of the project where the API key is scoped to.
	ProjectId pulumi.StringInput
	// Flag indicating whether the API key shoud be read-only
	ReadOnly pulumi.BoolInput
}

func (MetalProjectAPIKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalProjectAPIKeyArgs)(nil)).Elem()
}

type MetalProjectAPIKeyInput interface {
	pulumi.Input

	ToMetalProjectAPIKeyOutput() MetalProjectAPIKeyOutput
	ToMetalProjectAPIKeyOutputWithContext(ctx context.Context) MetalProjectAPIKeyOutput
}

func (*MetalProjectAPIKey) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalProjectAPIKey)(nil)).Elem()
}

func (i *MetalProjectAPIKey) ToMetalProjectAPIKeyOutput() MetalProjectAPIKeyOutput {
	return i.ToMetalProjectAPIKeyOutputWithContext(context.Background())
}

func (i *MetalProjectAPIKey) ToMetalProjectAPIKeyOutputWithContext(ctx context.Context) MetalProjectAPIKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalProjectAPIKeyOutput)
}

// MetalProjectAPIKeyArrayInput is an input type that accepts MetalProjectAPIKeyArray and MetalProjectAPIKeyArrayOutput values.
// You can construct a concrete instance of `MetalProjectAPIKeyArrayInput` via:
//
//          MetalProjectAPIKeyArray{ MetalProjectAPIKeyArgs{...} }
type MetalProjectAPIKeyArrayInput interface {
	pulumi.Input

	ToMetalProjectAPIKeyArrayOutput() MetalProjectAPIKeyArrayOutput
	ToMetalProjectAPIKeyArrayOutputWithContext(context.Context) MetalProjectAPIKeyArrayOutput
}

type MetalProjectAPIKeyArray []MetalProjectAPIKeyInput

func (MetalProjectAPIKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalProjectAPIKey)(nil)).Elem()
}

func (i MetalProjectAPIKeyArray) ToMetalProjectAPIKeyArrayOutput() MetalProjectAPIKeyArrayOutput {
	return i.ToMetalProjectAPIKeyArrayOutputWithContext(context.Background())
}

func (i MetalProjectAPIKeyArray) ToMetalProjectAPIKeyArrayOutputWithContext(ctx context.Context) MetalProjectAPIKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalProjectAPIKeyArrayOutput)
}

// MetalProjectAPIKeyMapInput is an input type that accepts MetalProjectAPIKeyMap and MetalProjectAPIKeyMapOutput values.
// You can construct a concrete instance of `MetalProjectAPIKeyMapInput` via:
//
//          MetalProjectAPIKeyMap{ "key": MetalProjectAPIKeyArgs{...} }
type MetalProjectAPIKeyMapInput interface {
	pulumi.Input

	ToMetalProjectAPIKeyMapOutput() MetalProjectAPIKeyMapOutput
	ToMetalProjectAPIKeyMapOutputWithContext(context.Context) MetalProjectAPIKeyMapOutput
}

type MetalProjectAPIKeyMap map[string]MetalProjectAPIKeyInput

func (MetalProjectAPIKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalProjectAPIKey)(nil)).Elem()
}

func (i MetalProjectAPIKeyMap) ToMetalProjectAPIKeyMapOutput() MetalProjectAPIKeyMapOutput {
	return i.ToMetalProjectAPIKeyMapOutputWithContext(context.Background())
}

func (i MetalProjectAPIKeyMap) ToMetalProjectAPIKeyMapOutputWithContext(ctx context.Context) MetalProjectAPIKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalProjectAPIKeyMapOutput)
}

type MetalProjectAPIKeyOutput struct{ *pulumi.OutputState }

func (MetalProjectAPIKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalProjectAPIKey)(nil)).Elem()
}

func (o MetalProjectAPIKeyOutput) ToMetalProjectAPIKeyOutput() MetalProjectAPIKeyOutput {
	return o
}

func (o MetalProjectAPIKeyOutput) ToMetalProjectAPIKeyOutputWithContext(ctx context.Context) MetalProjectAPIKeyOutput {
	return o
}

// Description string for the Project API Key resource.
// * `read-only` - (Optional) Flag indicating whether the API key shoud be read-only.
func (o MetalProjectAPIKeyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalProjectAPIKey) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// UUID of the project where the API key is scoped to.
func (o MetalProjectAPIKeyOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalProjectAPIKey) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Flag indicating whether the API key shoud be read-only
func (o MetalProjectAPIKeyOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v *MetalProjectAPIKey) pulumi.BoolOutput { return v.ReadOnly }).(pulumi.BoolOutput)
}

// API token which can be used in Equinix Metal API clients
func (o MetalProjectAPIKeyOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalProjectAPIKey) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

type MetalProjectAPIKeyArrayOutput struct{ *pulumi.OutputState }

func (MetalProjectAPIKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalProjectAPIKey)(nil)).Elem()
}

func (o MetalProjectAPIKeyArrayOutput) ToMetalProjectAPIKeyArrayOutput() MetalProjectAPIKeyArrayOutput {
	return o
}

func (o MetalProjectAPIKeyArrayOutput) ToMetalProjectAPIKeyArrayOutputWithContext(ctx context.Context) MetalProjectAPIKeyArrayOutput {
	return o
}

func (o MetalProjectAPIKeyArrayOutput) Index(i pulumi.IntInput) MetalProjectAPIKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalProjectAPIKey {
		return vs[0].([]*MetalProjectAPIKey)[vs[1].(int)]
	}).(MetalProjectAPIKeyOutput)
}

type MetalProjectAPIKeyMapOutput struct{ *pulumi.OutputState }

func (MetalProjectAPIKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalProjectAPIKey)(nil)).Elem()
}

func (o MetalProjectAPIKeyMapOutput) ToMetalProjectAPIKeyMapOutput() MetalProjectAPIKeyMapOutput {
	return o
}

func (o MetalProjectAPIKeyMapOutput) ToMetalProjectAPIKeyMapOutputWithContext(ctx context.Context) MetalProjectAPIKeyMapOutput {
	return o
}

func (o MetalProjectAPIKeyMapOutput) MapIndex(k pulumi.StringInput) MetalProjectAPIKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalProjectAPIKey {
		return vs[0].(map[string]*MetalProjectAPIKey)[vs[1].(string)]
	}).(MetalProjectAPIKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalProjectAPIKeyInput)(nil)).Elem(), &MetalProjectAPIKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalProjectAPIKeyArrayInput)(nil)).Elem(), MetalProjectAPIKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalProjectAPIKeyMapInput)(nil)).Elem(), MetalProjectAPIKeyMap{})
	pulumi.RegisterOutputType(MetalProjectAPIKeyOutput{})
	pulumi.RegisterOutputType(MetalProjectAPIKeyArrayOutput{})
	pulumi.RegisterOutputType(MetalProjectAPIKeyMapOutput{})
}
