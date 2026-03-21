LOCAL_BIN:=$(CURDIR)/bin
APP_NAME:=chat-service
BUILD_DIR:=$(LOCAL_BIN)

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

$(LOCAL_BIN)/golangci-lint: install-golangci-lint
	@:

lint: $(LOCAL_BIN)/golangci-lint
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

lint-fix: $(LOCAL_BIN)/golangci-lint
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml --fix

vet:
	go vet ./...

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

test:
	go test -v ./...

.PHONY: lint lint-fix vet build run test install-golangci-lint
