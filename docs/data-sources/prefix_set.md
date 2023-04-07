---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_prefix_set Data Source - terraform-provider-iosxr"
subcategory: "Route Policy"
description: |-
  This data source can read the Prefix Set configuration.
---

# iosxr_prefix_set (Data Source)

This data source can read the Prefix Set configuration.

## Example Usage

```terraform
data "iosxr_prefix_set" "example" {
  set_name = "PREFIX_SET_1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `set_name` (String) Set name

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The path of the retrieved object.
- `rpl` (String) prefix statements

