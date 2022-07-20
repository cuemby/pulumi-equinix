// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this datasource to retrieve Metal Gateway resources in Equinix Metal.
 *
 * > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@cuemby/equinix";
 * import * as equinix from "@pulumi/equinix";
 *
 * // Create Metal Gateway for a VLAN with a private IPv4 block with 8 IP addresses
 * const testMetalVlan = new equinix.MetalVlan("testMetalVlan", {
 *     description: "test VLAN in SV",
 *     metro: "sv",
 *     projectId: local.project_id,
 * });
 * const testGetMetalGateway = equinix.GetMetalGateway({
 *     gatewayId: local.gateway_id,
 * });
 * ```
 */
export function getMetalGateway(args: GetMetalGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetMetalGatewayResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("equinix:index/getMetalGateway:GetMetalGateway", {
        "gatewayId": args.gatewayId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetMetalGateway.
 */
export interface GetMetalGatewayArgs {
    /**
     * UUID of the metal gateway resource to retrieve.
     */
    gatewayId: string;
}

/**
 * A collection of values returned by GetMetalGateway.
 */
export interface GetMetalGatewayResult {
    readonly gatewayId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * UUID of IP reservation block bound to the gateway.
     */
    readonly ipReservationId: string;
    /**
     * Size of the private IPv4 subnet bound to this metal gateway. One of
     * `8`, `16`, `32`, `64`, `128`.
     */
    readonly privateIpv4SubnetSize: number;
    /**
     * UUID of the project where the gateway is scoped to.
     */
    readonly projectId: string;
    /**
     * Status of the gateway resource.
     */
    readonly state: string;
    /**
     * UUID of the VLAN where the gateway is scoped to.
     */
    readonly vlanId: string;
    /**
     * UUID of the VRF associated with the IP Reservation.
     */
    readonly vrfId: string;
}

export function getMetalGatewayOutput(args: GetMetalGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMetalGatewayResult> {
    return pulumi.output(args).apply(a => getMetalGateway(a, opts))
}

/**
 * A collection of arguments for invoking GetMetalGateway.
 */
export interface GetMetalGatewayOutputArgs {
    /**
     * UUID of the metal gateway resource to retrieve.
     */
    gatewayId: pulumi.Input<string>;
}
