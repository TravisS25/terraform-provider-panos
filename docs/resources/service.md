---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "panos_service Resource - panos"
subcategory: ""
description: |-
  
---

# panos_service (Resource)



## Example Usage

```terraform
resource "panos_service" "example" {
  location = {
    device_group = {
      name = panos_device_group.example.name
    }
  }

  name        = "example-service"
  description = "example service"

  protocol = {
    tcp = {
      destination_port = "80"
    }
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `location` (Attributes) The location of this object. (see [below for nested schema](#nestedatt--location))
- `name` (String) The name of the service.

### Optional

- `description` (String) The description.
- `protocol` (Attributes) (see [below for nested schema](#nestedatt--protocol))
- `tags` (List of String) The administrative tags.

### Read-Only

- `tfid` (String) The Terraform ID.

<a id="nestedatt--location"></a>
### Nested Schema for `location`

Optional:

- `device_group` (Attributes) Located in a specific device group. (see [below for nested schema](#nestedatt--location--device_group))
- `from_panorama_shared` (Boolean) Located in shared in the config pushed from Panorama.
- `from_panorama_vsys` (Attributes) Located in a specific vsys in the config pushed from Panorama. (see [below for nested schema](#nestedatt--location--from_panorama_vsys))
- `shared` (Boolean) Located in shared.
- `vsys` (Attributes) Located in a specific vsys. (see [below for nested schema](#nestedatt--location--vsys))

<a id="nestedatt--location--device_group"></a>
### Nested Schema for `location.device_group`

Optional:

- `name` (String) The device group.
- `panorama_device` (String) The panorama device.


<a id="nestedatt--location--from_panorama_vsys"></a>
### Nested Schema for `location.from_panorama_vsys`

Optional:

- `vsys` (String) The vsys.


<a id="nestedatt--location--vsys"></a>
### Nested Schema for `location.vsys`

Optional:

- `name` (String) The vsys.
- `ngfw_device` (String) The NGFW device.



<a id="nestedatt--protocol"></a>
### Nested Schema for `protocol`

Optional:

- `tcp` (Attributes) (see [below for nested schema](#nestedatt--protocol--tcp))
- `udp` (Attributes) (see [below for nested schema](#nestedatt--protocol--udp))

<a id="nestedatt--protocol--tcp"></a>
### Nested Schema for `protocol.tcp`

Optional:

- `destination_port` (Number) The destination port.
- `override` (Attributes) (see [below for nested schema](#nestedatt--protocol--tcp--override))
- `source_port` (Number) The source port.

<a id="nestedatt--protocol--tcp--override"></a>
### Nested Schema for `protocol.tcp.override`

Optional:

- `halfclose_timeout` (Number) TCP session half-close timeout value (in second)
- `timeout` (Number) TCP session timeout value (in second)
- `timewait_timeout` (Number) TCP session time-wait timeout value (in second)



<a id="nestedatt--protocol--udp"></a>
### Nested Schema for `protocol.udp`

Optional:

- `destination_port` (Number) The destination port.
- `override` (Attributes) (see [below for nested schema](#nestedatt--protocol--udp--override))
- `source_port` (Number) The source port.

<a id="nestedatt--protocol--udp--override"></a>
### Nested Schema for `protocol.udp.override`

Optional:

- `no` (String)
- `yes` (Attributes) (see [below for nested schema](#nestedatt--protocol--udp--override--yes))

<a id="nestedatt--protocol--udp--override--yes"></a>
### Nested Schema for `protocol.udp.override.yes`

Optional:

- `timeout` (Number) UDP session timeout value (in second)