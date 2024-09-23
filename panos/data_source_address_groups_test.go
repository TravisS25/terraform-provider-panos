package panos

import (
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/pango"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Data source listing tests.
func TestAccPanosDsAddressGroupList(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsAddressObjectConfig(name),
				Check:  checkDataSourceListing("panos_address_groups"),
			},
		},
	})
}

func testAccDsAddressGroupConfig(objectName, objectValue, groupName string) string {
	return fmt.Sprintf(
		`
		data "panos_address_groups" "test" {}

		resource "panos_address_object" "test" {
			name = %q
			description = "address object acctest"
			value = %q
		}

		resource "panos_address_group" "x" {
			name = %q
			description = "address group acctest"
			static_addresses = [
				panos_address_object.test.value
			]
		}
		`,
		objectName,
		objectValue,
		groupName,
	)
}

func testAccDsPanosAddressGroupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "panos_address_group" {
			continue
		}

		if rs.Primary.ID != "" {
			var err error

			switch con := testAccProvider.Meta().(type) {
			case *pango.Firewall:
				vsys, name := parseAddressGroupId(rs.Primary.ID)
				_, err = con.Objects.Address.Get(vsys, name)
			case *pango.Panorama:
				dg, name := parseAddressGroupId(rs.Primary.ID)
				_, err = con.Objects.Address.Get(dg, name)
			}
			if err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}
