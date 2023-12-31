BUILD_DIR=./.dist
APP_NAME=teman-petani

build:
	@CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	@$(BUILD_DIR)/$(APP_NAME)

test:
	@go test -v -timeout 30s -coverprofile=cover.out -cover ./apis/...
	@go tool cover -html=cover.out