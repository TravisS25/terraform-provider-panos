package panos

import (
	"github.com/PaloAltoNetworks/pango"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceEthernetInterfaces() *schema.Resource {
	s := listingSchema()
	s["vsys"] = vsysSchema("vsys1")
	s["device_group"] = deviceGroupSchema()

	return &schema.Resource{
		Read: dataSourceEthernetInterfacesRead,

		Schema: s,
	}
}

func dataSourceEthernetInterfacesRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	var listing []string
	var id string

	switch con := meta.(type) {
	case *pango.Firewall:
		listing, err = con.Network.EthernetInterface.GetList()
	}

	if err != nil {
		return err
	}

	d.SetId(id)
	saveListing(d, listing)
	return nil
}
