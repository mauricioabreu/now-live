.PHONY: help ingest package recorder

RECORDER_VERSION=0.0.1

ingest: ## Produce some video and ingest it in the packager
	./ingest.sh

recorder-build: ## Build recorder image
	docker build --tag recorder:${RECORDER_VERSION} -f Dockerfile-recorder .

recorder-run: ## Run recorder
	docker run --name recorder --rm -v MediaVolume:/app/media recorder:${RECORDER_VERSION}

help: ## Lists available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'