// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Resource `equinix.NetworkDeviceLink` allows creation and management of Equinix
 * Network Edge virtual network device links.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@cuemby/equinix";
 *
 * // Example of device link with HA device pair
 * // where each device is in different metro
 * const test = new equinix.NetworkDeviceLink("test", {
 *     subnet: "192.168.40.64/27",
 *     devices: [
 *         {
 *             id: equinix_network_device.test.uuid,
 *             asn: equinix_network_device.test.asn > 0 ? equinix_network_device.test.asn : 22111,
 *             interfaceId: 6,
 *         },
 *         {
 *             id: equinix_network_device.test.secondary_device[0].uuid,
 *             asn: equinix_network_device.test.secondary_device[0].asn > 0 ? equinix_network_device.test.secondary_device[0].asn : 22333,
 *             interfaceId: 7,
 *         },
 *     ],
 *     links: [{
 *         accountNumber: equinix_network_device.test.account_number,
 *         srcMetroCode: equinix_network_device.test.metro_code,
 *         dstMetroCode: equinix_network_device.test.secondary_device[0].metro_code,
 *         throughput: "50",
 *         throughputUnit: "Mbps",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using an existing ID
 *
 * ```sh
 *  $ pulumi import equinix:index/networkDeviceLink:NetworkDeviceLink example {existing_id}
 * ```
 */
export class NetworkDeviceLink extends pulumi.CustomResource {
    /**
     * Get an existing NetworkDeviceLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkDeviceLinkState, opts?: pulumi.CustomResourceOptions): NetworkDeviceLink {
        return new NetworkDeviceLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix:index/networkDeviceLink:NetworkDeviceLink';

    /**
     * Returns true if the given object is an instance of NetworkDeviceLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkDeviceLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkDeviceLink.__pulumiType;
    }

    /**
     * definition of one or more devices belonging to the
     * device link. See Device section below for more details.
     */
    public readonly devices!: pulumi.Output<outputs.NetworkDeviceLinkDevice[]>;
    /**
     * definition of one or more, inter metro, connections belonging
     * to the device link. See Link section below for more details.
     */
    public readonly links!: pulumi.Output<outputs.NetworkDeviceLinkLink[] | undefined>;
    /**
     * device link name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * device link provisioning status on a given device. One of `PROVISIONING`,
     * `PROVISIONED`, `DEPROVISIONING`, `DEPROVISIONED`, `FAILED`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * device link subnet in CIDR format. Not required for link
     * between self configured devices.
     */
    public readonly subnet!: pulumi.Output<string | undefined>;
    /**
     * Device link unique identifier.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a NetworkDeviceLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkDeviceLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkDeviceLinkArgs | NetworkDeviceLinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkDeviceLinkState | undefined;
            resourceInputs["devices"] = state ? state.devices : undefined;
            resourceInputs["links"] = state ? state.links : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnet"] = state ? state.subnet : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as NetworkDeviceLinkArgs | undefined;
            if ((!args || args.devices === undefined) && !opts.urn) {
                throw new Error("Missing required property 'devices'");
            }
            resourceInputs["devices"] = args ? args.devices : undefined;
            resourceInputs["links"] = args ? args.links : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["subnet"] = args ? args.subnet : undefined;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkDeviceLink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkDeviceLink resources.
 */
export interface NetworkDeviceLinkState {
    /**
     * definition of one or more devices belonging to the
     * device link. See Device section below for more details.
     */
    devices?: pulumi.Input<pulumi.Input<inputs.NetworkDeviceLinkDevice>[]>;
    /**
     * definition of one or more, inter metro, connections belonging
     * to the device link. See Link section below for more details.
     */
    links?: pulumi.Input<pulumi.Input<inputs.NetworkDeviceLinkLink>[]>;
    /**
     * device link name.
     */
    name?: pulumi.Input<string>;
    /**
     * device link provisioning status on a given device. One of `PROVISIONING`,
     * `PROVISIONED`, `DEPROVISIONING`, `DEPROVISIONED`, `FAILED`.
     */
    status?: pulumi.Input<string>;
    /**
     * device link subnet in CIDR format. Not required for link
     * between self configured devices.
     */
    subnet?: pulumi.Input<string>;
    /**
     * Device link unique identifier.
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkDeviceLink resource.
 */
export interface NetworkDeviceLinkArgs {
    /**
     * definition of one or more devices belonging to the
     * device link. See Device section below for more details.
     */
    devices: pulumi.Input<pulumi.Input<inputs.NetworkDeviceLinkDevice>[]>;
    /**
     * definition of one or more, inter metro, connections belonging
     * to the device link. See Link section below for more details.
     */
    links?: pulumi.Input<pulumi.Input<inputs.NetworkDeviceLinkLink>[]>;
    /**
     * device link name.
     */
    name?: pulumi.Input<string>;
    /**
     * device link subnet in CIDR format. Not required for link
     * between self configured devices.
     */
    subnet?: pulumi.Input<string>;
}
