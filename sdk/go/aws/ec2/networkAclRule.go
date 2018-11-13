// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an entry (a rule) in a network ACL with the specified rule number.
// 
// ~> **NOTE on Network ACLs and Network ACL Rules:** Terraform currently
// provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
// defined in-line. At this time you cannot use a Network ACL with in-line rules
// in conjunction with any Network ACL Rule resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
type NetworkAclRule struct {
	s *pulumi.ResourceState
}

// NewNetworkAclRule registers a new resource with the given unique name, arguments, and options.
func NewNetworkAclRule(ctx *pulumi.Context,
	name string, args *NetworkAclRuleArgs, opts ...pulumi.ResourceOpt) (*NetworkAclRule, error) {
	if args == nil || args.NetworkAclId == nil {
		return nil, errors.New("missing required argument 'NetworkAclId'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.RuleAction == nil {
		return nil, errors.New("missing required argument 'RuleAction'")
	}
	if args == nil || args.RuleNumber == nil {
		return nil, errors.New("missing required argument 'RuleNumber'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cidrBlock"] = nil
		inputs["egress"] = nil
		inputs["fromPort"] = nil
		inputs["icmpCode"] = nil
		inputs["icmpType"] = nil
		inputs["ipv6CidrBlock"] = nil
		inputs["networkAclId"] = nil
		inputs["protocol"] = nil
		inputs["ruleAction"] = nil
		inputs["ruleNumber"] = nil
		inputs["toPort"] = nil
	} else {
		inputs["cidrBlock"] = args.CidrBlock
		inputs["egress"] = args.Egress
		inputs["fromPort"] = args.FromPort
		inputs["icmpCode"] = args.IcmpCode
		inputs["icmpType"] = args.IcmpType
		inputs["ipv6CidrBlock"] = args.Ipv6CidrBlock
		inputs["networkAclId"] = args.NetworkAclId
		inputs["protocol"] = args.Protocol
		inputs["ruleAction"] = args.RuleAction
		inputs["ruleNumber"] = args.RuleNumber
		inputs["toPort"] = args.ToPort
	}
	s, err := ctx.RegisterResource("aws:ec2/networkAclRule:NetworkAclRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkAclRule{s: s}, nil
}

// GetNetworkAclRule gets an existing NetworkAclRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAclRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NetworkAclRuleState, opts ...pulumi.ResourceOpt) (*NetworkAclRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cidrBlock"] = state.CidrBlock
		inputs["egress"] = state.Egress
		inputs["fromPort"] = state.FromPort
		inputs["icmpCode"] = state.IcmpCode
		inputs["icmpType"] = state.IcmpType
		inputs["ipv6CidrBlock"] = state.Ipv6CidrBlock
		inputs["networkAclId"] = state.NetworkAclId
		inputs["protocol"] = state.Protocol
		inputs["ruleAction"] = state.RuleAction
		inputs["ruleNumber"] = state.RuleNumber
		inputs["toPort"] = state.ToPort
	}
	s, err := ctx.ReadResource("aws:ec2/networkAclRule:NetworkAclRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &NetworkAclRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *NetworkAclRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *NetworkAclRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
func (r *NetworkAclRule) CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cidrBlock"])
}

// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
func (r *NetworkAclRule) Egress() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["egress"])
}

// The from port to match.
func (r *NetworkAclRule) FromPort() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["fromPort"])
}

// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
func (r *NetworkAclRule) IcmpCode() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["icmpCode"])
}

// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
func (r *NetworkAclRule) IcmpType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["icmpType"])
}

// The IPv6 CIDR block to allow or deny.
func (r *NetworkAclRule) Ipv6CidrBlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipv6CidrBlock"])
}

// The ID of the network ACL.
func (r *NetworkAclRule) NetworkAclId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkAclId"])
}

// The protocol. A value of -1 means all protocols.
func (r *NetworkAclRule) Protocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocol"])
}

// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
func (r *NetworkAclRule) RuleAction() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ruleAction"])
}

// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
func (r *NetworkAclRule) RuleNumber() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["ruleNumber"])
}

// The to port to match.
func (r *NetworkAclRule) ToPort() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["toPort"])
}

// Input properties used for looking up and filtering NetworkAclRule resources.
type NetworkAclRuleState struct {
	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock interface{}
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress interface{}
	// The from port to match.
	FromPort interface{}
	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
	IcmpCode interface{}
	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
	IcmpType interface{}
	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock interface{}
	// The ID of the network ACL.
	NetworkAclId interface{}
	// The protocol. A value of -1 means all protocols.
	Protocol interface{}
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction interface{}
	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber interface{}
	// The to port to match.
	ToPort interface{}
}

// The set of arguments for constructing a NetworkAclRule resource.
type NetworkAclRuleArgs struct {
	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock interface{}
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress interface{}
	// The from port to match.
	FromPort interface{}
	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
	IcmpCode interface{}
	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
	IcmpType interface{}
	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock interface{}
	// The ID of the network ACL.
	NetworkAclId interface{}
	// The protocol. A value of -1 means all protocols.
	Protocol interface{}
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction interface{}
	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber interface{}
	// The to port to match.
	ToPort interface{}
}
