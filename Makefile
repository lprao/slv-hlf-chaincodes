.PHONY: all
all: slvint

.PHONY: dep
dep: ## Get the dependencies
	@go get -v -d ./...

.PHONY: slvint
slvint: ## Build slvint chaincode
	@CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-extldflags -static" -a -o bin/slvint ./cmd/slv_int_cc.go

.PHONY: clean
clean: ## cleanup
	@rm bin/slvint

.PHONY: test
test:
	@go test -v -coverprofile=cover.out ./...

.PHONY: show_coverage
show_coverage:
	@go tool cover -html=cover.out

help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
