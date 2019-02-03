package main

import (
	"errors"
	"fmt"
	"log"

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
	return nil
}
