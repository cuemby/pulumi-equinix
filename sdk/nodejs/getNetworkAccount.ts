// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get number and identifier of Equinix Network Edge
 * billing account in a given metro location.
 *
 * Billing account reference is required to create Network Edge virtual device
 * in corresponding metro location.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as equinix from "@pulumi/equinix";
 *
 * const dc = equinix.GetNetworkAccount({
 *     metroCode: "DC",
 *     status: "Active",
 * });
 * export const number = dc.then(dc => dc.number);
 * ```
 */
export function getNetworkAccount(args: GetNetworkAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkAccountResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("equinix:index/getNetworkAccount:GetNetworkAccount", {
        "metroCode": args.metroCode,
        "name": args.name,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNetworkAccount.
 */
export interface GetNetworkAccountArgs {
    /**
     * Account location metro code.
     */
    metroCode: string;
    /**
     * Account name for filtering.
     */
    name?: string;
    /**
     * Account status for filtering. Possible values are: `Active`, `Processing`,
     * `Submitted`, `Staged`.
     */
    status?: string;
}

/**
 * A collection of values returned by GetNetworkAccount.
 */
export interface GetNetworkAccountResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly metroCode: string;
    readonly name: string;
    /**
     * Account unique number.
     */
    readonly number: string;
    readonly status: string;
    /**
     * Account unique identifier.
     */
    readonly ucmId: string;
}

export function getNetworkAccountOutput(args: GetNetworkAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkAccountResult> {
    return pulumi.output(args).apply(a => getNetworkAccount(a, opts))
}

/**
 * A collection of arguments for invoking GetNetworkAccount.
 */
export interface GetNetworkAccountOutputArgs {
    /**
     * Account location metro code.
     */
    metroCode: pulumi.Input<string>;
    /**
     * Account name for filtering.
     */
    name?: pulumi.Input<string>;
    /**
     * Account status for filtering. Possible values are: `Active`, `Processing`,
     * `Submitted`, `Staged`.
     */
    status?: pulumi.Input<string>;
}
