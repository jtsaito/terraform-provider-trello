package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTrelloBoard() *schema.Resource {
	return &schema.Resource{
		Create: resourceTrelloBoardCreate,
		Read:   resourceTrelloBoardRead,
		Delete: resourceTrelloBoardDelete,

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceTrelloBoardCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	b := trello.Board{Name: d.Get("name").(string)}

	err := client.CreateBoard(&b, trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not create board: %s", err)
	}

	return resourceTrelloBoardRead(d, meta)
}

func resourceTrelloBoardRead(d *schema.ResourceData, meta interface{}) error {
	member := meta.(*Config).Member

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		return fmt.Errorf("failed to fetch boards for member: %s", err)
	}

	var board *trello.Board
	for _, b := range boards {
		if b.Name == d.Get("name") {
			board = b
			break
		}
	}
	if board == nil {
		return fmt.Errorf("Trello board not found. board %s", d.Get("name"))
	}

	d.SetId(board.ID)
	d.Set("name", board.Name)
	d.Set("description", board.Desc)

	return nil
}

func resourceTrelloBoardDelete(d *schema.ResourceData, meta interface{}) error {

	// The trello API does not allow to set a client on an empty board.
	// We need to run GetBoard on the client, not on the trello.Board struct.
	client := meta.(*Config).Client
	board, err := client.GetBoard(d.Id(), trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not get Trello board by id %s: %s", d.Id(), err)
	}

	err = board.Delete(trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not delete Trello board %s: %s", d.Id(), err)
	}

	return nil
}
