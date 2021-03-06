package main

import (
	"golang_grpc_gin_jaeger/grpcServer"
	"golang_grpc_gin_jaeger/httpServer"
)

func main() {

	ch := make(chan struct{})

	go grpcServer.Run()
	go httpServer.Run()

	<-ch
}
