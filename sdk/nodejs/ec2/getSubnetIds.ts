// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * `aws_subnet_ids` provides a list of ids for a vpc_id
 * 
 * This resource can be useful for getting back a list of subnet ids for a vpc.
 */
export function getSubnetIds(args: GetSubnetIdsArgs): Promise<GetSubnetIdsResult> {
    return pulumi.runtime.invoke("aws:ec2/getSubnetIds:getSubnetIds", {
        "filters": args.filters,
        "tags": args.tags,
        "vpcId": args.vpcId,
    });
}

/**
 * A collection of arguments for invoking getSubnetIds.
 */
export interface GetSubnetIdsArgs {
    readonly filters?: { name: string, values: string[] }[];
    /**
     * A mapping of tags, each pair of which must exactly match
     * a pair on the desired subnets.
     */
    readonly tags?: {[key: string]: any};
    /**
     * The VPC ID that you want to filter from.
     */
    readonly vpcId: string;
}

/**
 * A collection of values returned by getSubnetIds.
 */
export interface GetSubnetIdsResult {
    /**
     * A list of all the subnet ids found. This data source will fail if none are found.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
