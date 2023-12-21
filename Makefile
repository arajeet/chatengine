build:
	 /usr/local/go/bin/go build -o ./bin/chatengine

run: build
	./bin/chatengine

test: /usr/local/go/bin/go test -o ./...

clean:
	rm -rf ./bin/*