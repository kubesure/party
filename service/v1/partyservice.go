package v1

import (
	"context"
	"fmt"
	"os"
	"time"

	party "github.com/kubesure/party/api/v1"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type partyrec struct {
	FirstName    string `bson:"firstName"`
	LastName     string `bson:"lastName"`
	Gender       int    `bson:"gender"`
	DataOfBirth  string `bson:"dateOfBirth"`
	Email        string `bson:"email"`
	MobileNumber string `bson:"mobileNumber"`
	AddressLine1 string `bson:"addressLine1"`
	AddressLine2 string `bson:"addressLine2"`
	AddressLine3 string `bson:"addressLine3"`
	PinCode      int32  `bson:"pinCode"`
	City         string `bson:"city"`
	PanNumber    string `bson:"panNumber"`
	Aadhaar      int64  `bson:"aadhaar"`
}

//PartyService to host a party service
type PartyService struct{}

var mongopartysvc = os.Getenv("mongopartysvc")

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
}

func encode(request *party.PartyRequest) bson.M {
	return bson.M{
		"partyId": request.Party.Id, "firstName": request.Party.FirstName, "lastName": request.Party.FirstName,
		"gender": request.Party.Gender, "email": request.Party.Email, "dateOfBirth": request.Party.DataOfBirth,
		"mobileNumber": request.Party.Phones[0].Number, "addressLine1": request.Party.AddressLine1,
		"addressLine2": request.Party.AddressLine2, "addressLine3": request.Party.AddressLine3,
		"pinCode": request.Party.PinCode, "city": request.Party.City, "latitude": request.Party.Latitude,
		"longitude": request.Party.Longitude, "panNumber": request.Party.PanNumber,
		"aadhaar": request.Party.Aadhaar, "createdDate": time.Now().String()}
}

//CreateParty creates a party
func (s *PartyService) CreateParty(ctx context.Context, request *party.PartyRequest) (*party.Party, error) {
	client, err := conn()
	defer client.Disconnect(context.Background())

	id, err := nextid(client)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	request.Party.Id = int64(id)
	rec := encode(request)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	coll := client.Database("parties").Collection("party")
	_, errcol := coll.InsertOne(context.Background(), rec)

	if errcol != nil {
		log.Error(errcol)
		return nil, errcol
	}
	return request.Party, nil
}

//GetParty gets an individual party
func (s *PartyService) GetParty(ctx context.Context, request *party.PartyRequest) (*party.Party, error) {
	client, err := conn()
	defer client.Disconnect(context.Background())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	coll := client.Database("parties").Collection("party")
	filter := bson.M{"partyId": request.Party.Id}
	rec := partyrec{}
	result := coll.FindOne(context.Background(), filter)
	errdecode := result.Decode(&rec)
	if errdecode != nil {
		log.Error(err)
		return nil, errdecode
	}
	request.Party.FirstName = rec.FirstName
	request.Party.LastName = rec.LastName
	request.Party.DataOfBirth = rec.DataOfBirth
	request.Party.Email = rec.Email
	request.Party.AddressLine1 = rec.AddressLine1
	request.Party.AddressLine2 = rec.AddressLine2
	request.Party.AddressLine3 = rec.AddressLine3
	request.Party.City = rec.City
	request.Party.PinCode = int32(rec.PinCode)
	request.Party.PanNumber = rec.PanNumber
	request.Party.Aadhaar = rec.Aadhaar
	if rec.Gender == 0 {
		request.Party.Gender = party.Party_MALE
	}
	if rec.Gender == 1 {
		request.Party.Gender = party.Party_FEMALE
	}
	var phones []*party.Party_PhoneNumber
	phone := party.Party_PhoneNumber{Number: rec.MobileNumber}
	phone.Type = party.Party_MOBILE
	phones = append(phones, &phone)
	request.Party.Phones = phones
	return request.Party, nil
}

//UpdateParty updates an individual party
func (s *PartyService) UpdateParty(ctx context.Context, request *party.PartyRequest) (*party.Party, error) {
	client, err := conn()
	defer client.Disconnect(context.Background())
	if err != nil {
		log.Error(err)
		return nil, err
	}

	filter := bson.D{
		{"partyId", request.Party.Id},
	}
	data := bson.D{{"$set", bson.D{
		{"firstName", request.Party.FirstName},
		{"lastName", request.Party.LastName},
		{"gender", request.Party.Gender},
		{"email", request.Party.Email},
		{"dateOfBirth", request.Party.DataOfBirth},
		{"mobileNumber", request.Party.Phones[0].Number},
		{"addressLine1", request.Party.AddressLine1},
		{"addressLine2", request.Party.AddressLine2},
		{"addressLine3", request.Party.AddressLine3},
		{"pinCode", request.Party.PinCode},
		{"city", request.Party.City},
		{"latitude", request.Party.Latitude},
		{"longitude", request.Party.Longitude},
		{"panNumber", request.Party.PanNumber},
		{"aadhaar", request.Party.Aadhaar},
		{"updatedDate", time.Now().String()},
	}}}

	coll := client.Database("parties").Collection("party")
	result, errupdate := coll.UpdateOne(context.Background(), filter, data)

	if errupdate != nil {
		log.Error(err)
		return nil, errupdate
	}
	if result.ModifiedCount == 0 {
		log.Error("party not updated")
		return nil, fmt.Errorf("party not updated")
	}
	return request.Party, nil
}

func nextid(c *mongo.Client) (int, error) {
	coll := c.Database("parties").Collection("counter")
	filter := bson.M{"_id": "partyid"}
	update := bson.M{"$inc": bson.M{"value": 1}}
	aft := options.After
	opt := options.FindOneAndUpdateOptions{Upsert: new(bool), ReturnDocument: &aft}
	result := coll.FindOneAndUpdate(context.Background(), filter, update, &opt)
	type record struct {
		PartyID string `bson:"partytid"`
		Value   int    `bson:"value"`
	}
	var data record
	errdecode := result.Decode(&data)
	if errdecode != nil {
		log.Error(errdecode)
		return 0, errdecode
	}
	return data.Value, nil
}

func conn() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+mongopartysvc+":27017"))
	err := client.Ping(ctx, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return client, nil
}
