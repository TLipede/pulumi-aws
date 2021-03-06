// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an association between a WAF Regional WebACL and Application Load Balancer.
 * 
 * -> **Note:** An Application Load Balancer can only be associated with one WAF Regional WebACL.
 */
export class WebAclAssociation extends pulumi.CustomResource {
    /**
     * Get an existing WebAclAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAclAssociationState): WebAclAssociation {
        return new WebAclAssociation(name, <any>state, { id });
    }

    /**
     * Application Load Balancer ARN to associate with.
     */
    public readonly resourceArn: pulumi.Output<string>;
    /**
     * The ID of the WAF Regional WebACL to create an association.
     */
    public readonly webAclId: pulumi.Output<string>;

    /**
     * Create a WebAclAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAclAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAclAssociationArgs | WebAclAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: WebAclAssociationState = argsOrState as WebAclAssociationState | undefined;
            inputs["resourceArn"] = state ? state.resourceArn : undefined;
            inputs["webAclId"] = state ? state.webAclId : undefined;
        } else {
            const args = argsOrState as WebAclAssociationArgs | undefined;
            if (!args || args.resourceArn === undefined) {
                throw new Error("Missing required property 'resourceArn'");
            }
            if (!args || args.webAclId === undefined) {
                throw new Error("Missing required property 'webAclId'");
            }
            inputs["resourceArn"] = args ? args.resourceArn : undefined;
            inputs["webAclId"] = args ? args.webAclId : undefined;
        }
        super("aws:wafregional/webAclAssociation:WebAclAssociation", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebAclAssociation resources.
 */
export interface WebAclAssociationState {
    /**
     * Application Load Balancer ARN to associate with.
     */
    readonly resourceArn?: pulumi.Input<string>;
    /**
     * The ID of the WAF Regional WebACL to create an association.
     */
    readonly webAclId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebAclAssociation resource.
 */
export interface WebAclAssociationArgs {
    /**
     * Application Load Balancer ARN to associate with.
     */
    readonly resourceArn: pulumi.Input<string>;
    /**
     * The ID of the WAF Regional WebACL to create an association.
     */
    readonly webAclId: pulumi.Input<string>;
}
