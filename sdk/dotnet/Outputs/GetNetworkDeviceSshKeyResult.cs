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
    public sealed class GetNetworkDeviceSshKeyResult
    {
        public readonly string KeyName;
        public readonly string Username;

        [OutputConstructor]
        private GetNetworkDeviceSshKeyResult(
            string keyName,

            string username)
        {
            KeyName = keyName;
            Username = username;
        }
    }
}
