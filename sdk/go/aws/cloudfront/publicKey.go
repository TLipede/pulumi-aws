// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type PublicKey struct {
	s *pulumi.ResourceState
}

// NewPublicKey registers a new resource with the given unique name, arguments, and options.
func NewPublicKey(ctx *pulumi.Context,
	name string, args *PublicKeyArgs, opts ...pulumi.ResourceOpt) (*PublicKey, error) {
	if args == nil || args.EncodedKey == nil {
		return nil, errors.New("missing required argument 'EncodedKey'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["comment"] = nil
		inputs["encodedKey"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
	} else {
		inputs["comment"] = args.Comment
		inputs["encodedKey"] = args.EncodedKey
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
	}
	inputs["callerReference"] = nil
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("aws:cloudfront/publicKey:PublicKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PublicKey{s: s}, nil
}

// GetPublicKey gets an existing PublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PublicKeyState, opts ...pulumi.ResourceOpt) (*PublicKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["callerReference"] = state.CallerReference
		inputs["comment"] = state.Comment
		inputs["encodedKey"] = state.EncodedKey
		inputs["etag"] = state.Etag
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
	}
	s, err := ctx.ReadResource("aws:cloudfront/publicKey:PublicKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PublicKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PublicKey) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PublicKey) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Internal value used by CloudFront to allow future updates to the public key configuration.
func (r *PublicKey) CallerReference() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["callerReference"])
}

// An optional comment about the public key.
func (r *PublicKey) Comment() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["comment"])
}

// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
func (r *PublicKey) EncodedKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["encodedKey"])
}

// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
func (r *PublicKey) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The name for the public key. By default generated by Terraform.
func (r *PublicKey) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name for the public key. Conflicts with `name`.
func (r *PublicKey) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

// Input properties used for looking up and filtering PublicKey resources.
type PublicKeyState struct {
	// Internal value used by CloudFront to allow future updates to the public key configuration.
	CallerReference interface{}
	// An optional comment about the public key.
	Comment interface{}
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey interface{}
	// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
	Etag interface{}
	// The name for the public key. By default generated by Terraform.
	Name interface{}
	// The name for the public key. Conflicts with `name`.
	NamePrefix interface{}
}

// The set of arguments for constructing a PublicKey resource.
type PublicKeyArgs struct {
	// An optional comment about the public key.
	Comment interface{}
	// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
	EncodedKey interface{}
	// The name for the public key. By default generated by Terraform.
	Name interface{}
	// The name for the public key. Conflicts with `name`.
	NamePrefix interface{}
}
