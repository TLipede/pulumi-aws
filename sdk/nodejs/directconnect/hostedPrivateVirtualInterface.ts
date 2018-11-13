// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect hosted private virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
 * A hosted virtual interface is a virtual interface that is owned by another AWS account.
 */
export class HostedPrivateVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing HostedPrivateVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HostedPrivateVirtualInterfaceState): HostedPrivateVirtualInterface {
        return new HostedPrivateVirtualInterface(name, <any>state, { id });
    }

    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    public readonly addressFamily: pulumi.Output<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    public readonly amazonAddress: pulumi.Output<string>;
    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    public readonly bgpAsn: pulumi.Output<number>;
    /**
     * The authentication key for BGP configuration.
     */
    public readonly bgpAuthKey: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    public readonly connectionId: pulumi.Output<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    public readonly customerAddress: pulumi.Output<string>;
    /**
     * Indicates whether jumbo frames (9001 MTU) are supported.
     */
    public /*out*/ readonly jumboFrameCapable: pulumi.Output<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
     */
    public readonly mtu: pulumi.Output<number | undefined>;
    /**
     * The name for the virtual interface.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    public readonly ownerAccountId: pulumi.Output<string>;
    /**
     * The VLAN ID.
     */
    public readonly vlan: pulumi.Output<number>;

    /**
     * Create a HostedPrivateVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostedPrivateVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HostedPrivateVirtualInterfaceArgs | HostedPrivateVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: HostedPrivateVirtualInterfaceState = argsOrState as HostedPrivateVirtualInterfaceState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            inputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["customerAddress"] = state ? state.customerAddress : undefined;
            inputs["jumboFrameCapable"] = state ? state.jumboFrameCapable : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ownerAccountId"] = state ? state.ownerAccountId : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as HostedPrivateVirtualInterfaceArgs | undefined;
            if (!args || args.addressFamily === undefined) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if (!args || args.bgpAsn === undefined) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if (!args || args.connectionId === undefined) {
                throw new Error("Missing required property 'connectionId'");
            }
            if (!args || args.ownerAccountId === undefined) {
                throw new Error("Missing required property 'ownerAccountId'");
            }
            if (!args || args.vlan === undefined) {
                throw new Error("Missing required property 'vlan'");
            }
            inputs["addressFamily"] = args ? args.addressFamily : undefined;
            inputs["amazonAddress"] = args ? args.amazonAddress : undefined;
            inputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            inputs["bgpAuthKey"] = args ? args.bgpAuthKey : undefined;
            inputs["connectionId"] = args ? args.connectionId : undefined;
            inputs["customerAddress"] = args ? args.customerAddress : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["ownerAccountId"] = args ? args.ownerAccountId : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["jumboFrameCapable"] = undefined /*out*/;
        }
        super("aws:directconnect/hostedPrivateVirtualInterface:HostedPrivateVirtualInterface", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HostedPrivateVirtualInterface resources.
 */
export interface HostedPrivateVirtualInterfaceState {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    /**
     * The ARN of the virtual interface.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    readonly bgpAsn?: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    readonly bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    readonly connectionId?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * Indicates whether jumbo frames (9001 MTU) are supported.
     */
    readonly jumboFrameCapable?: pulumi.Input<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    readonly ownerAccountId?: pulumi.Input<string>;
    /**
     * The VLAN ID.
     */
    readonly vlan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a HostedPrivateVirtualInterface resource.
 */
export interface HostedPrivateVirtualInterfaceArgs {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    readonly bgpAsn: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    readonly bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    readonly connectionId: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS account that will own the new virtual interface.
     */
    readonly ownerAccountId: pulumi.Input<string>;
    /**
     * The VLAN ID.
     */
    readonly vlan: pulumi.Input<number>;
}
