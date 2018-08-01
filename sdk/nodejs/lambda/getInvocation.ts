// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Use this data source to invoke custom lambda functions as data source.
 * The lambda function is invoked with [RequestResponse](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax)
 * invocation type.
 */
export function getInvocation(args: GetInvocationArgs): Promise<GetInvocationResult> {
    return pulumi.runtime.invoke("aws:lambda/getInvocation:getInvocation", {
        "functionName": args.functionName,
        "input": args.input,
        "qualifier": args.qualifier,
    });
}

/**
 * A collection of arguments for invoking getInvocation.
 */
export interface GetInvocationArgs {
    /**
     * The name of the lambda function.
     */
    readonly functionName: string;
    /**
     * A string in JSON format that is passed as payload to the lambda function.
     */
    readonly input: string;
    /**
     * The qualifier (a.k.a version) of the lambda function. Defaults
     * to `$LATEST`.
     */
    readonly qualifier?: string;
}

/**
 * A collection of values returned by getInvocation.
 */
export interface GetInvocationResult {
    /**
     * A result of the lambda function invocation.
     */
    readonly result: string;
    /**
     * This field is set only if result is a map of primitive types.
     */
    readonly resultMap: {[key: string]: string};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
