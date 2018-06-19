# gRPC Usage Example

A simple client-server architecture showing how [gRPC](https://grpc.io/) works.

Contains:

- A Server, implemented in Go
- A Client, implemented in Python

## Running

In order to simplify the execution, a convenient Makefile is provided.

That said, [Docker](https://www.docker.com/community-edition) and `make` is all you need to run both server and client.

### Server

Run server first by opening a bash console and executing `make run-server`.

The command will download Go image, get the dependencies, build the application and run it.

### Client

Open a bash console and execute `make run-client`.

The command will download Python 3 image, get dependencies and run the application.

## Protobuf compilation

If you want to apply changes you will have to recompile the `.proto` file.

You will need:

- [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/downloads)
- [Go support for Protocol Buffers](https://github.com/golang/protobuf)

Execution:

Open a bash console and execute `make proto-generation`.

This command will generate both Go and Python code in the right place.