// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
 * in a given region for the purpose of using in an AWS Route53 Alias.
 */
export function getHostedZoneId(args?: GetHostedZoneIdArgs): Promise<GetHostedZoneIdResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId", {
        "region": args.region,
    });
}

/**
 * A collection of arguments for invoking getHostedZoneId.
 */
export interface GetHostedZoneIdArgs {
    /**
     * Name of the region whose AWS ELB HostedZoneId is desired.
     * Defaults to the region from the AWS provider configuration.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getHostedZoneId.
 */
export interface GetHostedZoneIdResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
