// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodeDeploy deployment config for an application
type DeploymentConfig struct {
	s *pulumi.ResourceState
}

// NewDeploymentConfig registers a new resource with the given unique name, arguments, and options.
func NewDeploymentConfig(ctx *pulumi.Context,
	name string, args *DeploymentConfigArgs, opts ...pulumi.ResourceOpt) (*DeploymentConfig, error) {
	if args == nil || args.DeploymentConfigName == nil {
		return nil, errors.New("missing required argument 'DeploymentConfigName'")
	}
	if args == nil || args.MinimumHealthyHosts == nil {
		return nil, errors.New("missing required argument 'MinimumHealthyHosts'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["deploymentConfigName"] = nil
		inputs["minimumHealthyHosts"] = nil
	} else {
		inputs["deploymentConfigName"] = args.DeploymentConfigName
		inputs["minimumHealthyHosts"] = args.MinimumHealthyHosts
	}
	inputs["deploymentConfigId"] = nil
	s, err := ctx.RegisterResource("aws:codedeploy/deploymentConfig:DeploymentConfig", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DeploymentConfig{s: s}, nil
}

// GetDeploymentConfig gets an existing DeploymentConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentConfig(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DeploymentConfigState, opts ...pulumi.ResourceOpt) (*DeploymentConfig, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["deploymentConfigId"] = state.DeploymentConfigId
		inputs["deploymentConfigName"] = state.DeploymentConfigName
		inputs["minimumHealthyHosts"] = state.MinimumHealthyHosts
	}
	s, err := ctx.ReadResource("aws:codedeploy/deploymentConfig:DeploymentConfig", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DeploymentConfig{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DeploymentConfig) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DeploymentConfig) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The AWS Assigned deployment config id
func (r *DeploymentConfig) DeploymentConfigId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deploymentConfigId"])
}

// The name of the deployment config.
func (r *DeploymentConfig) DeploymentConfigName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deploymentConfigName"])
}

// A minimum_healthy_hosts block. Minimum Healthy Hosts are documented below.
func (r *DeploymentConfig) MinimumHealthyHosts() *pulumi.Output {
	return r.s.State["minimumHealthyHosts"]
}

// Input properties used for looking up and filtering DeploymentConfig resources.
type DeploymentConfigState struct {
	// The AWS Assigned deployment config id
	DeploymentConfigId interface{}
	// The name of the deployment config.
	DeploymentConfigName interface{}
	// A minimum_healthy_hosts block. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts interface{}
}

// The set of arguments for constructing a DeploymentConfig resource.
type DeploymentConfigArgs struct {
	// The name of the deployment config.
	DeploymentConfigName interface{}
	// A minimum_healthy_hosts block. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts interface{}
}
