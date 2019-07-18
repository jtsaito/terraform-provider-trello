TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=trello
DIR=~/.terraform.d/plugins

default: build

build: build-linux

build-darwin:
		GO111MODULE=on GOOS=darwin GOARCH=amd64 go install

build-linux:
		GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install

install:
	mkdir -vp $(DIR)
	go build -o $(DIR)/terraform-provider-trello

uninstall:
	@rm -vf $(DIR)/terraform-provider-trello

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

.PHONY: build build-darwin build-linux testacc uninstall vet fmt
