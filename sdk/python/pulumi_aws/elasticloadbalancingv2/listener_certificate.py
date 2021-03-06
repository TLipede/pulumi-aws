# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class ListenerCertificate(pulumi.CustomResource):
    """
    Provides a Load Balancer Listener Certificate resource.
    
    This resource is for additional certificates and does not replace the default certificate on the listener.
    
    ~> **Note:** `aws_alb_listener_certificate` is known as `aws_lb_listener_certificate`. The functionality is identical.
    """
    def __init__(__self__, __name__, __opts__=None, certificate_arn=None, listener_arn=None):
        """Create a ListenerCertificate resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not certificate_arn:
            raise TypeError('Missing required property certificate_arn')
        __props__['certificate_arn'] = certificate_arn

        if not listener_arn:
            raise TypeError('Missing required property listener_arn')
        __props__['listener_arn'] = listener_arn

        super(ListenerCertificate, __self__).__init__(
            'aws:elasticloadbalancingv2/listenerCertificate:ListenerCertificate',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

