PROTO_DIR=protobuf
GEN_DIR=common/genproto/orders

PROTO_FILES=$(PROTO_DIR)/orders.proto

gen:
	mkdir -p $(GEN_DIR)
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

clean:
	rm -f $(GEN_DIR)/*.pb.go