.PHONY: build
build: install

.PHONY: install
install:
	go install ./cmd/...

.PHONY: run
run:
	go run ./cmd/...

.PHONY: test
test: test-lint test-unit

.PHONY: test-lint
test-lint:
	/bin/sh -c ' \
		gofmt -l ./cmd/ ./pkg/ \
	'

.PHONY: test-unit
test-unit:
	/bin/sh -c ' \
		CGO_ENABLED=0 go test \
		-installsuffix "static" \
		./cmd/... ./pkg/... \
	'
