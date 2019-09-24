BUILD_DIR			:= build
CC					:= go build
DFLAGS				:= -race
CROSS				:= GOOS=linux GOARCH=amd64
NAME				:= boilerplate
GITORG				:= github.com/stevenleroux

GITHASH				:= $(shell git rev-parse HEAD)
VERSION				:= $(shell git describe --tags --candidates 1)
BUILDDATE			:= $(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')
CFLAGS				:= -X $(GITORG)/$(NAME)/core.Githash=$(GITHASH) -X $(GITORG)/$(NAME)/core.Version=$(VERSION) -X $(GITORG)/$(NAME)/core.BuildDate=$(BUILDDATE)

rwildcard			:= $(foreach d,$(wildcard $1*),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))
MODULES_PATHS	:= $(call rwildcard, ./cmd, *.go) $(call rwildcard, ./core, *.go) $(call rwildcard, ./services, *.go)
LINT_PATHS		:= $(NAME).go ./core/... ./cmd/... ./services/...
FORMAT_PATHS	:= $(NAME).go ./core ./cmd ./services
VPATH					:= $(BUILD_DIR)

.SECONDEXPANSION:

.PHONY: all
all: init dep format lint test release

.PHONY: build
build: $(NAME).go $(MODULES_PATHS)
	$(CC) $(DFLAGS) -ldflags "$(CFLAGS)" -o $(BUILD_DIR)/$(NAME) $(NAME).go

.PHONY: init
init:
	curl -s https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	go get -u github.com/alecthomas/gometalinter

.PHONY: release
release: $(NAME).go $(MODULES_PATHS)
	$(CC) -ldflags "$(CFLAGS)" -o $(BUILD_DIR)/$(NAME) $(NAME).go

.PHONY: dist
dist: $(NAME).go $(MODULES_PATHS)
	$(CROSS) $(CC) -ldflags "$(CFLAGS) -s -w" -o $(BUILD_DIR)/$(NAME) $(NAME).go

.PHONY: lint
lint:
	gometalinter --config .gometalinter.json $(LINT_PATHS)

.PHONY: format
format:
	gofmt -w -s $(FORMAT_PATHS) $(NAME).go

.PHONY: test
test:
	go test -v ./...

.PHONY: dev
dev: format lint build

.PHONY: mod
mod:
	gometalinter --install
	dep ensure -v

.PHONY: clean
clean:
	rm -r -v build