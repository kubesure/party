package main

import (
	p "github.com/kubesure/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

const (
	address = "localhost:50051"
)

func TestPartyCreate(t *testing.T) {

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
}

/* func TestMarshalPropsal(t *testing.T) {
	pp := Party{FirstName: "Prashant", LastName: "Patel",
		MobileNumber: "1234567890", Email: "primary@gmail.com",
		PanNumber: "AJBDD12345G", Relationship: "self", Gender: "m",
		Address: "Mumbai", DataOfBirth: "14/01/1977", SumInsured: 12000}
	mother := Party{FirstName: "Usha", LastName: "Patel",
		MobileNumber: "1234567890", Email: "Nominee@gmail.com",
		PanNumber: "AJBDD12345G", Relationship: "mother", Gender: "f", Address: "Mumbai",
		DataOfBirth: "19/01/1956", SumInsured: 12000}
	proposal := Proposal{Premium: 3000, Pid: "12345678"}
	proposal.Party = append(proposal.Party, pp)
	proposal.Party = append(proposal.Party, mother)
	data, _ := json.Marshal(proposal)
	log.Println(string(data))
} */
