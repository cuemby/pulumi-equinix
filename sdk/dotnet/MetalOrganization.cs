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
    /// <summary>
    /// Provides a resource to manage organization resource in Equinix Metal.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Equinix = Cuemby.Equinix;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a new Organization
    ///         var tfOrganization1 = new Equinix.MetalOrganization("tfOrganization1", new Equinix.MetalOrganizationArgs
    ///         {
    ///             Description = "quux",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an existing organization ID
    /// 
    /// ```sh
    ///  $ pulumi import equinix:index/metalOrganization:MetalOrganization equinix_metal_organization {existing_organization_id}
    /// ```
    /// </summary>
    [EquinixResourceType("equinix:index/metalOrganization:MetalOrganization")]
    public partial class MetalOrganization : Pulumi.CustomResource
    {
        /// <summary>
        /// Postal address.
        /// </summary>
        [Output("address")]
        public Output<Outputs.MetalOrganizationAddress> Address { get; private set; } = null!;

        /// <summary>
        /// The timestamp for when the organization was created.
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// Description string.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Logo URL.
        /// </summary>
        [Output("logo")]
        public Output<string?> Logo { get; private set; } = null!;

        /// <summary>
        /// The name of the Organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Twitter handle.
        /// </summary>
        [Output("twitter")]
        public Output<string?> Twitter { get; private set; } = null!;

        /// <summary>
        /// The timestamp for the last time the organization was updated.
        /// </summary>
        [Output("updated")]
        public Output<string> Updated { get; private set; } = null!;

        /// <summary>
        /// Website link.
        /// </summary>
        [Output("website")]
        public Output<string?> Website { get; private set; } = null!;


        /// <summary>
        /// Create a MetalOrganization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetalOrganization(string name, MetalOrganizationArgs args, CustomResourceOptions? options = null)
            : base("equinix:index/metalOrganization:MetalOrganization", name, args ?? new MetalOrganizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetalOrganization(string name, Input<string> id, MetalOrganizationState? state = null, CustomResourceOptions? options = null)
            : base("equinix:index/metalOrganization:MetalOrganization", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/cuemby/pulumi-equinix/releases/downloads/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MetalOrganization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetalOrganization Get(string name, Input<string> id, MetalOrganizationState? state = null, CustomResourceOptions? options = null)
        {
            return new MetalOrganization(name, id, state, options);
        }
    }

    public sealed class MetalOrganizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Postal address.
        /// </summary>
        [Input("address", required: true)]
        public Input<Inputs.MetalOrganizationAddressArgs> Address { get; set; } = null!;

        /// <summary>
        /// Description string.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Logo URL.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// The name of the Organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Twitter handle.
        /// </summary>
        [Input("twitter")]
        public Input<string>? Twitter { get; set; }

        /// <summary>
        /// Website link.
        /// </summary>
        [Input("website")]
        public Input<string>? Website { get; set; }

        public MetalOrganizationArgs()
        {
        }
    }

    public sealed class MetalOrganizationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Postal address.
        /// </summary>
        [Input("address")]
        public Input<Inputs.MetalOrganizationAddressGetArgs>? Address { get; set; }

        /// <summary>
        /// The timestamp for when the organization was created.
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// Description string.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Logo URL.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// The name of the Organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Twitter handle.
        /// </summary>
        [Input("twitter")]
        public Input<string>? Twitter { get; set; }

        /// <summary>
        /// The timestamp for the last time the organization was updated.
        /// </summary>
        [Input("updated")]
        public Input<string>? Updated { get; set; }

        /// <summary>
        /// Website link.
        /// </summary>
        [Input("website")]
        public Input<string>? Website { get; set; }

        public MetalOrganizationState()
        {
        }
    }
}
