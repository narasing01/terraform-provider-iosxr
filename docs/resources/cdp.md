---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_cdp Resource - terraform-provider-iosxr"
subcategory: "System"
description: |-
  This resource can manage the CDP configuration.
---

# iosxr_cdp (Resource)

This resource can manage the CDP configuration.

## Example Usage

```terraform
resource "iosxr_cdp" "example" {
  enable                = true
  holdtime              = 12
  timer                 = 34
  advertise_v1          = true
  log_adjacency_changes = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `advertise_v1` (Boolean) CDP sends version-1 advertisements only
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `enable` (Boolean) Enable CDP
- `holdtime` (Number) Specify the holdtime (in sec) to be sent in packets
  - Range: `10`-`255`
- `log_adjacency_changes` (Boolean) Configure CDP to log changes to adjacency table
- `timer` (Number) Specify the rate at which CDP packets are sent (in sec)
  - Range: `5`-`254`

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_cdp.example "Cisco-IOS-XR-um-cdp-cfg:/cdp"
```
