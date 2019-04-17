package main

import (
	p "github.com/kubesure/party/api/v1"
	service "github.com/kubesure/party/service/v1"
	//"golang.org/x/net/context"
	//"google.golang.org/grpc"
	"log"
	"testing"
)

const (
	address = "localhost:50051"
)

func data() *p.PartyRequest {
	party := p.Party{FirstName: "Gopher bikertales", LastName: "Patel", Gender: p.Party_MALE,
		Email: "pras.p.in@gmail.com", PanNumber: "ABJPP2406G", Aadhaar: 123456789012, DataOfBirth: "14/01/1977",
		AddressLine1: "Ketaki", AddressLine2: "Maneklal", AddressLine3: "Ghatkopar",
		PinCode: 4000086, City: "Mumbai", Latitude: 1212122.333, Longitude: 32232232.33}
	var phones []*p.Party_PhoneNumber
	phone := p.Party_PhoneNumber{Number: "123456789"}
	phone.Type = p.Party_MOBILE
	phones = append(phones, &phone)
	party.Phones = phones

	var req = p.PartyRequest{}
	req.Party = &party
	return &req
}

func TestCreateParty(t *testing.T) {
	svc := service.PartyService{}
	pty, err := svc.CreateParty(nil, data())
	if err != nil && pty.Id < 0 {
		t.Errorf("wanted %b got %s", pty.Id, "Patel")
	}
}

func TestGetParty(t *testing.T) {
	svc := service.PartyService{}
	req := p.PartyRequest{}
	req.Party = &p.Party{}
	req.Party.Id = 7
	party, err := svc.GetParty(nil, &req)
	log.Println("gender ", party.Gender)
	log.Println("phone type ", party.Phones[0].Type)
	if err != nil && party.LastName == "Patel" {
		t.Errorf("wanted %s got %s", "Patel", "nothing")
	}
}

func TestUpdateParty(t *testing.T) {
	svc := service.PartyService{}
	req := data()
	req.Party.Id = 8
	req.Party.FirstName = "Thomas"
	req.Party.LastName = "Earl"
	_, err := svc.UpdateParty(nil, req)
	if err != nil {
		t.Errorf(err.Error(), nil)
	}
}

/* func TestPartyCreateRPC(t *testing.T) {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := p.NewPartyServiceClient(conn)

	party := p.Party{FirstName: "Gopher bikertales", LastName: "Patel", Aadhaar: "12222",
		Gender: p.Party_FEMALE}
	party.Address = "Ghatkopar"
	party.DataOfBirth = "12-09-1856"
	party.Email = "pras.p.in@gmail.com"
	party.PanNumber = "12322232fff"

	var phones []*p.Party_PhoneNumber
	phone := p.Party_PhoneNumber{Number: "123456789"}
	phone.Type = p.Party_MOBILE
	phones = append(phones, &phone)
	party.Phones = phones

	var req = p.PartyRequest{}
	req.Party = &party
	req.Api = "v1"

	res, err := c.CreateParty(context.Background(), &req)
	if err != nil {
		log.Fatalf("error calling: %v", err)
	}

	if res.PartyId < 0 {
		t.Errorf("wanted %b got %s", res.PartyId, "122222")
	}

	log.Println(res.PartyId)
} */
