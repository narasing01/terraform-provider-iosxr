// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrRouterOSPF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterOSPFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "mpls_ldp_sync", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "hello_interval", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "dead_interval", "40"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "priority", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "mtu_ignore_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "mtu_ignore_disable", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "passive_enable", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "passive_disable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "router_id", "10.11.12.13"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_connected", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_connected_tag", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_connected_metric_type", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_static", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_static_tag", "2"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_static_metric_type", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "bfd_fast_detect", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "bfd_minimum_interval", "300"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "bfd_multiplier", "3"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "default_information_originate", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "default_information_originate_always", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "default_information_originate_metric_type", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "areas.0.area_id", "0"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_bgp.0.as_number", "65001"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_bgp.0.tag", "3"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_bgp.0.metric_type", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_isis.0.instance_name", "P1"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_isis.0.level_1", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_isis.0.level_2", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_isis.0.level_1_2", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_isis.0.tag", "3"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_isis.0.metric_type", "1"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_ospf.0.instance_name", "OSPF2"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_ospf.0.match_internal", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_ospf.0.match_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_ospf.0.match_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_ospf.0.tag", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf.test", "redistribute_ospf.0.metric_type", "1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterOSPFConfig = `

resource "iosxr_router_ospf" "test" {
	delete_mode = "attributes"
	process_name = "OSPF1"
	mpls_ldp_sync = false
	hello_interval = 10
	dead_interval = 40
	priority = 10
	mtu_ignore_enable = true
	mtu_ignore_disable = false
	passive_enable = false
	passive_disable = true
	router_id = "10.11.12.13"
	redistribute_connected = true
	redistribute_connected_tag = 1
	redistribute_connected_metric_type = "1"
	redistribute_static = true
	redistribute_static_tag = 2
	redistribute_static_metric_type = "1"
	bfd_fast_detect = true
	bfd_minimum_interval = 300
	bfd_multiplier = 3
	default_information_originate = true
	default_information_originate_always = true
	default_information_originate_metric_type = 1
	areas = [{
		area_id = "0"
	}]
	redistribute_bgp = [{
		as_number = "65001"
		tag = 3
		metric_type = "1"
	}]
	redistribute_isis = [{
		instance_name = "P1"
		level_1 = true
		level_2 = false
		level_1_2 = false
		tag = 3
		metric_type = "1"
	}]
	redistribute_ospf = [{
		instance_name = "OSPF2"
		match_internal = true
		match_external = false
		match_nssa_external = false
		tag = 4
		metric_type = "1"
	}]
}

data "iosxr_router_ospf" "test" {
	process_name = "OSPF1"
	depends_on = [iosxr_router_ospf.test]
}
`
