lint:
	gofmt -e -d ./
	go vet ./...

test:
	go test ./...