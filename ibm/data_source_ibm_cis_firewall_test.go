package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisFirewallDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisFirewallDataSourceConfig_basic1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis_firewall.lockdown", "firewall_type", "lockdowns"),
				),
			},
		},
	})
}

func testAccCheckIBMCisFirewallDataSourceConfig_basic1() string {
	return fmt.Sprintf(`
	data "ibm_cis_domain" "cis_domain" {
		cis_id = data.ibm_cis.cis.id
		domain = "cis-terraform.com"
	}
	  
	data "ibm_resource_group" "test_acc" {
		name = "Default"
	}
	  
	data "ibm_cis" "cis" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "test-domain"
	}
	data "ibm_cis_firewall" "lockdown"{
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		firewall_type = "lockdowns"
	  }
	  
	`)
}
