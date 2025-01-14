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
	"regexp"

	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/client"
	"github.com/CiscoDevNet/terraform-provider-iosxr/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func NewPolicyMapQoSResource() resource.Resource {
	return &PolicyMapQoSResource{}
}

type PolicyMapQoSResource struct {
	client *client.Client
}

func (r *PolicyMapQoSResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_map_qos"
}

func (r *PolicyMapQoSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage the Policy Map QoS configuration.",

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The path of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"policy_map_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the policymap").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set description for this policy-map").String,
				Optional:            true,
			},
			"classes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the class-map").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}`), ""),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of class-map").AddStringEnumDescription("qos", "traffic").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("qos", "traffic"),
							},
						},
						"set_mpls_experimental_topmost": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets the experimental value of the MPLS packet top-most labels.").AddIntegerRangeDescription(0, 7).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 7),
							},
						},
						"set_dscp": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set IP DSCP (DiffServ CodePoint)").String,
							Optional:            true,
						},
						"priority_level": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure a priority level").AddIntegerRangeDescription(1, 7).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 7),
							},
						},
						"queue_limits": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure queue-limit (taildrop threshold) for this class").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("queue-limit value").String,
										Required:            true,
									},
									"unit": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("queue-limit unit").AddStringEnumDescription("bytes", "kbytes", "mbytes", "ms", "packets", "percent", "us").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("bytes", "kbytes", "mbytes", "ms", "packets", "percent", "us"),
										},
									},
								},
							},
						},
						"service_policy_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the child service policy").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}`), ""),
							},
						},
						"police_rate_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Committed Information Rate").String,
							Optional:            true,
						},
						"police_rate_unit": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Rate unit").AddStringEnumDescription("bps", "cellsps", "gbps", "kbps", "mbps", "per-million", "per-thousand", "percent", "pps").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bps", "cellsps", "gbps", "kbps", "mbps", "per-million", "per-thousand", "percent", "pps"),
							},
						},
						"police_conform_action_transmit": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Transmit packet").String,
							Optional:            true,
						},
						"police_conform_action_drop": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Drop packet").String,
							Optional:            true,
						},
						"police_exceed_action_transmit": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Transmit packet").String,
							Optional:            true,
						},
						"police_exceed_action_drop": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Drop packet").String,
							Optional:            true,
						},
						"police_violate_action_transmit": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Transmit packet").String,
							Optional:            true,
						},
						"police_violate_action_drop": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Drop packet").String,
							Optional:            true,
						},
						"shape_average_rate_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
						},
						"shape_average_rate_unit": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Shape rate unit").AddStringEnumDescription("bps", "cellsps", "gbps", "kbps", "mbps", "per-million", "per-thousand", "percent").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bps", "cellsps", "gbps", "kbps", "mbps", "per-million", "per-thousand", "percent"),
							},
						},
						"bandwidth_remaining_unit": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Bandwidth value unit").AddStringEnumDescription("percent", "ratio").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("percent", "ratio"),
							},
						},
						"bandwidth_remaining_value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Bandwidth value").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *PolicyMapQoSResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *PolicyMapQoSResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyMapQoS

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getPath()))

	// Create object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.Id = types.StringValue(plan.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getPath()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PolicyMapQoSResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyMapQoS

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.ValueString()))

	getResp, diags := r.client.Get(ctx, state.Device.ValueString(), state.Id.ValueString())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.updateFromBody(ctx, getResp.Notification[0].Update[0].Val.GetJsonIetfVal())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *PolicyMapQoSResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyMapQoS

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var ops []client.SetOperation

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	// Update object
	body := plan.toBody(ctx)
	ops = append(ops, client.SetOperation{Path: plan.getPath(), Body: body, Operation: client.Update})

	deletedListItems := plan.getDeletedListItems(ctx, state)
	tflog.Debug(ctx, fmt.Sprintf("List items to delete: %+v", deletedListItems))

	for _, i := range deletedListItems {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	emptyLeafsDelete := plan.getEmptyLeafsDelete(ctx)
	tflog.Debug(ctx, fmt.Sprintf("List of empty leafs to delete: %+v", emptyLeafsDelete))

	for _, i := range emptyLeafsDelete {
		ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
	}

	_, diags = r.client.Set(ctx, plan.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PolicyMapQoSResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyMapQoS

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	var ops []client.SetOperation
	deleteMode := "all"

	if deleteMode == "all" {
		ops = append(ops, client.SetOperation{Path: state.Id.ValueString(), Body: "", Operation: client.Delete})
	} else {
		deletePaths := state.getDeletePaths(ctx)
		tflog.Debug(ctx, fmt.Sprintf("Paths to delete: %+v", deletePaths))

		for _, i := range deletePaths {
			ops = append(ops, client.SetOperation{Path: i, Body: "", Operation: client.Delete})
		}
	}

	_, diags = r.client.Set(ctx, state.Device.ValueString(), ops...)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *PolicyMapQoSResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
