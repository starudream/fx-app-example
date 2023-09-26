PROJECT ?= $(shell basename $(CURDIR))
MODULE  ?= $(shell go list -m)

GO      ?= GO111MODULE=on go
VERSION ?= $(shell git describe --tags 2>/dev/null || echo "dev")
BIDTIME ?= $(shell date +%FT%T%z)

BITTAGS :=
LDFLAGS := -s -w
LDFLAGS += -X "$(MODULE)/version.gitVersion=$(VERSION)"

.PHONY: bin
bin:
	@$(MAKE) bin-app bin-config

.PHONY: bin-%
bin-%:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) build -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' -o bin/example-$* $(MODULE)/example/$*

.PHONY: run-%
run-%:
	@$(MAKE) tidy
	CGO_ENABLED=1 $(GO) run -race -tags '$(BITTAGS)' -ldflags '$(LDFLAGS)' $(MODULE)/example/$*

.PHONY: tidy
tidy:
	@$(GO) mod tidy
