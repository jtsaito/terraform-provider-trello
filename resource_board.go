package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBoard() *schema.Resource {
	return &schema.Resource{
		Create: resourceBoradCreate,
		Read:   resourceBoradRead,
		Update: resourceBoradUpdate,
		Delete: resourceBoradDelete,

		Schema: map[string]*schema.Schema{
			"board_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBoradCreate(d *schema.ResourceData, m interface{}) error {
	id := d.Get("boardId").(string)
	d.SetId(id)
	return resourceBoradRead(d, m)
}

func resourceBoradRead(d *schema.ResourceData, m interface{}) error {
	d.Set("boardId", "dummyID")

	return nil
}

func resourceBoradUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("board_id") {
		if err := updateBoard(); err != nil {
			return err
		}
	}

	return resourceBoradRead(d, m)
}

func updateBoard() error {
	return nil
}

func resourceBoradDelete(d *schema.ResourceData, m interface{}) error {
	if err := deleteBoard(); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteBoard() error {
	// return nil if resource already deleted.
	// otherwise run delete on client

	return nil
}
