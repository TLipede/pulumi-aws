# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class SecurityGroup(pulumi.CustomResource):
    """
    Creates a new Amazon Redshift security group. You use security groups to control access to non-VPC clusters
    """
    def __init__(__self__, __name__, __opts__=None, description=None, ingress=None, name=None):
        """Create a SecurityGroup resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        description = 'Managed by Pulumi'
        __props__['description'] = description

        if not ingress:
            raise TypeError('Missing required property ingress')
        __props__['ingress'] = ingress

        __props__['name'] = name

        super(SecurityGroup, __self__).__init__(
            'aws:redshift/securityGroup:SecurityGroup',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

