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

    public sealed class ECXL2ConnectionSecondaryConnectionGetArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionGetArgs>? _actions;

        /// <summary>
        /// One or more pending actions to complete connection provisioning.
        /// </summary>
        public InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionGetArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.ECXL2ConnectionSecondaryConnectionActionGetArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// Unique identifier authorizing Equinix to provision a connection
        /// towards a cloud service provider. If not specified primary `authorization_key` will be used. However,
        /// some service providers may require different keys for each connection. More information on
        /// [Equinix Fabric how to guide](https://developer.equinix.com/docs/ecx-how-to-guide).
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// Applicable with `device_uuid`, identifier of network interface
        /// on a given device. If not specified then first available interface will be selected.
        /// </summary>
        [Input("deviceInterfaceId")]
        public Input<int>? DeviceInterfaceId { get; set; }

        /// <summary>
        /// Applicable with primary `device_uuid`. Identifier of the Network Edge
        /// virtual device from which the secondary connection would originate. If not specified primary
        /// `device_uuid` will be used.
        /// </summary>
        [Input("deviceUuid")]
        public Input<string>? DeviceUuid { get; set; }

        /// <summary>
        /// secondary connection name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Applicable with primary `port_uuid`. Identifier of the Equinix Fabric Port from
        /// which the secondary connection would originate. If not specified primary `port_uuid` will be used.
        /// </summary>
        [Input("portUuid")]
        public Input<string>? PortUuid { get; set; }

        /// <summary>
        /// Unique identifier of the service provider's profile.
        /// </summary>
        [Input("profileUuid")]
        public Input<string>? ProfileUuid { get; set; }

        /// <summary>
        /// Connection provisioning status on service provider's side.
        /// </summary>
        [Input("providerStatus")]
        public Input<string>? ProviderStatus { get; set; }

        /// <summary>
        /// Unique identifier of group containing a primary and secondary connection.
        /// </summary>
        [Input("redundancyGroup")]
        public Input<string>? RedundancyGroup { get; set; }

        /// <summary>
        /// Connection redundancy type, applicable for HA connections. Valid values are
        /// `PRIMARY`, `SECONDARY`.
        /// </summary>
        [Input("redundancyType")]
        public Input<string>? RedundancyType { get; set; }

        /// <summary>
        /// Unique identifier of the redundant connection, applicable for HA connections.
        /// </summary>
        [Input("redundantUuid")]
        public Input<string>? RedundantUuid { get; set; }

        /// <summary>
        /// The metro code that denotes the secondary connection’s
        /// destination (Z side). .
        /// </summary>
        [Input("sellerMetroCode")]
        public Input<string>? SellerMetroCode { get; set; }

        /// <summary>
        /// The region in which the seller port resides. If not specified
        /// primary `seller_region` will be used.
        /// </summary>
        [Input("sellerRegion")]
        public Input<string>? SellerRegion { get; set; }

        /// <summary>
        /// Required with primary `service_token`. Unique Equinix Fabric key
        /// given by a provider that grants you authorization to enable connectivity from an Equinix Fabric Port or
        /// virtual device. Each connection (primary and secondary) requires a separate token.
        /// More details in [Fabric Service Tokens](https://docs.equinix.com/en-us/Content/Interconnection/Fabric/service%20tokens/Fabric-Service-Tokens.htm).
        /// </summary>
        [Input("serviceToken")]
        public Input<string>? ServiceToken { get; set; }

        /// <summary>
        /// Speed/Bandwidth to be allocated to the secondary connection. If not
        /// specified primary `speed` will be used.
        /// </summary>
        [Input("speed")]
        public Input<int>? Speed { get; set; }

        /// <summary>
        /// Unit of the speed/bandwidth to be allocated to the secondary
        /// connection. If not specified primary `speed_unit` will be used.
        /// </summary>
        [Input("speedUnit")]
        public Input<string>? SpeedUnit { get; set; }

        /// <summary>
        /// Connection provisioning status on Equinix Fabric side.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Unique identifier of the connection.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// The Equinix Fabric Token the connection was created with. Applicable if the
        /// connection was created with a `service_token` (a-side) or `zside_service_token` (z-side).
        /// * `secondary_connection`:
        /// * `zside_port_uuid`
        /// * `zside_vlan_stag`
        /// * `zside_vlan_ctag`
        /// * `redundancy_type`
        /// * `redundancy_group`
        /// * `vendor_token`
        /// </summary>
        [Input("vendorToken")]
        public Input<string>? VendorToken { get; set; }

        /// <summary>
        /// Applicable with `port_uuid`. C-Tag/Inner-Tag of the secondary
        /// connection, a numeric character ranging from 2 - 4094.
        /// </summary>
        [Input("vlanCtag")]
        public Input<int>? VlanCtag { get; set; }

        /// <summary>
        /// S-Tag/Outer-Tag of the secondary connection, a
        /// numeric character ranging from 2 - 4094.
        /// </summary>
        [Input("vlanStag")]
        public Input<int>? VlanStag { get; set; }

        /// <summary>
        /// Unique identifier of the port on the remote/destination side
        /// (z-side). Allows you to connect between your own ports or virtual devices across your company's
        /// Equinix Fabric deployment, with no need for a private service profile.
        /// </summary>
        [Input("zsidePortUuid")]
        public Input<string>? ZsidePortUuid { get; set; }

        /// <summary>
        /// C-Tag/Inner-Tag of the connection on the remote/destination
        /// side (z-side) - a numeric character ranging from 2 - 4094.
        /// `secondary_connection` is defined it will internally use same `zside_vlan_ctag` for the secondary
        /// connection.
        /// </summary>
        [Input("zsideVlanCtag")]
        public Input<int>? ZsideVlanCtag { get; set; }

        /// <summary>
        /// S-Tag/Outer-Tag of the connection on the remote/destination
        /// side (z-side) - a numeric character ranging from 2 - 4094.
        /// </summary>
        [Input("zsideVlanStag")]
        public Input<int>? ZsideVlanStag { get; set; }

        public ECXL2ConnectionSecondaryConnectionGetArgs()
        {
        }
    }
}
