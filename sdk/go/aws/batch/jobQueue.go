// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Batch Job Queue resource.
type JobQueue struct {
	s *pulumi.ResourceState
}

// NewJobQueue registers a new resource with the given unique name, arguments, and options.
func NewJobQueue(ctx *pulumi.Context,
	name string, args *JobQueueArgs, opts ...pulumi.ResourceOpt) (*JobQueue, error) {
	if args == nil || args.ComputeEnvironments == nil {
		return nil, errors.New("missing required argument 'ComputeEnvironments'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.State == nil {
		return nil, errors.New("missing required argument 'State'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["computeEnvironments"] = nil
		inputs["name"] = nil
		inputs["priority"] = nil
		inputs["state"] = nil
	} else {
		inputs["computeEnvironments"] = args.ComputeEnvironments
		inputs["name"] = args.Name
		inputs["priority"] = args.Priority
		inputs["state"] = args.State
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:batch/jobQueue:JobQueue", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobQueue{s: s}, nil
}

// GetJobQueue gets an existing JobQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobQueue(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JobQueueState, opts ...pulumi.ResourceOpt) (*JobQueue, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["computeEnvironments"] = state.ComputeEnvironments
		inputs["name"] = state.Name
		inputs["priority"] = state.Priority
		inputs["state"] = state.State
	}
	s, err := ctx.ReadResource("aws:batch/jobQueue:JobQueue", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobQueue{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *JobQueue) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *JobQueue) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name of the job queue.
func (r *JobQueue) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Specifies the set of compute environments
// mapped to a job queue and their order.  The position of the compute environments
// in the list will dictate the order. You can associate up to 3 compute environments
// with a job queue.
func (r *JobQueue) ComputeEnvironments() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["computeEnvironments"])
}

// Specifies the name of the job queue.
func (r *JobQueue) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The priority of the job queue. Job queues with a higher priority
// are evaluated first when associated with the same compute environment.
func (r *JobQueue) Priority() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["priority"])
}

// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
func (r *JobQueue) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// Input properties used for looking up and filtering JobQueue resources.
type JobQueueState struct {
	// The Amazon Resource Name of the job queue.
	Arn interface{}
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments interface{}
	// Specifies the name of the job queue.
	Name interface{}
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority interface{}
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State interface{}
}

// The set of arguments for constructing a JobQueue resource.
type JobQueueArgs struct {
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments interface{}
	// Specifies the name of the job queue.
	Name interface{}
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority interface{}
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State interface{}
}
