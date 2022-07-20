// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Resource `equinix.ECXL2Serviceprofile` is used to manage layer 2 service profiles
 * in Equinix Fabric.
 *
 * This resource relies on the Equinix Fabric API. The parameters
 * and attributes available map to the fields described at
 * <https://developer.equinix.com/catalog/sellerv3#operation/getProfileByIdOrNameUsingGET>.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@pulumi/equinix";
 *
 * const private_profile = new equinix.ECXL2Serviceprofile("private-profile", {
 *     bandwidthThresholdNotifications: [
 *         "John.Doe@example.com",
 *         "Marry.Doe@example.com",
 *     ],
 *     connectionNameLabel: "Connection",
 *     description: "my private profile",
 *     features: {
 *         allowRemoteConnections: true,
 *         testProfile: false,
 *     },
 *     ports: [
 *         {
 *             metroCode: "NY",
 *             uuid: "a867f685-422f-22f7-6de0-320a5c00abdd",
 *         },
 *         {
 *             metroCode: "NY",
 *             uuid: "a867f685-4231-2317-6de0-320a5c00abdd",
 *         },
 *     ],
 *     private: true,
 *     privateUserEmails: [
 *         "John.Doe@example.com",
 *         "Marry.Doe@example.com",
 *     ],
 *     profileStatuschangeNotifications: [
 *         "John.Doe@example.com",
 *         "Marry.Doe@example.com",
 *     ],
 *     speedBands: [
 *         {
 *             speed: 1000,
 *             speedUnit: "MB",
 *         },
 *         {
 *             speed: 500,
 *             speedUnit: "MB",
 *         },
 *         {
 *             speed: 100,
 *             speedUnit: "MB",
 *         },
 *     ],
 *     vcStatuschangeNotifications: [
 *         "John.Doe@example.com",
 *         "Marry.Doe@example.com",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using an existing ID
 *
 * ```sh
 *  $ pulumi import equinix:index/eCXL2Serviceprofile:ECXL2Serviceprofile example {existing_id}
 * ```
 */
export class ECXL2Serviceprofile extends pulumi.CustomResource {
    /**
     * Get an existing ECXL2Serviceprofile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ECXL2ServiceprofileState, opts?: pulumi.CustomResourceOptions): ECXL2Serviceprofile {
        return new ECXL2Serviceprofile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'equinix:index/eCXL2Serviceprofile:ECXL2Serviceprofile';

    /**
     * Returns true if the given object is an instance of ECXL2Serviceprofile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ECXL2Serviceprofile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ECXL2Serviceprofile.__pulumiType;
    }

    /**
     * Boolean value that determines if API integration is enabled. It
     * allows you to complete connection provisioning in less than five minutes. Without API Integration,
     * additional manual steps will be required and the provisioning will likely take longer.
     */
    public readonly apiIntegration!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the authentication key label to be used by the
     * Authentication Key service. It allows Service Providers with QinQ ports to accept groups of
     * connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.
     */
    public readonly authkeyLabel!: pulumi.Output<string | undefined>;
    /**
     * Specifies the port bandwidth threshold percentage. If
     * the bandwidth limit is met or exceeded, an alert is sent to the seller.
     */
    public readonly bandwidthAlertThreshold!: pulumi.Output<number | undefined>;
    /**
     * A list of email addresses that will receive
     * notifications about bandwidth thresholds.
     */
    public readonly bandwidthThresholdNotifications!: pulumi.Output<string[]>;
    /**
     * Custom name used for calling a connections
     * e.g. `circuit`. Defaults to `Connection`.
     */
    public readonly connectionNameLabel!: pulumi.Output<string | undefined>;
    /**
     * C-Tag/Inner-Tag label name for the connections.
     */
    public readonly ctagLabel!: pulumi.Output<string | undefined>;
    /**
     * Description of the service profile.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Applicable when `apiIntegration` is set to `true`. It
     * indicates whether the port and VLAN details are managed by Equinix.
     */
    public readonly equinixManagedPortVlan!: pulumi.Output<boolean | undefined>;
    /**
     * Block of profile features configuration. See Features below
     * for more details.
     */
    public readonly features!: pulumi.Output<outputs.ECXL2ServiceprofileFeatures>;
    /**
     * Specifies the API integration ID that was provided to the customer
     * during onboarding. You can validate your API integration ID using the validateIntegrationId API.
     */
    public readonly integrationId!: pulumi.Output<string | undefined>;
    /**
     * Name of the service profile. An alpha-numeric 50 characters string which can
     * include only hyphens and underscores.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * You can set an alert for when a percentage of your profile has
     * been sold. Service providers like to use this functionality to alert them when they need to add
     * more ports or when they need to create a new service profile. Required with
     * `oversubscriptionAllowed`, defaults to `1x`.
     */
    public readonly oversubscription!: pulumi.Output<string | undefined>;
    /**
     * Boolean value that determines if, regardless of the
     * utilization, Equinix Fabric will continue to add connections to your links until we reach the
     * oversubscription limit. By selecting this service, you acknowledge that you will manage decisions
     * on when to increase capacity on these link.
     */
    public readonly oversubscriptionAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * One or more definitions of ports residing in locations, from which your
     * customers will be able to access services using this service profile. See Port below for
     * more details.
     */
    public readonly ports!: pulumi.Output<outputs.ECXL2ServiceprofilePort[]>;
    /**
     * Boolean value that indicates whether or not this is a private profile,
     * i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for
     * creating connections if correct permissions are granted.
     */
    public readonly private!: pulumi.Output<boolean | undefined>;
    /**
     * An array of users email ids who have permission to access this
     * service profile. Argument is required when profile is set as private.
     */
    public readonly privateUserEmails!: pulumi.Output<string[] | undefined>;
    /**
     * A list of email addresses that will receive
     * notifications about profile status changes.
     */
    public readonly profileStatuschangeNotifications!: pulumi.Output<string[]>;
    /**
     * Boolean value that determines if your connections will require
     * redundancy. if yes, then users need to create a secondary redundant connection.
     */
    public readonly redundancyRequired!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether the VLAN ID of. the secondary
     * connection is the same as the primary connection.
     */
    public readonly secondaryVlanFromPrimary!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean value that indicates whether multiple connections
     * can be created with the same authorization key to connect to this service profile after the first
     * connection has been approved by the seller.
     */
    public readonly servicekeyAutogenerated!: pulumi.Output<boolean | undefined>;
    /**
     * One or more definitions of supported speed/bandwidth. Argument is
     * required when `speedFromApi` is set to `false`. See Speed Band below for more
     * details.
     */
    public readonly speedBands!: pulumi.Output<outputs.ECXL2ServiceprofileSpeedBand[] | undefined>;
    /**
     * Boolean value that determines if customer is allowed
     * to enter a custom connection speed.
     */
    public readonly speedCustomizationAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean valuta that determines if connection speed will be derived
     * from an API call. Argument has to be specified when `apiIntegration` is enabled.
     */
    public readonly speedFromApi!: pulumi.Output<boolean | undefined>;
    /**
     * Service profile provisioning status.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Specifies additional tagging information required by the seller profile
     * for Dot1Q to QinQ translation. See [Enhance Dot1q to QinQ translation support](https://docs.equinix.com/es/Content/Interconnection/Fabric/layer-2/Fabric-Create-Layer2-Service-Profile.htm#:~:text=Enhance%20Dot1q%20to%20QinQ%20translation%20support)
     * for additional information. Valid values are:
     */
    public readonly tagType!: pulumi.Output<string | undefined>;
    /**
     * Unique identifier of the port.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;
    /**
     * A list of email addresses that will receive
     * notifications about connections approvals and rejections.
     */
    public readonly vcStatuschangeNotifications!: pulumi.Output<string[]>;

    /**
     * Create a ECXL2Serviceprofile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ECXL2ServiceprofileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ECXL2ServiceprofileArgs | ECXL2ServiceprofileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ECXL2ServiceprofileState | undefined;
            resourceInputs["apiIntegration"] = state ? state.apiIntegration : undefined;
            resourceInputs["authkeyLabel"] = state ? state.authkeyLabel : undefined;
            resourceInputs["bandwidthAlertThreshold"] = state ? state.bandwidthAlertThreshold : undefined;
            resourceInputs["bandwidthThresholdNotifications"] = state ? state.bandwidthThresholdNotifications : undefined;
            resourceInputs["connectionNameLabel"] = state ? state.connectionNameLabel : undefined;
            resourceInputs["ctagLabel"] = state ? state.ctagLabel : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["equinixManagedPortVlan"] = state ? state.equinixManagedPortVlan : undefined;
            resourceInputs["features"] = state ? state.features : undefined;
            resourceInputs["integrationId"] = state ? state.integrationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["oversubscription"] = state ? state.oversubscription : undefined;
            resourceInputs["oversubscriptionAllowed"] = state ? state.oversubscriptionAllowed : undefined;
            resourceInputs["ports"] = state ? state.ports : undefined;
            resourceInputs["private"] = state ? state.private : undefined;
            resourceInputs["privateUserEmails"] = state ? state.privateUserEmails : undefined;
            resourceInputs["profileStatuschangeNotifications"] = state ? state.profileStatuschangeNotifications : undefined;
            resourceInputs["redundancyRequired"] = state ? state.redundancyRequired : undefined;
            resourceInputs["secondaryVlanFromPrimary"] = state ? state.secondaryVlanFromPrimary : undefined;
            resourceInputs["servicekeyAutogenerated"] = state ? state.servicekeyAutogenerated : undefined;
            resourceInputs["speedBands"] = state ? state.speedBands : undefined;
            resourceInputs["speedCustomizationAllowed"] = state ? state.speedCustomizationAllowed : undefined;
            resourceInputs["speedFromApi"] = state ? state.speedFromApi : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tagType"] = state ? state.tagType : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
            resourceInputs["vcStatuschangeNotifications"] = state ? state.vcStatuschangeNotifications : undefined;
        } else {
            const args = argsOrState as ECXL2ServiceprofileArgs | undefined;
            if ((!args || args.bandwidthThresholdNotifications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidthThresholdNotifications'");
            }
            if ((!args || args.features === undefined) && !opts.urn) {
                throw new Error("Missing required property 'features'");
            }
            if ((!args || args.ports === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ports'");
            }
            if ((!args || args.profileStatuschangeNotifications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'profileStatuschangeNotifications'");
            }
            if ((!args || args.vcStatuschangeNotifications === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vcStatuschangeNotifications'");
            }
            resourceInputs["apiIntegration"] = args ? args.apiIntegration : undefined;
            resourceInputs["authkeyLabel"] = args ? args.authkeyLabel : undefined;
            resourceInputs["bandwidthAlertThreshold"] = args ? args.bandwidthAlertThreshold : undefined;
            resourceInputs["bandwidthThresholdNotifications"] = args ? args.bandwidthThresholdNotifications : undefined;
            resourceInputs["connectionNameLabel"] = args ? args.connectionNameLabel : undefined;
            resourceInputs["ctagLabel"] = args ? args.ctagLabel : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["equinixManagedPortVlan"] = args ? args.equinixManagedPortVlan : undefined;
            resourceInputs["features"] = args ? args.features : undefined;
            resourceInputs["integrationId"] = args ? args.integrationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oversubscription"] = args ? args.oversubscription : undefined;
            resourceInputs["oversubscriptionAllowed"] = args ? args.oversubscriptionAllowed : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["private"] = args ? args.private : undefined;
            resourceInputs["privateUserEmails"] = args ? args.privateUserEmails : undefined;
            resourceInputs["profileStatuschangeNotifications"] = args ? args.profileStatuschangeNotifications : undefined;
            resourceInputs["redundancyRequired"] = args ? args.redundancyRequired : undefined;
            resourceInputs["secondaryVlanFromPrimary"] = args ? args.secondaryVlanFromPrimary : undefined;
            resourceInputs["servicekeyAutogenerated"] = args ? args.servicekeyAutogenerated : undefined;
            resourceInputs["speedBands"] = args ? args.speedBands : undefined;
            resourceInputs["speedCustomizationAllowed"] = args ? args.speedCustomizationAllowed : undefined;
            resourceInputs["speedFromApi"] = args ? args.speedFromApi : undefined;
            resourceInputs["tagType"] = args ? args.tagType : undefined;
            resourceInputs["vcStatuschangeNotifications"] = args ? args.vcStatuschangeNotifications : undefined;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ECXL2Serviceprofile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ECXL2Serviceprofile resources.
 */
export interface ECXL2ServiceprofileState {
    /**
     * Boolean value that determines if API integration is enabled. It
     * allows you to complete connection provisioning in less than five minutes. Without API Integration,
     * additional manual steps will be required and the provisioning will likely take longer.
     */
    apiIntegration?: pulumi.Input<boolean>;
    /**
     * Name of the authentication key label to be used by the
     * Authentication Key service. It allows Service Providers with QinQ ports to accept groups of
     * connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.
     */
    authkeyLabel?: pulumi.Input<string>;
    /**
     * Specifies the port bandwidth threshold percentage. If
     * the bandwidth limit is met or exceeded, an alert is sent to the seller.
     */
    bandwidthAlertThreshold?: pulumi.Input<number>;
    /**
     * A list of email addresses that will receive
     * notifications about bandwidth thresholds.
     */
    bandwidthThresholdNotifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom name used for calling a connections
     * e.g. `circuit`. Defaults to `Connection`.
     */
    connectionNameLabel?: pulumi.Input<string>;
    /**
     * C-Tag/Inner-Tag label name for the connections.
     */
    ctagLabel?: pulumi.Input<string>;
    /**
     * Description of the service profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Applicable when `apiIntegration` is set to `true`. It
     * indicates whether the port and VLAN details are managed by Equinix.
     */
    equinixManagedPortVlan?: pulumi.Input<boolean>;
    /**
     * Block of profile features configuration. See Features below
     * for more details.
     */
    features?: pulumi.Input<inputs.ECXL2ServiceprofileFeatures>;
    /**
     * Specifies the API integration ID that was provided to the customer
     * during onboarding. You can validate your API integration ID using the validateIntegrationId API.
     */
    integrationId?: pulumi.Input<string>;
    /**
     * Name of the service profile. An alpha-numeric 50 characters string which can
     * include only hyphens and underscores.
     */
    name?: pulumi.Input<string>;
    /**
     * You can set an alert for when a percentage of your profile has
     * been sold. Service providers like to use this functionality to alert them when they need to add
     * more ports or when they need to create a new service profile. Required with
     * `oversubscriptionAllowed`, defaults to `1x`.
     */
    oversubscription?: pulumi.Input<string>;
    /**
     * Boolean value that determines if, regardless of the
     * utilization, Equinix Fabric will continue to add connections to your links until we reach the
     * oversubscription limit. By selecting this service, you acknowledge that you will manage decisions
     * on when to increase capacity on these link.
     */
    oversubscriptionAllowed?: pulumi.Input<boolean>;
    /**
     * One or more definitions of ports residing in locations, from which your
     * customers will be able to access services using this service profile. See Port below for
     * more details.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.ECXL2ServiceprofilePort>[]>;
    /**
     * Boolean value that indicates whether or not this is a private profile,
     * i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for
     * creating connections if correct permissions are granted.
     */
    private?: pulumi.Input<boolean>;
    /**
     * An array of users email ids who have permission to access this
     * service profile. Argument is required when profile is set as private.
     */
    privateUserEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of email addresses that will receive
     * notifications about profile status changes.
     */
    profileStatuschangeNotifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean value that determines if your connections will require
     * redundancy. if yes, then users need to create a secondary redundant connection.
     */
    redundancyRequired?: pulumi.Input<boolean>;
    /**
     * Indicates whether the VLAN ID of. the secondary
     * connection is the same as the primary connection.
     */
    secondaryVlanFromPrimary?: pulumi.Input<boolean>;
    /**
     * Boolean value that indicates whether multiple connections
     * can be created with the same authorization key to connect to this service profile after the first
     * connection has been approved by the seller.
     */
    servicekeyAutogenerated?: pulumi.Input<boolean>;
    /**
     * One or more definitions of supported speed/bandwidth. Argument is
     * required when `speedFromApi` is set to `false`. See Speed Band below for more
     * details.
     */
    speedBands?: pulumi.Input<pulumi.Input<inputs.ECXL2ServiceprofileSpeedBand>[]>;
    /**
     * Boolean value that determines if customer is allowed
     * to enter a custom connection speed.
     */
    speedCustomizationAllowed?: pulumi.Input<boolean>;
    /**
     * Boolean valuta that determines if connection speed will be derived
     * from an API call. Argument has to be specified when `apiIntegration` is enabled.
     */
    speedFromApi?: pulumi.Input<boolean>;
    /**
     * Service profile provisioning status.
     */
    state?: pulumi.Input<string>;
    /**
     * Specifies additional tagging information required by the seller profile
     * for Dot1Q to QinQ translation. See [Enhance Dot1q to QinQ translation support](https://docs.equinix.com/es/Content/Interconnection/Fabric/layer-2/Fabric-Create-Layer2-Service-Profile.htm#:~:text=Enhance%20Dot1q%20to%20QinQ%20translation%20support)
     * for additional information. Valid values are:
     */
    tagType?: pulumi.Input<string>;
    /**
     * Unique identifier of the port.
     */
    uuid?: pulumi.Input<string>;
    /**
     * A list of email addresses that will receive
     * notifications about connections approvals and rejections.
     */
    vcStatuschangeNotifications?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ECXL2Serviceprofile resource.
 */
export interface ECXL2ServiceprofileArgs {
    /**
     * Boolean value that determines if API integration is enabled. It
     * allows you to complete connection provisioning in less than five minutes. Without API Integration,
     * additional manual steps will be required and the provisioning will likely take longer.
     */
    apiIntegration?: pulumi.Input<boolean>;
    /**
     * Name of the authentication key label to be used by the
     * Authentication Key service. It allows Service Providers with QinQ ports to accept groups of
     * connections or VLANs from Dot1q customers. This is similar to S-Tag/C-Tag capabilities.
     */
    authkeyLabel?: pulumi.Input<string>;
    /**
     * Specifies the port bandwidth threshold percentage. If
     * the bandwidth limit is met or exceeded, an alert is sent to the seller.
     */
    bandwidthAlertThreshold?: pulumi.Input<number>;
    /**
     * A list of email addresses that will receive
     * notifications about bandwidth thresholds.
     */
    bandwidthThresholdNotifications: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom name used for calling a connections
     * e.g. `circuit`. Defaults to `Connection`.
     */
    connectionNameLabel?: pulumi.Input<string>;
    /**
     * C-Tag/Inner-Tag label name for the connections.
     */
    ctagLabel?: pulumi.Input<string>;
    /**
     * Description of the service profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Applicable when `apiIntegration` is set to `true`. It
     * indicates whether the port and VLAN details are managed by Equinix.
     */
    equinixManagedPortVlan?: pulumi.Input<boolean>;
    /**
     * Block of profile features configuration. See Features below
     * for more details.
     */
    features: pulumi.Input<inputs.ECXL2ServiceprofileFeatures>;
    /**
     * Specifies the API integration ID that was provided to the customer
     * during onboarding. You can validate your API integration ID using the validateIntegrationId API.
     */
    integrationId?: pulumi.Input<string>;
    /**
     * Name of the service profile. An alpha-numeric 50 characters string which can
     * include only hyphens and underscores.
     */
    name?: pulumi.Input<string>;
    /**
     * You can set an alert for when a percentage of your profile has
     * been sold. Service providers like to use this functionality to alert them when they need to add
     * more ports or when they need to create a new service profile. Required with
     * `oversubscriptionAllowed`, defaults to `1x`.
     */
    oversubscription?: pulumi.Input<string>;
    /**
     * Boolean value that determines if, regardless of the
     * utilization, Equinix Fabric will continue to add connections to your links until we reach the
     * oversubscription limit. By selecting this service, you acknowledge that you will manage decisions
     * on when to increase capacity on these link.
     */
    oversubscriptionAllowed?: pulumi.Input<boolean>;
    /**
     * One or more definitions of ports residing in locations, from which your
     * customers will be able to access services using this service profile. See Port below for
     * more details.
     */
    ports: pulumi.Input<pulumi.Input<inputs.ECXL2ServiceprofilePort>[]>;
    /**
     * Boolean value that indicates whether or not this is a private profile,
     * i.e. not public like AWS/Azure/Oracle/Google, etc. If private, it can only be available for
     * creating connections if correct permissions are granted.
     */
    private?: pulumi.Input<boolean>;
    /**
     * An array of users email ids who have permission to access this
     * service profile. Argument is required when profile is set as private.
     */
    privateUserEmails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of email addresses that will receive
     * notifications about profile status changes.
     */
    profileStatuschangeNotifications: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean value that determines if your connections will require
     * redundancy. if yes, then users need to create a secondary redundant connection.
     */
    redundancyRequired?: pulumi.Input<boolean>;
    /**
     * Indicates whether the VLAN ID of. the secondary
     * connection is the same as the primary connection.
     */
    secondaryVlanFromPrimary?: pulumi.Input<boolean>;
    /**
     * Boolean value that indicates whether multiple connections
     * can be created with the same authorization key to connect to this service profile after the first
     * connection has been approved by the seller.
     */
    servicekeyAutogenerated?: pulumi.Input<boolean>;
    /**
     * One or more definitions of supported speed/bandwidth. Argument is
     * required when `speedFromApi` is set to `false`. See Speed Band below for more
     * details.
     */
    speedBands?: pulumi.Input<pulumi.Input<inputs.ECXL2ServiceprofileSpeedBand>[]>;
    /**
     * Boolean value that determines if customer is allowed
     * to enter a custom connection speed.
     */
    speedCustomizationAllowed?: pulumi.Input<boolean>;
    /**
     * Boolean valuta that determines if connection speed will be derived
     * from an API call. Argument has to be specified when `apiIntegration` is enabled.
     */
    speedFromApi?: pulumi.Input<boolean>;
    /**
     * Specifies additional tagging information required by the seller profile
     * for Dot1Q to QinQ translation. See [Enhance Dot1q to QinQ translation support](https://docs.equinix.com/es/Content/Interconnection/Fabric/layer-2/Fabric-Create-Layer2-Service-Profile.htm#:~:text=Enhance%20Dot1q%20to%20QinQ%20translation%20support)
     * for additional information. Valid values are:
     */
    tagType?: pulumi.Input<string>;
    /**
     * A list of email addresses that will receive
     * notifications about connections approvals and rejections.
     */
    vcStatuschangeNotifications: pulumi.Input<pulumi.Input<string>[]>;
}
