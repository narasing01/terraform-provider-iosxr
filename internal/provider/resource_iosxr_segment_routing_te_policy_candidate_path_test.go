// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrSegmentRoutingTEPolicyCandidatePath(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrSegmentRoutingTEPolicyCandidatePathConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_index", "100"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.type", "dynamic"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.pcep", "true"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.metric_type", "igp"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.hop_type", "mpls"),
					resource.TestCheckResourceAttr("iosxr_segment_routing_te_policy_candidate_path.test", "path_infos.0.segment_list_name", "dynamic"),
				),
			},
			{
				ResourceName:  "iosxr_segment_routing_te_policy_candidate_path.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy[policy-name=POLICY1]/candidate-paths/preferences/preference[path-index=%!d(string=100)]",
			},
		},
	})
}

func testAccIosxrSegmentRoutingTEPolicyCandidatePathConfig_minimum() string {
	return `
	resource "iosxr_segment_routing_te_policy_candidate_path" "test" {
		policy_name = "POLICY1"
		path_index = 100
	}
	`
}

func testAccIosxrSegmentRoutingTEPolicyCandidatePathConfig_all() string {
	return `
	resource "iosxr_segment_routing_te_policy_candidate_path" "test" {
		policy_name = "POLICY1"
		path_index = 100
		path_infos = [{
			type = "dynamic"
			pcep = true
			metric_type = "igp"
			hop_type = "mpls"
			segment_list_name = "dynamic"
		}]
	}
	`
}
