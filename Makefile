proto-generation:
	protoc --go_out=plugins=grpc:server *.proto
	python -m grpc_tools.protoc -I . --python_out=client --grpc_python_out=client demorpcservice.proto

run-server:
	docker build -t grpc-usage-example-server server
	docker run --rm --name grpc-server -p 6000:6000 grpc-usage-example-server

run-client:
	docker build -t grpc-usage-example-client client
	docker run --rm --link grpc-server grpc-usage-example-client