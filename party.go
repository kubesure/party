package main

import (
	"context"
	"github.com/kubesure/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) CreateParty(ctx context.Context, in *api.PartyRequest) (*api.PartyResponse, error) {
	log.Println(in.Party.FirstName)
	log.Println(in.Party.Gender)
	log.Println(in.Party.Phones[0].Number)
	log.Println(in.Party.Phones[0].Type)
	return &api.PartyResponse{PartyId: 122345}, nil
}

func main() {
	log.Println("party server on")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterPartyServiceServer(s, &server{})
	s.Serve(lis)
}
