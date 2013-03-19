.PHONY: test build clean

test: build
	./test
	./test < /dev/null; [ $$? -eq 1 ]


build:
	go build -v

clean:
	rm -f test
