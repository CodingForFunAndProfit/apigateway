proto:
	protoc pkg/**/pb/*.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --proto_path=.

server:
	go run cmd/main.go