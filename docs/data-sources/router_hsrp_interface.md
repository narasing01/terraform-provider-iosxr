---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_router_hsrp_interface Data Source - terraform-provider-iosxr"
subcategory: "HSRP"
description: |-
  This data source can read the Router HSRP Interface configuration.
---

# iosxr_router_hsrp_interface (Data Source)

This data source can read the Router HSRP Interface configuration.

## Example Usage

```terraform
data "iosxr_router_hsrp_interface" "example" {
  interface_name = "GigabitEthernet0/0/0/1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `interface_name` (String) HSRP interface configuration subcommands

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `hsrp_bfd_minimum_interval` (Number) Hello interval
- `hsrp_bfd_multiplier` (Number) Detect multiplier
- `hsrp_delay_minimum` (Number) Set minimum delay on every interface up event
- `hsrp_delay_reload` (Number) Set reload delay for first interface up event
- `hsrp_mac_refresh` (Number) HSRP MGO subordinate MAC refresh rate
- `hsrp_redirects_disable` (Boolean) Disable HSRP filtered ICMP redirects
- `hsrp_use_bia` (Boolean) Use burned-in address
- `id` (String) The path of the retrieved object.