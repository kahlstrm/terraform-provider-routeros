# routeros_ip_dhcp_server_option_matcher (Resource)


## Example Usage
```terraform
resource "routeros_ip_dhcp_server_option_matcher" "dhcp1_ip_by_vendor_class" {
  name = "dhcp1_ip_by_vendor_class"
  server = "dhcp1"
  address_pool = "pool1"

  code  = 60  # Vendor Class Identifier
  value = "android-dhcp-11"
  matching_type = "exact"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `code` (Number) DHCP option code. All codes are available at http://www.iana.org/assignments/bootp-dhcp-parameters
- `name` (String) Changing the name of this resource will force it to be recreated.
	> The links of other configuration properties to this resource may be lost!
	> Changing the name of the resource outside of a Terraform will result in a loss of control integrity for that resource!

### Optional

- `address_pool` (String) IP pool, from which to take IP addresses for the clients. If set to static-only, then only the clients that have a static lease (added in lease submenu) will be allowed.
- `comment` (String)
- `disabled` (Boolean)
- `matching_type` (String) Matching method:

- exact: option should match exactly to value
- substring: value can match anywhere in the option string; at the start, middle, or end.
- `option_set` (String) A custom set of DHCP options defined in the Option Sets menu.
- `server` (String) Server name which serves option matcher.
- `value` (String) A value that will be searched for in option.
Available data types for value are:

- string
- HEX

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-server/matcher/get [print show-ids]]
terraform import routeros_ip_dhcp_server_option_matcher.test "*1"
```
