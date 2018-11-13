// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Step Function State Machine resource
type StateMachine struct {
	s *pulumi.ResourceState
}

// NewStateMachine registers a new resource with the given unique name, arguments, and options.
func NewStateMachine(ctx *pulumi.Context,
	name string, args *StateMachineArgs, opts ...pulumi.ResourceOpt) (*StateMachine, error) {
	if args == nil || args.Definition == nil {
		return nil, errors.New("missing required argument 'Definition'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["definition"] = nil
		inputs["name"] = nil
		inputs["roleArn"] = nil
	} else {
		inputs["definition"] = args.Definition
		inputs["name"] = args.Name
		inputs["roleArn"] = args.RoleArn
	}
	inputs["creationDate"] = nil
	inputs["status"] = nil
	s, err := ctx.RegisterResource("aws:sfn/stateMachine:StateMachine", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StateMachine{s: s}, nil
}

// GetStateMachine gets an existing StateMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStateMachine(ctx *pulumi.Context,
	name string, id pulumi.ID, state *StateMachineState, opts ...pulumi.ResourceOpt) (*StateMachine, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationDate"] = state.CreationDate
		inputs["definition"] = state.Definition
		inputs["name"] = state.Name
		inputs["roleArn"] = state.RoleArn
		inputs["status"] = state.Status
	}
	s, err := ctx.ReadResource("aws:sfn/stateMachine:StateMachine", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &StateMachine{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *StateMachine) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *StateMachine) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The date the state machine was created.
func (r *StateMachine) CreationDate() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationDate"])
}

// The Amazon States Language definition of the state machine.
func (r *StateMachine) Definition() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["definition"])
}

// The name of the state machine.
func (r *StateMachine) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
func (r *StateMachine) RoleArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["roleArn"])
}

// The current status of the state machine. Either "ACTIVE" or "DELETING".
func (r *StateMachine) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Input properties used for looking up and filtering StateMachine resources.
type StateMachineState struct {
	// The date the state machine was created.
	CreationDate interface{}
	// The Amazon States Language definition of the state machine.
	Definition interface{}
	// The name of the state machine.
	Name interface{}
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn interface{}
	// The current status of the state machine. Either "ACTIVE" or "DELETING".
	Status interface{}
}

// The set of arguments for constructing a StateMachine resource.
type StateMachineArgs struct {
	// The Amazon States Language definition of the state machine.
	Definition interface{}
	// The name of the state machine.
	Name interface{}
	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn interface{}
}
