// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type LLDP struct {
	Device                             types.String `tfsdk:"device"`
	Id                                 types.String `tfsdk:"id"`
	DeleteMode                         types.String `tfsdk:"delete_mode"`
	Holdtime                           types.Int64  `tfsdk:"holdtime"`
	Timer                              types.Int64  `tfsdk:"timer"`
	Reinit                             types.Int64  `tfsdk:"reinit"`
	SubinterfacesEnable                types.Bool   `tfsdk:"subinterfaces_enable"`
	ManagementEnable                   types.Bool   `tfsdk:"management_enable"`
	PriorityaddrEnable                 types.Bool   `tfsdk:"priorityaddr_enable"`
	ExtendedShowWidthEnable            types.Bool   `tfsdk:"extended_show_width_enable"`
	TlvSelectManagementAddressDisable  types.Bool   `tfsdk:"tlv_select_management_address_disable"`
	TlvSelectPortDescriptionDisable    types.Bool   `tfsdk:"tlv_select_port_description_disable"`
	TlvSelectSystemCapabilitiesDisable types.Bool   `tfsdk:"tlv_select_system_capabilities_disable"`
	TlvSelectSystemDescriptionDisable  types.Bool   `tfsdk:"tlv_select_system_description_disable"`
	TlvSelectSystemNameDisable         types.Bool   `tfsdk:"tlv_select_system_name_disable"`
}

type LLDPData struct {
	Device                             types.String `tfsdk:"device"`
	Id                                 types.String `tfsdk:"id"`
	Holdtime                           types.Int64  `tfsdk:"holdtime"`
	Timer                              types.Int64  `tfsdk:"timer"`
	Reinit                             types.Int64  `tfsdk:"reinit"`
	SubinterfacesEnable                types.Bool   `tfsdk:"subinterfaces_enable"`
	ManagementEnable                   types.Bool   `tfsdk:"management_enable"`
	PriorityaddrEnable                 types.Bool   `tfsdk:"priorityaddr_enable"`
	ExtendedShowWidthEnable            types.Bool   `tfsdk:"extended_show_width_enable"`
	TlvSelectManagementAddressDisable  types.Bool   `tfsdk:"tlv_select_management_address_disable"`
	TlvSelectPortDescriptionDisable    types.Bool   `tfsdk:"tlv_select_port_description_disable"`
	TlvSelectSystemCapabilitiesDisable types.Bool   `tfsdk:"tlv_select_system_capabilities_disable"`
	TlvSelectSystemDescriptionDisable  types.Bool   `tfsdk:"tlv_select_system_description_disable"`
	TlvSelectSystemNameDisable         types.Bool   `tfsdk:"tlv_select_system_name_disable"`
}

func (data LLDP) getPath() string {
	return "Cisco-IOS-XR-um-lldp-cfg:/lldp"
}

func (data LLDPData) getPath() string {
	return "Cisco-IOS-XR-um-lldp-cfg:/lldp"
}

func (data LLDP) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Holdtime.IsNull() && !data.Holdtime.IsUnknown() {
		body, _ = sjson.Set(body, "holdtime", strconv.FormatInt(data.Holdtime.ValueInt64(), 10))
	}
	if !data.Timer.IsNull() && !data.Timer.IsUnknown() {
		body, _ = sjson.Set(body, "timer", strconv.FormatInt(data.Timer.ValueInt64(), 10))
	}
	if !data.Reinit.IsNull() && !data.Reinit.IsUnknown() {
		body, _ = sjson.Set(body, "reinit", strconv.FormatInt(data.Reinit.ValueInt64(), 10))
	}
	if !data.SubinterfacesEnable.IsNull() && !data.SubinterfacesEnable.IsUnknown() {
		if data.SubinterfacesEnable.ValueBool() {
			body, _ = sjson.Set(body, "subinterfaces.enable", map[string]string{})
		}
	}
	if !data.ManagementEnable.IsNull() && !data.ManagementEnable.IsUnknown() {
		if data.ManagementEnable.ValueBool() {
			body, _ = sjson.Set(body, "management.enable", map[string]string{})
		}
	}
	if !data.PriorityaddrEnable.IsNull() && !data.PriorityaddrEnable.IsUnknown() {
		if data.PriorityaddrEnable.ValueBool() {
			body, _ = sjson.Set(body, "priorityaddr.enable", map[string]string{})
		}
	}
	if !data.ExtendedShowWidthEnable.IsNull() && !data.ExtendedShowWidthEnable.IsUnknown() {
		if data.ExtendedShowWidthEnable.ValueBool() {
			body, _ = sjson.Set(body, "extended-show-width.enable", map[string]string{})
		}
	}
	if !data.TlvSelectManagementAddressDisable.IsNull() && !data.TlvSelectManagementAddressDisable.IsUnknown() {
		if data.TlvSelectManagementAddressDisable.ValueBool() {
			body, _ = sjson.Set(body, "tlv-select.management-address.disable", map[string]string{})
		}
	}
	if !data.TlvSelectPortDescriptionDisable.IsNull() && !data.TlvSelectPortDescriptionDisable.IsUnknown() {
		if data.TlvSelectPortDescriptionDisable.ValueBool() {
			body, _ = sjson.Set(body, "tlv-select.port-description.disable", map[string]string{})
		}
	}
	if !data.TlvSelectSystemCapabilitiesDisable.IsNull() && !data.TlvSelectSystemCapabilitiesDisable.IsUnknown() {
		if data.TlvSelectSystemCapabilitiesDisable.ValueBool() {
			body, _ = sjson.Set(body, "tlv-select.system-capabilities.disable", map[string]string{})
		}
	}
	if !data.TlvSelectSystemDescriptionDisable.IsNull() && !data.TlvSelectSystemDescriptionDisable.IsUnknown() {
		if data.TlvSelectSystemDescriptionDisable.ValueBool() {
			body, _ = sjson.Set(body, "tlv-select.system-description.disable", map[string]string{})
		}
	}
	if !data.TlvSelectSystemNameDisable.IsNull() && !data.TlvSelectSystemNameDisable.IsUnknown() {
		if data.TlvSelectSystemNameDisable.ValueBool() {
			body, _ = sjson.Set(body, "tlv-select.system-name.disable", map[string]string{})
		}
	}
	return body
}

func (data *LLDP) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "holdtime"); value.Exists() && !data.Holdtime.IsNull() {
		data.Holdtime = types.Int64Value(value.Int())
	} else {
		data.Holdtime = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "timer"); value.Exists() && !data.Timer.IsNull() {
		data.Timer = types.Int64Value(value.Int())
	} else {
		data.Timer = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "reinit"); value.Exists() && !data.Reinit.IsNull() {
		data.Reinit = types.Int64Value(value.Int())
	} else {
		data.Reinit = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "subinterfaces.enable"); !data.SubinterfacesEnable.IsNull() {
		if value.Exists() {
			data.SubinterfacesEnable = types.BoolValue(true)
		} else {
			data.SubinterfacesEnable = types.BoolValue(false)
		}
	} else {
		data.SubinterfacesEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "management.enable"); !data.ManagementEnable.IsNull() {
		if value.Exists() {
			data.ManagementEnable = types.BoolValue(true)
		} else {
			data.ManagementEnable = types.BoolValue(false)
		}
	} else {
		data.ManagementEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "priorityaddr.enable"); !data.PriorityaddrEnable.IsNull() {
		if value.Exists() {
			data.PriorityaddrEnable = types.BoolValue(true)
		} else {
			data.PriorityaddrEnable = types.BoolValue(false)
		}
	} else {
		data.PriorityaddrEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "extended-show-width.enable"); !data.ExtendedShowWidthEnable.IsNull() {
		if value.Exists() {
			data.ExtendedShowWidthEnable = types.BoolValue(true)
		} else {
			data.ExtendedShowWidthEnable = types.BoolValue(false)
		}
	} else {
		data.ExtendedShowWidthEnable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "tlv-select.management-address.disable"); !data.TlvSelectManagementAddressDisable.IsNull() {
		if value.Exists() {
			data.TlvSelectManagementAddressDisable = types.BoolValue(true)
		} else {
			data.TlvSelectManagementAddressDisable = types.BoolValue(false)
		}
	} else {
		data.TlvSelectManagementAddressDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "tlv-select.port-description.disable"); !data.TlvSelectPortDescriptionDisable.IsNull() {
		if value.Exists() {
			data.TlvSelectPortDescriptionDisable = types.BoolValue(true)
		} else {
			data.TlvSelectPortDescriptionDisable = types.BoolValue(false)
		}
	} else {
		data.TlvSelectPortDescriptionDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "tlv-select.system-capabilities.disable"); !data.TlvSelectSystemCapabilitiesDisable.IsNull() {
		if value.Exists() {
			data.TlvSelectSystemCapabilitiesDisable = types.BoolValue(true)
		} else {
			data.TlvSelectSystemCapabilitiesDisable = types.BoolValue(false)
		}
	} else {
		data.TlvSelectSystemCapabilitiesDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "tlv-select.system-description.disable"); !data.TlvSelectSystemDescriptionDisable.IsNull() {
		if value.Exists() {
			data.TlvSelectSystemDescriptionDisable = types.BoolValue(true)
		} else {
			data.TlvSelectSystemDescriptionDisable = types.BoolValue(false)
		}
	} else {
		data.TlvSelectSystemDescriptionDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "tlv-select.system-name.disable"); !data.TlvSelectSystemNameDisable.IsNull() {
		if value.Exists() {
			data.TlvSelectSystemNameDisable = types.BoolValue(true)
		} else {
			data.TlvSelectSystemNameDisable = types.BoolValue(false)
		}
	} else {
		data.TlvSelectSystemNameDisable = types.BoolNull()
	}
}

func (data *LLDPData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "holdtime"); value.Exists() {
		data.Holdtime = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "timer"); value.Exists() {
		data.Timer = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "reinit"); value.Exists() {
		data.Reinit = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "subinterfaces.enable"); value.Exists() {
		data.SubinterfacesEnable = types.BoolValue(true)
	} else {
		data.SubinterfacesEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "management.enable"); value.Exists() {
		data.ManagementEnable = types.BoolValue(true)
	} else {
		data.ManagementEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "priorityaddr.enable"); value.Exists() {
		data.PriorityaddrEnable = types.BoolValue(true)
	} else {
		data.PriorityaddrEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "extended-show-width.enable"); value.Exists() {
		data.ExtendedShowWidthEnable = types.BoolValue(true)
	} else {
		data.ExtendedShowWidthEnable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "tlv-select.management-address.disable"); value.Exists() {
		data.TlvSelectManagementAddressDisable = types.BoolValue(true)
	} else {
		data.TlvSelectManagementAddressDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "tlv-select.port-description.disable"); value.Exists() {
		data.TlvSelectPortDescriptionDisable = types.BoolValue(true)
	} else {
		data.TlvSelectPortDescriptionDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "tlv-select.system-capabilities.disable"); value.Exists() {
		data.TlvSelectSystemCapabilitiesDisable = types.BoolValue(true)
	} else {
		data.TlvSelectSystemCapabilitiesDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "tlv-select.system-description.disable"); value.Exists() {
		data.TlvSelectSystemDescriptionDisable = types.BoolValue(true)
	} else {
		data.TlvSelectSystemDescriptionDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "tlv-select.system-name.disable"); value.Exists() {
		data.TlvSelectSystemNameDisable = types.BoolValue(true)
	} else {
		data.TlvSelectSystemNameDisable = types.BoolValue(false)
	}
}

func (data *LLDP) getDeletedListItems(ctx context.Context, state LLDP) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *LLDP) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.SubinterfacesEnable.IsNull() && !data.SubinterfacesEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/subinterfaces/enable", data.getPath()))
	}
	if !data.ManagementEnable.IsNull() && !data.ManagementEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/management/enable", data.getPath()))
	}
	if !data.PriorityaddrEnable.IsNull() && !data.PriorityaddrEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/priorityaddr/enable", data.getPath()))
	}
	if !data.ExtendedShowWidthEnable.IsNull() && !data.ExtendedShowWidthEnable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/extended-show-width/enable", data.getPath()))
	}
	if !data.TlvSelectManagementAddressDisable.IsNull() && !data.TlvSelectManagementAddressDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tlv-select/management-address/disable", data.getPath()))
	}
	if !data.TlvSelectPortDescriptionDisable.IsNull() && !data.TlvSelectPortDescriptionDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tlv-select/port-description/disable", data.getPath()))
	}
	if !data.TlvSelectSystemCapabilitiesDisable.IsNull() && !data.TlvSelectSystemCapabilitiesDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tlv-select/system-capabilities/disable", data.getPath()))
	}
	if !data.TlvSelectSystemDescriptionDisable.IsNull() && !data.TlvSelectSystemDescriptionDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tlv-select/system-description/disable", data.getPath()))
	}
	if !data.TlvSelectSystemNameDisable.IsNull() && !data.TlvSelectSystemNameDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/tlv-select/system-name/disable", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *LLDP) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Holdtime.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/holdtime", data.getPath()))
	}
	if !data.Timer.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/timer", data.getPath()))
	}
	if !data.Reinit.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/reinit", data.getPath()))
	}
	if !data.SubinterfacesEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/subinterfaces/enable", data.getPath()))
	}
	if !data.ManagementEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/management/enable", data.getPath()))
	}
	if !data.PriorityaddrEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/priorityaddr/enable", data.getPath()))
	}
	if !data.ExtendedShowWidthEnable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/extended-show-width/enable", data.getPath()))
	}
	if !data.TlvSelectManagementAddressDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tlv-select/management-address/disable", data.getPath()))
	}
	if !data.TlvSelectPortDescriptionDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tlv-select/port-description/disable", data.getPath()))
	}
	if !data.TlvSelectSystemCapabilitiesDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tlv-select/system-capabilities/disable", data.getPath()))
	}
	if !data.TlvSelectSystemDescriptionDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tlv-select/system-description/disable", data.getPath()))
	}
	if !data.TlvSelectSystemNameDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/tlv-select/system-name/disable", data.getPath()))
	}
	return deletePaths
}
