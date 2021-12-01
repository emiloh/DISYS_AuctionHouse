genproto:
	protoc --proto_path=./ --go_out=./ --go-grpc_out=./ proto/ah.proto