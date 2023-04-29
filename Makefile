build:
	go build -o bin/

run: build
	./bin/malma-blog-be

vendor:
	go mod vendor