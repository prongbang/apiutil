build:
	mkdir -p functions
	go get ./...
	go build -o functions/apiutil ./...
	./functions/apiutil