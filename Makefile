.PHONY: build
build: install

.PHONY: install
install:
	go install ./cmd/...

.PHONY: run
run:
	go run ./cmd/...
