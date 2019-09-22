# party


#### biz design

Creates parties insured and nominee by calling an internal party grpc service. 

#### components

Mongodb v4, GRPC, Golang 

#### Dev setup and test

1. create db party 
    ```
       use parties
       db.counter.insert({"_id" : "partyid" , "value": 0 })
       db.counter.find({}).pretty()
       db.party.find({}).pretty()
    ```
2. Run party and quote 
   ``` go run ../quote/quote.go
       go run party.go
   ```

3. Run curl to create quote and party

```
curl -i -X POST http://localhost:8000/api/v1/healths/quotes  -H 'Content-Type: application/json' -d '{
    "code": "1A",
    "SumInsured": 12000,
    "dateOfBirth" : "14/01/1977",
    "Premium": 3000,
    "parties": [
        {
            "firstName": "Bhavesh",
            "lastName": "Yadav",
            "gender": "MALE",
            "dateOfBirth": "14/01/1977",
            "mobileNumber": "1234567890",
            "email": "primary@gmail.com",
            "panNumber": "AJBDD12345G",
            "aadhaar": 123456789012,
            "addressLine1": "ketaki",
            "addressLine2": "maneklal",
            "addressLine3": "Ghatkopar",
            "city": "mumbai",
            "pinCode": 400086,
            "latitude": 123223232,
            "longitude": 12345643,
            "relationship": "self",
            "isPrimary": true
        },
        {
            "firstName": "Usha",
            "lastName": "Patel",
            "gender": "FEMALE",
            "dateOfBirth": "14/01/1977",
            "mobileNumber": "1234567890",
            "email": "nominiee@gmail.com",
            "panNumber": "AJBDD12345G",
            "aadhaar": 123456789012,
            "addressLine1": "ketaki",
            "addressLine2": "maneklal",
            "addressLine3": "Ghatkopar",
            "city": "mumbai",
            "pincode": 400086,
            "latitude": 123223232,
            "Longitude": 12345643,
            "relationship": "self",
            "IsPrimary": false
        }
    ]
}'
```

rs.initiate({ _id: "rs0", members:[ 
        { _id: 0, host: "mongo-party-0.mongopartysvc:27017" },
        { _id: 1, host: "mongo-party-1.mongopartysvc:27017" },
        { _id: 2, host: "mongo-party-2.mongopartysvc:27017" },
]});
