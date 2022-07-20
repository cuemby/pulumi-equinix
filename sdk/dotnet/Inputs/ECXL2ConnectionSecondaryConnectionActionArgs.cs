// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix.Inputs
{

    public sealed class ECXL2ConnectionSecondaryConnectionActionArgs : Pulumi.ResourceArgs
    {
        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("operationId")]
        public Input<string>? OperationId { get; set; }

        [Input("requiredDatas")]
        private InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionRequiredDataArgs>? _requiredDatas;
        public InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionRequiredDataArgs> RequiredDatas
        {
            get => _requiredDatas ?? (_requiredDatas = new InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionRequiredDataArgs>());
            set => _requiredDatas = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ECXL2ConnectionSecondaryConnectionActionArgs()
        {
        }
    }
}
