cmd_path_server = ./cmd/server.go
cmd_path_client = ./cmd/client.go

build:
	go build $(cmd_path_server)
	go build $(cmd_path_client)

run-server:
	go run $(cmd_path_server)

run-client:
	go run $(cmd_path_client)