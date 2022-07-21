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
    /// Use this resource to request the creation an Interconnection asset to connect with other parties using [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/).
    /// 
    /// &gt; Equinix Metal connection with service_token_type `a_side` is not generally available and may not be enabled yet for your organization.
    /// 
    /// ## Example Usage
    /// ### Shared Connection with z_side token - Non-redundant Connection from your own Equinix Fabric Port to Equinix Metal
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Equinix = Cuemby.Equinix;
    /// using Equinix = Pulumi.Equinix;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleMetalConnection = new Equinix.MetalConnection("exampleMetalConnection", new Equinix.MetalConnectionArgs
    ///         {
    ///             ProjectId = local.Project_id,
    ///             Type = "shared",
    ///             Redundancy = "primary",
    ///             Metro = "FR",
    ///             Speed = "200Mbps",
    ///             ServiceTokenType = "z_side",
    ///         });
    ///         var exampleGetECXPort = Output.Create(Equinix.GetECXPort.InvokeAsync(new Equinix.GetECXPortArgs
    ///         {
    ///             Name = "CX-FR5-NL-Dot1q-BO-1G-PRI",
    ///         }));
    ///         var exampleECXL2Connection = new Equinix.ECXL2Connection("exampleECXL2Connection", new Equinix.ECXL2ConnectionArgs
    ///         {
    ///             ZsideServiceToken = exampleMetalConnection.ServiceTokens.Apply(serviceTokens =&gt; serviceTokens[0].Id),
    ///             Speed = 200,
    ///             SpeedUnit = "MB",
    ///             Notifications = 
    ///             {
    ///                 "example@equinix.com",
    ///             },
    ///             SellerMetroCode = "FR",
    ///             PortUuid = exampleGetECXPort.Apply(exampleGetECXPort =&gt; exampleGetECXPort.Id),
    ///             VlanStag = 1020,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [EquinixResourceType("equinix:index/metalConnection:MetalConnection")]
    public partial class MetalConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Description for the connection resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Facility where the connection will be created.
        /// </summary>
        [Output("facility")]
        public Output<string> Facility { get; private set; } = null!;

        /// <summary>
        /// Metro where the connection will be created.
        /// </summary>
        [Output("metro")]
        public Output<string> Metro { get; private set; } = null!;

        /// <summary>
        /// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// Name of the connection resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization where the connection is scoped to.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of
        /// port is described in documentation of the
        /// equinix.MetalConnection datasource.
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<Outputs.MetalConnectionPort>> Ports { get; private set; } = null!;

        /// <summary>
        /// ID of the project where the connection is scoped to, must be set for.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Connection redundancy - redundant or primary.
        /// </summary>
        [Output("redundancy")]
        public Output<string> Redundancy { get; private set; } = null!;

        /// <summary>
        /// Only used with shared connection. Type of service token to use for the connection, a_side or z_side.
        /// </summary>
        [Output("serviceTokenType")]
        public Output<string?> ServiceTokenType { get; private set; } = null!;

        /// <summary>
        /// List of connection service tokens with attributes. Scehma of service_token is described in documentation of the equinix.MetalConnection datasource.
        /// </summary>
        [Output("serviceTokens")]
        public Output<ImmutableArray<Outputs.MetalConnectionServiceToken>> ServiceTokens { get; private set; } = null!;

        /// <summary>
        /// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
        /// </summary>
        [Output("speed")]
        public Output<string> Speed { get; private set; } = null!;

        /// <summary>
        /// Status of the connection resource.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// String list of tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// Connection type - dedicated or shared.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
        /// </summary>
        [Output("vlans")]
        public Output<ImmutableArray<int>> Vlans { get; private set; } = null!;


        /// <summary>
        /// Create a MetalConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetalConnection(string name, MetalConnectionArgs args, CustomResourceOptions? options = null)
            : base("equinix:index/metalConnection:MetalConnection", name, args ?? new MetalConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetalConnection(string name, Input<string> id, MetalConnectionState? state = null, CustomResourceOptions? options = null)
            : base("equinix:index/metalConnection:MetalConnection", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/cuemby",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MetalConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetalConnection Get(string name, Input<string> id, MetalConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new MetalConnection(name, id, state, options);
        }
    }

    public sealed class MetalConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the connection resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Facility where the connection will be created.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Metro where the connection will be created.
        /// </summary>
        [Input("metro")]
        public Input<string>? Metro { get; set; }

        /// <summary>
        /// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of the connection resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization where the connection is scoped to.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// ID of the project where the connection is scoped to, must be set for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Connection redundancy - redundant or primary.
        /// </summary>
        [Input("redundancy", required: true)]
        public Input<string> Redundancy { get; set; } = null!;

        /// <summary>
        /// Only used with shared connection. Type of service token to use for the connection, a_side or z_side.
        /// </summary>
        [Input("serviceTokenType")]
        public Input<string>? ServiceTokenType { get; set; }

        /// <summary>
        /// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
        /// </summary>
        [Input("speed", required: true)]
        public Input<string> Speed { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// String list of tags.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Connection type - dedicated or shared.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("vlans")]
        private InputList<int>? _vlans;

        /// <summary>
        /// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
        /// </summary>
        public InputList<int> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<int>());
            set => _vlans = value;
        }

        public MetalConnectionArgs()
        {
        }
    }

    public sealed class MetalConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the connection resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Facility where the connection will be created.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Metro where the connection will be created.
        /// </summary>
        [Input("metro")]
        public Input<string>? Metro { get; set; }

        /// <summary>
        /// Mode for connections in IBX facilities with the dedicated type - standard or tunnel. Default is standard.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Name of the connection resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization where the connection is scoped to.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("ports")]
        private InputList<Inputs.MetalConnectionPortGetArgs>? _ports;

        /// <summary>
        /// List of connection ports - primary (`ports[0]`) and secondary (`ports[1]`). Schema of
        /// port is described in documentation of the
        /// equinix.MetalConnection datasource.
        /// </summary>
        public InputList<Inputs.MetalConnectionPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.MetalConnectionPortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// ID of the project where the connection is scoped to, must be set for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Connection redundancy - redundant or primary.
        /// </summary>
        [Input("redundancy")]
        public Input<string>? Redundancy { get; set; }

        /// <summary>
        /// Only used with shared connection. Type of service token to use for the connection, a_side or z_side.
        /// </summary>
        [Input("serviceTokenType")]
        public Input<string>? ServiceTokenType { get; set; }

        [Input("serviceTokens")]
        private InputList<Inputs.MetalConnectionServiceTokenGetArgs>? _serviceTokens;

        /// <summary>
        /// List of connection service tokens with attributes. Scehma of service_token is described in documentation of the equinix.MetalConnection datasource.
        /// </summary>
        public InputList<Inputs.MetalConnectionServiceTokenGetArgs> ServiceTokens
        {
            get => _serviceTokens ?? (_serviceTokens = new InputList<Inputs.MetalConnectionServiceTokenGetArgs>());
            set => _serviceTokens = value;
        }

        /// <summary>
        /// Connection speed - one of 50Mbps, 200Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, 10Gbps.
        /// </summary>
        [Input("speed")]
        public Input<string>? Speed { get; set; }

        /// <summary>
        /// Status of the connection resource.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// String list of tags.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Fabric Token from the [Equinix Fabric Portal](https://ecxfabric.equinix.com/dashboard)
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Connection type - dedicated or shared.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("vlans")]
        private InputList<int>? _vlans;

        /// <summary>
        /// Only used with shared connection. Vlans to attach. Pass one vlan for Primary/Single connection and two vlans for Redundant connection.
        /// </summary>
        public InputList<int> Vlans
        {
            get => _vlans ?? (_vlans = new InputList<int>());
            set => _vlans = value;
        }

        public MetalConnectionState()
        {
        }
    }
}
