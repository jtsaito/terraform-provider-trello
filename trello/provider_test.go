package trello

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const TrelloAPINotFoundMessage = "The requested resource was not found."
const TrelloSDKNotFoundMessage = "404: model not found"

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"trello": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

// validates the test API keys exist for testing
func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TRELLO_API_KEY"); v == "" {
		t.Fatal("TRELLO_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("TRELLO_MEMBER_ID"); v == "" {
		t.Fatal("TRELLO_MEMBER_ID must be set for acceptance tests")
	}
	if v := os.Getenv("TRELLO_TOKEN"); v == "" {
		t.Fatal("TRELLO_TOKEN must be set for acceptance tests")
	}
}
