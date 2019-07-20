package trello

import (
	"fmt"
	"strings"
	"testing"

	"github.com/adlio/trello"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

// TestAccTrelloListBasic is an integration test for create and delete
func TestAccTrelloListBasic(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTrelloListResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTrelloListResourceBasic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrelloListResourceExists(rName),
					resource.TestCheckResourceAttr("trello_board.test-board", "name", rName),
				),
			},
		},
	})
}

// verify all Trello boards have been destroyed remote
func testAccCheckTrelloListResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Config).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "trello_board" {
			_, err := client.GetBoard(rs.Primary.ID, trello.Defaults())

			if err == nil {
				return fmt.Errorf("Trello board %s still exists", rs.Primary.ID)
			}

			if !strings.Contains(err.Error(), TrelloAPINotFoundMessage) {
				return err
			}
		} else if rs.Type == "trello_list" {
			l, err := client.GetList(rs.Primary.ID, trello.Defaults())
			if !strings.Contains(err.Error(), TrelloSDKNotFoundMessage) {
				return err
			}
		}
	}

	return nil
}

func testAccTrelloListResourceBasic(resourceName string) string {
	return fmt.Sprintf(testAccTrelloListResourceBasicTemplate, resourceName, resourceName)
}

const testAccTrelloListResourceBasicTemplate = `
  resource "trello_board" "test-board" {
	  name        = "%s"
	  description = "A test description."
	}

  resource "trello_list" "test-list" {
	  name     = "tf-test-list-%s"
    board_id = "${trello_board.test-board.id}"
	}
`

func testAccCheckTrelloListResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}
