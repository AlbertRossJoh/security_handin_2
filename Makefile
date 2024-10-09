PROTO_FILE := grpc/client.proto
OUT_DIR := .

# To run `make generate`
.PHONY: generate
generate:
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative $(PROTO_FILE)
