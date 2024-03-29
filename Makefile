PKG       := github.com/ChrisWiegman/goodhosts/v4
VERSION   := $(shell git describe --tags || echo "0.0.1")
TIMESTAMP := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

.PHONY: change
change:
	docker run \
		--rm \
		--platform linux/amd64 \
		--mount type=bind,source=$(PWD),target=/src \
		-w /src \
		-it \
		ghcr.io/miniscruff/changie \
		new

.PHONY: changelog
changelog:
	docker run \
		--rm \
		--platform linux/amd64 \
		--mount type=bind,source=$(PWD),target=/src \
		-w /src \
		-it \
		ghcr.io/miniscruff/changie \
		batch $(call ARGS,defaultstring)
	docker run \
		--rm \
		--platform linux/amd64 \
		--mount type=bind,source=$(PWD),target=/src \
		-w /src \
		-it \
		ghcr.io/miniscruff/changie \
		merge

.PHONY: clean
clean:
	rm -rf \
		dist \
		vendor

.PHONY: install
install:
	go mod vendor
	go install \
		-ldflags "-s -w -X $(PKG)/internal/commands.Version=$(VERSION) -X $(PKG)/internal/commands.Timestamp=$(TIMESTAMP)" \
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

.PHONY: snapshot
snapshot:
	docker run --rm \
	--privileged \
	-v $(PWD):/go/src/$(PKG) \
	-w /go/src/$(PKG) \
	goreleaser/goreleaser \
		release \
		--clean \
		--release-notes=./.changes/$(VERSION).md \
		--snapshot

.PHONY: test
test:
	go \
		test \
		-timeout 30s\
		-cover \
		./...
