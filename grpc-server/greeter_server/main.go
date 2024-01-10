package main

import (
	"context"
	"fmt"
	"hello-grpc/greeter_server/entities"
	pb "hello-grpc/sayhello"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
	db *gorm.DB
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	var u entities.User
	s.db.Model(&u).Where("name = ?", in.GetName()).Find(&u)

	ret := fmt.Sprintf("Say Hello " + in.GetName() + "has no Email")
	if u.ID != 0 {
		ret = fmt.Sprintf("Say Hello "+in.GetName()+"has Email: %v", u.Email)
	}
	return &pb.HelloReply{Message: ret}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{
		db: entities.InitDB(),
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
