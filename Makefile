NAME := Topicgram
GO?= go

# Paths & packages
OUT_DIR ?= dist/$(NAME)

.PHONY: help tidy test snapshot release goreleaser-check clean ci-info

all: release


run:
	$(GO) run ./cmd/Topicgram

run-release:
	$(GO) run -ldflags "-X main.version=release" ./cmd/Topicgram

dev: 
	@mkdir -p $(dir $(OUT_DIR))
	$(GO) build -o $(OUT_DIR) ./cmd/Topicgram
	@echo "Build completed: $(OUT_DIR)"


release:
	@mkdir -p $(dir $(OUT_DIR))
	$(GO) build -trimpath -buildvcs=false -ldflags "-s -w -X main.version=release" -o $(OUT_DIR) ./cmd/Topicgram
	@echo "Build completed: $(OUT_DIR)"


clean:
	rm -rf dist


