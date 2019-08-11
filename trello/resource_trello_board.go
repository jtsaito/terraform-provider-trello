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
		Update: resourceTrelloBoardUpdate,
		Delete: resourceTrelloBoardDelete,

		Schema: map[string]*schema.Schema{
			"closed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"pinned": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"short_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTrelloBoardCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	b := trello.Board{
		Closed:         d.Get("closed").(bool),
		Desc:           d.Get("description").(string),
		IDOrganization: d.Get("organization_id").(string),
		Name:           d.Get("name").(string),
		Pinned:         d.Get("pinned").(bool),
	}
	args := trello.Arguments{"defaultLists": "false"}

	err := client.CreateBoard(&b, args)
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
	d.Set("closed", board.Closed)
	d.Set("description", board.Desc)
	d.Set("name", board.Name)
	d.Set("organization_id", board.IDOrganization)
	d.Set("pinned", board.Pinned)
	d.Set("url", board.URL)
	d.Set("short_url", board.ShortURL)

	return nil
}

// resourceTrelloBoardUpdate will overwrite with board default values!
func resourceTrelloBoardUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Config).Client

	b := trello.Board{
		ID:   d.Id(),
		Desc: d.Get("description").(string),
		Name: d.Get("name").(string),
	}

	err := client.PutBoard(&b, trello.Defaults())
	if err != nil {
		return fmt.Errorf("could not put board: %s", err)
	}

	return resourceTrelloBoardRead(d, meta)
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
