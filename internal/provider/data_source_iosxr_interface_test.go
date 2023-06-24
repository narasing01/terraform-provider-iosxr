// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrInterfacePrerequisitesConfig + testAccDataSourceIosxrInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "l2transport", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "point_to_point", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "multipoint", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "dampening_decay_half_life_value", "2"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv4_point_to_point", "true"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "service_policy_input.0.name", "PMAP-IN"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "service_policy_output.0.name", "PMAP-OUT"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "shutdown", "true"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "mtu", "9000"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "bandwidth", "100000"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "description", "My Interface Description"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "load_interval", "30"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv4_address", "1.1.1.1"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv4_netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv6_link_local_address", "fe80::1"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv6_link_local_zone", "0"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv6_autoconfig", "false"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv6_enable", "true"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv6_addresses.0.address", "2001::1"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv6_addresses.0.prefix_length", "64"),
					resource.TestCheckResourceAttr("data.iosxr_interface.test", "ipv6_addresses.0.zone", "0"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrInterfacePrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name=PMAP-IN]"
	attributes = {
		"policy-map-name" = "PMAP-IN"
	}
	lists = [
		{
			name = "class"
			key = "name,type"
			items = [
				{
					"name" = "class-default"
					"type" = "qos"
					"set/qos-group" = "0"
				},
			]
		},
	]
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name=PMAP-OUT]"
	attributes = {
		"policy-map-name" = "PMAP-OUT"
	}
	lists = [
		{
			name = "class"
			key = "name,type"
			items = [
				{
					"name" = "class-default"
					"type" = "qos"
					"set/dscp" = "0"
				},
			]
		},
	]
}

`

const testAccDataSourceIosxrInterfaceConfig = `

resource "iosxr_interface" "test" {
	interface_name = "GigabitEthernet0/0/0/1"
	l2transport = false
	point_to_point = false
	multipoint = false
	dampening_decay_half_life_value = 2
	ipv4_point_to_point = true
	service_policy_input = [{
		name = "PMAP-IN"
	}]
	service_policy_output = [{
		name = "PMAP-OUT"
	}]
	shutdown = true
	mtu = 9000
	bandwidth = 100000
	description = "My Interface Description"
	load_interval = 30
	vrf = "VRF1"
	ipv4_address = "1.1.1.1"
	ipv4_netmask = "255.255.255.0"
	ipv6_link_local_address = "fe80::1"
	ipv6_link_local_zone = "0"
	ipv6_autoconfig = false
	ipv6_enable = true
	ipv6_addresses = [{
		address = "2001::1"
		prefix_length = 64
		zone = "0"
	}]
	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]
}

data "iosxr_interface" "test" {
	interface_name = "GigabitEthernet0/0/0/1"
	depends_on = [iosxr_interface.test]
}
`
