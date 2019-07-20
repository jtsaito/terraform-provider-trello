## terraform-provider-trello

This is a Terraform provider for Trello.

## Requirements

This implementation requires Go >= v1.11 and uses [go Modules](https://github.com/golang/go/wiki/Modules) for packaging.

Currently, package `github.com/armon/go-radix` seems to cause problems for `github.com/hashicorp/terraform v0.12`. Therefore, we use the older `github.com/hashicorp/terraform v0.11.5`.

## Usage

The Trello provider requires three environment variables for authenticating with the Trello API: `TRELLO_API_KEY`, `TRELLO_MEMBER_ID`, `TRELLO_TOKEN`.
(Alternatively, you can set the values in the provider block.) The [Trello developers' documentation](https://trello.com/app-key/) explains how to obtain the set of credentials.

### Example

Below is an example using using a Trello board and list.

```
provider "trello" {
}

resource "trello_board" "my-board" {
  name        = "My Board"
  description = "This is my Trello board"
}

resource "trello_list" "my-backlog" {
  name     = "My Backlog"
  board_id = "${trello_board.my-board.id}"
}
```

## Acceptance tests

Acceptance tests are run like this:

```
TRELLO_MEMBER_ID=<your-user-id> TRELLO_TOKEN=<your-token> TRELLO_API_KEY=<your-api-ke> make testacc
```

## Releasing a version

Release builds are automatically created and uploaded from Travis CI to GitHub by [Travis CI integration](https://github.com/jtsaito/terraform-provider-trello/blob/master/.travis.yml). To trigger the build, just [create a release](https://github.com/jtsaito/terraform-provider-trello/releases/new). The release name has to match the semantic version naming scheme.
