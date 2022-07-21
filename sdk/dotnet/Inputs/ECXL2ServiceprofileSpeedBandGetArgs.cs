// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Cuemby.Equinix.Inputs
{

    public sealed class ECXL2ServiceprofileSpeedBandGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Speed/bandwidth supported by this service profile.
        /// </summary>
        [Input("speed", required: true)]
        public Input<int> Speed { get; set; } = null!;

        /// <summary>
        /// Unit of the speed/bandwidth supported by this service profile. One of
        /// `MB`, `GB`.
        /// </summary>
        [Input("speedUnit", required: true)]
        public Input<string> SpeedUnit { get; set; } = null!;

        public ECXL2ServiceprofileSpeedBandGetArgs()
        {
        }
    }
}
