package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTrelloList() *schema.Resource {
	return &schema.Resource{
		Create: resourceTrelloListCreate,
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

func resourceTrelloListCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	board, err := client.GetBoard(d.Get("board_id").(string), trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not get board: %s", err)
	}

	list, err := board.CreateList(d.Get("name").(string), trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not create list: %s", err)
	}

	d.SetId(list.ID)

	return resourceTrelloListRead(d, meta)
}
func resourceTrelloListDelete(d *schema.ResourceData, meta interface{}) error {
	// the TrelloAPI does not supported deleting lists. lists may only be archived.
	// https://developers.trello.com/reference/#lists
	return nil
}
