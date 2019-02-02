package main

import "github.com/hashicorp/terraform/helper/schema"

// Provider defines the schema for the Trello provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
		},
		// ConfigureFunc: providerConfigure,
	}
}
