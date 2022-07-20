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
    public sealed class GetNetworkDeviceClusterDetailNode0Result
    {
        /// <summary>
        /// Unique identifier of applied license file
        /// </summary>
        public readonly string LicenseFileId;
        public readonly string LicenseToken;
        /// <summary>
        /// Name of an existing Equinix Network Edge device
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// UUID of an existing Equinix Network Edge device
        /// </summary>
        public readonly string Uuid;
        public readonly ImmutableArray<Outputs.GetNetworkDeviceClusterDetailNode0VendorConfigurationResult> VendorConfigurations;

        [OutputConstructor]
        private GetNetworkDeviceClusterDetailNode0Result(
            string licenseFileId,

            string licenseToken,

            string name,

            string uuid,

            ImmutableArray<Outputs.GetNetworkDeviceClusterDetailNode0VendorConfigurationResult> vendorConfigurations)
        {
            LicenseFileId = licenseFileId;
            LicenseToken = licenseToken;
            Name = name;
            Uuid = uuid;
            VendorConfigurations = vendorConfigurations;
        }
    }
}
