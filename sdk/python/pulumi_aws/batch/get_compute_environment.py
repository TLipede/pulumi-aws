# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetComputeEnvironmentResult(object):
    """
    A collection of values returned by getComputeEnvironment.
    """
    def __init__(__self__, arn=None, ecs_cluster_arn=None, service_role=None, state=None, status=None, status_reason=None, type=None):
        if not arn:
            raise TypeError('Missing required argument arn')
        elif not isinstance(arn, basestring):
            raise TypeError('Expected argument arn to be a basestring')
        __self__.arn = arn
        """
        The ARN of the compute environment.
        """
        if not ecs_cluster_arn:
            raise TypeError('Missing required argument ecs_cluster_arn')
        elif not isinstance(ecs_cluster_arn, basestring):
            raise TypeError('Expected argument ecs_cluster_arn to be a basestring')
        __self__.ecs_cluster_arn = ecs_cluster_arn
        """
        The ARN of the underlying Amazon ECS cluster used by the compute environment.
        """
        if not service_role:
            raise TypeError('Missing required argument service_role')
        elif not isinstance(service_role, basestring):
            raise TypeError('Expected argument service_role to be a basestring')
        __self__.service_role = service_role
        """
        The ARN of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
        """
        if not state:
            raise TypeError('Missing required argument state')
        elif not isinstance(state, basestring):
            raise TypeError('Expected argument state to be a basestring')
        __self__.state = state
        """
        The state of the compute environment (for example, `ENABLED` or `DISABLED`). If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues.
        """
        if not status:
            raise TypeError('Missing required argument status')
        elif not isinstance(status, basestring):
            raise TypeError('Expected argument status to be a basestring')
        __self__.status = status
        """
        The current status of the compute environment (for example, `CREATING` or `VALID`).
        """
        if not status_reason:
            raise TypeError('Missing required argument status_reason')
        elif not isinstance(status_reason, basestring):
            raise TypeError('Expected argument status_reason to be a basestring')
        __self__.status_reason = status_reason
        """
        A short, human-readable string to provide additional details about the current status of the compute environment.
        """
        if not type:
            raise TypeError('Missing required argument type')
        elif not isinstance(type, basestring):
            raise TypeError('Expected argument type to be a basestring')
        __self__.type = type
        """
        The type of the compute environment (for example, `MANAGED` or `UNMANAGED`).
        """

def get_compute_environment(compute_environment_name=None):
    """
    The Batch Compute Environment data source allows access to details of a specific
    compute environment within AWS Batch.
    """
    __args__ = dict()

    __args__['computeEnvironmentName'] = compute_environment_name
    __ret__ = pulumi.runtime.invoke('aws:batch/getComputeEnvironment:getComputeEnvironment', __args__)

    return GetComputeEnvironmentResult(
        arn=__ret__['arn'],
        ecs_cluster_arn=__ret__['ecsClusterArn'],
        service_role=__ret__['serviceRole'],
        state=__ret__['state'],
        status=__ret__['status'],
        status_reason=__ret__['statusReason'],
        type=__ret__['type'])