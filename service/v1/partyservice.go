package v1

import (
	"context"
	"github.com/kubesure/party/api/v1"
	"log"
)

type Server struct{}

//SearchParty search a party by party details supplied
func (s *Server) CreateParty(ctx context.Context, in *party.PartyRequest) (*party.PartyResponse, error) {
	log.Println(in.Party.FirstName)
	log.Println(in.Party.Gender)
	log.Println(in.Party.Phones[0].Number)
	log.Println(in.Party.Phones[0].Type)
	in.Party.PartyId = 1234566
	return &party.PartyResponse{Party: in.Party}, nil
}

//SearchParty search a party by party details supplied
func (s *Server) SearchParty(ctx context.Context, in *party.PartyRequest) (*party.PartyResponse, error) {
	log.Println(in.Party.FirstName)
	log.Println(in.Party.Gender)
	log.Println(in.Party.Phones[0].Number)
	log.Println(in.Party.Phones[0].Type)
	in.Party.PartyId = 1234566
	return &party.PartyResponse{Party: in.Party}, nil
}

//UpdateParty updates a party by party details supplied
func (s *Server) UpdateParty(ctx context.Context, in *party.PartyRequest) (*party.PartyResponse, error) {
	log.Println(in.Party.FirstName)
	log.Println(in.Party.Gender)
	log.Println(in.Party.Phones[0].Number)
	log.Println(in.Party.Phones[0].Type)
	in.Party.PartyId = 1234566
	return &party.PartyResponse{Party: in.Party}, nil
}
