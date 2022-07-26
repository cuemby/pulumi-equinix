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
    public sealed class GetECXL2SellerprofileMetroResult
    {
        /// <summary>
        /// Location metro code.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// List of IBXes supported within given metro.
        /// </summary>
        public readonly ImmutableArray<string> Ibxes;
        /// <summary>
        /// Name of the seller profile.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of regions supported within given.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Regions;

        [OutputConstructor]
        private GetECXL2SellerprofileMetroResult(
            string code,

            ImmutableArray<string> ibxes,

            string name,

            ImmutableDictionary<string, string> regions)
        {
            Code = code;
            Ibxes = ibxes;
            Name = name;
            Regions = regions;
        }
    }
}
