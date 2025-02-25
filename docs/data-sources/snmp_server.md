---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_snmp_server Data Source - terraform-provider-iosxr"
subcategory: "Management"
description: |-
  This data source can read the SNMP Server configuration.
---

# iosxr_snmp_server (Data Source)

This data source can read the SNMP Server configuration.

## Example Usage

```terraform
data "iosxr_snmp_server" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `bfd` (Boolean) Enable BFD traps
- `bgp_bgp4_mib_updown` (Boolean) Enable CISCO-BGP4-MIB v2 up/down traps
- `bgp_cbgp2_updown` (Boolean) Enable CISCO-BGP4-MIB v2 up/down traps
- `bridgemib` (Boolean) Enable SNMP Trap for Bridge MIB
- `config` (Boolean) Enable SNMP config traps
- `contact` (String) Text for mib Object sysContact
- `copy_complete` (Boolean) Enable CISCO-CONFIG-COPY-MIB ccCopyCompletion traps
- `entity` (Boolean) Enable SNMP entity traps
- `entity_redundancy_all` (Boolean) Enable all CISCO-ENTITY-REDUNDANCY-MIB traps
- `entity_state_operstatus` (Boolean) Enable entity oper status enable notification
- `ethernet_oam_events` (Boolean) Enable all OAM event traps
- `fru_ctrl` (Boolean) Enable SNMP entity FRU control traps
- `groups` (Attributes List) Name of the group (see [below for nested schema](#nestedatt--groups))
- `id` (String) The path of the retrieved object.
- `isis_adjacency_change` (String) isisAdjacencyChange
- `isis_all` (String) Enable all IS-IS traps
- `isis_area_mismatch` (String) isisAreaMismatch
- `isis_attempt_to_exceed_max_sequence` (String) isisAttemptToExceedMaxSequence
- `isis_authentication_failure` (String) isisAuthenticationFailure
- `isis_authentication_type_failure` (String) isisAuthenticationTypeFailure
- `isis_corrupted_lsp_detected` (String) isisCorruptedLSPDetected
- `isis_database_overload` (String) isisDatabaseOverload
- `isis_id_len_mismatch` (String) isisIDLenMismatch
- `isis_lsp_error_detected` (String) isisLSPErrorDetected
- `isis_lsp_too_large_to_propagate` (String) isisLSPTooLargeToPropagate
- `isis_manual_address_drops` (String) isisManualAddressDrops
- `isis_max_area_addresses_mismatch` (String) isisMaxAreaAddressesMismatch
- `isis_orig_lsp_buff_size_mismatch` (String) isisOrigLSPBuffSizeMismatch
- `isis_own_lsp_purge` (String) isisOwnLSPPurge
- `isis_protocols_supported_mismatch` (String) isisProtocolsSupportedMismatch
- `isis_rejected_adjacency` (String) isisRejectedAdjacency
- `isis_sequence_number_skip` (String) isisSequenceNumberSkip
- `isis_version_skew` (String) isisVersionSkew
- `l2vpn_all` (Boolean) Enable all L2VPN traps
- `l2vpn_vc_down` (Boolean) Enable VC down traps
- `l2vpn_vc_up` (Boolean) Enable VC up traps
- `location` (String) Text for mib Object sysLocation
- `ntp` (Boolean) Enable SNMP Cisco Ntp traps
- `power` (Boolean) Enable SNMP entity power traps
- `rf` (Boolean) Enable SNMP RF-MIB traps
- `sensor` (Boolean) Enable SNMP entity sensor traps
- `system` (Boolean) Enable SNMP SYSTEMMIB-MIB traps
- `trap_source_both` (String) Assign an interface for the source address of all traps
- `traps_snmp_linkdown` (Boolean) Enable SNMPv2-MIB linDownp traps
- `traps_snmp_linkup` (Boolean) Enable SNMPv2-MIB linkUp traps
- `users` (Attributes List) Name of the user (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--groups"></a>
### Nested Schema for `groups`

Read-Only:

- `group_name` (String) Name of the group
- `v3_context` (String) Attach a SNMP context
- `v3_ipv4` (String) Type of Access-list
- `v3_ipv6` (String) Type of Access-list
- `v3_notify` (String) specify a notify view for the group
- `v3_priv` (Boolean) group using authPriv Security Level
- `v3_read` (String) specify a read view for this group
- `v3_write` (String) specify a write view for this group


<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `group_name` (String) Group to which the user belongs
- `user_name` (String) Name of the user
- `v3_auth_md5_encryption_aes` (String) Specifies an aes-128 ENCRYPTED authentication password
- `v3_auth_md5_encryption_default` (String) Specifies an default ENCRYPTED authentication password
