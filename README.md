## terraform-provider-trello

This is a Terraform provider for Trello.

## Requirements

This implementation requires Go v1.12 and uses [go Modules](https://github.com/golang/go/wiki/Modules) for packaging.

Currently, package `github.com/armon/go-radix` seems to cause problems for `github.com/hashicorp/terraform v0.12`. Therefore, we use the older `github.com/hashicorp/terraform v0.11.5`.

## Acceptance tests

Acceptance tests are run like this:

```
TRELLO_MEMBER_ID=<your-user-id> TRELLO_TOKEN=<your-token> TRELLO_API_KEY=<your-api-ke> make testacc
```
