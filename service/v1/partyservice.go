package v1

import (
	"context"
	party "github.com/kubesure/party/api/v1"
	"log"
)

//PartyService to host a party service
type PartyService struct{}

//CreateParty creates a party
func (s *PartyService) CreateParty(ctx context.Context, request *party.PartyRequest) (*party.Party, error) {
	log.Println(request.Party.FirstName)
	log.Println(request.Party.Gender)
	log.Println(request.Party.Phones[0].Number)
	log.Println(request.Party.Phones[0].Type)
	request.Party.Id = 1234566
	return request.Party, nil
}

//GetParty gets an individual party
func (s *PartyService) GetParty(ctx context.Context, request *party.PartyRequest) (*party.Party, error) {
	log.Println(request.Party.FirstName)
	log.Println(request.Party.Gender)
	log.Println(request.Party.Phones[0].Number)
	log.Println(request.Party.Phones[0].Type)
	request.Party.Id = 1234566
	return request.Party, nil
}

//UpdateParty updates an individual party
func (s *PartyService) UpdateParty(ctx context.Context, request *party.PartyRequest) (*party.Party, error) {
	log.Println(request.Party.FirstName)
	log.Println(request.Party.Gender)
	log.Println(request.Party.Phones[0].Number)
	log.Println(request.Party.Phones[0].Type)
	request.Party.Id = 1234566
	return request.Party, nil
}