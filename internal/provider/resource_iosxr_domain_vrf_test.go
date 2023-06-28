// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrDomainVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrDomainVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "vrf_name", "TEST-VRF"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "domains.0.domain_name", "DOMAIN11"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "domains.0.order", "12345"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "lookup_disable", "true"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "lookup_source_interface", "Loopback2147483647"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "name", "DNAME"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv4_hosts.0.host_name", "HOST-AGC"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv4_hosts.0.ip_address.0", "10.0.0.0"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "name_servers.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "name_servers.0.order", "0"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv6_hosts.0.host_name", "HOST-ACC"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv6_hosts.0.ipv6_address.0", "10::10"),
					resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "multicast", "TESTACC"),
				),
			},
			{
				ResourceName:  "iosxr_domain_vrf.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-domain-cfg:/domain/vrfs/vrf[vrf-name=TEST-VRF]",
			},
		},
	})
}

func testAccIosxrDomainVRFConfig_minimum() string {
	return `
	resource "iosxr_domain_vrf" "test" {
		vrf_name = "TEST-VRF"
	}
	`
}

func testAccIosxrDomainVRFConfig_all() string {
	return `
	resource "iosxr_domain_vrf" "test" {
		vrf_name = "TEST-VRF"
		domains = [{
			domain_name = "DOMAIN11"
			order = 12345
		}]
		lookup_disable = true
		lookup_source_interface = "Loopback2147483647"
		name = "DNAME"
		ipv4_hosts = [{
			host_name = "HOST-AGC"
			ip_address = ["10.0.0.0"]
		}]
		name_servers = [{
			address = "10.0.0.1"
			order = 0
		}]
		ipv6_hosts = [{
			host_name = "HOST-ACC"
			ipv6_address = ["10::10"]
		}]
		multicast = "TESTACC"
	}
	`
}
