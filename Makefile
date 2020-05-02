.PHONY: clean help

clean: ## Remove media files
	find media -mindepth 1 ! -iname '.keep' -exec rm -r {} \;

help: ## Lists available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'