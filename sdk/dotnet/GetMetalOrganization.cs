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
    public static class GetMetalOrganization
    {
        /// <summary>
        /// Provides an Equinix Metal organization datasource.
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
        ///         var test = Output.Create(Equinix.GetMetalOrganization.InvokeAsync(new Equinix.GetMetalOrganizationArgs
        ///         {
        ///             OrganizationId = local.Org_id,
        ///         }));
        ///         this.ProjectsInTheOrg = test.Apply(test =&gt; test.ProjectIds);
        ///     }
        /// 
        ///     [Output("projectsInTheOrg")]
        ///     public Output&lt;string&gt; ProjectsInTheOrg { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMetalOrganizationResult> InvokeAsync(GetMetalOrganizationArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMetalOrganizationResult>("equinix:index/getMetalOrganization:GetMetalOrganization", args ?? new GetMetalOrganizationArgs(), options.WithDefaults());

        /// <summary>
        /// Provides an Equinix Metal organization datasource.
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
        ///         var test = Output.Create(Equinix.GetMetalOrganization.InvokeAsync(new Equinix.GetMetalOrganizationArgs
        ///         {
        ///             OrganizationId = local.Org_id,
        ///         }));
        ///         this.ProjectsInTheOrg = test.Apply(test =&gt; test.ProjectIds);
        ///     }
        /// 
        ///     [Output("projectsInTheOrg")]
        ///     public Output&lt;string&gt; ProjectsInTheOrg { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMetalOrganizationResult> Invoke(GetMetalOrganizationInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMetalOrganizationResult>("equinix:index/getMetalOrganization:GetMetalOrganization", args ?? new GetMetalOrganizationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetalOrganizationArgs : Pulumi.InvokeArgs
    {
        [Input("addresses")]
        private List<Inputs.GetMetalOrganizationAddressArgs>? _addresses;

        /// <summary>
        /// Postal address.
        /// </summary>
        public List<Inputs.GetMetalOrganizationAddressArgs> Addresses
        {
            get => _addresses ?? (_addresses = new List<Inputs.GetMetalOrganizationAddressArgs>());
            set => _addresses = value;
        }

        /// <summary>
        /// The organization name.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The UUID of the organization resource.
        /// </summary>
        [Input("organizationId")]
        public string? OrganizationId { get; set; }

        public GetMetalOrganizationArgs()
        {
        }
    }

    public sealed class GetMetalOrganizationInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("addresses")]
        private InputList<Inputs.GetMetalOrganizationAddressInputArgs>? _addresses;

        /// <summary>
        /// Postal address.
        /// </summary>
        public InputList<Inputs.GetMetalOrganizationAddressInputArgs> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<Inputs.GetMetalOrganizationAddressInputArgs>());
            set => _addresses = value;
        }

        /// <summary>
        /// The organization name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The UUID of the organization resource.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        public GetMetalOrganizationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMetalOrganizationResult
    {
        /// <summary>
        /// Postal address.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMetalOrganizationAddressResult> Addresses;
        /// <summary>
        /// Description string.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Logo URL.
        /// </summary>
        public readonly string Logo;
        public readonly string Name;
        public readonly string OrganizationId;
        /// <summary>
        /// UUIDs of project resources which belong to this organization.
        /// </summary>
        public readonly ImmutableArray<string> ProjectIds;
        /// <summary>
        /// Twitter handle.
        /// </summary>
        public readonly string Twitter;
        /// <summary>
        /// Website link.
        /// </summary>
        public readonly string Website;

        [OutputConstructor]
        private GetMetalOrganizationResult(
            ImmutableArray<Outputs.GetMetalOrganizationAddressResult> addresses,

            string description,

            string id,

            string logo,

            string name,

            string organizationId,

            ImmutableArray<string> projectIds,

            string twitter,

            string website)
        {
            Addresses = addresses;
            Description = description;
            Id = id;
            Logo = logo;
            Name = name;
            OrganizationId = organizationId;
            ProjectIds = projectIds;
            Twitter = twitter;
            Website = website;
        }
    }
}
