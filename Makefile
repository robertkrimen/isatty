.PHONY: test clean

test:
	$(MAKE) -C test $@

clean:
	$(MAKE) -C test $@

release: test
	godocdown --signature > README.markdown
