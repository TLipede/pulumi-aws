// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
 * By using this data source, you can reference SQS queues without having to hardcode
 * the ARNs as input.
 */
export function getQueue(args: GetQueueArgs): Promise<GetQueueResult> {
    return pulumi.runtime.invoke("aws:sqs/getQueue:getQueue", {
        "name": args.name,
    });
}

/**
 * A collection of arguments for invoking getQueue.
 */
export interface GetQueueArgs {
    /**
     * The name of the queue to match.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getQueue.
 */
export interface GetQueueResult {
    /**
     * The Amazon Resource Name (ARN) of the queue.
     */
    readonly arn: string;
    /**
     * The URL of the queue.
     */
    readonly url: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
