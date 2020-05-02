.PHONY: clean

clean:
	find media -mindepth 1 ! -iname '.keep' -exec rm -r {} \;