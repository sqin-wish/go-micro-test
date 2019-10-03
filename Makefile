test:
	protoc -Iapi/proto/v1 --micro_out=api/protoc_gen/v1 --go_out=plugins=grpc:api/protoc_gen/v1 api/proto/v1/*