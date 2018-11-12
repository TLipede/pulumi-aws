# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Route(pulumi.CustomResource):
    """
    Provides a resource to create a routing table entry (a route) in a VPC routing table.
    
    ~> **NOTE on Route Tables and Routes:** Terraform currently
    provides both a standalone Route resource and a Route Table resource with routes
    defined in-line. At this time you cannot use a Route Table with in-line routes
    in conjunction with any Route resources. Doing so will cause
    a conflict of rule settings and will overwrite rules.
    """
    def __init__(__self__, __name__, __opts__=None, destination_cidr_block=None, destination_ipv6_cidr_block=None, egress_only_gateway_id=None, gateway_id=None, instance_id=None, nat_gateway_id=None, network_interface_id=None, route_table_id=None, vpc_peering_connection_id=None):
        """Create a Route resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['destination_cidr_block'] = destination_cidr_block

        __props__['destination_ipv6_cidr_block'] = destination_ipv6_cidr_block

        __props__['egress_only_gateway_id'] = egress_only_gateway_id

        __props__['gateway_id'] = gateway_id

        __props__['instance_id'] = instance_id

        __props__['nat_gateway_id'] = nat_gateway_id

        __props__['network_interface_id'] = network_interface_id

        if not route_table_id:
            raise TypeError('Missing required property route_table_id')
        __props__['route_table_id'] = route_table_id

        __props__['vpc_peering_connection_id'] = vpc_peering_connection_id

        __props__['destination_prefix_list_id'] = None
        __props__['instance_owner_id'] = None
        __props__['origin'] = None
        __props__['state'] = None

        super(Route, __self__).__init__(
            'aws:ec2/route:Route',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

