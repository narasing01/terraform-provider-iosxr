// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrExtcommunitySet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrExtcommunitySetConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_extcommunity_set.test", "rpl_extended_community_opaque_set", "extcommunity-set opaque BLUE\n  100\nend-set\n"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrExtcommunitySetConfig = `

resource "iosxr_extcommunity_set" "test" {
	set_name = "BLUE"
	rpl_extended_community_opaque_set = "extcommunity-set opaque BLUE\n  100\nend-set\n"
}

data "iosxr_extcommunity_set" "test" {
	set_name = "BLUE"
	depends_on = [iosxr_extcommunity_set.test]
}
`
