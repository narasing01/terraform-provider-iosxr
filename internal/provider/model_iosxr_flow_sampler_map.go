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

type FlowSamplerMap struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Random types.Int64  `tfsdk:"random"`
	OutOf  types.Int64  `tfsdk:"out_of"`
}

type FlowSamplerMapData struct {
	Device types.String `tfsdk:"device"`
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Random types.Int64  `tfsdk:"random"`
	OutOf  types.Int64  `tfsdk:"out_of"`
}

func (data FlowSamplerMap) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-flow-cfg:/sampler-maps/sampler-map[sampler-map-name=%s]", data.Name.ValueString())
}

func (data FlowSamplerMapData) getPath() string {
	return fmt.Sprintf("Cisco-IOS-XR-um-flow-cfg:/sampler-maps/sampler-map[sampler-map-name=%s]", data.Name.ValueString())
}

func (data FlowSamplerMap) toBody(ctx context.Context) string {
	body := "{}"
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "sampler-map-name", data.Name.ValueString())
	}
	if !data.Random.IsNull() && !data.Random.IsUnknown() {
		body, _ = sjson.Set(body, "random", strconv.FormatInt(data.Random.ValueInt64(), 10))
	}
	if !data.OutOf.IsNull() && !data.OutOf.IsUnknown() {
		body, _ = sjson.Set(body, "out-of", strconv.FormatInt(data.OutOf.ValueInt64(), 10))
	}
	return body
}

func (data *FlowSamplerMap) updateFromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "random"); value.Exists() && !data.Random.IsNull() {
		data.Random = types.Int64Value(value.Int())
	} else {
		data.Random = types.Int64Null()
	}
	if value := gjson.GetBytes(res, "out-of"); value.Exists() && !data.OutOf.IsNull() {
		data.OutOf = types.Int64Value(value.Int())
	} else {
		data.OutOf = types.Int64Null()
	}
}

func (data *FlowSamplerMapData) fromBody(ctx context.Context, res []byte) {
	if value := gjson.GetBytes(res, "random"); value.Exists() {
		data.Random = types.Int64Value(value.Int())
	}
	if value := gjson.GetBytes(res, "out-of"); value.Exists() {
		data.OutOf = types.Int64Value(value.Int())
	}
}

func (data *FlowSamplerMap) getDeletedListItems(ctx context.Context, state FlowSamplerMap) []string {
	deletedListItems := make([]string, 0)
	return deletedListItems
}

func (data *FlowSamplerMap) getEmptyLeafsDelete(ctx context.Context) []string {
	emptyLeafsDelete := make([]string, 0)
	return emptyLeafsDelete
}

func (data *FlowSamplerMap) getDeletePaths(ctx context.Context) []string {
	var deletePaths []string
	if !data.Random.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/random", data.getPath()))
	}
	if !data.OutOf.IsNull() {
		deletePaths = append(deletePaths, fmt.Sprintf("%v/out-of", data.getPath()))
	}
	return deletePaths
}
