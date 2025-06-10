build:
	@go build -o ./bin/pricefetcher

run: build 
	@./bin/pricefetcher

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		protobuf/service.proto

.PHONY: proto