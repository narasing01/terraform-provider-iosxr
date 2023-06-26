// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrL2VPNBridgeGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrL2VPNBridgeGroupConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceIosxrL2VPNBridgeGroupConfig = `

resource "iosxr_l2vpn_bridge_group" "test" {
	delete_mode = "attributes"
	group_name = "BG123"
}

data "iosxr_l2vpn_bridge_group" "test" {
	group_name = "BG123"
	depends_on = [iosxr_l2vpn_bridge_group.test]
}
`
