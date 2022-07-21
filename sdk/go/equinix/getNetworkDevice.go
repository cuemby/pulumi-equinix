// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get Equinix Network Edge device details.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/cuemby/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi-equinix/sdk/go/equinix"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := equinix.GetNetworkDevice(ctx, &GetNetworkDeviceArgs{
// 			Uuid: pulumi.StringRef("f0b5c553-cdeb-4bc3-95b8-23db9ccfd5ee"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = equinix.GetNetworkDevice(ctx, &GetNetworkDeviceArgs{
// 			Name: pulumi.StringRef("Arcus-Gateway-A1"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNetworkDevice(ctx *pulumi.Context, args *LookupNetworkDeviceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkDeviceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupNetworkDeviceResult
	err := ctx.Invoke("equinix:index/getNetworkDevice:GetNetworkDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetNetworkDevice.
type LookupNetworkDeviceArgs struct {
	// Name of an existing Equinix Network Edge device
	Name *string `pulumi:"name"`
	// UUID of an existing Equinix Network Edge device
	Uuid *string `pulumi:"uuid"`
	// Device states to be considered valid when searching for a device by name
	ValidStatusList *string `pulumi:"validStatusList"`
}

// A collection of values returned by GetNetworkDevice.
type LookupNetworkDeviceResult struct {
	AccountNumber string `pulumi:"accountNumber"`
	// Unique identifier of applied ACL template
	AclTemplateId       string `pulumi:"aclTemplateId"`
	AdditionalBandwidth int    `pulumi:"additionalBandwidth"`
	// Autonomous system number
	Asn            int                             `pulumi:"asn"`
	Byol           bool                            `pulumi:"byol"`
	ClusterDetails []GetNetworkDeviceClusterDetail `pulumi:"clusterDetails"`
	CoreCount      int                             `pulumi:"coreCount"`
	Hostname       string                          `pulumi:"hostname"`
	// Device location Equinix Business Exchange name
	Ibx string `pulumi:"ibx"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	InterfaceCount int    `pulumi:"interfaceCount"`
	// List of device interfaces
	// * `interface.#.id` - interface identifier
	// * `interface.#.name` - interface name
	// * `interface.#.status` -  interface status (AVAILABLE, RESERVED, ASSIGNED)
	// * `interface.#.operational_status` - interface operational status (up or down)
	// * `interface.#.mac_address` - interface MAC address
	// * `interface.#.ip_address` - interface IP address
	// * `interface.#.assigned_type` - interface management type (Equinix Managed or empty)
	// * `interface.#.type` - interface type
	Interfaces  []GetNetworkDeviceInterface `pulumi:"interfaces"`
	LicenseFile string                      `pulumi:"licenseFile"`
	// Unique identifier of applied license file
	LicenseFileId string `pulumi:"licenseFileId"`
	// Device license registration status
	// * APPLYING_LICENSE
	// * REGISTERED
	// * APPLIED
	// * WAITING_FOR_CLUSTER_SETUP
	// * REGISTRATION_FAILED
	LicenseStatus       string   `pulumi:"licenseStatus"`
	LicenseToken        string   `pulumi:"licenseToken"`
	MetroCode           string   `pulumi:"metroCode"`
	MgmtAclTemplateUuid string   `pulumi:"mgmtAclTemplateUuid"`
	Name                string   `pulumi:"name"`
	Notifications       []string `pulumi:"notifications"`
	OrderReference      string   `pulumi:"orderReference"`
	PackageCode         string   `pulumi:"packageCode"`
	PurchaseOrderNumber string   `pulumi:"purchaseOrderNumber"`
	// Device redundancy type applicable for HA devices, either
	// primary or secondary
	RedundancyType string `pulumi:"redundancyType"`
	// Unique identifier for a redundant device applicable for HA devices
	RedundantId string `pulumi:"redundantId"`
	// Device location region
	Region           string                            `pulumi:"region"`
	SecondaryDevices []GetNetworkDeviceSecondaryDevice `pulumi:"secondaryDevices"`
	SelfManaged      bool                              `pulumi:"selfManaged"`
	// IP address of SSH enabled interface on the device
	SshIpAddress string `pulumi:"sshIpAddress"`
	// FQDN of SSH enabled interface on the device
	SshIpFqdn string                   `pulumi:"sshIpFqdn"`
	SshKeys   []GetNetworkDeviceSshKey `pulumi:"sshKeys"`
	// Device provisioning status
	// * INITIALIZING
	// * PROVISIONING
	// * PROVISIONED  (**NOTE: By default data source will only return devices in this state.  To include other states see `validStateList`**)
	// * WAITING_FOR_PRIMARY
	// * WAITING_FOR_SECONDARY
	// * WAITING_FOR_REPLICA_CLUSTER_NODES
	// * CLUSTER_SETUP_IN_PROGRESS
	// * FAILED
	// * DEPROVISIONING
	// * DEPROVISIONED
	Status         string `pulumi:"status"`
	TermLength     int    `pulumi:"termLength"`
	Throughput     int    `pulumi:"throughput"`
	ThroughputUnit string `pulumi:"throughputUnit"`
	TypeCode       string `pulumi:"typeCode"`
	// Device unique identifier
	Uuid string `pulumi:"uuid"`
	// Comma separated list of device states (from see `status` for full list) to be considered valid. Default is 'PROVISIONED'.  Case insensitive.
	ValidStatusList     *string           `pulumi:"validStatusList"`
	VendorConfiguration map[string]string `pulumi:"vendorConfiguration"`
	Version             string            `pulumi:"version"`
	WanInterfaceId      string            `pulumi:"wanInterfaceId"`
	// Device location zone code
	ZoneCode string `pulumi:"zoneCode"`
}

func LookupNetworkDeviceOutput(ctx *pulumi.Context, args LookupNetworkDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkDeviceResult, error) {
			args := v.(LookupNetworkDeviceArgs)
			r, err := LookupNetworkDevice(ctx, &args, opts...)
			var s LookupNetworkDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkDeviceResultOutput)
}

// A collection of arguments for invoking GetNetworkDevice.
type LookupNetworkDeviceOutputArgs struct {
	// Name of an existing Equinix Network Edge device
	Name pulumi.StringPtrInput `pulumi:"name"`
	// UUID of an existing Equinix Network Edge device
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
	// Device states to be considered valid when searching for a device by name
	ValidStatusList pulumi.StringPtrInput `pulumi:"validStatusList"`
}

func (LookupNetworkDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkDeviceArgs)(nil)).Elem()
}

// A collection of values returned by GetNetworkDevice.
type LookupNetworkDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkDeviceResult)(nil)).Elem()
}

func (o LookupNetworkDeviceResultOutput) ToLookupNetworkDeviceResultOutput() LookupNetworkDeviceResultOutput {
	return o
}

func (o LookupNetworkDeviceResultOutput) ToLookupNetworkDeviceResultOutputWithContext(ctx context.Context) LookupNetworkDeviceResultOutput {
	return o
}

func (o LookupNetworkDeviceResultOutput) AccountNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.AccountNumber }).(pulumi.StringOutput)
}

// Unique identifier of applied ACL template
func (o LookupNetworkDeviceResultOutput) AclTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.AclTemplateId }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) AdditionalBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) int { return v.AdditionalBandwidth }).(pulumi.IntOutput)
}

// Autonomous system number
func (o LookupNetworkDeviceResultOutput) Asn() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) int { return v.Asn }).(pulumi.IntOutput)
}

func (o LookupNetworkDeviceResultOutput) Byol() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) bool { return v.Byol }).(pulumi.BoolOutput)
}

func (o LookupNetworkDeviceResultOutput) ClusterDetails() GetNetworkDeviceClusterDetailArrayOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) []GetNetworkDeviceClusterDetail { return v.ClusterDetails }).(GetNetworkDeviceClusterDetailArrayOutput)
}

func (o LookupNetworkDeviceResultOutput) CoreCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) int { return v.CoreCount }).(pulumi.IntOutput)
}

func (o LookupNetworkDeviceResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// Device location Equinix Business Exchange name
func (o LookupNetworkDeviceResultOutput) Ibx() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Ibx }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) InterfaceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) int { return v.InterfaceCount }).(pulumi.IntOutput)
}

// List of device interfaces
// * `interface.#.id` - interface identifier
// * `interface.#.name` - interface name
// * `interface.#.status` -  interface status (AVAILABLE, RESERVED, ASSIGNED)
// * `interface.#.operational_status` - interface operational status (up or down)
// * `interface.#.mac_address` - interface MAC address
// * `interface.#.ip_address` - interface IP address
// * `interface.#.assigned_type` - interface management type (Equinix Managed or empty)
// * `interface.#.type` - interface type
func (o LookupNetworkDeviceResultOutput) Interfaces() GetNetworkDeviceInterfaceArrayOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) []GetNetworkDeviceInterface { return v.Interfaces }).(GetNetworkDeviceInterfaceArrayOutput)
}

func (o LookupNetworkDeviceResultOutput) LicenseFile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.LicenseFile }).(pulumi.StringOutput)
}

// Unique identifier of applied license file
func (o LookupNetworkDeviceResultOutput) LicenseFileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.LicenseFileId }).(pulumi.StringOutput)
}

// Device license registration status
// * APPLYING_LICENSE
// * REGISTERED
// * APPLIED
// * WAITING_FOR_CLUSTER_SETUP
// * REGISTRATION_FAILED
func (o LookupNetworkDeviceResultOutput) LicenseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.LicenseStatus }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) LicenseToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.LicenseToken }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) MetroCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.MetroCode }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) MgmtAclTemplateUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.MgmtAclTemplateUuid }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) Notifications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) []string { return v.Notifications }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkDeviceResultOutput) OrderReference() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.OrderReference }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) PackageCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.PackageCode }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) PurchaseOrderNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.PurchaseOrderNumber }).(pulumi.StringOutput)
}

// Device redundancy type applicable for HA devices, either
// primary or secondary
func (o LookupNetworkDeviceResultOutput) RedundancyType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.RedundancyType }).(pulumi.StringOutput)
}

// Unique identifier for a redundant device applicable for HA devices
func (o LookupNetworkDeviceResultOutput) RedundantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.RedundantId }).(pulumi.StringOutput)
}

// Device location region
func (o LookupNetworkDeviceResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) SecondaryDevices() GetNetworkDeviceSecondaryDeviceArrayOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) []GetNetworkDeviceSecondaryDevice { return v.SecondaryDevices }).(GetNetworkDeviceSecondaryDeviceArrayOutput)
}

func (o LookupNetworkDeviceResultOutput) SelfManaged() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) bool { return v.SelfManaged }).(pulumi.BoolOutput)
}

// IP address of SSH enabled interface on the device
func (o LookupNetworkDeviceResultOutput) SshIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.SshIpAddress }).(pulumi.StringOutput)
}

// FQDN of SSH enabled interface on the device
func (o LookupNetworkDeviceResultOutput) SshIpFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.SshIpFqdn }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) SshKeys() GetNetworkDeviceSshKeyArrayOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) []GetNetworkDeviceSshKey { return v.SshKeys }).(GetNetworkDeviceSshKeyArrayOutput)
}

// Device provisioning status
// * INITIALIZING
// * PROVISIONING
// * PROVISIONED  (**NOTE: By default data source will only return devices in this state.  To include other states see `validStateList`**)
// * WAITING_FOR_PRIMARY
// * WAITING_FOR_SECONDARY
// * WAITING_FOR_REPLICA_CLUSTER_NODES
// * CLUSTER_SETUP_IN_PROGRESS
// * FAILED
// * DEPROVISIONING
// * DEPROVISIONED
func (o LookupNetworkDeviceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) TermLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) int { return v.TermLength }).(pulumi.IntOutput)
}

func (o LookupNetworkDeviceResultOutput) Throughput() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) int { return v.Throughput }).(pulumi.IntOutput)
}

func (o LookupNetworkDeviceResultOutput) ThroughputUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.ThroughputUnit }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) TypeCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.TypeCode }).(pulumi.StringOutput)
}

// Device unique identifier
func (o LookupNetworkDeviceResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Uuid }).(pulumi.StringOutput)
}

// Comma separated list of device states (from see `status` for full list) to be considered valid. Default is 'PROVISIONED'.  Case insensitive.
func (o LookupNetworkDeviceResultOutput) ValidStatusList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) *string { return v.ValidStatusList }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkDeviceResultOutput) VendorConfiguration() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) map[string]string { return v.VendorConfiguration }).(pulumi.StringMapOutput)
}

func (o LookupNetworkDeviceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupNetworkDeviceResultOutput) WanInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.WanInterfaceId }).(pulumi.StringOutput)
}

// Device location zone code
func (o LookupNetworkDeviceResultOutput) ZoneCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkDeviceResult) string { return v.ZoneCode }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkDeviceResultOutput{})
}
