# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class AmiCopy(pulumi.CustomResource):
    """
    The "AMI copy" resource allows duplication of an Amazon Machine Image (AMI),
    including cross-region copies.
    
    If the source AMI has associated EBS snapshots, those will also be duplicated
    along with the AMI.
    
    This is useful for taking a single AMI provisioned in one region and making
    it available in another for a multi-region deployment.
    
    Copying an AMI can take several minutes. The creation of this resource will
    block until the new AMI is available for use on new instances.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, ebs_block_devices=None, encrypted=None, ephemeral_block_devices=None, kms_key_id=None, name=None, source_ami_id=None, source_ami_region=None, tags=None):
        """Create a AmiCopy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['ebs_block_devices'] = ebs_block_devices

        __props__['encrypted'] = encrypted

        __props__['ephemeral_block_devices'] = ephemeral_block_devices

        __props__['kms_key_id'] = kms_key_id

        __props__['name'] = name

        if not source_ami_id:
            raise TypeError('Missing required property source_ami_id')
        __props__['source_ami_id'] = source_ami_id

        if not source_ami_region:
            raise TypeError('Missing required property source_ami_region')
        __props__['source_ami_region'] = source_ami_region

        __props__['tags'] = tags

        __props__['architecture'] = None
        __props__['ena_support'] = None
        __props__['image_location'] = None
        __props__['kernel_id'] = None
        __props__['manage_ebs_snapshots'] = None
        __props__['ramdisk_id'] = None
        __props__['root_device_name'] = None
        __props__['root_snapshot_id'] = None
        __props__['sriov_net_support'] = None
        __props__['virtualization_type'] = None

        super(AmiCopy, __self__).__init__(
            'aws:ec2/amiCopy:AmiCopy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

