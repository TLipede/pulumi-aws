# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class SmsPreferences(pulumi.CustomResource):
    """
    Provides a way to set SNS SMS preferences.
    """
    def __init__(__self__, __name__, __opts__=None, default_sender_id=None, default_sms_type=None, delivery_status_iam_role_arn=None, delivery_status_success_sampling_rate=None, monthly_spend_limit=None, usage_report_s3_bucket=None):
        """Create a SmsPreferences resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['default_sender_id'] = default_sender_id

        __props__['default_sms_type'] = default_sms_type

        __props__['delivery_status_iam_role_arn'] = delivery_status_iam_role_arn

        __props__['delivery_status_success_sampling_rate'] = delivery_status_success_sampling_rate

        __props__['monthly_spend_limit'] = monthly_spend_limit

        __props__['usage_report_s3_bucket'] = usage_report_s3_bucket

        super(SmsPreferences, __self__).__init__(
            'aws:sns/smsPreferences:SmsPreferences',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

