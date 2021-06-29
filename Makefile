.PHONY: all generate-go

all: generate-go

generate-go:
	protoc --proto_path=. mdt_dialout.proto --go-grpc_out=./mdt_dialout --go_out=./mdt_dialout 
	protoc --proto_path=. telemetry_bis.proto --go_out=./telemetry_bis
