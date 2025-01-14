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
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Telnet struct {
	Device                    types.String     `tfsdk:"device"`
	Id                        types.String     `tfsdk:"id"`
	DeleteMode                types.String     `tfsdk:"delete_mode"`
	Ipv4ClientSourceInterface types.String     `tfsdk:"ipv4_client_source_interface"`
	Ipv6ClientSourceInterface types.String     `tfsdk:"ipv6_client_source_interface"`
	Vrfs                      []TelnetVrfs     `tfsdk:"vrfs"`
	VrfsDscp                  []TelnetVrfsDscp `tfsdk:"vrfs_dscp"`
}

type TelnetData struct {
	Device                    types.String     `tfsdk:"device"`
	Id                        types.String     `tfsdk:"id"`
	Ipv4ClientSourceInterface types.String     `tfsdk:"ipv4_client_source_interface"`
	Ipv6ClientSourceInterface types.String     `tfsdk:"ipv6_client_source_interface"`
	Vrfs                      []TelnetVrfs     `tfsdk:"vrfs"`
	VrfsDscp                  []TelnetVrfsDscp `tfsdk:"vrfs_dscp"`
}
type TelnetVrfs struct {
	VrfName              types.String `tfsdk:"vrf_name"`
	Ipv4ServerMaxServers types.Int64  `tfsdk:"ipv4_server_max_servers"`
	Ipv4ServerAccessList types.String `tfsdk:"ipv4_server_access_list"`
	Ipv6ServerMaxServers types.Int64  `tfsdk:"ipv6_server_max_servers"`
	Ipv6ServerAccessList types.String `tfsdk:"ipv6_server_access_list"`
}
type TelnetVrfsDscp struct {
	VrfName  types.String `tfsdk:"vrf_name"`
	Ipv4Dscp types.Int64  `tfsdk:"ipv4_dscp"`
}

func (data Telnet) getPath() string {
	return "Cisco-IOS-XR-um-telnet-cfg:/telnet"
}

func (data TelnetData) getPath() string {
	return "Cisco-IOS-XR-um-telnet-cfg:/telnet"
}

func (data Telnet) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Ipv4ClientSourceInterface.IsNull() && !data.Ipv4ClientSourceInterface.IsUnknown() {
		body, _ = sjson.Set(body, "ipv4.client.source-interface", data.Ipv4ClientSourceInterface.ValueString())
	}
	if !data.Ipv6ClientSourceInterface.IsNull() && !data.Ipv6ClientSourceInterface.IsUnknown() {
		body, _ = sjson.Set(body, "ipv6.client.source-interface", data.Ipv6ClientSourceInterface.ValueString())
	}
	if len(data.Vrfs) > 0 {
		body, _ = sjson.Set(body, "vrfs.vrf", []interface{}{})
		for index, item := range data.Vrfs {
			if !item.VrfName.IsNull() && !item.VrfName.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf"+"."+strconv.Itoa(index)+"."+"vrf-name", item.VrfName.ValueString())
			}
			if !item.Ipv4ServerMaxServers.IsNull() && !item.Ipv4ServerMaxServers.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf"+"."+strconv.Itoa(index)+"."+"ipv4.server.max-servers", strconv.FormatInt(item.Ipv4ServerMaxServers.ValueInt64(), 10))
			}
			if !item.Ipv4ServerAccessList.IsNull() && !item.Ipv4ServerAccessList.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf"+"."+strconv.Itoa(index)+"."+"ipv4.server.access-list", item.Ipv4ServerAccessList.ValueString())
			}
			if !item.Ipv6ServerMaxServers.IsNull() && !item.Ipv6ServerMaxServers.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf"+"."+strconv.Itoa(index)+"."+"ipv6.server.max-servers", strconv.FormatInt(item.Ipv6ServerMaxServers.ValueInt64(), 10))
			}
			if !item.Ipv6ServerAccessList.IsNull() && !item.Ipv6ServerAccessList.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf"+"."+strconv.Itoa(index)+"."+"ipv6.server.access-list", item.Ipv6ServerAccessList.ValueString())
			}
		}
	}
	if len(data.VrfsDscp) > 0 {
		body, _ = sjson.Set(body, "vrfs.vrf-dscp", []interface{}{})
		for index, item := range data.VrfsDscp {
			if !item.VrfName.IsNull() && !item.VrfName.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf-dscp"+"."+strconv.Itoa(index)+"."+"vrf-name", item.VrfName.ValueString())
			}
			if !item.Ipv4Dscp.IsNull() && !item.Ipv4Dscp.IsUnknown() {
				body, _ = sjson.Set(body, "vrfs.vrf-dscp"+"."+strconv.Itoa(index)+"."+"ipv4.dscp", strconv.FormatInt(item.Ipv4Dscp.ValueInt64(), 10))
			}
		}
	}
	return body
}

func (data *Telnet) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "ipv4.client.source-interface"); value.Exists() && !data.Ipv4ClientSourceInterface.IsNull() {
		data.Ipv4ClientSourceInterface = types.StringValue(value.String())
	} else {
		data.Ipv4ClientSourceInterface = types.StringNull()
	}
	if value := gjson.GetBytes(res, "ipv6.client.source-interface"); value.Exists() && !data.Ipv6ClientSourceInterface.IsNull() {
		data.Ipv6ClientSourceInterface = types.StringValue(value.String())
	} else {
		data.Ipv6ClientSourceInterface = types.StringNull()
	}
	for i := range data.Vrfs {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.Vrfs[i].VrfName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "vrfs.vrf").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("vrf-name"); value.Exists() && !data.Vrfs[i].VrfName.IsNull() {
			data.Vrfs[i].VrfName = types.StringValue(value.String())
		} else {
			data.Vrfs[i].VrfName = types.StringNull()
		}
		if value := r.Get("ipv4.server.max-servers"); value.Exists() && !data.Vrfs[i].Ipv4ServerMaxServers.IsNull() {
			data.Vrfs[i].Ipv4ServerMaxServers = types.Int64Value(value.Int())
		} else {
			data.Vrfs[i].Ipv4ServerMaxServers = types.Int64Null()
		}
		if value := r.Get("ipv4.server.access-list"); value.Exists() && !data.Vrfs[i].Ipv4ServerAccessList.IsNull() {
			data.Vrfs[i].Ipv4ServerAccessList = types.StringValue(value.String())
		} else {
			data.Vrfs[i].Ipv4ServerAccessList = types.StringNull()
		}
		if value := r.Get("ipv6.server.max-servers"); value.Exists() && !data.Vrfs[i].Ipv6ServerMaxServers.IsNull() {
			data.Vrfs[i].Ipv6ServerMaxServers = types.Int64Value(value.Int())
		} else {
			data.Vrfs[i].Ipv6ServerMaxServers = types.Int64Null()
		}
		if value := r.Get("ipv6.server.access-list"); value.Exists() && !data.Vrfs[i].Ipv6ServerAccessList.IsNull() {
			data.Vrfs[i].Ipv6ServerAccessList = types.StringValue(value.String())
		} else {
			data.Vrfs[i].Ipv6ServerAccessList = types.StringNull()
		}
	}
	for i := range data.VrfsDscp {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.VrfsDscp[i].VrfName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "vrfs.vrf-dscp").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("vrf-name"); value.Exists() && !data.VrfsDscp[i].VrfName.IsNull() {
			data.VrfsDscp[i].VrfName = types.StringValue(value.String())
		} else {
			data.VrfsDscp[i].VrfName = types.StringNull()
		}
		if value := r.Get("ipv4.dscp"); value.Exists() && !data.VrfsDscp[i].Ipv4Dscp.IsNull() {
			data.VrfsDscp[i].Ipv4Dscp = types.Int64Value(value.Int())
		} else {
			data.VrfsDscp[i].Ipv4Dscp = types.Int64Null()
		}
	}
}

func (data *TelnetData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "ipv4.client.source-interface"); value.Exists() {
		data.Ipv4ClientSourceInterface = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "ipv6.client.source-interface"); value.Exists() {
		data.Ipv6ClientSourceInterface = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "vrfs.vrf"); value.Exists() {
		data.Vrfs = make([]TelnetVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TelnetVrfs{}
			if cValue := v.Get("vrf-name"); cValue.Exists() {
				item.VrfName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ipv4.server.max-servers"); cValue.Exists() {
				item.Ipv4ServerMaxServers = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("ipv4.server.access-list"); cValue.Exists() {
				item.Ipv4ServerAccessList = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ipv6.server.max-servers"); cValue.Exists() {
				item.Ipv6ServerMaxServers = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("ipv6.server.access-list"); cValue.Exists() {
				item.Ipv6ServerAccessList = types.StringValue(cValue.String())
			}
			data.Vrfs = append(data.Vrfs, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "vrfs.vrf-dscp"); value.Exists() {
		data.VrfsDscp = make([]TelnetVrfsDscp, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TelnetVrfsDscp{}
			if cValue := v.Get("vrf-name"); cValue.Exists() {
				item.VrfName = types.StringValue(cValue.String())
			}
			if cValue := v.Get("ipv4.dscp"); cValue.Exists() {
				item.Ipv4Dscp = types.Int64Value(cValue.Int())
			}
			data.VrfsDscp = append(data.VrfsDscp, item)
			return true
		})
	}
}

func (data *Telnet) getDeletedListItems(ctx context.Context, state Telnet) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Vrfs {
		keys := [...]string{"vrf-name"}
		stateKeyValues := [...]string{state.Vrfs[i].VrfName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Vrfs[i].VrfName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Vrfs {
			found = true
			if state.Vrfs[i].VrfName.ValueString() != data.Vrfs[j].VrfName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/vrfs/vrf%v", state.getPath(), keyString))
		}
	}
	for i := range state.VrfsDscp {
		keys := [...]string{"vrf-name"}
		stateKeyValues := [...]string{state.VrfsDscp[i].VrfName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.VrfsDscp[i].VrfName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.VrfsDscp {
			found = true
			if state.VrfsDscp[i].VrfName.ValueString() != data.VrfsDscp[j].VrfName.ValueString() {
				found = false
			}
			if found {
				break
			}
		}
		if !found {
			keyString := ""
			for ki := range keys {
				keyString += "[" + keys[ki] + "=" + stateKeyValues[ki] + "]"
			}
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/vrfs/vrf-dscp%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *Telnet) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.Vrfs {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.Vrfs[i].VrfName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.VrfsDscp {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.VrfsDscp[i].VrfName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}

func (data *Telnet) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Ipv4ClientSourceInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv4/client/source-interface", data.getPath()))
	}
	if !data.Ipv6ClientSourceInterface.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/ipv6/client/source-interface", data.getPath()))
	}
	for i := range data.Vrfs {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.Vrfs[i].VrfName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vrfs/vrf%v", data.getPath(), keyString))
	}
	for i := range data.VrfsDscp {
		keys := [...]string{"vrf-name"}
		keyValues := [...]string{data.VrfsDscp[i].VrfName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/vrfs/vrf-dscp%v", data.getPath(), keyString))
	}
	return deletePaths
}
