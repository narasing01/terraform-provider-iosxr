// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrMPLSOAM(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrMPLSOAMConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam", "true"),
					resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_echo_disable_vendor_extension", "false"),
					resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_echo_reply_mode_control_channel_allow_reverse_lsp", "false"),
					resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_dpm_pps", "10"),
					resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_dpm_interval", "60"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrMPLSOAMConfig = `

resource "iosxr_mpls_oam" "test" {
	oam = true
	oam_echo_disable_vendor_extension = false
	oam_echo_reply_mode_control_channel_allow_reverse_lsp = false
	oam_dpm_pps = 10
	oam_dpm_interval = 60
}

data "iosxr_mpls_oam" "test" {
	depends_on = [iosxr_mpls_oam.test]
}
`