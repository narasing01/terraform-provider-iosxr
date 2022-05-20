// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type InterfaceIPv4 struct {
	Device        types.String `tfsdk:"device"`
	Id            types.String `tfsdk:"id"`
	InterfaceName types.String `tfsdk:"interface_name"`
	Address       types.String `tfsdk:"address"`
	Netmask       types.String `tfsdk:"netmask"`
	Unnumbered    types.String `tfsdk:"unnumbered"`
}

func (data InterfaceIPv4) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name=%s]/ipv4", data.InterfaceName.Value)
}

func (data InterfaceIPv4) toBody() string {
	body := "{}"

	if !data.Address.Null && !data.Address.Unknown {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.address", data.Address.Value)
	}
	if !data.Netmask.Null && !data.Netmask.Unknown {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.netmask", data.Netmask.Value)
	}
	if !data.Unnumbered.Null && !data.Unnumbered.Unknown {
		body, _ = sjson.Set(body, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.unnumbered", data.Unnumbered.Value)
	}

	return body
}

func (data *InterfaceIPv4) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.address"); value.Exists() {
		data.Address.Value = value.String()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.address.netmask"); value.Exists() {
		data.Netmask.Value = value.String()
	}
	if value := gjson.GetBytes(res, "Cisco-IOS-XR-um-if-ip-address-cfg:addresses.unnumbered"); value.Exists() {
		data.Unnumbered.Value = value.String()
	}
}

func (data *InterfaceIPv4) fromPlan(plan InterfaceIPv4) {
	data.Device = plan.Device
	data.InterfaceName.Value = plan.InterfaceName.Value
}
