// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tidwall/sjson"

	"github.com/tidwall/gjson"
)

type RouterISISInterface struct {
	Device                types.String `tfsdk:"device"`
	Id                    types.String `tfsdk:"id"`
	ProcessId             types.String `tfsdk:"process_id"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	CircuitType           types.String `tfsdk:"circuit_type"`
	HelloPaddingDisable   types.Bool   `tfsdk:"hello_padding_disable"`
	HelloPaddingSometimes types.Bool   `tfsdk:"hello_padding_sometimes"`
	Priority              types.Int64  `tfsdk:"priority"`
	PointToPoint          types.Bool   `tfsdk:"point_to_point"`
	Passive               types.Bool   `tfsdk:"passive"`
	Suppressed            types.Bool   `tfsdk:"suppressed"`
	Shutdown              types.Bool   `tfsdk:"shutdown"`
}

func (data RouterISISInterface) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=%s]/interfaces/interface[interface-name=%s]", data.ProcessId.Value, data.InterfaceName.Value)
}

func (data RouterISISInterface) toBody() string {
	body := "{}"

	if !data.CircuitType.Null && !data.CircuitType.Unknown {
		body, _ = sjson.Set(body, "circuit-type", data.CircuitType.Value)
	}
	if !data.HelloPaddingDisable.Null && !data.HelloPaddingDisable.Unknown {
		if data.HelloPaddingDisable.Value {
			body, _ = sjson.Set(body, "hello-padding.disable", map[string]string{})
		}
	}
	if !data.HelloPaddingSometimes.Null && !data.HelloPaddingSometimes.Unknown {
		if data.HelloPaddingSometimes.Value {
			body, _ = sjson.Set(body, "hello-padding.sometimes", map[string]string{})
		}
	}
	if !data.Priority.Null && !data.Priority.Unknown {
		body, _ = sjson.Set(body, "priority.priority-value", strconv.FormatInt(data.Priority.Value, 10))
	}
	if !data.PointToPoint.Null && !data.PointToPoint.Unknown {
		if data.PointToPoint.Value {
			body, _ = sjson.Set(body, "point-to-point", map[string]string{})
		}
	}
	if !data.Passive.Null && !data.Passive.Unknown {
		if data.Passive.Value {
			body, _ = sjson.Set(body, "passive", map[string]string{})
		}
	}
	if !data.Suppressed.Null && !data.Suppressed.Unknown {
		if data.Suppressed.Value {
			body, _ = sjson.Set(body, "suppressed", map[string]string{})
		}
	}
	if !data.Shutdown.Null && !data.Shutdown.Unknown {
		if data.Shutdown.Value {
			body, _ = sjson.Set(body, "shutdown", map[string]string{})
		}
	}

	return body
}

func (data *RouterISISInterface) fromBody(res []byte) {
	if value := gjson.GetBytes(res, "circuit-type"); value.Exists() {
		data.CircuitType.Value = value.String()
	}
	if value := gjson.GetBytes(res, "hello-padding.disable"); value.Exists() {
		data.HelloPaddingDisable.Value = true
	}
	if value := gjson.GetBytes(res, "hello-padding.sometimes"); value.Exists() {
		data.HelloPaddingSometimes.Value = true
	}
	if value := gjson.GetBytes(res, "priority.priority-value"); value.Exists() {
		data.Priority.Value = value.Int()
	}
	if value := gjson.GetBytes(res, "point-to-point"); value.Exists() {
		data.PointToPoint.Value = true
	}
	if value := gjson.GetBytes(res, "passive"); value.Exists() {
		data.Passive.Value = true
	}
	if value := gjson.GetBytes(res, "suppressed"); value.Exists() {
		data.Suppressed.Value = true
	}
	if value := gjson.GetBytes(res, "shutdown"); value.Exists() {
		data.Shutdown.Value = true
	}
}

func (data *RouterISISInterface) fromPlan(plan RouterISISInterface) {
	data.Device = plan.Device
	data.ProcessId.Value = plan.ProcessId.Value
	data.InterfaceName.Value = plan.InterfaceName.Value
}
