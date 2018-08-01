// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get IDs and VPC membership of Security Groups that are created
 * outside of Terraform.
 */
export function getSecurityGroups(args?: GetSecurityGroupsArgs): Promise<GetSecurityGroupsResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:ec2/getSecurityGroups:getSecurityGroups", {
        "filters": args.filters,
        "tags": args.tags,
    });
}

/**
 * A collection of arguments for invoking getSecurityGroups.
 */
export interface GetSecurityGroupsArgs {
    /**
     * One or more name/value pairs to use as filters. There are
     * several valid keys, for a full reference, check out
     * [describe-security-groups in the AWS CLI reference][1].
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * A mapping of tags, each pair of which must exactly match for
     * desired security groups.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getSecurityGroups.
 */
export interface GetSecurityGroupsResult {
    /**
     * IDs of the matches security groups.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: any};
    /**
     * The VPC IDs of the matched security groups. The data source's tag or filter *will span VPCs*
     * unless the `vpc-id` filter is also used.
     */
    readonly vpcIds: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
