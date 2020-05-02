.PHONY: clean help ingest package

clean: ## Remove media files
	find media -mindepth 1 ! -iname '.keep' -exec rm -r {} \;

ingest: ## Produce some video and ingest it in the packager
	./ingest.sh

package: ## Package content coming from the ingest
	./package.sh

help: ## Lists available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'