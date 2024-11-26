module grpc_server

go 1.22.7

toolchain go1.23.3

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/matsuridayo/libneko v0.0.0-20240702024904-1c47a3af7199
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
)

replace cloud.google.com/go => cloud.google.com/go v0.115.1
