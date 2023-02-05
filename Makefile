PKG       := github.com/ChrisWiegman/goodhosts
VERSION   := $(shell git describe --tags || echo "0.0.1")
TIMESTAMP := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

.PHONY: clean
clean:
	rm -rf \
		dist \
		vendor

.PHONY: install
install:
	go mod vendor
	go install \
		-ldflags "-s -w -X $(PKG)/internal/cmd.Version=$(VERSION) -X $(PKG)/internal/cmd.Timestamp=$(TIMESTAMP)" \
		./cmd/...

.PHONY: lint
lint:
	docker \
		run \
		-t \
		--rm \
		-v $(PWD):/app \
		-w /app \
		golangci/golangci-lint:latest \
		golangci-lint \
			run \
			-v \
			./...

.PHONY: run
run:
	go run ./cmd/...

.PHONY: test
test:
	go \
		test \
		-timeout 30s\
		-cover \
		./...
