BINARY_FILE=terraform-provider-trello

build:
	go build -o $(BINARY_FILE)

terraform_init:
	terraform init

plan: terraform_init
	terraform plan

apply: plan
	terraform apply
