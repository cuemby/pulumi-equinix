# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetNetworkAccountResult',
    'AwaitableGetNetworkAccountResult',
    'get_network_account',
    'get_network_account_output',
]

@pulumi.output_type
class GetNetworkAccountResult:
    """
    A collection of values returned by GetNetworkAccount.
    """
    def __init__(__self__, id=None, metro_code=None, name=None, number=None, status=None, ucm_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metro_code and not isinstance(metro_code, str):
            raise TypeError("Expected argument 'metro_code' to be a str")
        pulumi.set(__self__, "metro_code", metro_code)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number and not isinstance(number, str):
            raise TypeError("Expected argument 'number' to be a str")
        pulumi.set(__self__, "number", number)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if ucm_id and not isinstance(ucm_id, str):
            raise TypeError("Expected argument 'ucm_id' to be a str")
        pulumi.set(__self__, "ucm_id", ucm_id)

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
        return pulumi.get(self, "metro_code")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def number(self) -> str:
        """
        Account unique number.
        """
        return pulumi.get(self, "number")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="ucmId")
    def ucm_id(self) -> str:
        """
        Account unique identifier.
        """
        return pulumi.get(self, "ucm_id")


class AwaitableGetNetworkAccountResult(GetNetworkAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkAccountResult(
            id=self.id,
            metro_code=self.metro_code,
            name=self.name,
            number=self.number,
            status=self.status,
            ucm_id=self.ucm_id)


def get_network_account(metro_code: Optional[str] = None,
                        name: Optional[str] = None,
                        status: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkAccountResult:
    """
    Use this data source to get number and identifier of Equinix Network Edge
    billing account in a given metro location.

    Billing account reference is required to create Network Edge virtual device
    in corresponding metro location.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix as equinix

    dc = equinix.get_network_account(metro_code="DC",
        status="Active")
    pulumi.export("number", dc.number)
    ```


    :param str metro_code: Account location metro code.
    :param str name: Account name for filtering.
    :param str status: Account status for filtering. Possible values are: `Active`, `Processing`,
           `Submitted`, `Staged`.
    """
    __args__ = dict()
    __args__['metroCode'] = metro_code
    __args__['name'] = name
    __args__['status'] = status
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('equinix:index/getNetworkAccount:GetNetworkAccount', __args__, opts=opts, typ=GetNetworkAccountResult).value

    return AwaitableGetNetworkAccountResult(
        id=__ret__.id,
        metro_code=__ret__.metro_code,
        name=__ret__.name,
        number=__ret__.number,
        status=__ret__.status,
        ucm_id=__ret__.ucm_id)


@_utilities.lift_output_func(get_network_account)
def get_network_account_output(metro_code: Optional[pulumi.Input[str]] = None,
                               name: Optional[pulumi.Input[Optional[str]]] = None,
                               status: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkAccountResult]:
    """
    Use this data source to get number and identifier of Equinix Network Edge
    billing account in a given metro location.

    Billing account reference is required to create Network Edge virtual device
    in corresponding metro location.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix as equinix

    dc = equinix.get_network_account(metro_code="DC",
        status="Active")
    pulumi.export("number", dc.number)
    ```


    :param str metro_code: Account location metro code.
    :param str name: Account name for filtering.
    :param str status: Account status for filtering. Possible values are: `Active`, `Processing`,
           `Submitted`, `Staged`.
    """
    ...
