---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_prefix_set Resource - terraform-provider-iosxr"
subcategory: "Route Policy"
description: |-
  This resource can manage the Prefix Set configuration.
---

# iosxr_prefix_set (Resource)

This resource can manage the Prefix Set configuration.

## Example Usage

```terraform
resource "iosxr_prefix_set" "example" {
  set_name = "PREFIX_SET_1"
  rpl      = "prefix-set PREFIX_SET_1\n  10.1.1.0/26 ge 26,\n  10.1.2.0/26 ge 26\nend-set\n"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `rpl` (String) prefix statements
- `set_name` (String) Set name

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_prefix_set.example "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/prefix-sets/prefix-set[set-name=PREFIX_SET_1]"
```
