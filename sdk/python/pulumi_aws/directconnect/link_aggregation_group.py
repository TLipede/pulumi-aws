# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class LinkAggregationGroup(pulumi.CustomResource):
    """
    Provides a Direct Connect LAG.
    """
    def __init__(__self__, __name__, __opts__=None, connections_bandwidth=None, force_destroy=None, location=None, name=None, number_of_connections=None, tags=None):
        """Create a LinkAggregationGroup resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not connections_bandwidth:
            raise TypeError('Missing required property connections_bandwidth')
        __props__['connections_bandwidth'] = connections_bandwidth

        __props__['force_destroy'] = force_destroy

        if not location:
            raise TypeError('Missing required property location')
        __props__['location'] = location

        __props__['name'] = name

        __props__['number_of_connections'] = number_of_connections

        __props__['tags'] = tags

        __props__['arn'] = None

        super(LinkAggregationGroup, __self__).__init__(
            'aws:directconnect/linkAggregationGroup:LinkAggregationGroup',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

