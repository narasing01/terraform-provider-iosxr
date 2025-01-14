---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_logging Resource - terraform-provider-iosxr"
subcategory: "Management"
description: |-
  This resource can manage the Logging configuration.
---

# iosxr_logging (Resource)

This resource can manage the Logging configuration.

## Example Usage

```terraform
resource "iosxr_logging" "example" {
  ipv4_dscp                    = "cs6"
  trap                         = "informational"
  events_display_location      = true
  events_level                 = "informational"
  console                      = "disable"
  monitor                      = "disable"
  buffered_logging_buffer_size = 4000000
  buffered_level               = "debugging"
  facility_level               = "local7"
  hostnameprefix               = "HOSTNAME01"
  suppress_duplicates          = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `buffered_level` (String) configure this node
  - Choices: `alerts`, `critical`, `debugging`, `emergencies`, `errors`, `informational`, `notifications`, `warnings`
- `buffered_logging_buffer_size` (Number) Logging buffer size
  - Range: `307200`-`125000000`
- `console` (String) Set console logging
  - Choices: `alerts`, `critical`, `debugging`, `disable`, `emergencies`, `errors`, `informational`, `notifications`, `warning`
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `events_display_location` (Boolean) Include alarm source location in message text
- `events_level` (String) Log all events with equal or higher (lower level) severity
  - Choices: `alerts`, `critical`, `emergencies`, `errors`, `informational`, `notifications`, `warnings`
- `facility_level` (String) configure this node
  - Choices: `auth`, `cron`, `daemon`, `kern`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`, `lpr`, `mail`, `news`, `sys10`, `sys11`, `sys12`, `sys13`, `sys14`, `sys9`, `syslog`, `user`, `uucp`
- `hostnameprefix` (String) Hostname prefix to add on msgs to servers
- `ipv4_dscp` (String) Set IP DSCP (DiffServ CodePoint)
- `monitor` (String) Set monitor logging
  - Choices: `alerts`, `critical`, `debugging`, `disable`, `emergencies`, `errors`, `informational`, `notifications`, `warning`
- `suppress_duplicates` (Boolean) Suppress consecutive duplicate messages
- `trap` (String) Set trap logging
  - Choices: `alerts`, `critical`, `debugging`, `disable`, `emergencies`, `errors`, `informational`, `notifications`, `warning`

### Read-Only

- `id` (String) The path of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_logging.example "Cisco-IOS-XR-um-logging-cfg:/logging"
```
