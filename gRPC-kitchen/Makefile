run-orders:
	@go run $(wildcard services/orders/*.go)

run-kitchen:
	@go run $(wildcard services/kitchen/*.go)

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders
		--go-grpc_opt=paths=source_realtive