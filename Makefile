.PHONY: help ingest now-live package recorder

RECORDER_VERSION=0.0.1

ingest: ## Produce some video and ingest it in the packager
	./ingest.sh

now-live: ## Run Now Live platform
	docker-compose build
	docker-compose up

recorder-build: ## Build recorder image
	docker build --tag recorder:${RECORDER_VERSION} -f Dockerfile-recorder .

help: ## Lists available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'