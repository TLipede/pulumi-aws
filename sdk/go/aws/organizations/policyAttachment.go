// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to attach an AWS Organizations policy to an organization account, root, or unit.
type PolicyAttachment struct {
	s *pulumi.ResourceState
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOpt) (*PolicyAttachment, error) {
	if args == nil || args.PolicyId == nil {
		return nil, errors.New("missing required argument 'PolicyId'")
	}
	if args == nil || args.TargetId == nil {
		return nil, errors.New("missing required argument 'TargetId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["policyId"] = nil
		inputs["targetId"] = nil
	} else {
		inputs["policyId"] = args.PolicyId
		inputs["targetId"] = args.TargetId
	}
	s, err := ctx.RegisterResource("aws:organizations/policyAttachment:PolicyAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PolicyAttachment{s: s}, nil
}

// GetPolicyAttachment gets an existing PolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyAttachmentState, opts ...pulumi.ResourceOpt) (*PolicyAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["policyId"] = state.PolicyId
		inputs["targetId"] = state.TargetId
	}
	s, err := ctx.ReadResource("aws:organizations/policyAttachment:PolicyAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PolicyAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PolicyAttachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PolicyAttachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The unique identifier (ID) of the policy that you want to attach to the target.
func (r *PolicyAttachment) PolicyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyId"])
}

// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
func (r *PolicyAttachment) TargetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["targetId"])
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type PolicyAttachmentState struct {
	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId interface{}
	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId interface{}
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The unique identifier (ID) of the policy that you want to attach to the target.
	PolicyId interface{}
	// The unique identifier (ID) of the root, organizational unit, or account number that you want to attach the policy to.
	TargetId interface{}
}
