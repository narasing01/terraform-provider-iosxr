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

type L2VPNXconnectGroupP2P struct {
	Device                                types.String                                                 `tfsdk:"device"`
	Id                                    types.String                                                 `tfsdk:"id"`
	DeleteMode                            types.String                                                 `tfsdk:"delete_mode"`
	GroupName                             types.String                                                 `tfsdk:"group_name"`
	P2pXconnectName                       types.String                                                 `tfsdk:"p2p_xconnect_name"`
	Description                           types.String                                                 `tfsdk:"description"`
	Interfaces                            []L2VPNXconnectGroupP2PInterfaces                            `tfsdk:"interfaces"`
	Ipv4Neighbors                         []L2VPNXconnectGroupP2PIpv4Neighbors                         `tfsdk:"ipv4_neighbors"`
	Ipv6Neighbors                         []L2VPNXconnectGroupP2PIpv6Neighbors                         `tfsdk:"ipv6_neighbors"`
	EvpnTargetNeighbors                   []L2VPNXconnectGroupP2PEvpnTargetNeighbors                   `tfsdk:"evpn_target_neighbors"`
	EvpnServiceNeighbors                  []L2VPNXconnectGroupP2PEvpnServiceNeighbors                  `tfsdk:"evpn_service_neighbors"`
	NeighborEvpnEviSegmentRoutingServices []L2VPNXconnectGroupP2PNeighborEvpnEviSegmentRoutingServices `tfsdk:"neighbor_evpn_evi_segment_routing_services"`
}

type L2VPNXconnectGroupP2PData struct {
	Device                                types.String                                                 `tfsdk:"device"`
	Id                                    types.String                                                 `tfsdk:"id"`
	GroupName                             types.String                                                 `tfsdk:"group_name"`
	P2pXconnectName                       types.String                                                 `tfsdk:"p2p_xconnect_name"`
	Description                           types.String                                                 `tfsdk:"description"`
	Interfaces                            []L2VPNXconnectGroupP2PInterfaces                            `tfsdk:"interfaces"`
	Ipv4Neighbors                         []L2VPNXconnectGroupP2PIpv4Neighbors                         `tfsdk:"ipv4_neighbors"`
	Ipv6Neighbors                         []L2VPNXconnectGroupP2PIpv6Neighbors                         `tfsdk:"ipv6_neighbors"`
	EvpnTargetNeighbors                   []L2VPNXconnectGroupP2PEvpnTargetNeighbors                   `tfsdk:"evpn_target_neighbors"`
	EvpnServiceNeighbors                  []L2VPNXconnectGroupP2PEvpnServiceNeighbors                  `tfsdk:"evpn_service_neighbors"`
	NeighborEvpnEviSegmentRoutingServices []L2VPNXconnectGroupP2PNeighborEvpnEviSegmentRoutingServices `tfsdk:"neighbor_evpn_evi_segment_routing_services"`
}
type L2VPNXconnectGroupP2PInterfaces struct {
	InterfaceName types.String `tfsdk:"interface_name"`
}
type L2VPNXconnectGroupP2PIpv4Neighbors struct {
	Address types.String `tfsdk:"address"`
	PwId    types.Int64  `tfsdk:"pw_id"`
	PwClass types.String `tfsdk:"pw_class"`
}
type L2VPNXconnectGroupP2PIpv6Neighbors struct {
	Address types.String `tfsdk:"address"`
	PwId    types.Int64  `tfsdk:"pw_id"`
	PwClass types.String `tfsdk:"pw_class"`
}
type L2VPNXconnectGroupP2PEvpnTargetNeighbors struct {
	VpnId      types.Int64  `tfsdk:"vpn_id"`
	RemoteAcId types.Int64  `tfsdk:"remote_ac_id"`
	Source     types.Int64  `tfsdk:"source"`
	PwClass    types.String `tfsdk:"pw_class"`
}
type L2VPNXconnectGroupP2PEvpnServiceNeighbors struct {
	VpnId     types.Int64  `tfsdk:"vpn_id"`
	ServiceId types.Int64  `tfsdk:"service_id"`
	PwClass   types.String `tfsdk:"pw_class"`
}
type L2VPNXconnectGroupP2PNeighborEvpnEviSegmentRoutingServices struct {
	VpnId                     types.Int64  `tfsdk:"vpn_id"`
	ServiceId                 types.Int64  `tfsdk:"service_id"`
	SegmentRoutingSrv6Locator types.String `tfsdk:"segment_routing_srv6_locator"`
}

func (data L2VPNXconnectGroupP2P) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]/p2ps/p2p[p2p-xconnect-name=%s]", data.GroupName.ValueString(), data.P2pXconnectName.ValueString())
}

func (data L2VPNXconnectGroupP2PData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/xconnect/groups/group[group-name=%s]/p2ps/p2p[p2p-xconnect-name=%s]", data.GroupName.ValueString(), data.P2pXconnectName.ValueString())
}

func (data L2VPNXconnectGroupP2P) toBody(ctx context.Context) string {
	body := "{}"
	if !data.P2pXconnectName.IsNull() && !data.P2pXconnectName.IsUnknown() {
		body, _ = sjson.Set(body, "p2p-xconnect-name", data.P2pXconnectName.ValueString())
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces.interface", []interface{}{})
		for index, item := range data.Interfaces {
			if !item.InterfaceName.IsNull() && !item.InterfaceName.IsUnknown() {
				body, _ = sjson.Set(body, "interfaces.interface"+"."+strconv.Itoa(index)+"."+"interface-name", item.InterfaceName.ValueString())
			}
		}
	}
	if len(data.Ipv4Neighbors) > 0 {
		body, _ = sjson.Set(body, "neighbor.ipv4s.ipv4", []interface{}{})
		for index, item := range data.Ipv4Neighbors {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.ipv4s.ipv4"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.PwId.IsNull() && !item.PwId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.ipv4s.ipv4"+"."+strconv.Itoa(index)+"."+"pw-id", strconv.FormatInt(item.PwId.ValueInt64(), 10))
			}
			if !item.PwClass.IsNull() && !item.PwClass.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.ipv4s.ipv4"+"."+strconv.Itoa(index)+"."+"pw-class", item.PwClass.ValueString())
			}
		}
	}
	if len(data.Ipv6Neighbors) > 0 {
		body, _ = sjson.Set(body, "neighbor.ipv6s.ipv6", []interface{}{})
		for index, item := range data.Ipv6Neighbors {
			if !item.Address.IsNull() && !item.Address.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.ipv6s.ipv6"+"."+strconv.Itoa(index)+"."+"address", item.Address.ValueString())
			}
			if !item.PwId.IsNull() && !item.PwId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.ipv6s.ipv6"+"."+strconv.Itoa(index)+"."+"pw-id", strconv.FormatInt(item.PwId.ValueInt64(), 10))
			}
			if !item.PwClass.IsNull() && !item.PwClass.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.ipv6s.ipv6"+"."+strconv.Itoa(index)+"."+"pw-class", item.PwClass.ValueString())
			}
		}
	}
	if len(data.EvpnTargetNeighbors) > 0 {
		body, _ = sjson.Set(body, "neighbor.evpn.evi.targets.target", []interface{}{})
		for index, item := range data.EvpnTargetNeighbors {
			if !item.VpnId.IsNull() && !item.VpnId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.targets.target"+"."+strconv.Itoa(index)+"."+"vpn-id", strconv.FormatInt(item.VpnId.ValueInt64(), 10))
			}
			if !item.RemoteAcId.IsNull() && !item.RemoteAcId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.targets.target"+"."+strconv.Itoa(index)+"."+"remote-ac-id", strconv.FormatInt(item.RemoteAcId.ValueInt64(), 10))
			}
			if !item.Source.IsNull() && !item.Source.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.targets.target"+"."+strconv.Itoa(index)+"."+"source", strconv.FormatInt(item.Source.ValueInt64(), 10))
			}
			if !item.PwClass.IsNull() && !item.PwClass.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.targets.target"+"."+strconv.Itoa(index)+"."+"pw-class", item.PwClass.ValueString())
			}
		}
	}
	if len(data.EvpnServiceNeighbors) > 0 {
		body, _ = sjson.Set(body, "neighbor.evpn.evi.services.service", []interface{}{})
		for index, item := range data.EvpnServiceNeighbors {
			if !item.VpnId.IsNull() && !item.VpnId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.services.service"+"."+strconv.Itoa(index)+"."+"vpn-id", strconv.FormatInt(item.VpnId.ValueInt64(), 10))
			}
			if !item.ServiceId.IsNull() && !item.ServiceId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.services.service"+"."+strconv.Itoa(index)+"."+"service-id", strconv.FormatInt(item.ServiceId.ValueInt64(), 10))
			}
			if !item.PwClass.IsNull() && !item.PwClass.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.services.service"+"."+strconv.Itoa(index)+"."+"pw-class", item.PwClass.ValueString())
			}
		}
	}
	if len(data.NeighborEvpnEviSegmentRoutingServices) > 0 {
		body, _ = sjson.Set(body, "neighbor.evpn.evi.segment-routing-services.service", []interface{}{})
		for index, item := range data.NeighborEvpnEviSegmentRoutingServices {
			if !item.VpnId.IsNull() && !item.VpnId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.segment-routing-services.service"+"."+strconv.Itoa(index)+"."+"vpn-id", strconv.FormatInt(item.VpnId.ValueInt64(), 10))
			}
			if !item.ServiceId.IsNull() && !item.ServiceId.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.segment-routing-services.service"+"."+strconv.Itoa(index)+"."+"service-id", strconv.FormatInt(item.ServiceId.ValueInt64(), 10))
			}
			if !item.SegmentRoutingSrv6Locator.IsNull() && !item.SegmentRoutingSrv6Locator.IsUnknown() {
				body, _ = sjson.Set(body, "neighbor.evpn.evi.segment-routing-services.service"+"."+strconv.Itoa(index)+"."+"segment-routing.srv6.locator", item.SegmentRoutingSrv6Locator.ValueString())
			}
		}
	}
	return body
}

func (data *L2VPNXconnectGroupP2P) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.Interfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString()}

		var r gjson.Result
		gjson.GetBytes(res, "interfaces.interface").ForEach(
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
		if value := r.Get("interface-name"); value.Exists() && !data.Interfaces[i].InterfaceName.IsNull() {
			data.Interfaces[i].InterfaceName = types.StringValue(value.String())
		} else {
			data.Interfaces[i].InterfaceName = types.StringNull()
		}
	}
	for i := range data.Ipv4Neighbors {
		keys := [...]string{"address", "pw-id"}
		keyValues := [...]string{data.Ipv4Neighbors[i].Address.ValueString(), strconv.FormatInt(data.Ipv4Neighbors[i].PwId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "neighbor.ipv4s.ipv4").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.Ipv4Neighbors[i].Address.IsNull() {
			data.Ipv4Neighbors[i].Address = types.StringValue(value.String())
		} else {
			data.Ipv4Neighbors[i].Address = types.StringNull()
		}
		if value := r.Get("pw-id"); value.Exists() && !data.Ipv4Neighbors[i].PwId.IsNull() {
			data.Ipv4Neighbors[i].PwId = types.Int64Value(value.Int())
		} else {
			data.Ipv4Neighbors[i].PwId = types.Int64Null()
		}
		if value := r.Get("pw-class"); value.Exists() && !data.Ipv4Neighbors[i].PwClass.IsNull() {
			data.Ipv4Neighbors[i].PwClass = types.StringValue(value.String())
		} else {
			data.Ipv4Neighbors[i].PwClass = types.StringNull()
		}
	}
	for i := range data.Ipv6Neighbors {
		keys := [...]string{"address", "pw-id"}
		keyValues := [...]string{data.Ipv6Neighbors[i].Address.ValueString(), strconv.FormatInt(data.Ipv6Neighbors[i].PwId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "neighbor.ipv6s.ipv6").ForEach(
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
		if value := r.Get("address"); value.Exists() && !data.Ipv6Neighbors[i].Address.IsNull() {
			data.Ipv6Neighbors[i].Address = types.StringValue(value.String())
		} else {
			data.Ipv6Neighbors[i].Address = types.StringNull()
		}
		if value := r.Get("pw-id"); value.Exists() && !data.Ipv6Neighbors[i].PwId.IsNull() {
			data.Ipv6Neighbors[i].PwId = types.Int64Value(value.Int())
		} else {
			data.Ipv6Neighbors[i].PwId = types.Int64Null()
		}
		if value := r.Get("pw-class"); value.Exists() && !data.Ipv6Neighbors[i].PwClass.IsNull() {
			data.Ipv6Neighbors[i].PwClass = types.StringValue(value.String())
		} else {
			data.Ipv6Neighbors[i].PwClass = types.StringNull()
		}
	}
	for i := range data.EvpnTargetNeighbors {
		keys := [...]string{"vpn-id", "remote-ac-id", "source"}
		keyValues := [...]string{strconv.FormatInt(data.EvpnTargetNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.EvpnTargetNeighbors[i].RemoteAcId.ValueInt64(), 10), strconv.FormatInt(data.EvpnTargetNeighbors[i].Source.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "neighbor.evpn.evi.targets.target").ForEach(
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
		if value := r.Get("vpn-id"); value.Exists() && !data.EvpnTargetNeighbors[i].VpnId.IsNull() {
			data.EvpnTargetNeighbors[i].VpnId = types.Int64Value(value.Int())
		} else {
			data.EvpnTargetNeighbors[i].VpnId = types.Int64Null()
		}
		if value := r.Get("remote-ac-id"); value.Exists() && !data.EvpnTargetNeighbors[i].RemoteAcId.IsNull() {
			data.EvpnTargetNeighbors[i].RemoteAcId = types.Int64Value(value.Int())
		} else {
			data.EvpnTargetNeighbors[i].RemoteAcId = types.Int64Null()
		}
		if value := r.Get("source"); value.Exists() && !data.EvpnTargetNeighbors[i].Source.IsNull() {
			data.EvpnTargetNeighbors[i].Source = types.Int64Value(value.Int())
		} else {
			data.EvpnTargetNeighbors[i].Source = types.Int64Null()
		}
		if value := r.Get("pw-class"); value.Exists() && !data.EvpnTargetNeighbors[i].PwClass.IsNull() {
			data.EvpnTargetNeighbors[i].PwClass = types.StringValue(value.String())
		} else {
			data.EvpnTargetNeighbors[i].PwClass = types.StringNull()
		}
	}
	for i := range data.EvpnServiceNeighbors {
		keys := [...]string{"vpn-id", "service-id"}
		keyValues := [...]string{strconv.FormatInt(data.EvpnServiceNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.EvpnServiceNeighbors[i].ServiceId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "neighbor.evpn.evi.services.service").ForEach(
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
		if value := r.Get("vpn-id"); value.Exists() && !data.EvpnServiceNeighbors[i].VpnId.IsNull() {
			data.EvpnServiceNeighbors[i].VpnId = types.Int64Value(value.Int())
		} else {
			data.EvpnServiceNeighbors[i].VpnId = types.Int64Null()
		}
		if value := r.Get("service-id"); value.Exists() && !data.EvpnServiceNeighbors[i].ServiceId.IsNull() {
			data.EvpnServiceNeighbors[i].ServiceId = types.Int64Value(value.Int())
		} else {
			data.EvpnServiceNeighbors[i].ServiceId = types.Int64Null()
		}
		if value := r.Get("pw-class"); value.Exists() && !data.EvpnServiceNeighbors[i].PwClass.IsNull() {
			data.EvpnServiceNeighbors[i].PwClass = types.StringValue(value.String())
		} else {
			data.EvpnServiceNeighbors[i].PwClass = types.StringNull()
		}
	}
	for i := range data.NeighborEvpnEviSegmentRoutingServices {
		keys := [...]string{"vpn-id", "service-id"}
		keyValues := [...]string{strconv.FormatInt(data.NeighborEvpnEviSegmentRoutingServices[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.NeighborEvpnEviSegmentRoutingServices[i].ServiceId.ValueInt64(), 10)}

		var r gjson.Result
		gjson.GetBytes(res, "neighbor.evpn.evi.segment-routing-services.service").ForEach(
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
		if value := r.Get("vpn-id"); value.Exists() && !data.NeighborEvpnEviSegmentRoutingServices[i].VpnId.IsNull() {
			data.NeighborEvpnEviSegmentRoutingServices[i].VpnId = types.Int64Value(value.Int())
		} else {
			data.NeighborEvpnEviSegmentRoutingServices[i].VpnId = types.Int64Null()
		}
		if value := r.Get("service-id"); value.Exists() && !data.NeighborEvpnEviSegmentRoutingServices[i].ServiceId.IsNull() {
			data.NeighborEvpnEviSegmentRoutingServices[i].ServiceId = types.Int64Value(value.Int())
		} else {
			data.NeighborEvpnEviSegmentRoutingServices[i].ServiceId = types.Int64Null()
		}
		if value := r.Get("segment-routing.srv6.locator"); value.Exists() && !data.NeighborEvpnEviSegmentRoutingServices[i].SegmentRoutingSrv6Locator.IsNull() {
			data.NeighborEvpnEviSegmentRoutingServices[i].SegmentRoutingSrv6Locator = types.StringValue(value.String())
		} else {
			data.NeighborEvpnEviSegmentRoutingServices[i].SegmentRoutingSrv6Locator = types.StringNull()
		}
	}
}

func (data *L2VPNXconnectGroupP2PData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	}
	if value := gjson.GetBytes(res, "interfaces.interface"); value.Exists() {
		data.Interfaces = make([]L2VPNXconnectGroupP2PInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroupP2PInterfaces{}
			if cValue := v.Get("interface-name"); cValue.Exists() {
				item.InterfaceName = types.StringValue(cValue.String())
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "neighbor.ipv4s.ipv4"); value.Exists() {
		data.Ipv4Neighbors = make([]L2VPNXconnectGroupP2PIpv4Neighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroupP2PIpv4Neighbors{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pw-id"); cValue.Exists() {
				item.PwId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("pw-class"); cValue.Exists() {
				item.PwClass = types.StringValue(cValue.String())
			}
			data.Ipv4Neighbors = append(data.Ipv4Neighbors, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "neighbor.ipv6s.ipv6"); value.Exists() {
		data.Ipv6Neighbors = make([]L2VPNXconnectGroupP2PIpv6Neighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroupP2PIpv6Neighbors{}
			if cValue := v.Get("address"); cValue.Exists() {
				item.Address = types.StringValue(cValue.String())
			}
			if cValue := v.Get("pw-id"); cValue.Exists() {
				item.PwId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("pw-class"); cValue.Exists() {
				item.PwClass = types.StringValue(cValue.String())
			}
			data.Ipv6Neighbors = append(data.Ipv6Neighbors, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "neighbor.evpn.evi.targets.target"); value.Exists() {
		data.EvpnTargetNeighbors = make([]L2VPNXconnectGroupP2PEvpnTargetNeighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroupP2PEvpnTargetNeighbors{}
			if cValue := v.Get("vpn-id"); cValue.Exists() {
				item.VpnId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("remote-ac-id"); cValue.Exists() {
				item.RemoteAcId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("source"); cValue.Exists() {
				item.Source = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("pw-class"); cValue.Exists() {
				item.PwClass = types.StringValue(cValue.String())
			}
			data.EvpnTargetNeighbors = append(data.EvpnTargetNeighbors, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "neighbor.evpn.evi.services.service"); value.Exists() {
		data.EvpnServiceNeighbors = make([]L2VPNXconnectGroupP2PEvpnServiceNeighbors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroupP2PEvpnServiceNeighbors{}
			if cValue := v.Get("vpn-id"); cValue.Exists() {
				item.VpnId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("service-id"); cValue.Exists() {
				item.ServiceId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("pw-class"); cValue.Exists() {
				item.PwClass = types.StringValue(cValue.String())
			}
			data.EvpnServiceNeighbors = append(data.EvpnServiceNeighbors, item)
			return true
		})
	}
	if value := gjson.GetBytes(res, "neighbor.evpn.evi.segment-routing-services.service"); value.Exists() {
		data.NeighborEvpnEviSegmentRoutingServices = make([]L2VPNXconnectGroupP2PNeighborEvpnEviSegmentRoutingServices, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := L2VPNXconnectGroupP2PNeighborEvpnEviSegmentRoutingServices{}
			if cValue := v.Get("vpn-id"); cValue.Exists() {
				item.VpnId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("service-id"); cValue.Exists() {
				item.ServiceId = types.Int64Value(cValue.Int())
			}
			if cValue := v.Get("segment-routing.srv6.locator"); cValue.Exists() {
				item.SegmentRoutingSrv6Locator = types.StringValue(cValue.String())
			}
			data.NeighborEvpnEviSegmentRoutingServices = append(data.NeighborEvpnEviSegmentRoutingServices, item)
			return true
		})
	}
}

func (data *L2VPNXconnectGroupP2P) getDeletedListItems(ctx context.Context, state L2VPNXconnectGroupP2P) []string {
	deletedListItems := make([]string, 0)
	for i := range state.Interfaces {
		keys := [...]string{"interface-name"}
		stateKeyValues := [...]string{state.Interfaces[i].InterfaceName.ValueString()}

		emptyKeys := true
		if !reflect.ValueOf(state.Interfaces[i].InterfaceName.ValueString()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Interfaces {
			found = true
			if state.Interfaces[i].InterfaceName.ValueString() != data.Interfaces[j].InterfaceName.ValueString() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/interfaces/interface%v", state.getPath(), keyString))
		}
	}
	for i := range state.Ipv4Neighbors {
		keys := [...]string{"address", "pw-id"}
		stateKeyValues := [...]string{state.Ipv4Neighbors[i].Address.ValueString(), strconv.FormatInt(state.Ipv4Neighbors[i].PwId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv4Neighbors[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Ipv4Neighbors[i].PwId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv4Neighbors {
			found = true
			if state.Ipv4Neighbors[i].Address.ValueString() != data.Ipv4Neighbors[j].Address.ValueString() {
				found = false
			}
			if state.Ipv4Neighbors[i].PwId.ValueInt64() != data.Ipv4Neighbors[j].PwId.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbor/ipv4s/ipv4%v", state.getPath(), keyString))
		}
	}
	for i := range state.Ipv6Neighbors {
		keys := [...]string{"address", "pw-id"}
		stateKeyValues := [...]string{state.Ipv6Neighbors[i].Address.ValueString(), strconv.FormatInt(state.Ipv6Neighbors[i].PwId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.Ipv6Neighbors[i].Address.ValueString()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.Ipv6Neighbors[i].PwId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.Ipv6Neighbors {
			found = true
			if state.Ipv6Neighbors[i].Address.ValueString() != data.Ipv6Neighbors[j].Address.ValueString() {
				found = false
			}
			if state.Ipv6Neighbors[i].PwId.ValueInt64() != data.Ipv6Neighbors[j].PwId.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbor/ipv6s/ipv6%v", state.getPath(), keyString))
		}
	}
	for i := range state.EvpnTargetNeighbors {
		keys := [...]string{"vpn-id", "remote-ac-id", "source"}
		stateKeyValues := [...]string{strconv.FormatInt(state.EvpnTargetNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(state.EvpnTargetNeighbors[i].RemoteAcId.ValueInt64(), 10), strconv.FormatInt(state.EvpnTargetNeighbors[i].Source.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.EvpnTargetNeighbors[i].VpnId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.EvpnTargetNeighbors[i].RemoteAcId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.EvpnTargetNeighbors[i].Source.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.EvpnTargetNeighbors {
			found = true
			if state.EvpnTargetNeighbors[i].VpnId.ValueInt64() != data.EvpnTargetNeighbors[j].VpnId.ValueInt64() {
				found = false
			}
			if state.EvpnTargetNeighbors[i].RemoteAcId.ValueInt64() != data.EvpnTargetNeighbors[j].RemoteAcId.ValueInt64() {
				found = false
			}
			if state.EvpnTargetNeighbors[i].Source.ValueInt64() != data.EvpnTargetNeighbors[j].Source.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbor/evpn/evi/targets/target%v", state.getPath(), keyString))
		}
	}
	for i := range state.EvpnServiceNeighbors {
		keys := [...]string{"vpn-id", "service-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.EvpnServiceNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(state.EvpnServiceNeighbors[i].ServiceId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.EvpnServiceNeighbors[i].VpnId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.EvpnServiceNeighbors[i].ServiceId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.EvpnServiceNeighbors {
			found = true
			if state.EvpnServiceNeighbors[i].VpnId.ValueInt64() != data.EvpnServiceNeighbors[j].VpnId.ValueInt64() {
				found = false
			}
			if state.EvpnServiceNeighbors[i].ServiceId.ValueInt64() != data.EvpnServiceNeighbors[j].ServiceId.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbor/evpn/evi/services/service%v", state.getPath(), keyString))
		}
	}
	for i := range state.NeighborEvpnEviSegmentRoutingServices {
		keys := [...]string{"vpn-id", "service-id"}
		stateKeyValues := [...]string{strconv.FormatInt(state.NeighborEvpnEviSegmentRoutingServices[i].VpnId.ValueInt64(), 10), strconv.FormatInt(state.NeighborEvpnEviSegmentRoutingServices[i].ServiceId.ValueInt64(), 10)}

		emptyKeys := true
		if !reflect.ValueOf(state.NeighborEvpnEviSegmentRoutingServices[i].VpnId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if !reflect.ValueOf(state.NeighborEvpnEviSegmentRoutingServices[i].ServiceId.ValueInt64()).IsZero() {
			emptyKeys = false
		}
		if emptyKeys {
			continue
		}

		found := false
		for j := range data.NeighborEvpnEviSegmentRoutingServices {
			found = true
			if state.NeighborEvpnEviSegmentRoutingServices[i].VpnId.ValueInt64() != data.NeighborEvpnEviSegmentRoutingServices[j].VpnId.ValueInt64() {
				found = false
			}
			if state.NeighborEvpnEviSegmentRoutingServices[i].ServiceId.ValueInt64() != data.NeighborEvpnEviSegmentRoutingServices[j].ServiceId.ValueInt64() {
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
			deletedListItems = append(deletedListItems, fmt.Sprintf("%v/neighbor/evpn/evi/segment-routing-services/service%v", state.getPath(), keyString))
		}
	}
	return deletedListItems
}

func (data *L2VPNXconnectGroupP2P) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	for i := range data.Interfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString()}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.Ipv4Neighbors {
		keys := [...]string{"address", "pw-id"}
		keyValues := [...]string{data.Ipv4Neighbors[i].Address.ValueString(), strconv.FormatInt(data.Ipv4Neighbors[i].PwId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.Ipv6Neighbors {
		keys := [...]string{"address", "pw-id"}
		keyValues := [...]string{data.Ipv6Neighbors[i].Address.ValueString(), strconv.FormatInt(data.Ipv6Neighbors[i].PwId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.EvpnTargetNeighbors {
		keys := [...]string{"vpn-id", "remote-ac-id", "source"}
		keyValues := [...]string{strconv.FormatInt(data.EvpnTargetNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.EvpnTargetNeighbors[i].RemoteAcId.ValueInt64(), 10), strconv.FormatInt(data.EvpnTargetNeighbors[i].Source.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.EvpnServiceNeighbors {
		keys := [...]string{"vpn-id", "service-id"}
		keyValues := [...]string{strconv.FormatInt(data.EvpnServiceNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.EvpnServiceNeighbors[i].ServiceId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	for i := range data.NeighborEvpnEviSegmentRoutingServices {
		keys := [...]string{"vpn-id", "service-id"}
		keyValues := [...]string{strconv.FormatInt(data.NeighborEvpnEviSegmentRoutingServices[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.NeighborEvpnEviSegmentRoutingServices[i].ServiceId.ValueInt64(), 10)}
		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
	}
	return emptyLeafsDelete
}

func (data *L2VPNXconnectGroupP2P) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Description.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/description", data.getPath()))
	}
	for i := range data.Interfaces {
		keys := [...]string{"interface-name"}
		keyValues := [...]string{data.Interfaces[i].InterfaceName.ValueString()}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/interfaces/interface%v", data.getPath(), keyString))
	}
	for i := range data.Ipv4Neighbors {
		keys := [...]string{"address", "pw-id"}
		keyValues := [...]string{data.Ipv4Neighbors[i].Address.ValueString(), strconv.FormatInt(data.Ipv4Neighbors[i].PwId.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/neighbor/ipv4s/ipv4%v", data.getPath(), keyString))
	}
	for i := range data.Ipv6Neighbors {
		keys := [...]string{"address", "pw-id"}
		keyValues := [...]string{data.Ipv6Neighbors[i].Address.ValueString(), strconv.FormatInt(data.Ipv6Neighbors[i].PwId.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/neighbor/ipv6s/ipv6%v", data.getPath(), keyString))
	}
	for i := range data.EvpnTargetNeighbors {
		keys := [...]string{"vpn-id", "remote-ac-id", "source"}
		keyValues := [...]string{strconv.FormatInt(data.EvpnTargetNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.EvpnTargetNeighbors[i].RemoteAcId.ValueInt64(), 10), strconv.FormatInt(data.EvpnTargetNeighbors[i].Source.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/neighbor/evpn/evi/targets/target%v", data.getPath(), keyString))
	}
	for i := range data.EvpnServiceNeighbors {
		keys := [...]string{"vpn-id", "service-id"}
		keyValues := [...]string{strconv.FormatInt(data.EvpnServiceNeighbors[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.EvpnServiceNeighbors[i].ServiceId.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/neighbor/evpn/evi/services/service%v", data.getPath(), keyString))
	}
	for i := range data.NeighborEvpnEviSegmentRoutingServices {
		keys := [...]string{"vpn-id", "service-id"}
		keyValues := [...]string{strconv.FormatInt(data.NeighborEvpnEviSegmentRoutingServices[i].VpnId.ValueInt64(), 10), strconv.FormatInt(data.NeighborEvpnEviSegmentRoutingServices[i].ServiceId.ValueInt64(), 10)}

		keyString := ""
		for ki := range keys {
			keyString += "[" + keys[ki] + "=" + keyValues[ki] + "]"
		}
		deletePaths = append(deletePaths, fmt.Sprintf("%v/neighbor/evpn/evi/segment-routing-services/service%v", data.getPath(), keyString))
	}
	return deletePaths
}
