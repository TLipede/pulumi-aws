// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glacier

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Gamelift Build resource.
type Build struct {
	s *pulumi.ResourceState
}

// NewBuild registers a new resource with the given unique name, arguments, and options.
func NewBuild(ctx *pulumi.Context,
	name string, args *BuildArgs, opts ...pulumi.ResourceOpt) (*Build, error) {
	if args == nil || args.OperatingSystem == nil {
		return nil, errors.New("missing required argument 'OperatingSystem'")
	}
	if args == nil || args.StorageLocation == nil {
		return nil, errors.New("missing required argument 'StorageLocation'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["operatingSystem"] = nil
		inputs["storageLocation"] = nil
		inputs["version"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["operatingSystem"] = args.OperatingSystem
		inputs["storageLocation"] = args.StorageLocation
		inputs["version"] = args.Version
	}
	s, err := ctx.RegisterResource("aws:glacier/build:Build", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Build{s: s}, nil
}

// GetBuild gets an existing Build resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBuild(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BuildState, opts ...pulumi.ResourceOpt) (*Build, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["operatingSystem"] = state.OperatingSystem
		inputs["storageLocation"] = state.StorageLocation
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("aws:glacier/build:Build", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Build{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Build) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Build) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Name of the build
func (r *Build) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
func (r *Build) OperatingSystem() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["operatingSystem"])
}

// Information indicating where your game build files are stored. See below.
func (r *Build) StorageLocation() *pulumi.Output {
	return r.s.State["storageLocation"]
}

// Version that is associated with this build.
func (r *Build) Version() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering Build resources.
type BuildState struct {
	// Name of the build
	Name interface{}
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem interface{}
	// Information indicating where your game build files are stored. See below.
	StorageLocation interface{}
	// Version that is associated with this build.
	Version interface{}
}

// The set of arguments for constructing a Build resource.
type BuildArgs struct {
	// Name of the build
	Name interface{}
	// Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
	OperatingSystem interface{}
	// Information indicating where your game build files are stored. See below.
	StorageLocation interface{}
	// Version that is associated with this build.
	Version interface{}
}
