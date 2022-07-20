// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package equinix

import (
	"fmt"
	"path/filepath"

	"github.com/cuemby/pulumi-equinix/provider/pkg/version"
	"github.com/equinix/terraform-provider-equinix/equinix"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "equinix"
	// modules:
	mainMod = "index" // the equinix module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(equinix.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "equinix",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Equinix",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Cuemby",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://upload.wikimedia.org/wikipedia/commons/thumb/f/f7/Equinix_logo.svg/1200px-Equinix_logo.svg.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/cuemby/pulumi-equinix/releases/downloads/v${VERSION}",
		Description:       "A Pulumi package for creating and managing equinix cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "equinix", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.cuemby.com",
		Repository: "https://github.com/cuemby/pulumi-equinix",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "equinix",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
			"equinix_ecx_l2_connection":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ECXL2Connection")},
			"equinix_ecx_l2_connection_accepter": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ECXL2ConnectionAccepter")},
			"equinix_ecx_l2_serviceprofile":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ECXL2Serviceprofile")},
			"equinix_network_device":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkDevice")},
			"equinix_network_ssh_user":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkSSHUser")},
			"equinix_network_bgp":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkBGP")},
			"equinix_network_ssh_key":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkSSHKey")},
			"equinix_network_acl_template":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkACLTemplate")},
			"equinix_network_device_link":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NetworkDeviceLink")},
			"equinix_metal_user_api_key":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalUserAPIKey")},
			"equinix_metal_project_api_key":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalProjectAPIKey")},
			"equinix_metal_connection":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalConnection")},
			"equinix_metal_device":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalDevice")},
			"equinix_metal_device_network_type":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalDeviceNetworkType")},
			"equinix_metal_ssh_key":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalSSHKey")},
			"equinix_metal_port":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalPort")},
			"equinix_metal_project_ssh_key":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalProjectSSHKey")},
			"equinix_metal_project":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalProject")},
			"equinix_metal_organization":         {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalOrganization")},
			"equinix_metal_reserved_ip_block":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalReservedIPBlock")},
			"equinix_metal_ip_attachment":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalIPAttachment")},
			"equinix_metal_spot_market_request":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalSpotMarketRequest")},
			"equinix_metal_vlan":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalVlan")},
			"equinix_metal_virtual_circuit":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalVirtualCircuit")},
			"equinix_metal_vrf":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalVRF")},
			"equinix_metal_bgp_session":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalBGPSession")},
			"equinix_metal_port_vlan_attachment": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalPortVlanAttachment")},
			"equinix_metal_gateway":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "MetalGateway")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"equinix_ecx_port":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetECXPort")},
			"equinix_ecx_l2_sellerprofile":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetECXL2Sellerprofile")},
			"equinix_ecx_l2_sellerprofiles":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetECXL2Sellerprofiles")},
			"equinix_network_account":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNetworkAccount")},
			"equinix_network_device":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNetworkDevice")},
			"equinix_network_device_type":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNetworkDeviceType")},
			"equinix_network_device_software":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNetworkDeviceSoftware")},
			"equinix_network_device_platform":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNetworkDevicePlatform")},
			"equinix_metal_hardware_reservation": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalHardwareReservation")},
			"equinix_metal_metro":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalMetro")},
			"equinix_metal_facility":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalFacility")},
			"equinix_metal_connection":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalConnection")},
			"equinix_metal_ip_block_ranges":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalIPBlockRanges")},
			"equinix_metal_gateway":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalGateway")},
			"equinix_metal_precreated_ip_block":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalPrecreatedIPBlock")},
			"equinix_metal_operating_system":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalOperatingSystem")},
			"equinix_metal_organization":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalOrganization")},
			"equinix_metal_spot_market_price":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalSpotMarketPrice")},
			"equinix_metal_device":               {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalDevice")},
			"equinix_metal_device_bgp_neighbors": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalDeviceBGPNeighbors")},
			"equinix_metal_plans":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalPlans")},
			"equinix_metal_port":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalPort")},
			"equinix_metal_project":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalProject")},
			"equinix_metal_project_ssh_key":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalProjectSSHKey")},
			"equinix_metal_reserved_ip_block":    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalReservedIPBlock")},
			"equinix_metal_spot_market_request":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalSpotMarketRequest")},
			"equinix_metal_virtual_circuit":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalVirtualCircuit")},
			"equinix_metal_vlan":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalVlan")},
			"equinix_metal_vrf":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetMetalVRF")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
			PackageName: fmt.Sprintf("@cuemby/%s", mainPkg),
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
