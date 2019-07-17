package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTrelloList() *schema.Resource {
	return &schema.Resource{
		Delete: resourceTrelloListDelete,
		Schema: map[string]*schema.Schema{
			"board_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"closed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"pos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
		},
	}
}
func resourceTrelloListDelete(d *schema.ResourceData, meta interface{}) error {
	// the TrelloAPI does not supported deleting lists. lists may only be archived.
	// https://developers.trello.com/reference/#lists
	return nil
}
