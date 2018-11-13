# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSecretVersionResult(object):
    """
    A collection of values returned by getSecretVersion.
    """
    def __init__(__self__, arn=None, secret_binary=None, secret_string=None, version_id=None, version_stages=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        """
        The ARN of the secret.
        """
        if secret_binary and not isinstance(secret_binary, str):
            raise TypeError('Expected argument secret_binary to be a str')
        __self__.secret_binary = secret_binary
        """
        The decrypted part of the protected secret information that was originally provided as a binary. Base64 encoded.
        """
        if secret_string and not isinstance(secret_string, str):
            raise TypeError('Expected argument secret_string to be a str')
        __self__.secret_string = secret_string
        """
        The decrypted part of the protected secret information that was originally provided as a string.
        """
        if version_id and not isinstance(version_id, str):
            raise TypeError('Expected argument version_id to be a str')
        __self__.version_id = version_id
        """
        The unique identifier of this version of the secret.
        """
        if version_stages and not isinstance(version_stages, list):
            raise TypeError('Expected argument version_stages to be a list')
        __self__.version_stages = version_stages
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_secret_version(secret_id=None, version_id=None, version_stage=None):
    """
    Retrieve information about a Secrets Manager secret version, including its secret value. To retrieve secret metadata, see the [`aws_secretsmanager_secret` data source](https://www.terraform.io/docs/providers/aws/d/secretsmanager_secret.html).
    """
    __args__ = dict()

    __args__['secretId'] = secret_id
    __args__['versionId'] = version_id
    __args__['versionStage'] = version_stage
    __ret__ = await pulumi.runtime.invoke('aws:secretsmanager/getSecretVersion:getSecretVersion', __args__)

    return GetSecretVersionResult(
        arn=__ret__.get('arn'),
        secret_binary=__ret__.get('secretBinary'),
        secret_string=__ret__.get('secretString'),
        version_id=__ret__.get('versionId'),
        version_stages=__ret__.get('versionStages'),
        id=__ret__.get('id'))
