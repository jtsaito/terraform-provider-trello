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

// TestAccTrelloBoardBasic is an integration test for create and delete
func TestAccTrelloBoardBasic(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTrelloBoardResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTrelloResourceBasic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrelloResourceExists(rName),
					resource.TestCheckResourceAttr("trello_board.test-board", "closed", "false"),
					resource.TestCheckResourceAttr("trello_board.test-board", "description", "A test description."),
					resource.TestCheckResourceAttr("trello_board.test-board", "name", rName),
					resource.TestCheckResourceAttr("trello_board.test-board", "organization_id", ""),
					resource.TestCheckResourceAttr("trello_board.test-board", "pinned", "false"),
					resource.TestCheckResourceAttrSet("trello_board.test-board", "short_url"),
					resource.TestCheckResourceAttrSet("trello_board.test-board", "url"),
				),
			},
		},
	})
}

// verify all Trello boards have been destroyed remote
func testAccCheckTrelloBoardResourceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*Config).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "trello_board" {
			continue
		}

		_, err := client.GetBoard(rs.Primary.ID, trello.Defaults())

		if err == nil {
			return fmt.Errorf("Trello board %s still exists", rs.Primary.ID)
		}

		if !strings.Contains(err.Error(), TrelloAPINotFoundMessage) {
			return err
		}
	}

	return nil
}

func testAccTrelloResourceBasic(resourceName string) string {
	return fmt.Sprintf(testAccTrelloResourceBasicTemplate, resourceName)
}

const testAccTrelloResourceBasicTemplate = `
  resource "trello_board" "test-board" {
	  name        = "%s"
	  description = "A test description."
	}
`

func testAccCheckTrelloResourceExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		member := testAccProvider.Meta().(*Config).Member
		boards, err := member.GetBoards(trello.Defaults())
		if err != nil {
			return fmt.Errorf("Trello board %s does not exist", resourceName)
		}

		var board *trello.Board
		for _, b := range boards {
			if b.Name == resourceName {
				board = b
				break
			}
		}
		if board == nil {
			return fmt.Errorf("check failed: no remote board named %s", resourceName)
		}

		return nil
	}
}
