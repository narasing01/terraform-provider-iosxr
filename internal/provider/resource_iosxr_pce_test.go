// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrPCE(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrPCEConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_pce.test", "address_ipv4", "77.77.77.1"),
					resource.TestCheckResourceAttr("iosxr_pce.test", "state_sync_ipv4s.0.address", "100.100.100.11"),
					resource.TestCheckResourceAttr("iosxr_pce.test", "peer_filter_ipv4_access_list", "Accesslist1"),
					resource.TestCheckResourceAttr("iosxr_pce.test", "api_authentication_digest", "true"),
					resource.TestCheckResourceAttr("iosxr_pce.test", "api_sibling_ipv4", "100.100.100.2"),
					resource.TestCheckResourceAttr("iosxr_pce.test", "api_users.0.user_name", "rest-user"),
					resource.TestCheckResourceAttr("iosxr_pce.test", "api_users.0.password_encrypted", "00141215174C04140B"),
				),
			},
			{
				ResourceName:  "iosxr_pce.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-pce-cfg:/pce",
			},
		},
	})
}

func testAccIosxrPCEConfig_minimum() string {
	return `
	resource "iosxr_pce" "test" {
	}
	`
}

func testAccIosxrPCEConfig_all() string {
	return `
	resource "iosxr_pce" "test" {
		address_ipv4 = "77.77.77.1"
		state_sync_ipv4s = [{
			address = "100.100.100.11"
		}]
		peer_filter_ipv4_access_list = "Accesslist1"
		api_authentication_digest = true
		api_sibling_ipv4 = "100.100.100.2"
		api_users = [{
			user_name = "rest-user"
			password_encrypted = "00141215174C04140B"
		}]
	}
	`
}
