// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * The ECS task definition data source allows access to details of
 * a specific AWS ECS task definition.
 * 
 */
export function getTaskDefinition(args: GetTaskDefinitionArgs): Promise<GetTaskDefinitionResult> {
    return pulumi.runtime.invoke("aws:ecs/getTaskDefinition:getTaskDefinition", {
        "taskDefinition": args.taskDefinition,
    });
}

/**
 * A collection of arguments for invoking getTaskDefinition.
 */
export interface GetTaskDefinitionArgs {
    /**
     * The family for the latest ACTIVE revision, family and revision (family:revision) for a specific revision in the family, the ARN of the task definition to access to.
     */
    readonly taskDefinition: string;
}

/**
 * A collection of values returned by getTaskDefinition.
 */
export interface GetTaskDefinitionResult {
    /**
     * The family of this task definition
     */
    readonly family: string;
    /**
     * The Docker networking mode to use for the containers in this task.
     */
    readonly networkMode: string;
    /**
     * The revision of this task definition
     */
    readonly revision: number;
    /**
     * The status of this task definition
     */
    readonly status: string;
    /**
     * The ARN of the IAM role that containers in this task can assume
     */
    readonly taskRoleArn: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
