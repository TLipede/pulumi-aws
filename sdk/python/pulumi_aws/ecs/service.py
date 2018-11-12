# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Service(pulumi.CustomResource):
    """
    -> **Note:** To prevent a race condition during service deletion, make sure to set `depends_on` to the related `aws_iam_role_policy`; otherwise, the policy may be destroyed too soon and the ECS service will then get stuck in the `DRAINING` state.
    
    Provides an ECS service - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).
    
    See [ECS Services section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html).
    """
    def __init__(__self__, __name__, __opts__=None, cluster=None, deployment_maximum_percent=None, deployment_minimum_healthy_percent=None, desired_count=None, health_check_grace_period_seconds=None, iam_role=None, launch_type=None, load_balancers=None, name=None, network_configuration=None, ordered_placement_strategies=None, placement_constraints=None, placement_strategies=None, scheduling_strategy=None, service_registries=None, task_definition=None, wait_for_steady_state=None):
        """Create a Service resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cluster'] = cluster

        __props__['deployment_maximum_percent'] = deployment_maximum_percent

        __props__['deployment_minimum_healthy_percent'] = deployment_minimum_healthy_percent

        __props__['desired_count'] = desired_count

        __props__['health_check_grace_period_seconds'] = health_check_grace_period_seconds

        __props__['iam_role'] = iam_role

        __props__['launch_type'] = launch_type

        __props__['load_balancers'] = load_balancers

        __props__['name'] = name

        __props__['network_configuration'] = network_configuration

        __props__['ordered_placement_strategies'] = ordered_placement_strategies

        __props__['placement_constraints'] = placement_constraints

        __props__['placement_strategies'] = placement_strategies

        __props__['scheduling_strategy'] = scheduling_strategy

        __props__['service_registries'] = service_registries

        if not task_definition:
            raise TypeError('Missing required property task_definition')
        __props__['task_definition'] = task_definition

        __props__['wait_for_steady_state'] = wait_for_steady_state

        super(Service, __self__).__init__(
            'aws:ecs/service:Service',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

