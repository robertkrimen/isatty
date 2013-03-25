.PHONY: test clean

test:
	$(MAKE) -C example $@

clean:
	$(MAKE) -C example $@

release: test
	godocdown --signature > README.markdown
