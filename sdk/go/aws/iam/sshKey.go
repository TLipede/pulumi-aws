// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Uploads an SSH public key and associates it with the specified IAM user.
type SshKey struct {
	s *pulumi.ResourceState
}

// NewSshKey registers a new resource with the given unique name, arguments, and options.
func NewSshKey(ctx *pulumi.Context,
	name string, args *SshKeyArgs, opts ...pulumi.ResourceOpt) (*SshKey, error) {
	if args == nil || args.Encoding == nil {
		return nil, errors.New("missing required argument 'Encoding'")
	}
	if args == nil || args.PublicKey == nil {
		return nil, errors.New("missing required argument 'PublicKey'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["encoding"] = nil
		inputs["publicKey"] = nil
		inputs["status"] = nil
		inputs["username"] = nil
	} else {
		inputs["encoding"] = args.Encoding
		inputs["publicKey"] = args.PublicKey
		inputs["status"] = args.Status
		inputs["username"] = args.Username
	}
	inputs["fingerprint"] = nil
	inputs["sshPublicKeyId"] = nil
	s, err := ctx.RegisterResource("aws:iam/sshKey:SshKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SshKey{s: s}, nil
}

// GetSshKey gets an existing SshKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SshKeyState, opts ...pulumi.ResourceOpt) (*SshKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["encoding"] = state.Encoding
		inputs["fingerprint"] = state.Fingerprint
		inputs["publicKey"] = state.PublicKey
		inputs["sshPublicKeyId"] = state.SshPublicKeyId
		inputs["status"] = state.Status
		inputs["username"] = state.Username
	}
	s, err := ctx.ReadResource("aws:iam/sshKey:SshKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SshKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SshKey) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SshKey) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
func (r *SshKey) Encoding() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["encoding"])
}

// The MD5 message digest of the SSH public key.
func (r *SshKey) Fingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fingerprint"])
}

// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
func (r *SshKey) PublicKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicKey"])
}

// The unique identifier for the SSH public key.
func (r *SshKey) SshPublicKeyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sshPublicKeyId"])
}

// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
func (r *SshKey) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// The name of the IAM user to associate the SSH public key with.
func (r *SshKey) Username() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["username"])
}

// Input properties used for looking up and filtering SshKey resources.
type SshKeyState struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding interface{}
	// The MD5 message digest of the SSH public key.
	Fingerprint interface{}
	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey interface{}
	// The unique identifier for the SSH public key.
	SshPublicKeyId interface{}
	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status interface{}
	// The name of the IAM user to associate the SSH public key with.
	Username interface{}
}

// The set of arguments for constructing a SshKey resource.
type SshKeyArgs struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding interface{}
	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey interface{}
	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status interface{}
	// The name of the IAM user to associate the SSH public key with.
	Username interface{}
}
