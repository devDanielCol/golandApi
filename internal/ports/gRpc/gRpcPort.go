package grpcPort

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}
type Response struct {
	Message string
}

func (s *Server) YourGRPCMethod(ctx context.Context, request *Response) (*Response, error) {
	// Lógica de tu adaptador aquí
	response := &Response{
		Message: "Respuesta desde el adaptador gRPC",
	}

	return response, nil
}

func InitRpc() {

	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error al escuchar en el puerto %s: %v", port, err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar el servidor gRPC: %v", err)
	}
}
