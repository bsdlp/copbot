lint:
	gofmt -e -d ./
	go vet ./...

test:
	go test ./...

proto:
	protoc --proto_path=./commands --go_out=./commands ./commands/*.proto
