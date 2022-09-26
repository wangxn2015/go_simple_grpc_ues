gen:
	protoc --proto_path=api \
		--go_out=:api \
		--go-grpc_out=:api \
		api/*.proto
	#protoc --proto_path=api --go_out=:api api/*.proto

server:
	go run cmd/server/server.go --port 50051

client:
	go run cmd/client/client.go

clean:
	rm api/*.go


