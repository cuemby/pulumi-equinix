// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Cuemby.Equinix
{
    public static class GetMetalHardwareReservation
    {
        /// <summary>
        /// Use this data source to retrieve a [hardware reservation resource from Equinix Metal](https://metal.equinix.com/developers/docs/deploy/reserved/).
        /// 
        /// You can look up hardware reservation by its ID or by ID of device which occupies it.
        /// </summary>
        public static Task<GetMetalHardwareReservationResult> InvokeAsync(GetMetalHardwareReservationArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMetalHardwareReservationResult>("equinix:index/getMetalHardwareReservation:GetMetalHardwareReservation", args ?? new GetMetalHardwareReservationArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve a [hardware reservation resource from Equinix Metal](https://metal.equinix.com/developers/docs/deploy/reserved/).
        /// 
        /// You can look up hardware reservation by its ID or by ID of device which occupies it.
        /// </summary>
        public static Output<GetMetalHardwareReservationResult> Invoke(GetMetalHardwareReservationInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMetalHardwareReservationResult>("equinix:index/getMetalHardwareReservation:GetMetalHardwareReservation", args ?? new GetMetalHardwareReservationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetalHardwareReservationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of device occupying the reservation.
        /// </summary>
        [Input("deviceId")]
        public string? DeviceId { get; set; }

        /// <summary>
        /// ID of the hardware reservation.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        public GetMetalHardwareReservationArgs()
        {
        }
    }

    public sealed class GetMetalHardwareReservationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// UUID of device occupying the reservation.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// ID of the hardware reservation.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GetMetalHardwareReservationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMetalHardwareReservationResult
    {
        /// <summary>
        /// UUID of device occupying the reservation.
        /// </summary>
        public readonly string DeviceId;
        /// <summary>
        /// Plan type for the reservation.
        /// </summary>
        public readonly string Facility;
        /// <summary>
        /// ID of the hardware reservation to look up.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Plan type for the reservation.
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// UUID of project this reservation is scoped to.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Flag indicating whether the reserved server is provisionable or not. Spare
        /// devices can't be provisioned unless they are activated first.
        /// </summary>
        public readonly bool Provisionable;
        /// <summary>
        /// Reservation short ID.
        /// </summary>
        public readonly string ShortId;
        /// <summary>
        /// Flag indicating whether the Hardware Reservation is a spare. Spare Hardware
        /// Reservations are used when a Hardware Reservations requires service from Metal Equinix.
        /// </summary>
        public readonly bool Spare;
        /// <summary>
        /// Switch short ID, can be used to determine if two devices are connected to the
        /// same switch.
        /// </summary>
        public readonly string SwitchUuid;

        [OutputConstructor]
        private GetMetalHardwareReservationResult(
            string deviceId,

            string facility,

            string id,

            string plan,

            string projectId,

            bool provisionable,

            string shortId,

            bool spare,

            string switchUuid)
        {
            DeviceId = deviceId;
            Facility = facility;
            Id = id;
            Plan = plan;
            ProjectId = projectId;
            Provisionable = provisionable;
            ShortId = shortId;
            Spare = spare;
            SwitchUuid = switchUuid;
        }
    }
}
