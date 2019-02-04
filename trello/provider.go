package trello

import (
	"log"

	trello "github.com/adlio/trello"
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider defines the schema for the Trello provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRELLO_API_KEY", nil),
			},
			"member_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRELLO_MEMBER_ID", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRELLO_TOKEN", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"trello_board": dataSourceTrelloBoard(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// Config is a container for all Trello related configuration
type Config struct {
	Client *trello.Client
	Member *trello.Member
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := trello.NewClient(d.Get("api_key").(string), d.Get("token").(string))

	member, err := client.GetMember(d.Get("member_id").(string), trello.Defaults())

	if err != nil {
		return nil, err
	}

	log.Println("[INFO] Trello client and member initialised")

	return &Config{client, member}, nil
}
