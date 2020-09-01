package fastly

import (
	// "log"
	// "github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type dataSourceFastlyACLEntriesResult struct {
	Addresses []string
}

func dataSourceFastlyACLEntries() *schema.Resource {
	return &schema.Resource{
		Read: resourceServiceAclEntriesV1Read,

		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"acl_id": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
