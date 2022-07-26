// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Cuemby.Equinix.Outputs
{

    [OutputType]
    public sealed class GetMetalPlansPlanResult
    {
        public readonly ImmutableArray<string> AvailableInMetros;
        public readonly ImmutableArray<string> AvailableIns;
        public readonly string Class;
        public readonly ImmutableArray<string> DeploymentTypes;
        public readonly string Description;
        /// <summary>
        /// id of the plan
        /// </summary>
        public readonly string Id;
        public readonly bool Legacy;
        public readonly string Line;
        /// <summary>
        /// name of the plan
        /// - `slug`- plan slug
        /// - `description`- description of the plan
        /// - `line`- plan line, e.g. baremetal
        /// - `legacy`- flag showing if it's a legacy plan
        /// - `class`- plan class
        /// - `pricing_hour`- plan hourly price
        /// - `pricing_month`- plan monthly price
        /// - `deployment_types`- list of deployment types, e.g. on_demand, spot_market
        /// - `available_in`- list of facilities where the plan is available
        /// - `available_in_metros`- list of facilities where the plan is available
        /// </summary>
        public readonly string Name;
        public readonly double PricingHour;
        public readonly double PricingMonth;
        public readonly string Slug;

        [OutputConstructor]
        private GetMetalPlansPlanResult(
            ImmutableArray<string> availableInMetros,

            ImmutableArray<string> availableIns,

            string @class,

            ImmutableArray<string> deploymentTypes,

            string description,

            string id,

            bool legacy,

            string line,

            string name,

            double pricingHour,

            double pricingMonth,

            string slug)
        {
            AvailableInMetros = availableInMetros;
            AvailableIns = availableIns;
            Class = @class;
            DeploymentTypes = deploymentTypes;
            Description = description;
            Id = id;
            Legacy = legacy;
            Line = line;
            Name = name;
            PricingHour = pricingHour;
            PricingMonth = pricingMonth;
            Slug = slug;
        }
    }
}
