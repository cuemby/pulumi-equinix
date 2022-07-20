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
    public sealed class MetalConnectionServiceToken
    {
        public readonly string? ExpiresAt;
        public readonly string? Id;
        public readonly string? MaxAllowedSpeed;
        public readonly string? Role;
        public readonly string? State;
        /// <summary>
        /// Connection type - dedicated or shared.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private MetalConnectionServiceToken(
            string? expiresAt,

            string? id,

            string? maxAllowedSpeed,

            string? role,

            string? state,

            string? type)
        {
            ExpiresAt = expiresAt;
            Id = id;
            MaxAllowedSpeed = maxAllowedSpeed;
            Role = role;
            State = state;
            Type = type;
        }
    }
}
