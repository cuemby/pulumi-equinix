// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage User SSH keys on your Equinix Metal user account. If you create a new device in a project, all the keys of the project's collaborators will be injected to the device.
//
// The link between User SSH key and device is implicit. If you want to make sure that a key will be copied to a device, you must ensure that the device resource `dependsOn` the key resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := equinix.NewMetalSSHKey(ctx, "key1", &equinix.MetalSSHKeyArgs{
// 			PublicKey: readFileOrPanic("/home/terraform/.ssh/id_rsa.pub"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.NewMetalDevice(ctx, "test", &equinix.MetalDeviceArgs{
// 			Hostname:        pulumi.String("test-device"),
// 			Plan:            pulumi.String("c3.small.x86"),
// 			Metro:           pulumi.String("sv"),
// 			OperatingSystem: pulumi.String("ubuntu_20_04"),
// 			BillingCycle:    pulumi.String("hourly"),
// 			ProjectId:       pulumi.Any(local.Project_id),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			pulumi.Resource("equinix_metal_ssh_key.key1"),
// 		}))
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
// This resource can be imported using an existing SSH Key ID
//
// ```sh
//  $ pulumi import equinix:index/metalSSHKey:MetalSSHKey equinix_metal_ssh_key {existing_sshkey_id}
// ```
type MetalSSHKey struct {
	pulumi.CustomResourceState

	// The timestamp for when the SSH key was created.
	Created pulumi.StringOutput `pulumi:"created"`
	// The fingerprint of the SSH key.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUID of the Equinix Metal API User who owns this key.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The timestamp for the last time the SSH key was updated.
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewMetalSSHKey registers a new resource with the given unique name, arguments, and options.
func NewMetalSSHKey(ctx *pulumi.Context,
	name string, args *MetalSSHKeyArgs, opts ...pulumi.ResourceOption) (*MetalSSHKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	var resource MetalSSHKey
	err := ctx.RegisterResource("equinix:index/metalSSHKey:MetalSSHKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetalSSHKey gets an existing MetalSSHKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetalSSHKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetalSSHKeyState, opts ...pulumi.ResourceOption) (*MetalSSHKey, error) {
	var resource MetalSSHKey
	err := ctx.ReadResource("equinix:index/metalSSHKey:MetalSSHKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetalSSHKey resources.
type metalSSHKeyState struct {
	// The timestamp for when the SSH key was created.
	Created *string `pulumi:"created"`
	// The fingerprint of the SSH key.
	Fingerprint *string `pulumi:"fingerprint"`
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The UUID of the Equinix Metal API User who owns this key.
	OwnerId *string `pulumi:"ownerId"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey *string `pulumi:"publicKey"`
	// The timestamp for the last time the SSH key was updated.
	Updated *string `pulumi:"updated"`
}

type MetalSSHKeyState struct {
	// The timestamp for when the SSH key was created.
	Created pulumi.StringPtrInput
	// The fingerprint of the SSH key.
	Fingerprint pulumi.StringPtrInput
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The UUID of the Equinix Metal API User who owns this key.
	OwnerId pulumi.StringPtrInput
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringPtrInput
	// The timestamp for the last time the SSH key was updated.
	Updated pulumi.StringPtrInput
}

func (MetalSSHKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*metalSSHKeyState)(nil)).Elem()
}

type metalSSHKeyArgs struct {
	// The name of the SSH key for identification
	Name *string `pulumi:"name"`
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a MetalSSHKey resource.
type MetalSSHKeyArgs struct {
	// The name of the SSH key for identification
	Name pulumi.StringPtrInput
	// The public key. If this is a file, it
	// can be read using the file interpolation function
	PublicKey pulumi.StringInput
}

func (MetalSSHKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metalSSHKeyArgs)(nil)).Elem()
}

type MetalSSHKeyInput interface {
	pulumi.Input

	ToMetalSSHKeyOutput() MetalSSHKeyOutput
	ToMetalSSHKeyOutputWithContext(ctx context.Context) MetalSSHKeyOutput
}

func (*MetalSSHKey) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalSSHKey)(nil)).Elem()
}

func (i *MetalSSHKey) ToMetalSSHKeyOutput() MetalSSHKeyOutput {
	return i.ToMetalSSHKeyOutputWithContext(context.Background())
}

func (i *MetalSSHKey) ToMetalSSHKeyOutputWithContext(ctx context.Context) MetalSSHKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalSSHKeyOutput)
}

// MetalSSHKeyArrayInput is an input type that accepts MetalSSHKeyArray and MetalSSHKeyArrayOutput values.
// You can construct a concrete instance of `MetalSSHKeyArrayInput` via:
//
//          MetalSSHKeyArray{ MetalSSHKeyArgs{...} }
type MetalSSHKeyArrayInput interface {
	pulumi.Input

	ToMetalSSHKeyArrayOutput() MetalSSHKeyArrayOutput
	ToMetalSSHKeyArrayOutputWithContext(context.Context) MetalSSHKeyArrayOutput
}

type MetalSSHKeyArray []MetalSSHKeyInput

func (MetalSSHKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalSSHKey)(nil)).Elem()
}

func (i MetalSSHKeyArray) ToMetalSSHKeyArrayOutput() MetalSSHKeyArrayOutput {
	return i.ToMetalSSHKeyArrayOutputWithContext(context.Background())
}

func (i MetalSSHKeyArray) ToMetalSSHKeyArrayOutputWithContext(ctx context.Context) MetalSSHKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalSSHKeyArrayOutput)
}

// MetalSSHKeyMapInput is an input type that accepts MetalSSHKeyMap and MetalSSHKeyMapOutput values.
// You can construct a concrete instance of `MetalSSHKeyMapInput` via:
//
//          MetalSSHKeyMap{ "key": MetalSSHKeyArgs{...} }
type MetalSSHKeyMapInput interface {
	pulumi.Input

	ToMetalSSHKeyMapOutput() MetalSSHKeyMapOutput
	ToMetalSSHKeyMapOutputWithContext(context.Context) MetalSSHKeyMapOutput
}

type MetalSSHKeyMap map[string]MetalSSHKeyInput

func (MetalSSHKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalSSHKey)(nil)).Elem()
}

func (i MetalSSHKeyMap) ToMetalSSHKeyMapOutput() MetalSSHKeyMapOutput {
	return i.ToMetalSSHKeyMapOutputWithContext(context.Background())
}

func (i MetalSSHKeyMap) ToMetalSSHKeyMapOutputWithContext(ctx context.Context) MetalSSHKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetalSSHKeyMapOutput)
}

type MetalSSHKeyOutput struct{ *pulumi.OutputState }

func (MetalSSHKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetalSSHKey)(nil)).Elem()
}

func (o MetalSSHKeyOutput) ToMetalSSHKeyOutput() MetalSSHKeyOutput {
	return o
}

func (o MetalSSHKeyOutput) ToMetalSSHKeyOutputWithContext(ctx context.Context) MetalSSHKeyOutput {
	return o
}

// The timestamp for when the SSH key was created.
func (o MetalSSHKeyOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalSSHKey) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The fingerprint of the SSH key.
func (o MetalSSHKeyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalSSHKey) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The name of the SSH key for identification
func (o MetalSSHKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalSSHKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The UUID of the Equinix Metal API User who owns this key.
func (o MetalSSHKeyOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalSSHKey) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The public key. If this is a file, it
// can be read using the file interpolation function
func (o MetalSSHKeyOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalSSHKey) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// The timestamp for the last time the SSH key was updated.
func (o MetalSSHKeyOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v *MetalSSHKey) pulumi.StringOutput { return v.Updated }).(pulumi.StringOutput)
}

type MetalSSHKeyArrayOutput struct{ *pulumi.OutputState }

func (MetalSSHKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetalSSHKey)(nil)).Elem()
}

func (o MetalSSHKeyArrayOutput) ToMetalSSHKeyArrayOutput() MetalSSHKeyArrayOutput {
	return o
}

func (o MetalSSHKeyArrayOutput) ToMetalSSHKeyArrayOutputWithContext(ctx context.Context) MetalSSHKeyArrayOutput {
	return o
}

func (o MetalSSHKeyArrayOutput) Index(i pulumi.IntInput) MetalSSHKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetalSSHKey {
		return vs[0].([]*MetalSSHKey)[vs[1].(int)]
	}).(MetalSSHKeyOutput)
}

type MetalSSHKeyMapOutput struct{ *pulumi.OutputState }

func (MetalSSHKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetalSSHKey)(nil)).Elem()
}

func (o MetalSSHKeyMapOutput) ToMetalSSHKeyMapOutput() MetalSSHKeyMapOutput {
	return o
}

func (o MetalSSHKeyMapOutput) ToMetalSSHKeyMapOutputWithContext(ctx context.Context) MetalSSHKeyMapOutput {
	return o
}

func (o MetalSSHKeyMapOutput) MapIndex(k pulumi.StringInput) MetalSSHKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetalSSHKey {
		return vs[0].(map[string]*MetalSSHKey)[vs[1].(string)]
	}).(MetalSSHKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetalSSHKeyInput)(nil)).Elem(), &MetalSSHKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalSSHKeyArrayInput)(nil)).Elem(), MetalSSHKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetalSSHKeyMapInput)(nil)).Elem(), MetalSSHKeyMap{})
	pulumi.RegisterOutputType(MetalSSHKeyOutput{})
	pulumi.RegisterOutputType(MetalSSHKeyArrayOutput{})
	pulumi.RegisterOutputType(MetalSSHKeyMapOutput{})
}
