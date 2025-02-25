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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type RouterBGPNeighborAddressFamily struct {
	Device                                             types.String `tfsdk:"device"`
	Id                                                 types.String `tfsdk:"id"`
	DeleteMode                                         types.String `tfsdk:"delete_mode"`
	AsNumber                                           types.String `tfsdk:"as_number"`
	NeighborAddress                                    types.String `tfsdk:"neighbor_address"`
	AfName                                             types.String `tfsdk:"af_name"`
	ImportStitchingRtReOriginateStitchingRt            types.Bool   `tfsdk:"import_stitching_rt_re_originate_stitching_rt"`
	RouteReflectorClientInheritanceDisable             types.Bool   `tfsdk:"route_reflector_client_inheritance_disable"`
	AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt types.Bool   `tfsdk:"advertise_vpnv4_unicast_enable_re_originated_stitching_rt"`
	NextHopSelfInheritanceDisable                      types.Bool   `tfsdk:"next_hop_self_inheritance_disable"`
	EncapsulationTypeSrv6                              types.Bool   `tfsdk:"encapsulation_type_srv6"`
}

type RouterBGPNeighborAddressFamilyData struct {
	Device                                             types.String `tfsdk:"device"`
	Id                                                 types.String `tfsdk:"id"`
	AsNumber                                           types.String `tfsdk:"as_number"`
	NeighborAddress                                    types.String `tfsdk:"neighbor_address"`
	AfName                                             types.String `tfsdk:"af_name"`
	ImportStitchingRtReOriginateStitchingRt            types.Bool   `tfsdk:"import_stitching_rt_re_originate_stitching_rt"`
	RouteReflectorClientInheritanceDisable             types.Bool   `tfsdk:"route_reflector_client_inheritance_disable"`
	AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt types.Bool   `tfsdk:"advertise_vpnv4_unicast_enable_re_originated_stitching_rt"`
	NextHopSelfInheritanceDisable                      types.Bool   `tfsdk:"next_hop_self_inheritance_disable"`
	EncapsulationTypeSrv6                              types.Bool   `tfsdk:"encapsulation_type_srv6"`
}

func (data RouterBGPNeighborAddressFamily) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/neighbors/neighbor[neighbor-address=%s]/address-families/address-family[af-name=%s]", data.AsNumber.ValueString(), data.NeighborAddress.ValueString(), data.AfName.ValueString())
}

func (data RouterBGPNeighborAddressFamilyData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=%s]/neighbors/neighbor[neighbor-address=%s]/address-families/address-family[af-name=%s]", data.AsNumber.ValueString(), data.NeighborAddress.ValueString(), data.AfName.ValueString())
}

func (data RouterBGPNeighborAddressFamily) toBody(ctx context.Context) string {
	body := "{}"
	if !data.AfName.IsNull() && !data.AfName.IsUnknown() {
		body, _ = sjson.Set(body, "af-name", data.AfName.ValueString())
	}
	if !data.ImportStitchingRtReOriginateStitchingRt.IsNull() && !data.ImportStitchingRtReOriginateStitchingRt.IsUnknown() {
		if data.ImportStitchingRtReOriginateStitchingRt.ValueBool() {
			body, _ = sjson.Set(body, "import.stitching-rt.re-originate.stitching-rt", map[string]string{})
		}
	}
	if !data.RouteReflectorClientInheritanceDisable.IsNull() && !data.RouteReflectorClientInheritanceDisable.IsUnknown() {
		if data.RouteReflectorClientInheritanceDisable.ValueBool() {
			body, _ = sjson.Set(body, "route-reflector-client.inheritance-disable", map[string]string{})
		}
	}
	if !data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt.IsNull() && !data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt.IsUnknown() {
		if data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt.ValueBool() {
			body, _ = sjson.Set(body, "advertise.vpnv4.unicast.enable.re-originated.stitching-rt", map[string]string{})
		}
	}
	if !data.NextHopSelfInheritanceDisable.IsNull() && !data.NextHopSelfInheritanceDisable.IsUnknown() {
		if data.NextHopSelfInheritanceDisable.ValueBool() {
			body, _ = sjson.Set(body, "next-hop-self.inheritance-disable", map[string]string{})
		}
	}
	if !data.EncapsulationTypeSrv6.IsNull() && !data.EncapsulationTypeSrv6.IsUnknown() {
		if data.EncapsulationTypeSrv6.ValueBool() {
			body, _ = sjson.Set(body, "encapsulation-type.srv6", map[string]string{})
		}
	}
	return body
}

func (data *RouterBGPNeighborAddressFamily) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "import.stitching-rt.re-originate.stitching-rt"); !data.ImportStitchingRtReOriginateStitchingRt.IsNull() {
		if value.Exists() {
			data.ImportStitchingRtReOriginateStitchingRt = types.BoolValue(true)
		} else {
			data.ImportStitchingRtReOriginateStitchingRt = types.BoolValue(false)
		}
	} else {
		data.ImportStitchingRtReOriginateStitchingRt = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "route-reflector-client.inheritance-disable"); !data.RouteReflectorClientInheritanceDisable.IsNull() {
		if value.Exists() {
			data.RouteReflectorClientInheritanceDisable = types.BoolValue(true)
		} else {
			data.RouteReflectorClientInheritanceDisable = types.BoolValue(false)
		}
	} else {
		data.RouteReflectorClientInheritanceDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "advertise.vpnv4.unicast.enable.re-originated.stitching-rt"); !data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt.IsNull() {
		if value.Exists() {
			data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt = types.BoolValue(true)
		} else {
			data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt = types.BoolValue(false)
		}
	} else {
		data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "next-hop-self.inheritance-disable"); !data.NextHopSelfInheritanceDisable.IsNull() {
		if value.Exists() {
			data.NextHopSelfInheritanceDisable = types.BoolValue(true)
		} else {
			data.NextHopSelfInheritanceDisable = types.BoolValue(false)
		}
	} else {
		data.NextHopSelfInheritanceDisable = types.BoolNull()
	}
	if value := gjson.GetBytes(res, "encapsulation-type.srv6"); !data.EncapsulationTypeSrv6.IsNull() {
		if value.Exists() {
			data.EncapsulationTypeSrv6 = types.BoolValue(true)
		} else {
			data.EncapsulationTypeSrv6 = types.BoolValue(false)
		}
	} else {
		data.EncapsulationTypeSrv6 = types.BoolNull()
	}
}

func (data *RouterBGPNeighborAddressFamilyData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "import.stitching-rt.re-originate.stitching-rt"); value.Exists() {
		data.ImportStitchingRtReOriginateStitchingRt = types.BoolValue(true)
	} else {
		data.ImportStitchingRtReOriginateStitchingRt = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "route-reflector-client.inheritance-disable"); value.Exists() {
		data.RouteReflectorClientInheritanceDisable = types.BoolValue(true)
	} else {
		data.RouteReflectorClientInheritanceDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "advertise.vpnv4.unicast.enable.re-originated.stitching-rt"); value.Exists() {
		data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt = types.BoolValue(true)
	} else {
		data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "next-hop-self.inheritance-disable"); value.Exists() {
		data.NextHopSelfInheritanceDisable = types.BoolValue(true)
	} else {
		data.NextHopSelfInheritanceDisable = types.BoolValue(false)
	}
	if value := gjson.GetBytes(res, "encapsulation-type.srv6"); value.Exists() {
		data.EncapsulationTypeSrv6 = types.BoolValue(true)
	} else {
		data.EncapsulationTypeSrv6 = types.BoolValue(false)
	}
}

func (data *RouterBGPNeighborAddressFamily) getDeletedListItems(ctx context.Context, state RouterBGPNeighborAddressFamily) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *RouterBGPNeighborAddressFamily) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	if !data.ImportStitchingRtReOriginateStitchingRt.IsNull() && !data.ImportStitchingRtReOriginateStitchingRt.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/import/stitching-rt/re-originate/stitching-rt", data.getPath()))
	}
	if !data.RouteReflectorClientInheritanceDisable.IsNull() && !data.RouteReflectorClientInheritanceDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/route-reflector-client/inheritance-disable", data.getPath()))
	}
	if !data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt.IsNull() && !data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/advertise/vpnv4/unicast/enable/re-originated/stitching-rt", data.getPath()))
	}
	if !data.NextHopSelfInheritanceDisable.IsNull() && !data.NextHopSelfInheritanceDisable.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/next-hop-self/inheritance-disable", data.getPath()))
	}
	if !data.EncapsulationTypeSrv6.IsNull() && !data.EncapsulationTypeSrv6.ValueBool() {
		emptyLeafsDelete = append(emptyLeafsDelete, fmt.Sprintf("%v/encapsulation-type/srv6", data.getPath()))
	}
	return emptyLeafsDelete
}

func (data *RouterBGPNeighborAddressFamily) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.ImportStitchingRtReOriginateStitchingRt.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/import/stitching-rt/re-originate/stitching-rt", data.getPath()))
	}
	if !data.RouteReflectorClientInheritanceDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/route-reflector-client/inheritance-disable", data.getPath()))
	}
	if !data.AdvertiseVpnv4UnicastEnableReOriginatedStitchingRt.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/advertise/vpnv4/unicast/enable/re-originated/stitching-rt", data.getPath()))
	}
	if !data.NextHopSelfInheritanceDisable.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/next-hop-self/inheritance-disable", data.getPath()))
	}
	if !data.EncapsulationTypeSrv6.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/encapsulation-type/srv6", data.getPath()))
	}
	return deletePaths
}
