# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetECXPortResult',
    'AwaitableGetECXPortResult',
    'get_ecx_port',
    'get_ecx_port_output',
]

@pulumi.output_type
class GetECXPortResult:
    """
    A collection of values returned by GetECXPort.
    """
    def __init__(__self__, bandwidth=None, buyout=None, encapsulation=None, ibx=None, id=None, metro_code=None, name=None, priority=None, region=None, status=None, uuid=None):
        if bandwidth and not isinstance(bandwidth, str):
            raise TypeError("Expected argument 'bandwidth' to be a str")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if buyout and not isinstance(buyout, bool):
            raise TypeError("Expected argument 'buyout' to be a bool")
        pulumi.set(__self__, "buyout", buyout)
        if encapsulation and not isinstance(encapsulation, str):
            raise TypeError("Expected argument 'encapsulation' to be a str")
        pulumi.set(__self__, "encapsulation", encapsulation)
        if ibx and not isinstance(ibx, str):
            raise TypeError("Expected argument 'ibx' to be a str")
        pulumi.set(__self__, "ibx", ibx)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metro_code and not isinstance(metro_code, str):
            raise TypeError("Expected argument 'metro_code' to be a str")
        pulumi.set(__self__, "metro_code", metro_code)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, str):
            raise TypeError("Expected argument 'priority' to be a str")
        pulumi.set(__self__, "priority", priority)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter
    def bandwidth(self) -> str:
        """
        Port Bandwidth in bytes.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def buyout(self) -> bool:
        """
        Boolean value that indicates whether the port supports unlimited connections. If
        `false`, the port is a standard port with limited connections. If `true`, the port is an
        `unlimited connections` port that allows multiple connections at no additional charge.
        """
        return pulumi.get(self, "buyout")

    @property
    @pulumi.getter
    def encapsulation(self) -> str:
        """
        The VLAN encapsulation of the port (Dot1q or QinQ).
        """
        return pulumi.get(self, "encapsulation")

    @property
    @pulumi.getter
    def ibx(self) -> str:
        """
        Port location Equinix Business Exchange (IBX).
        """
        return pulumi.get(self, "ibx")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="metroCode")
    def metro_code(self) -> str:
        """
        Port location metro code.
        """
        return pulumi.get(self, "metro_code")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> str:
        """
        The priority of the device (primary / secondary) where the port
        resides.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Port location region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Port status that indicates whether a port has been assigned or is ready for
        connection.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        """
        Unique identifier of the port.
        """
        return pulumi.get(self, "uuid")


class AwaitableGetECXPortResult(GetECXPortResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetECXPortResult(
            bandwidth=self.bandwidth,
            buyout=self.buyout,
            encapsulation=self.encapsulation,
            ibx=self.ibx,
            id=self.id,
            metro_code=self.metro_code,
            name=self.name,
            priority=self.priority,
            region=self.region,
            status=self.status,
            uuid=self.uuid)


def get_ecx_port(name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetECXPortResult:
    """
    Use this data source to get details of Equinix Fabric port with a given name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix as equinix

    tf_pri_dot1q = equinix.get_ecx_port(name="sit-001-CX-NY5-NL-Dot1q-BO-10G-PRI-JP-157")
    pulumi.export("id", tf_pri_dot1q.id)
    ```


    :param str name: Name of the port.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('equinix:index/getECXPort:GetECXPort', __args__, opts=opts, typ=GetECXPortResult).value

    return AwaitableGetECXPortResult(
        bandwidth=__ret__.bandwidth,
        buyout=__ret__.buyout,
        encapsulation=__ret__.encapsulation,
        ibx=__ret__.ibx,
        id=__ret__.id,
        metro_code=__ret__.metro_code,
        name=__ret__.name,
        priority=__ret__.priority,
        region=__ret__.region,
        status=__ret__.status,
        uuid=__ret__.uuid)


@_utilities.lift_output_func(get_ecx_port)
def get_ecx_port_output(name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetECXPortResult]:
    """
    Use this data source to get details of Equinix Fabric port with a given name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix as equinix

    tf_pri_dot1q = equinix.get_ecx_port(name="sit-001-CX-NY5-NL-Dot1q-BO-10G-PRI-JP-157")
    pulumi.export("id", tf_pri_dot1q.id)
    ```


    :param str name: Name of the port.
    """
    ...
