// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSNMPServerMIB(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSNMPServerMIBConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_snmp_server_mib.test", "ifmib_ifalias_long", "true"),
					resource.TestCheckResourceAttr("iosxr_snmp_server_mib.test", "ifindex_persist", "true"),
				),
			},
			{
				ResourceName:  "iosxr_snmp_server_mib.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server-mibs",
			},
		},
	})
}

func testAccIosxrSNMPServerMIBConfig_minimum() string {
	return `
	resource "iosxr_snmp_server_mib" "test" {
	}
	`
}

func testAccIosxrSNMPServerMIBConfig_all() string {
	return `
	resource "iosxr_snmp_server_mib" "test" {
		ifmib_ifalias_long = true
		ifindex_persist = true
	}
	`
}