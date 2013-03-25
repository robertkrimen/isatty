.PHONY: test build clean

test: build
	./example
	./example < /dev/null; [ $$? -eq 1 ]


build:
	go build -v

clean:
	rm -f example
