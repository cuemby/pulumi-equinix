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

    public sealed class MetalSpotMarketRequestInstanceParametersArgs : Pulumi.ResourceArgs
    {
        [Input("alwaysPxe")]
        public Input<bool>? AlwaysPxe { get; set; }

        [Input("billingCycle", required: true)]
        public Input<string> BillingCycle { get; set; } = null!;

        [Input("customdata")]
        public Input<string>? Customdata { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("features")]
        private InputList<string>? _features;
        public InputList<string> Features
        {
            get => _features ?? (_features = new InputList<string>());
            set => _features = value;
        }

        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        [Input("ipxeScriptUrl")]
        public Input<string>? IpxeScriptUrl { get; set; }

        /// <summary>
        /// Blocks deletion of the SpotMarketRequest device until the lock is disabled.
        /// </summary>
        [Input("locked")]
        public Input<bool>? Locked { get; set; }

        [Input("operatingSystem", required: true)]
        public Input<string> OperatingSystem { get; set; } = null!;

        [Input("plan", required: true)]
        public Input<string> Plan { get; set; } = null!;

        [Input("projectSshKeys")]
        private InputList<string>? _projectSshKeys;
        public InputList<string> ProjectSshKeys
        {
            get => _projectSshKeys ?? (_projectSshKeys = new InputList<string>());
            set => _projectSshKeys = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("termintationTime")]
        public Input<string>? TermintationTime { get; set; }

        [Input("userSshKeys")]
        private InputList<string>? _userSshKeys;
        public InputList<string> UserSshKeys
        {
            get => _userSshKeys ?? (_userSshKeys = new InputList<string>());
            set => _userSshKeys = value;
        }

        [Input("userdata")]
        public Input<string>? Userdata { get; set; }

        public MetalSpotMarketRequestInstanceParametersArgs()
        {
        }
    }
}
