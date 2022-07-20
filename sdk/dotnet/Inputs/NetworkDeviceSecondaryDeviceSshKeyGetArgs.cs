// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Inputs
{

    public sealed class NetworkDeviceSecondaryDeviceSshKeyGetArgs : Pulumi.ResourceArgs
    {
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// username associated with given key.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public NetworkDeviceSecondaryDeviceSshKeyGetArgs()
        {
        }
    }
}
