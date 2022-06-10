package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"

	"overengineered-groupmebot/protos"
	"overengineered-groupmebot/services/mvp"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type mvpService struct {
	mvp.UnimplementedMvpServer
}

func (s *mvpService) GetMvpMessage(c context.Context, p *protos.Player) (*protos.Message, error) {
	return &protos.Message{
		Body: fmt.Sprintf("MVP is %s", p.GetName()),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	mvp.RegisterMvpServer(grpcServer, &mvpService{})
	fmt.Println("listening...")
	grpcServer.Serve(lis)
}
