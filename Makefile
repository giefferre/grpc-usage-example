proto-generation:
	protoc --go_out=plugins=grpc:server *.proto
	python -m grpc_tools.protoc -I . --python_out=client --grpc_python_out=client demorpcservice.proto
