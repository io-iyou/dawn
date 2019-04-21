package main

import (
	"log"
	"net"

	impl "github.com/iyou/dawn/internal/impl/service"
	pb "github.com/iyou/dawn/service/sdk/go"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterArticlesServer(s, new(impl.ArticlesImpl))
	pb.RegisterUsersServer(s, new(impl.UsersImpl))
	pb.RegisterMessagesServer(s, new(impl.MessageImpl))

	log.Println("begin..." + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
