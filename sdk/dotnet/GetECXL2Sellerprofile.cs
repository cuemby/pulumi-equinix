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
    public static class GetECXL2Sellerprofile
    {
        /// <summary>
        /// Use this data source to get details of Equinix Fabric layer 2 seller profile with a given name
        /// and / or organization.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var aws = Output.Create(Equinix.GetECXL2Sellerprofile.InvokeAsync(new Equinix.GetECXL2SellerprofileArgs
        ///         {
        ///             Name = "AWS Direct Connect",
        ///         }));
        ///         this.Id = aws.Apply(aws =&gt; aws.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetECXL2SellerprofileResult> InvokeAsync(GetECXL2SellerprofileArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetECXL2SellerprofileResult>("equinix:index/getECXL2Sellerprofile:GetECXL2Sellerprofile", args ?? new GetECXL2SellerprofileArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get details of Equinix Fabric layer 2 seller profile with a given name
        /// and / or organization.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Equinix = Pulumi.Equinix;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var aws = Output.Create(Equinix.GetECXL2Sellerprofile.InvokeAsync(new Equinix.GetECXL2SellerprofileArgs
        ///         {
        ///             Name = "AWS Direct Connect",
        ///         }));
        ///         this.Id = aws.Apply(aws =&gt; aws.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetECXL2SellerprofileResult> Invoke(GetECXL2SellerprofileInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetECXL2SellerprofileResult>("equinix:index/getECXL2Sellerprofile:GetECXL2Sellerprofile", args ?? new GetECXL2SellerprofileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetECXL2SellerprofileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the seller profile.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Name of seller's global organization.
        /// </summary>
        [Input("organizationGlobalName")]
        public string? OrganizationGlobalName { get; set; }

        /// <summary>
        /// Name of seller's organization.
        /// </summary>
        [Input("organizationName")]
        public string? OrganizationName { get; set; }

        public GetECXL2SellerprofileArgs()
        {
        }
    }

    public sealed class GetECXL2SellerprofileInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the seller profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of seller's global organization.
        /// </summary>
        [Input("organizationGlobalName")]
        public Input<string>? OrganizationGlobalName { get; set; }

        /// <summary>
        /// Name of seller's organization.
        /// </summary>
        [Input("organizationName")]
        public Input<string>? OrganizationName { get; set; }

        public GetECXL2SellerprofileInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetECXL2SellerprofileResult
    {
        /// <summary>
        /// One or more specifications of additional buyer information attributes that
        /// can be provided in connection definition that uses given seller profile.
        /// See Additional Info Attribute below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetECXL2SellerprofileAdditionalInfoResult> AdditionalInfos;
        /// <summary>
        /// Textual description of additional information attribute.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Seller profile's encapsulation (either Dot1q or QinQ).
        /// </summary>
        public readonly string Encapsulation;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// One or more specifications of metro locations supported by seller profile.
        /// See Metro Attribute below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetECXL2SellerprofileMetroResult> Metros;
        /// <summary>
        /// Name of additional information attribute.
        /// </summary>
        public readonly string Name;
        public readonly string OrganizationGlobalName;
        public readonly string OrganizationName;
        /// <summary>
        /// Boolean that indicate if seller requires connections to be redundant
        /// </summary>
        public readonly bool RedundancyRequired;
        /// <summary>
        /// One or more specifications of speed/bandwidth supported by given seller profile.
        /// See Speed Band Attribute below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetECXL2SellerprofileSpeedBandResult> SpeedBands;
        /// <summary>
        /// Boolean that indicates if seller allows customer to enter a
        /// custom connection speed.
        /// </summary>
        public readonly bool SpeedCustomizationAllowed;
        /// <summary>
        /// Boolean that indicates if seller is deriving connection speed from an API call.
        /// </summary>
        public readonly bool SpeedFromApi;
        /// <summary>
        /// Unique identifier of the seller profile.
        /// </summary>
        public readonly string Uuid;

        [OutputConstructor]
        private GetECXL2SellerprofileResult(
            ImmutableArray<Outputs.GetECXL2SellerprofileAdditionalInfoResult> additionalInfos,

            string description,

            string encapsulation,

            string id,

            ImmutableArray<Outputs.GetECXL2SellerprofileMetroResult> metros,

            string name,

            string organizationGlobalName,

            string organizationName,

            bool redundancyRequired,

            ImmutableArray<Outputs.GetECXL2SellerprofileSpeedBandResult> speedBands,

            bool speedCustomizationAllowed,

            bool speedFromApi,

            string uuid)
        {
            AdditionalInfos = additionalInfos;
            Description = description;
            Encapsulation = encapsulation;
            Id = id;
            Metros = metros;
            Name = name;
            OrganizationGlobalName = organizationGlobalName;
            OrganizationName = organizationName;
            RedundancyRequired = redundancyRequired;
            SpeedBands = speedBands;
            SpeedCustomizationAllowed = speedCustomizationAllowed;
            SpeedFromApi = speedFromApi;
            Uuid = uuid;
        }
    }
}
