package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTrelloList() *schema.Resource {
	return &schema.Resource{
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
