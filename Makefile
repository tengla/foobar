
.PHONY: clean

build-kv:
	go build -o dist/kv cmd/kv/main.go

build-server:
	go build -o dist/server cmd/server/main.go

clean:
	rm -rf dist
