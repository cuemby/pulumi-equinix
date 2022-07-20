# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetMetalFacilityResult',
    'AwaitableGetMetalFacilityResult',
    'get_metal_facility',
    'get_metal_facility_output',
]

@pulumi.output_type
class GetMetalFacilityResult:
    """
    A collection of values returned by GetMetalFacility.
    """
    def __init__(__self__, capacities=None, code=None, features=None, features_requireds=None, id=None, metro=None, name=None):
        if capacities and not isinstance(capacities, list):
            raise TypeError("Expected argument 'capacities' to be a list")
        pulumi.set(__self__, "capacities", capacities)
        if code and not isinstance(code, str):
            raise TypeError("Expected argument 'code' to be a str")
        pulumi.set(__self__, "code", code)
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        pulumi.set(__self__, "features", features)
        if features_requireds and not isinstance(features_requireds, list):
            raise TypeError("Expected argument 'features_requireds' to be a list")
        pulumi.set(__self__, "features_requireds", features_requireds)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metro and not isinstance(metro, str):
            raise TypeError("Expected argument 'metro' to be a str")
        pulumi.set(__self__, "metro", metro)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacities(self) -> Optional[Sequence['outputs.GetMetalFacilityCapacityResult']]:
        return pulumi.get(self, "capacities")

    @property
    @pulumi.getter
    def code(self) -> str:
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def features(self) -> Sequence[str]:
        """
        The features of the facility.
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="featuresRequireds")
    def features_requireds(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "features_requireds")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metro(self) -> str:
        """
        The metro code the facility is part of.
        """
        return pulumi.get(self, "metro")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the facility.
        """
        return pulumi.get(self, "name")


class AwaitableGetMetalFacilityResult(GetMetalFacilityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMetalFacilityResult(
            capacities=self.capacities,
            code=self.code,
            features=self.features,
            features_requireds=self.features_requireds,
            id=self.id,
            metro=self.metro,
            name=self.name)


def get_metal_facility(capacities: Optional[Sequence[pulumi.InputType['GetMetalFacilityCapacityArgs']]] = None,
                       code: Optional[str] = None,
                       features_requireds: Optional[Sequence[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMetalFacilityResult:
    """
    Provides an Equinix Metal facility datasource.


    :param Sequence[pulumi.InputType['GetMetalFacilityCapacityArgs']] capacities: One or more device plans for which the facility must have capacity.
    :param str code: The facility code to search for facilities.
    :param Sequence[str] features_requireds: Set of feature strings that the facility must have. Some
           possible values are `baremetal`, `ibx`, `storage`, `global_ipv4`, `backend_transfer`, `layer_2`.
    """
    __args__ = dict()
    __args__['capacities'] = capacities
    __args__['code'] = code
    __args__['featuresRequireds'] = features_requireds
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('equinix:index/getMetalFacility:GetMetalFacility', __args__, opts=opts, typ=GetMetalFacilityResult).value

    return AwaitableGetMetalFacilityResult(
        capacities=__ret__.capacities,
        code=__ret__.code,
        features=__ret__.features,
        features_requireds=__ret__.features_requireds,
        id=__ret__.id,
        metro=__ret__.metro,
        name=__ret__.name)


@_utilities.lift_output_func(get_metal_facility)
def get_metal_facility_output(capacities: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetMetalFacilityCapacityArgs']]]]] = None,
                              code: Optional[pulumi.Input[str]] = None,
                              features_requireds: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMetalFacilityResult]:
    """
    Provides an Equinix Metal facility datasource.


    :param Sequence[pulumi.InputType['GetMetalFacilityCapacityArgs']] capacities: One or more device plans for which the facility must have capacity.
    :param str code: The facility code to search for facilities.
    :param Sequence[str] features_requireds: Set of feature strings that the facility must have. Some
           possible values are `baremetal`, `ibx`, `storage`, `global_ipv4`, `backend_transfer`, `layer_2`.
    """
    ...
