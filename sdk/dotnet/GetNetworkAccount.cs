// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Cuemby.Equinix
{
    public static class GetNetworkAccount
    {
        /// <summary>
        /// Use this data source to get number and identifier of Equinix Network Edge
        /// billing account in a given metro location.
        /// 
        /// Billing account reference is required to create Network Edge virtual device
        /// in corresponding metro location.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var dc = Output.Create(Equinix.GetNetworkAccount.InvokeAsync(new Equinix.GetNetworkAccountArgs
        ///         {
        ///             MetroCode = "DC",
        ///             Status = "Active",
        ///         }));
        ///         this.Number = dc.Apply(dc =&gt; dc.Number);
        ///     }
        /// 
        ///     [Output("number")]
        ///     public Output&lt;string&gt; Number { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNetworkAccountResult> InvokeAsync(GetNetworkAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkAccountResult>("equinix:index/getNetworkAccount:GetNetworkAccount", args ?? new GetNetworkAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get number and identifier of Equinix Network Edge
        /// billing account in a given metro location.
        /// 
        /// Billing account reference is required to create Network Edge virtual device
        /// in corresponding metro location.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var dc = Output.Create(Equinix.GetNetworkAccount.InvokeAsync(new Equinix.GetNetworkAccountArgs
        ///         {
        ///             MetroCode = "DC",
        ///             Status = "Active",
        ///         }));
        ///         this.Number = dc.Apply(dc =&gt; dc.Number);
        ///     }
        /// 
        ///     [Output("number")]
        ///     public Output&lt;string&gt; Number { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNetworkAccountResult> Invoke(GetNetworkAccountInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkAccountResult>("equinix:index/getNetworkAccount:GetNetworkAccount", args ?? new GetNetworkAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Account location metro code.
        /// </summary>
        [Input("metroCode", required: true)]
        public string MetroCode { get; set; } = null!;

        /// <summary>
        /// Account name for filtering.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Account status for filtering. Possible values are: `Active`, `Processing`,
        /// `Submitted`, `Staged`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetNetworkAccountArgs()
        {
        }
    }

    public sealed class GetNetworkAccountInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Account location metro code.
        /// </summary>
        [Input("metroCode", required: true)]
        public Input<string> MetroCode { get; set; } = null!;

        /// <summary>
        /// Account name for filtering.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Account status for filtering. Possible values are: `Active`, `Processing`,
        /// `Submitted`, `Staged`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetNetworkAccountInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkAccountResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MetroCode;
        public readonly string Name;
        /// <summary>
        /// Account unique number.
        /// </summary>
        public readonly string Number;
        public readonly string Status;
        /// <summary>
        /// Account unique identifier.
        /// </summary>
        public readonly string UcmId;

        [OutputConstructor]
        private GetNetworkAccountResult(
            string id,

            string metroCode,

            string name,

            string number,

            string status,

            string ucmId)
        {
            Id = id;
            MetroCode = metroCode;
            Name = name;
            Number = number;
            Status = status;
            UcmId = ucmId;
        }
    }
}
