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
    public sealed class GetECXL2SellerprofileSpeedBandResult
    {
        /// <summary>
        /// Speed/bandwidth supported by given service profile.
        /// </summary>
        public readonly int Speed;
        /// <summary>
        /// Unit of the speed/bandwidth supported by given service profile.
        /// </summary>
        public readonly string SpeedUnit;

        [OutputConstructor]
        private GetECXL2SellerprofileSpeedBandResult(
            int speed,

            string speedUnit)
        {
            Speed = speed;
            SpeedUnit = speedUnit;
        }
    }
}
