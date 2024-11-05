start:
	go build && ./backend

lint:
	$(GOBIN)/golangci-lint run ./...

lint-ci:
	$(GOBIN)/golangci-lint run ./... --out-format=colored-line-number --timeout=5m

generate:
	go generate ./...

test:
	go test ./...

tidy:
	go mod tidy

tidy-ci:
	tidied -verbose