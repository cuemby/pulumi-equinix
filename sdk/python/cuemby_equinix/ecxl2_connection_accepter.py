# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ECXL2ConnectionAccepterArgs', 'ECXL2ConnectionAccepter']

@pulumi.input_type
class ECXL2ConnectionAccepterArgs:
    def __init__(__self__, *,
                 connection_id: pulumi.Input[str],
                 access_key: Optional[pulumi.Input[str]] = None,
                 aws_profile: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ECXL2ConnectionAccepter resource.
        :param pulumi.Input[str] connection_id: Identifier of Layer 2 connection that will be accepted.
        :param pulumi.Input[str] access_key: Access Key used to accept connection on provider side.
        :param pulumi.Input[str] aws_profile: AWS Profile Name for retrieving credentials from.
               shared credentials file
        :param pulumi.Input[str] secret_key: Secret Key used to accept connection on provider side.
        """
        pulumi.set(__self__, "connection_id", connection_id)
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if aws_profile is not None:
            pulumi.set(__self__, "aws_profile", aws_profile)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Input[str]:
        """
        Identifier of Layer 2 connection that will be accepted.
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Access Key used to accept connection on provider side.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="awsProfile")
    def aws_profile(self) -> Optional[pulumi.Input[str]]:
        """
        AWS Profile Name for retrieving credentials from.
        shared credentials file
        """
        return pulumi.get(self, "aws_profile")

    @aws_profile.setter
    def aws_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_profile", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret Key used to accept connection on provider side.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)


@pulumi.input_type
class _ECXL2ConnectionAccepterState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 aws_connection_id: Optional[pulumi.Input[str]] = None,
                 aws_profile: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ECXL2ConnectionAccepter resources.
        :param pulumi.Input[str] access_key: Access Key used to accept connection on provider side.
        :param pulumi.Input[str] aws_connection_id: Identifier of a hosted Direct Connect connection on AWS side,
               applicable for accepter resource with connections to AWS only.
        :param pulumi.Input[str] aws_profile: AWS Profile Name for retrieving credentials from.
               shared credentials file
        :param pulumi.Input[str] connection_id: Identifier of Layer 2 connection that will be accepted.
        :param pulumi.Input[str] secret_key: Secret Key used to accept connection on provider side.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if aws_connection_id is not None:
            pulumi.set(__self__, "aws_connection_id", aws_connection_id)
        if aws_profile is not None:
            pulumi.set(__self__, "aws_profile", aws_profile)
        if connection_id is not None:
            pulumi.set(__self__, "connection_id", connection_id)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Access Key used to accept connection on provider side.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="awsConnectionId")
    def aws_connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of a hosted Direct Connect connection on AWS side,
        applicable for accepter resource with connections to AWS only.
        """
        return pulumi.get(self, "aws_connection_id")

    @aws_connection_id.setter
    def aws_connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_connection_id", value)

    @property
    @pulumi.getter(name="awsProfile")
    def aws_profile(self) -> Optional[pulumi.Input[str]]:
        """
        AWS Profile Name for retrieving credentials from.
        shared credentials file
        """
        return pulumi.get(self, "aws_profile")

    @aws_profile.setter
    def aws_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_profile", value)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of Layer 2 connection that will be accepted.
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_id", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret Key used to accept connection on provider side.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)


class ECXL2ConnectionAccepter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 aws_profile: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        This resource can be imported using an existing ID

        ```sh
         $ pulumi import equinix:index/eCXL2ConnectionAccepter:ECXL2ConnectionAccepter example {existing_id}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Access Key used to accept connection on provider side.
        :param pulumi.Input[str] aws_profile: AWS Profile Name for retrieving credentials from.
               shared credentials file
        :param pulumi.Input[str] connection_id: Identifier of Layer 2 connection that will be accepted.
        :param pulumi.Input[str] secret_key: Secret Key used to accept connection on provider side.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ECXL2ConnectionAccepterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        This resource can be imported using an existing ID

        ```sh
         $ pulumi import equinix:index/eCXL2ConnectionAccepter:ECXL2ConnectionAccepter example {existing_id}
        ```

        :param str resource_name: The name of the resource.
        :param ECXL2ConnectionAccepterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ECXL2ConnectionAccepterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 aws_profile: Optional[pulumi.Input[str]] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ECXL2ConnectionAccepterArgs.__new__(ECXL2ConnectionAccepterArgs)

            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["aws_profile"] = aws_profile
            if connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'connection_id'")
            __props__.__dict__["connection_id"] = connection_id
            __props__.__dict__["secret_key"] = secret_key
            __props__.__dict__["aws_connection_id"] = None
        super(ECXL2ConnectionAccepter, __self__).__init__(
            'equinix:index/eCXL2ConnectionAccepter:ECXL2ConnectionAccepter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            aws_connection_id: Optional[pulumi.Input[str]] = None,
            aws_profile: Optional[pulumi.Input[str]] = None,
            connection_id: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None) -> 'ECXL2ConnectionAccepter':
        """
        Get an existing ECXL2ConnectionAccepter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: Access Key used to accept connection on provider side.
        :param pulumi.Input[str] aws_connection_id: Identifier of a hosted Direct Connect connection on AWS side,
               applicable for accepter resource with connections to AWS only.
        :param pulumi.Input[str] aws_profile: AWS Profile Name for retrieving credentials from.
               shared credentials file
        :param pulumi.Input[str] connection_id: Identifier of Layer 2 connection that will be accepted.
        :param pulumi.Input[str] secret_key: Secret Key used to accept connection on provider side.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ECXL2ConnectionAccepterState.__new__(_ECXL2ConnectionAccepterState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["aws_connection_id"] = aws_connection_id
        __props__.__dict__["aws_profile"] = aws_profile
        __props__.__dict__["connection_id"] = connection_id
        __props__.__dict__["secret_key"] = secret_key
        return ECXL2ConnectionAccepter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        Access Key used to accept connection on provider side.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="awsConnectionId")
    def aws_connection_id(self) -> pulumi.Output[str]:
        """
        Identifier of a hosted Direct Connect connection on AWS side,
        applicable for accepter resource with connections to AWS only.
        """
        return pulumi.get(self, "aws_connection_id")

    @property
    @pulumi.getter(name="awsProfile")
    def aws_profile(self) -> pulumi.Output[Optional[str]]:
        """
        AWS Profile Name for retrieving credentials from.
        shared credentials file
        """
        return pulumi.get(self, "aws_profile")

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Output[str]:
        """
        Identifier of Layer 2 connection that will be accepted.
        """
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        Secret Key used to accept connection on provider side.
        """
        return pulumi.get(self, "secret_key")

