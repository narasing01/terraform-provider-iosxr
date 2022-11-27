// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &VRFDataSource{}
	_ datasource.DataSourceWithConfigure = &VRFDataSource{}
)

func NewVRFDataSource() datasource.DataSource {
	return &VRFDataSource{}
}

type VRFDataSource struct {
	client *client.Client
}

func (d *VRFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrf"
}

func (d *VRFDataSource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the VRF configuration.",

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The path of the retrieved object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"vrf_name": {
				MarkdownDescription: "VRF name",
				Type:                types.StringType,
				Required:            true,
			},
			"description": {
				MarkdownDescription: "A description for the VRF",
				Type:                types.StringType,
				Computed:            true,
			},
			"vpn_id": {
				MarkdownDescription: "VPN ID, (OUI:VPN-Index) format(hex), 4 bytes VPN_Index Part",
				Type:                types.StringType,
				Computed:            true,
			},
			"address_family_ipv4_unicast": {
				MarkdownDescription: "Unicast sub address family",
				Type:                types.BoolType,
				Computed:            true,
			},
			"address_family_ipv4_multicast": {
				MarkdownDescription: "Multicast topology",
				Type:                types.BoolType,
				Computed:            true,
			},
			"address_family_ipv4_flowspec": {
				MarkdownDescription: "Flowspec sub address family",
				Type:                types.BoolType,
				Computed:            true,
			},
			"address_family_ipv6_unicast": {
				MarkdownDescription: "Unicast sub address family",
				Type:                types.BoolType,
				Computed:            true,
			},
			"address_family_ipv6_multicast": {
				MarkdownDescription: "Multicast topology",
				Type:                types.BoolType,
				Computed:            true,
			},
			"address_family_ipv6_flowspec": {
				MarkdownDescription: "Flowspec sub address family",
				Type:                types.BoolType,
				Computed:            true,
			},
			"rd_two_byte_as_as_number": {
				MarkdownDescription: "bgp as-number",
				Type:                types.StringType,
				Computed:            true,
			},
			"rd_two_byte_as_index": {
				MarkdownDescription: "ASN2:index (hex or decimal format)",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"rd_four_byte_as_as_number": {
				MarkdownDescription: "4-byte AS number",
				Type:                types.StringType,
				Computed:            true,
			},
			"rd_four_byte_as_index": {
				MarkdownDescription: "ASN2:index (hex or decimal format)",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"rd_ip_address_ipv4_address": {
				MarkdownDescription: "configure this node",
				Type:                types.StringType,
				Computed:            true,
			},
			"rd_ip_address_index": {
				MarkdownDescription: "IPv4Address:index (hex or decimal format)",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"address_family_ipv4_unicast_import_route_target_two_byte_as_format": {
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Two Byte AS Number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv4_unicast_import_route_target_four_byte_as_format": {
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Four Byte AS number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv4_unicast_import_route_target_ip_address_format": {
				MarkdownDescription: "IP address",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip_address": {
						MarkdownDescription: "IP address",
						Type:                types.StringType,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "IPv4Address:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv4_unicast_export_route_target_two_byte_as_format": {
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Two Byte AS Number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv4_unicast_export_route_target_four_byte_as_format": {
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Four Byte AS number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv4_unicast_export_route_target_ip_address_format": {
				MarkdownDescription: "IP address",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip_address": {
						MarkdownDescription: "IP address",
						Type:                types.StringType,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "IPv4Address:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv6_unicast_import_route_target_two_byte_as_format": {
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Two Byte AS Number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv6_unicast_import_route_target_four_byte_as_format": {
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Four Byte AS number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv6_unicast_import_route_target_ip_address_format": {
				MarkdownDescription: "IP address",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip_address": {
						MarkdownDescription: "IP address",
						Type:                types.StringType,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "IPv4Address:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv6_unicast_export_route_target_two_byte_as_format": {
				MarkdownDescription: "Two Byte AS Number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Two Byte AS Number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv6_unicast_export_route_target_four_byte_as_format": {
				MarkdownDescription: "Four Byte AS number Route Target",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"as_number": {
						MarkdownDescription: "Four Byte AS number",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "ASN2:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
			"address_family_ipv6_unicast_export_route_target_ip_address_format": {
				MarkdownDescription: "IP address",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"ip_address": {
						MarkdownDescription: "IP address",
						Type:                types.StringType,
						Computed:            true,
					},
					"index": {
						MarkdownDescription: "IPv4Address:index (hex or decimal format)",
						Type:                types.Int64Type,
						Computed:            true,
					},
					"stitching": {
						MarkdownDescription: "These are stitching RTs",
						Type:                types.BoolType,
						Computed:            true,
					},
				}),
			},
		},
	}, nil
}

func (d *VRFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*client.Client)
}

func (d *VRFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config VRF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getPath()))

	getResp, diags := d.client.Get(ctx, config.Device.ValueString(), config.getPath())
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.fromBody(getResp.Notification[0].Update[0].Val.GetJsonIetfVal())
	config.Id = types.StringValue(config.getPath())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getPath()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
