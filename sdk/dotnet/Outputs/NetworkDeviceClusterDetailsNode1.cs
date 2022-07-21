// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Cuemby.Equinix.Outputs
{

    [OutputType]
    public sealed class NetworkDeviceClusterDetailsNode1
    {
        /// <summary>
        /// License file id. This is necessary for Fortinet and Juniper clusters.
        /// </summary>
        public readonly string? LicenseFileId;
        /// <summary>
        /// License token. This is necessary for Palo Alto clusters.
        /// </summary>
        public readonly string? LicenseToken;
        /// <summary>
        /// Device name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Device unique identifier.
        /// </summary>
        public readonly string? Uuid;
        /// <summary>
        /// An object that has fields relevant to the vendor of the
        /// cluster device. See Cluster Details - Nodes - Vendor Configuration
        /// below for more details.
        /// </summary>
        public readonly Outputs.NetworkDeviceClusterDetailsNode1VendorConfiguration? VendorConfiguration;

        [OutputConstructor]
        private NetworkDeviceClusterDetailsNode1(
            string? licenseFileId,

            string? licenseToken,

            string? name,

            string? uuid,

            Outputs.NetworkDeviceClusterDetailsNode1VendorConfiguration? vendorConfiguration)
        {
            LicenseFileId = licenseFileId;
            LicenseToken = licenseToken;
            Name = name;
            Uuid = uuid;
            VendorConfiguration = vendorConfiguration;
        }
    }
}
