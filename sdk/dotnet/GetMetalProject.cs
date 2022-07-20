// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Equinix
{
    public static class GetMetalProject
    {
        /// <summary>
        /// Use this datasource to retrieve attributes of the Project API resource.
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
        ///         var tfProject1 = Output.Create(Equinix.GetMetalProject.InvokeAsync(new Equinix.GetMetalProjectArgs
        ///         {
        ///             Name = "Terraform Fun",
        ///         }));
        ///         this.UsersOfTerraformFun = tfProject1.Apply(tfProject1 =&gt; tfProject1.UserIds);
        ///     }
        /// 
        ///     [Output("usersOfTerraformFun")]
        ///     public Output&lt;string&gt; UsersOfTerraformFun { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMetalProjectResult> InvokeAsync(GetMetalProjectArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMetalProjectResult>("equinix:index/getMetalProject:GetMetalProject", args ?? new GetMetalProjectArgs(), options.WithDefaults());

        /// <summary>
        /// Use this datasource to retrieve attributes of the Project API resource.
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
        ///         var tfProject1 = Output.Create(Equinix.GetMetalProject.InvokeAsync(new Equinix.GetMetalProjectArgs
        ///         {
        ///             Name = "Terraform Fun",
        ///         }));
        ///         this.UsersOfTerraformFun = tfProject1.Apply(tfProject1 =&gt; tfProject1.UserIds);
        ///     }
        /// 
        ///     [Output("usersOfTerraformFun")]
        ///     public Output&lt;string&gt; UsersOfTerraformFun { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMetalProjectResult> Invoke(GetMetalProjectInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMetalProjectResult>("equinix:index/getMetalProject:GetMetalProject", args ?? new GetMetalProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetalProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name which is used to look up the project.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The UUID by which to look up the project.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        public GetMetalProjectArgs()
        {
        }
    }

    public sealed class GetMetalProjectInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name which is used to look up the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The UUID by which to look up the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public GetMetalProjectInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMetalProjectResult
    {
        /// <summary>
        /// Whether Backend Transfer is enabled for this project.
        /// </summary>
        public readonly bool BackendTransfer;
        /// <summary>
        /// Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMetalProjectBgpConfigResult> BgpConfigs;
        /// <summary>
        /// The timestamp for when the project was created.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The UUID of this project's parent organization.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The UUID of payment method for this project.
        /// </summary>
        public readonly string PaymentMethodId;
        public readonly string ProjectId;
        /// <summary>
        /// The timestamp for the last time the project was updated.
        /// </summary>
        public readonly string Updated;
        /// <summary>
        /// List of UUIDs of user accounts which belong to this project.
        /// </summary>
        public readonly ImmutableArray<string> UserIds;

        [OutputConstructor]
        private GetMetalProjectResult(
            bool backendTransfer,

            ImmutableArray<Outputs.GetMetalProjectBgpConfigResult> bgpConfigs,

            string created,

            string id,

            string name,

            string organizationId,

            string paymentMethodId,

            string projectId,

            string updated,

            ImmutableArray<string> userIds)
        {
            BackendTransfer = backendTransfer;
            BgpConfigs = bgpConfigs;
            Created = created;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            PaymentMethodId = paymentMethodId;
            ProjectId = projectId;
            Updated = updated;
            UserIds = userIds;
        }
    }
}
