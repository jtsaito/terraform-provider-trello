package main

import "github.com/hashicorp/terraform/helper/schema"
	"log"


// Provider defines the schema for the Trello provider
func Provider() *schema.Provider {
	return &schema.Provider{
		},
		DataSourcesMap: map[string]*schema.Resource{
			"trello_board": dataSourceTrelloBoard(),
		},
	}
}
