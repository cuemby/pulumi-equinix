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

    public sealed class ECXL2ConnectionSecondaryConnectionActionGetArgs : Pulumi.ResourceArgs
    {
        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("operationId")]
        public Input<string>? OperationId { get; set; }

        [Input("requiredDatas")]
        private InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionRequiredDataGetArgs>? _requiredDatas;
        public InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionRequiredDataGetArgs> RequiredDatas
        {
            get => _requiredDatas ?? (_requiredDatas = new InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionRequiredDataGetArgs>());
            set => _requiredDatas = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ECXL2ConnectionSecondaryConnectionActionGetArgs()
        {
        }
    }
}
