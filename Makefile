.PHONY: clean help ingest package recorder

RECORDER_VERSION=0.0.1

clean: ## Remove media files
	find media -mindepth 1 ! -iname '.keep' -exec rm -r {} +;

ingest: ## Produce some video and ingest it in the packager
	./ingest.sh

package: ## Package content coming from the ingest
	./package.sh

recorder-build: ## Build recorder image
	docker build --tag recorder:${RECORDER_VERSION} -f Dockerfile-recorder .

recorder-run: ## Run recorder
	docker run --name recorder --rm -v $(pwd)/media/:/app/media recorder:${RECORDER_VERSION}

help: ## Lists available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'