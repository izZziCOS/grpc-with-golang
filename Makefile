generate_grpc_code:
	protoc \
	--go_out=ticketing \
	--go_opt=paths=source_relative \
	--go-grpc_out=ticketing \
	--go-grpc_opt=paths=source_relative \
	ticket.proto