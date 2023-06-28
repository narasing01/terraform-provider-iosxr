// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrServiceTimestamps(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrServiceTimestampsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "debug_datetime_localtime", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "debug_datetime_msec", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "debug_datetime_show_timezone", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "debug_datetime_year", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "debug_uptime", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "debug_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "log_datetime_localtime", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "log_datetime_msec", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "log_datetime_show_timezone", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "log_datetime_year", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "log_uptime", "true"),
					resource.TestCheckResourceAttr("data.iosxr_service_timestamps.test", "log_disable", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrServiceTimestampsConfig = `

resource "iosxr_service_timestamps" "test" {
	debug_datetime_localtime = true
	debug_datetime_msec = true
	debug_datetime_show_timezone = true
	debug_datetime_year = true
	debug_uptime = true
	debug_disable = true
	log_datetime_localtime = true
	log_datetime_msec = true
	log_datetime_show_timezone = true
	log_datetime_year = true
	log_uptime = true
	log_disable = true
}

data "iosxr_service_timestamps" "test" {
	depends_on = [iosxr_service_timestamps.test]
}
`