---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_domain Resource - terraform-provider-iosxr"
subcategory: "System"
description: |-
  This resource can manage the Domain configuration.
---

# iosxr_domain (Resource)

This resource can manage the Domain configuration.

## Example Usage

```terraform
resource "iosxr_domain" "example" {
  domains = [
    {
      domain_name = "DOMAIN1"
      order       = 0
    }
  ]
  lookup_disable          = true
  lookup_source_interface = "Loopback2147483647"
  name                    = "DOMAIN"
  ipv4_hosts = [
    {
      host_name  = "HOST_NAME"
      ip_address = ["10.0.0.0"]
    }
  ]
  name_servers = [
    {
      address = "10.0.0.1"
      order   = 345
    }
  ]
  ipv6_hosts = [
    {
      host_name    = "HOST_NAME_IPV6"
      ipv6_address = ["10::10"]
    }
  ]
  multicast             = "DOMAIN1_ACC"
  default_flows_disable = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `default_flows_disable` (Boolean) disables default flows programming
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `domains` (Attributes List) A domain name (see [below for nested schema](#nestedatt--domains))
- `ipv4_hosts` (Attributes List) Name of host (see [below for nested schema](#nestedatt--ipv4_hosts))
- `ipv6_hosts` (Attributes List) Name of host (see [below for nested schema](#nestedatt--ipv6_hosts))
- `lookup_disable` (Boolean) Disable Domain Name System hostname translation
- `lookup_source_interface` (String) Specify source interface for DNS resolver
- `multicast` (String) Define the domain name for multicast address lookups
- `name` (String) Define the default domain name
- `name_servers` (Attributes List) Specify address of name server to use (see [below for nested schema](#nestedatt--name_servers))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--domains"></a>
### Nested Schema for `domains`

Required:

- `domain_name` (String) A domain name
- `order` (Number) This is used to sort the servers in the order of precedence
  - Range: `0`-`4294967295`


<a id="nestedatt--ipv4_hosts"></a>
### Nested Schema for `ipv4_hosts`

Required:

- `host_name` (String) Name of host

Optional:

- `ip_address` (List of String) Host IP address (maximum of 8)


<a id="nestedatt--ipv6_hosts"></a>
### Nested Schema for `ipv6_hosts`

Required:

- `host_name` (String) Name of host

Optional:

- `ipv6_address` (List of String) IPv6 name or address (maximum four addresses)


<a id="nestedatt--name_servers"></a>
### Nested Schema for `name_servers`

Required:

- `address` (String) Specify address of name server to use
- `order` (Number) This is used to sort the servers in the order of precedence
  - Range: `0`-`4294967295`

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_domain.example "Cisco-IOS-XR-um-domain-cfg:/domain"
```
