package main

import (
	"context"
	api "github.com/kubesure/party/api/v1"
	service "github.com/kubesure/party/service/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

const (
	port = ":50051"
)

func main() {
	log.Println("party server on...")
	ctx := context.Background()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	svc := &service.PartyService{}
	api.RegisterPartyServiceServer(s, svc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Print("shutting down party server...")
			s.GracefulStop()
			<-ctx.Done()
		}
	}()
	s.Serve(lis)
}
