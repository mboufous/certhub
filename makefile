SHELL := /bin/bash

SERVICE_NAME 	:= certhub
# GO_FLAGS   		?= CGO_ENABLED=0 GOOS=linux GOARCH=amd64
OUTPUT_BIN 		?= bin/${SERVICE_NAME}
PACKAGE    		:= github.com/mboufous/$(SERVICE_NAME)
VERSION    		?= v1.0
IMAGE      		:= ${SERVICE_NAME}:${VERSION}

default: help

run:
	@go run cmd/service/main.go

test: # Run all tests
	go clean --testcache && go test ./... -count=1 --race

clean: ## Clean compiled service
	rm ${OUTPUT_BIN} || true

tidy: # tidy
	go mod tidy
	go mod vendor

.PHONY: build
build:  ## Builds the service
	@${GO_FLAGS} go build -o ${OUTPUT_BIN} \
	-ldflags "-w -s -X ${PACKAGE}/main.build=${VERSION}" \
	-a main.go

#--------------------------------
# FrontEnd

.PHONY: frontend
frontend: ## Start frontapp react app
	npm run dev --prefix frontend/

#--------------------------------
# Docker
img: clean build ## Build Docker image
	docker build \
		-f build/Dockerfile.dev \
		-t $(IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
up:
	@docker-compose up -d
down:
	@docker-compose down
	
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;38m %s\033[0m\n", $$1, $$2}'
