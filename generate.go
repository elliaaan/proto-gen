package main

//go:generate protoc --proto_path=../protos/inventory --go_out=. --go-grpc_out=. ../protos/inventory/inventory.proto
//go:generate protoc --proto_path=../protos/order --go_out=. --go-grpc_out=. ../protos/order/order.proto

func main() {}
