# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetLocalDiskResult(object):
    """
    A collection of values returned by getLocalDisk.
    """
    def __init__(__self__, disk_id=None, id=None):
        if disk_id and not isinstance(disk_id, str):
            raise TypeError('Expected argument disk_id to be a str')
        __self__.disk_id = disk_id
        """
        The disk identifier. e.g. `pci-0000:03:00.0-scsi-0:0:0:0`
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_local_disk(disk_node=None, disk_path=None, gateway_arn=None):
    """
    Retrieve information about a Storage Gateway local disk. The disk identifier is useful for adding the disk as a cache or upload buffer to a gateway.
    """
    __args__ = dict()

    __args__['diskNode'] = disk_node
    __args__['diskPath'] = disk_path
    __args__['gatewayArn'] = gateway_arn
    __ret__ = await pulumi.runtime.invoke('aws:storagegateway/getLocalDisk:getLocalDisk', __args__)

    return GetLocalDiskResult(
        disk_id=__ret__.get('diskId'),
        id=__ret__.get('id'))
