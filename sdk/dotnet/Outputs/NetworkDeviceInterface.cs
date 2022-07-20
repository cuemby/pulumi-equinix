// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Outputs
{

    [OutputType]
    public sealed class NetworkDeviceInterface
    {
        /// <summary>
        /// interface management type (Equinix Managed or empty).
        /// </summary>
        public readonly string? AssignedType;
        /// <summary>
        /// interface identifier.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// interface IP address.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// interface MAC address.
        /// </summary>
        public readonly string? MacAddress;
        /// <summary>
        /// Device name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// interface operational status. One of `up`, `down`.
        /// </summary>
        public readonly string? OperationalStatus;
        /// <summary>
        /// interface status. One of `AVAILABLE`, `RESERVED`, `ASSIGNED`.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// interface type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private NetworkDeviceInterface(
            string? assignedType,

            int? id,

            string? ipAddress,

            string? macAddress,

            string? name,

            string? operationalStatus,

            string? status,

            string? type)
        {
            AssignedType = assignedType;
            Id = id;
            IpAddress = ipAddress;
            MacAddress = macAddress;
            Name = name;
            OperationalStatus = operationalStatus;
            Status = status;
            Type = type;
        }
    }
}
