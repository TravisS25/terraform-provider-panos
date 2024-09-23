package panos

import (
	"github.com/PaloAltoNetworks/pango"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServiceGroups() *schema.Resource {
	s := listingSchema()
	s["vsys"] = vsysSchema("vsys1")
	s["device_group"] = deviceGroupSchema()

	return &schema.Resource{
		Read: dataSourceServiceGroupsRead,

		Schema: s,
	}
}

func dataSourceServiceGroupsRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	var listing []string
	var id string

	switch con := meta.(type) {
	case *pango.Firewall:
		id = d.Get("vsys").(string)
		listing, err = con.Objects.ServiceGroup.GetList(id)
	case *pango.Panorama:
		id = d.Get("device_group").(string)
		listing, err = con.Objects.ServiceGroup.GetList(id)
	}

	if err != nil {
		return err
	}

	d.SetId(id)
	saveListing(d, listing)
	return nil
}
