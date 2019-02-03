package main

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceTrelloBoard() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTrelloBoardRead,

		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceTrelloBoardRead(d *schema.ResourceData, m interface{}) error {
	member := m.(*TrelloConfig).Member

	boards, err := member.GetBoards(trello.Defaults())
	if err != nil {
		return err
	}

	var board *trello.Board
	for _, b := range boards {
		if b.Name == d.Get("name") {
			board = b
			break
		}
	}

	if board == nil {
		return fmt.Errorf("not found. board %s", d.Get("name"))
	}

	d.SetId(board.ID)
	d.Set("name", board.Name)
	d.Set("description", board.Desc)

	return nil
}
