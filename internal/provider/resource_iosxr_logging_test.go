// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrLogging(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrLoggingConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_logging.test", "ipv4_dscp", "cs6"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "trap", "informational"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "events_display_location", "true"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "events_level", "informational"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "console", "disable"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "monitor", "disable"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "buffered_logging_buffer_size", "4000000"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "buffered_level", "debugging"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "facility_level", "local7"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "hostnameprefix", "HOSTNAME01"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "suppress_duplicates", "true"),
					resource.TestCheckResourceAttr("iosxr_logging.test", "source_interfaces.0.source_interface_name", "Loopback10"),
				),
			},
			{
				ResourceName:  "iosxr_logging.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-logging-cfg:logging",
			},
		},
	})
}

func testAccIosxrLoggingConfig_minimum() string {
	return `
	resource "iosxr_logging" "test" {
	}
	`
}

func testAccIosxrLoggingConfig_all() string {
	return `
	resource "iosxr_logging" "test" {
		ipv4_dscp = "cs6"
		trap = "informational"
		events_display_location = true
		events_level = "informational"
		console = "disable"
		monitor = "disable"
		buffered_logging_buffer_size = 4000000
		buffered_level = "debugging"
		facility_level = "local7"
		hostnameprefix = "HOSTNAME01"
		suppress_duplicates = true
		source_interfaces = [{
			source_interface_name = "Loopback10"
		}]
	}
	`
}