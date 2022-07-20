# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetMetalVirtualCircuitResult',
    'AwaitableGetMetalVirtualCircuitResult',
    'get_metal_virtual_circuit',
    'get_metal_virtual_circuit_output',
]

@pulumi.output_type
class GetMetalVirtualCircuitResult:
    """
    A collection of values returned by GetMetalVirtualCircuit.
    """
    def __init__(__self__, connection_id=None, customer_ip=None, description=None, id=None, md5=None, metal_ip=None, name=None, nni_vlan=None, nni_vnid=None, peer_asn=None, port_id=None, project_id=None, speed=None, status=None, subnet=None, tags=None, virtual_circuit_id=None, vlan_id=None, vnid=None, vrf_id=None):
        if connection_id and not isinstance(connection_id, str):
            raise TypeError("Expected argument 'connection_id' to be a str")
        pulumi.set(__self__, "connection_id", connection_id)
        if customer_ip and not isinstance(customer_ip, str):
            raise TypeError("Expected argument 'customer_ip' to be a str")
        pulumi.set(__self__, "customer_ip", customer_ip)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if md5 and not isinstance(md5, str):
            raise TypeError("Expected argument 'md5' to be a str")
        pulumi.set(__self__, "md5", md5)
        if metal_ip and not isinstance(metal_ip, str):
            raise TypeError("Expected argument 'metal_ip' to be a str")
        pulumi.set(__self__, "metal_ip", metal_ip)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nni_vlan and not isinstance(nni_vlan, int):
            raise TypeError("Expected argument 'nni_vlan' to be a int")
        pulumi.set(__self__, "nni_vlan", nni_vlan)
        if nni_vnid and not isinstance(nni_vnid, int):
            raise TypeError("Expected argument 'nni_vnid' to be a int")
        pulumi.set(__self__, "nni_vnid", nni_vnid)
        if peer_asn and not isinstance(peer_asn, int):
            raise TypeError("Expected argument 'peer_asn' to be a int")
        pulumi.set(__self__, "peer_asn", peer_asn)
        if port_id and not isinstance(port_id, str):
            raise TypeError("Expected argument 'port_id' to be a str")
        pulumi.set(__self__, "port_id", port_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if speed and not isinstance(speed, str):
            raise TypeError("Expected argument 'speed' to be a str")
        pulumi.set(__self__, "speed", speed)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subnet and not isinstance(subnet, str):
            raise TypeError("Expected argument 'subnet' to be a str")
        pulumi.set(__self__, "subnet", subnet)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if virtual_circuit_id and not isinstance(virtual_circuit_id, str):
            raise TypeError("Expected argument 'virtual_circuit_id' to be a str")
        pulumi.set(__self__, "virtual_circuit_id", virtual_circuit_id)
        if vlan_id and not isinstance(vlan_id, str):
            raise TypeError("Expected argument 'vlan_id' to be a str")
        pulumi.set(__self__, "vlan_id", vlan_id)
        if vnid and not isinstance(vnid, int):
            raise TypeError("Expected argument 'vnid' to be a int")
        pulumi.set(__self__, "vnid", vnid)
        if vrf_id and not isinstance(vrf_id, str):
            raise TypeError("Expected argument 'vrf_id' to be a str")
        pulumi.set(__self__, "vrf_id", vrf_id)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> str:
        """
        UUID of Connection where the VC is scoped to.
        """
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="customerIp")
    def customer_ip(self) -> str:
        """
        The Customer IP address which the CSR switch will peer with. Will default to the other usable IP in the subnet.
        """
        return pulumi.get(self, "customer_ip")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description for the Virtual Circuit resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def md5(self) -> str:
        """
        The password that can be set for the VRF BGP peer
        """
        return pulumi.get(self, "md5")

    @property
    @pulumi.getter(name="metalIp")
    def metal_ip(self) -> str:
        """
        The Metal IP address for the SVI (Switch Virtual Interface) of the VirtualCircuit. Will default to the first usable IP in the subnet.
        """
        return pulumi.get(self, "metal_ip")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the virtual circuit resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nniVlan")
    def nni_vlan(self) -> int:
        return pulumi.get(self, "nni_vlan")

    @property
    @pulumi.getter(name="nniVnid")
    def nni_vnid(self) -> int:
        return pulumi.get(self, "nni_vnid")

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> int:
        """
        The BGP ASN of the peer. The same ASN may be the used across several VCs, but it cannot be the same as the local_asn of the VRF.
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="portId")
    def port_id(self) -> str:
        """
        UUID of the Connection Port where the VC is scoped to.
        """
        return pulumi.get(self, "port_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        ID of project to which the VC belongs.
        * `vnid`, `nni_vlan`, `nni_nvid` - VLAN parameters, see the
        [documentation for Equinix Fabric](https://metal.equinix.com/developers/docs/networking/fabric/).
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def speed(self) -> str:
        """
        Speed of the Virtual Circuit resource.
        """
        return pulumi.get(self, "speed")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the virtal circuit.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def subnet(self) -> str:
        """
        A subnet from one of the IP
        blocks associated with the VRF that we will help create an IP reservation for. Can only be either a /30 or /31.
        * For a /31 block, it will only have two IP addresses, which will be used for
        the metal_ip and customer_ip.
        * For a /30 block, it will have four IP addresses, but the first and last IP addresses are not usable. We will default to the first usable IP address for the metal_ip.
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags for the Virtual Circuit resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualCircuitId")
    def virtual_circuit_id(self) -> str:
        return pulumi.get(self, "virtual_circuit_id")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> str:
        return pulumi.get(self, "vlan_id")

    @property
    @pulumi.getter
    def vnid(self) -> int:
        return pulumi.get(self, "vnid")

    @property
    @pulumi.getter(name="vrfId")
    def vrf_id(self) -> str:
        """
        UUID of the VLAN to associate.
        """
        return pulumi.get(self, "vrf_id")


class AwaitableGetMetalVirtualCircuitResult(GetMetalVirtualCircuitResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMetalVirtualCircuitResult(
            connection_id=self.connection_id,
            customer_ip=self.customer_ip,
            description=self.description,
            id=self.id,
            md5=self.md5,
            metal_ip=self.metal_ip,
            name=self.name,
            nni_vlan=self.nni_vlan,
            nni_vnid=self.nni_vnid,
            peer_asn=self.peer_asn,
            port_id=self.port_id,
            project_id=self.project_id,
            speed=self.speed,
            status=self.status,
            subnet=self.subnet,
            tags=self.tags,
            virtual_circuit_id=self.virtual_circuit_id,
            vlan_id=self.vlan_id,
            vnid=self.vnid,
            vrf_id=self.vrf_id)


def get_metal_virtual_circuit(virtual_circuit_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMetalVirtualCircuitResult:
    """
    Use this data source to retrieve a virtual circuit resource from
    [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)

    > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix as equinix

    example_connection = equinix.get_metal_connection(connection_id="4347e805-eb46-4699-9eb9-5c116e6a017d")
    example_vc = equinix.get_metal_virtual_circuit(virtual_circuit_id=example_connection.ports[1].virtual_circuit_ids[0])
    ```


    :param str virtual_circuit_id: ID of the virtual circuit resource
    """
    __args__ = dict()
    __args__['virtualCircuitId'] = virtual_circuit_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('equinix:index/getMetalVirtualCircuit:GetMetalVirtualCircuit', __args__, opts=opts, typ=GetMetalVirtualCircuitResult).value

    return AwaitableGetMetalVirtualCircuitResult(
        connection_id=__ret__.connection_id,
        customer_ip=__ret__.customer_ip,
        description=__ret__.description,
        id=__ret__.id,
        md5=__ret__.md5,
        metal_ip=__ret__.metal_ip,
        name=__ret__.name,
        nni_vlan=__ret__.nni_vlan,
        nni_vnid=__ret__.nni_vnid,
        peer_asn=__ret__.peer_asn,
        port_id=__ret__.port_id,
        project_id=__ret__.project_id,
        speed=__ret__.speed,
        status=__ret__.status,
        subnet=__ret__.subnet,
        tags=__ret__.tags,
        virtual_circuit_id=__ret__.virtual_circuit_id,
        vlan_id=__ret__.vlan_id,
        vnid=__ret__.vnid,
        vrf_id=__ret__.vrf_id)


@_utilities.lift_output_func(get_metal_virtual_circuit)
def get_metal_virtual_circuit_output(virtual_circuit_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMetalVirtualCircuitResult]:
    """
    Use this data source to retrieve a virtual circuit resource from
    [Equinix Fabric - software-defined interconnections](https://metal.equinix.com/developers/docs/networking/fabric/)

    > VRF features are not generally available. The interfaces related to VRF resources may change ahead of general availability.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_equinix as equinix

    example_connection = equinix.get_metal_connection(connection_id="4347e805-eb46-4699-9eb9-5c116e6a017d")
    example_vc = equinix.get_metal_virtual_circuit(virtual_circuit_id=example_connection.ports[1].virtual_circuit_ids[0])
    ```


    :param str virtual_circuit_id: ID of the virtual circuit resource
    """
    ...
