package main

//go:generate protoc --proto_path=../protos/inventory --go_out=pb/inventory --go-grpc_out=pb/inventory ../protos/inventory/inventory.proto
//go:generate protoc --proto_path=../protos/order --go_out=pb/order --go-grpc_out=pb/order ../protos/order/order.proto

func main() {}
