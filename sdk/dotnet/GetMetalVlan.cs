// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Cuemby.Equinix
{
    public static class GetMetalVlan
    {
        /// <summary>
        /// Provides an Equinix Metal Virtual Network datasource. VLANs data sources can be
        /// searched by VLAN UUID, or project UUID and vxlan number.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Fetch a vlan by ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Cuemby.Equinix;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foovlan = new Equinix.MetalVlan("foovlan", new Equinix.MetalVlanArgs
        ///         {
        ///             ProjectId = local.Project_id,
        ///             Metro = "sv",
        ///             Vxlan = 5,
        ///         });
        ///         var dsvlan = Equinix.GetMetalVlan.Invoke(new Equinix.GetMetalVlanInvokeArgs
        ///         {
        ///             VlanId = foovlan.Id,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Fetch a vlan by project ID, vxlan and metro
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var dsvlan = Output.Create(Equinix.GetMetalVlan.InvokeAsync(new Equinix.GetMetalVlanArgs
        ///         {
        ///             ProjectId = local.Project_id,
        ///             Vxlan = 5,
        ///             Metro = "sv",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMetalVlanResult> InvokeAsync(GetMetalVlanArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMetalVlanResult>("equinix:index/getMetalVlan:GetMetalVlan", args ?? new GetMetalVlanArgs(), options.WithDefaults());

        /// <summary>
        /// Provides an Equinix Metal Virtual Network datasource. VLANs data sources can be
        /// searched by VLAN UUID, or project UUID and vxlan number.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// Fetch a vlan by ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Cuemby.Equinix;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var foovlan = new Equinix.MetalVlan("foovlan", new Equinix.MetalVlanArgs
        ///         {
        ///             ProjectId = local.Project_id,
        ///             Metro = "sv",
        ///             Vxlan = 5,
        ///         });
        ///         var dsvlan = Equinix.GetMetalVlan.Invoke(new Equinix.GetMetalVlanInvokeArgs
        ///         {
        ///             VlanId = foovlan.Id,
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Fetch a vlan by project ID, vxlan and metro
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var dsvlan = Output.Create(Equinix.GetMetalVlan.InvokeAsync(new Equinix.GetMetalVlanArgs
        ///         {
        ///             ProjectId = local.Project_id,
        ///             Vxlan = 5,
        ///             Metro = "sv",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetMetalVlanResult> Invoke(GetMetalVlanInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMetalVlanResult>("equinix:index/getMetalVlan:GetMetalVlan", args ?? new GetMetalVlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetalVlanArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Facility where the VLAN is deployed.
        /// </summary>
        [Input("facility")]
        public string? Facility { get; set; }

        /// <summary>
        /// Metro where the VLAN is deployed.
        /// </summary>
        [Input("metro")]
        public string? Metro { get; set; }

        /// <summary>
        /// UUID of parent project of the VLAN. Use together with the vxlan number and metro or facility.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// Metal UUID of the VLAN resource to look up.
        /// </summary>
        [Input("vlanId")]
        public string? VlanId { get; set; }

        /// <summary>
        /// vxlan number of the VLAN to look up. Use together with the project_id and metro or facility.
        /// </summary>
        [Input("vxlan")]
        public int? Vxlan { get; set; }

        public GetMetalVlanArgs()
        {
        }
    }

    public sealed class GetMetalVlanInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Facility where the VLAN is deployed.
        /// </summary>
        [Input("facility")]
        public Input<string>? Facility { get; set; }

        /// <summary>
        /// Metro where the VLAN is deployed.
        /// </summary>
        [Input("metro")]
        public Input<string>? Metro { get; set; }

        /// <summary>
        /// UUID of parent project of the VLAN. Use together with the vxlan number and metro or facility.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Metal UUID of the VLAN resource to look up.
        /// </summary>
        [Input("vlanId")]
        public Input<string>? VlanId { get; set; }

        /// <summary>
        /// vxlan number of the VLAN to look up. Use together with the project_id and metro or facility.
        /// </summary>
        [Input("vxlan")]
        public Input<int>? Vxlan { get; set; }

        public GetMetalVlanInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMetalVlanResult
    {
        /// <summary>
        /// List of device ID to which this VLAN is assigned.
        /// </summary>
        public readonly ImmutableArray<string> AssignedDevicesIds;
        /// <summary>
        /// Description text of the VLAN resource.
        /// </summary>
        public readonly string Description;
        public readonly string Facility;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Metro;
        public readonly string ProjectId;
        public readonly string VlanId;
        public readonly int Vxlan;

        [OutputConstructor]
        private GetMetalVlanResult(
            ImmutableArray<string> assignedDevicesIds,

            string description,

            string facility,

            string id,

            string metro,

            string projectId,

            string vlanId,

            int vxlan)
        {
            AssignedDevicesIds = assignedDevicesIds;
            Description = description;
            Facility = facility;
            Id = id;
            Metro = metro;
            ProjectId = projectId;
            VlanId = vlanId;
            Vxlan = vxlan;
        }
    }
}
