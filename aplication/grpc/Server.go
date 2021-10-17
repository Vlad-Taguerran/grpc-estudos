package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartServerGRPC() {
	list, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Não foi possivel conectar ao servidor", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	log.Println("Servidor iniciado")

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Impossivel se conectar", err)
	}
}
