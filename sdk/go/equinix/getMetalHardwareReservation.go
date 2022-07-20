// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a [hardware reservation resource from Equinix Metal](https://metal.equinix.com/developers/docs/deploy/reserved/).
//
// You can look up hardware reservation by its ID or by ID of device which occupies it.
func GetMetalHardwareReservation(ctx *pulumi.Context, args *GetMetalHardwareReservationArgs, opts ...pulumi.InvokeOption) (*GetMetalHardwareReservationResult, error) {
	var rv GetMetalHardwareReservationResult
	err := ctx.Invoke("equinix:index/getMetalHardwareReservation:GetMetalHardwareReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetMetalHardwareReservation.
type GetMetalHardwareReservationArgs struct {
	// UUID of device occupying the reservation.
	DeviceId *string `pulumi:"deviceId"`
	// ID of the hardware reservation.
	Id *string `pulumi:"id"`
}

// A collection of values returned by GetMetalHardwareReservation.
type GetMetalHardwareReservationResult struct {
	// UUID of device occupying the reservation.
	DeviceId string `pulumi:"deviceId"`
	// Plan type for the reservation.
	Facility string `pulumi:"facility"`
	// ID of the hardware reservation to look up.
	Id string `pulumi:"id"`
	// Plan type for the reservation.
	Plan string `pulumi:"plan"`
	// UUID of project this reservation is scoped to.
	ProjectId string `pulumi:"projectId"`
	// Flag indicating whether the reserved server is provisionable or not. Spare
	// devices can't be provisioned unless they are activated first.
	Provisionable bool `pulumi:"provisionable"`
	// Reservation short ID.
	ShortId string `pulumi:"shortId"`
	// Flag indicating whether the Hardware Reservation is a spare. Spare Hardware
	// Reservations are used when a Hardware Reservations requires service from Metal Equinix.
	Spare bool `pulumi:"spare"`
	// Switch short ID, can be used to determine if two devices are connected to the
	// same switch.
	SwitchUuid string `pulumi:"switchUuid"`
}

func GetMetalHardwareReservationOutput(ctx *pulumi.Context, args GetMetalHardwareReservationOutputArgs, opts ...pulumi.InvokeOption) GetMetalHardwareReservationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMetalHardwareReservationResult, error) {
			args := v.(GetMetalHardwareReservationArgs)
			r, err := GetMetalHardwareReservation(ctx, &args, opts...)
			var s GetMetalHardwareReservationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMetalHardwareReservationResultOutput)
}

// A collection of arguments for invoking GetMetalHardwareReservation.
type GetMetalHardwareReservationOutputArgs struct {
	// UUID of device occupying the reservation.
	DeviceId pulumi.StringPtrInput `pulumi:"deviceId"`
	// ID of the hardware reservation.
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (GetMetalHardwareReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMetalHardwareReservationArgs)(nil)).Elem()
}

// A collection of values returned by GetMetalHardwareReservation.
type GetMetalHardwareReservationResultOutput struct{ *pulumi.OutputState }

func (GetMetalHardwareReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMetalHardwareReservationResult)(nil)).Elem()
}

func (o GetMetalHardwareReservationResultOutput) ToGetMetalHardwareReservationResultOutput() GetMetalHardwareReservationResultOutput {
	return o
}

func (o GetMetalHardwareReservationResultOutput) ToGetMetalHardwareReservationResultOutputWithContext(ctx context.Context) GetMetalHardwareReservationResultOutput {
	return o
}

// UUID of device occupying the reservation.
func (o GetMetalHardwareReservationResultOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) string { return v.DeviceId }).(pulumi.StringOutput)
}

// Plan type for the reservation.
func (o GetMetalHardwareReservationResultOutput) Facility() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) string { return v.Facility }).(pulumi.StringOutput)
}

// ID of the hardware reservation to look up.
func (o GetMetalHardwareReservationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Plan type for the reservation.
func (o GetMetalHardwareReservationResultOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) string { return v.Plan }).(pulumi.StringOutput)
}

// UUID of project this reservation is scoped to.
func (o GetMetalHardwareReservationResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Flag indicating whether the reserved server is provisionable or not. Spare
// devices can't be provisioned unless they are activated first.
func (o GetMetalHardwareReservationResultOutput) Provisionable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) bool { return v.Provisionable }).(pulumi.BoolOutput)
}

// Reservation short ID.
func (o GetMetalHardwareReservationResultOutput) ShortId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) string { return v.ShortId }).(pulumi.StringOutput)
}

// Flag indicating whether the Hardware Reservation is a spare. Spare Hardware
// Reservations are used when a Hardware Reservations requires service from Metal Equinix.
func (o GetMetalHardwareReservationResultOutput) Spare() pulumi.BoolOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) bool { return v.Spare }).(pulumi.BoolOutput)
}

// Switch short ID, can be used to determine if two devices are connected to the
// same switch.
func (o GetMetalHardwareReservationResultOutput) SwitchUuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetMetalHardwareReservationResult) string { return v.SwitchUuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMetalHardwareReservationResultOutput{})
}
