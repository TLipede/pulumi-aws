# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetFileSystemResult(object):
    """
    A collection of values returned by getFileSystem.
    """
    def __init__(__self__, creation_token=None, dns_name=None, encrypted=None, file_system_id=None, kms_key_id=None, performance_mode=None, tags=None, id=None):
        if creation_token and not isinstance(creation_token, str):
            raise TypeError('Expected argument creation_token to be a str')
        __self__.creation_token = creation_token
        if dns_name and not isinstance(dns_name, str):
            raise TypeError('Expected argument dns_name to be a str')
        __self__.dns_name = dns_name
        """
        The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        """
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError('Expected argument encrypted to be a bool')
        __self__.encrypted = encrypted
        """
        Whether EFS is encrypted.
        """
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError('Expected argument file_system_id to be a str')
        __self__.file_system_id = file_system_id
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError('Expected argument kms_key_id to be a str')
        __self__.kms_key_id = kms_key_id
        """
        The ARN for the KMS encryption key.
        """
        if performance_mode and not isinstance(performance_mode, str):
            raise TypeError('Expected argument performance_mode to be a str')
        __self__.performance_mode = performance_mode
        """
        The PerformanceMode of the file system.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        """
        The list of tags assigned to the file system.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_file_system(creation_token=None, file_system_id=None, tags=None):
    """
    Provides information about an Elastic File System (EFS).
    """
    __args__ = dict()

    __args__['creationToken'] = creation_token
    __args__['fileSystemId'] = file_system_id
    __args__['tags'] = tags
    __ret__ = await pulumi.runtime.invoke('aws:efs/getFileSystem:getFileSystem', __args__)

    return GetFileSystemResult(
        creation_token=__ret__.get('creationToken'),
        dns_name=__ret__.get('dnsName'),
        encrypted=__ret__.get('encrypted'),
        file_system_id=__ret__.get('fileSystemId'),
        kms_key_id=__ret__.get('kmsKeyId'),
        performance_mode=__ret__.get('performanceMode'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))
