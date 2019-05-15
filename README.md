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
    ```
2. Run party and quote 
   ``` go run ../quote/quote.go
       go run party.go
   ```

3. Run curl to create quote and party

```
curl -i -X POST \
  http://localhost:8000/api/v1/healths/quotes   \
  -d '{                                         \
    "code": "1A",                               \
    "SumInsured": 12000,                        \
    "Premium": 3000,                            \
    "parties": [                                \
        {                                       \
            "firstName": "Bhavesh",             \
            "lastName": "Yadav",                \ 
            "gender": "MALE",                   \ 
            "dataOfBirth": "14/01/1977",        \ 
            "mobileNumber": "1234567890",       \
            "email": "primary@gmail.com",       \
            "panNumber": "AJBDD12345G",         \
            "aadhaar": 123456789012,            \
            "addressLine1": "ketaki",           \ 
            "addressLine2": "maneklal",         \
            "addressLine3": "Ghatkopar",        \ 
            "city": "mumbai",                   \
            "pinCode": 400086,                  \    
            "latitude": 123223232,              \
            "longitude": 12345643,              \
            "relationship": "self",             \
            "isPrimary": true                   \ 
        },                                      \
        {                                       \
            "firstName": "Usha",                \
            "lastName": "Patel",                \
            "gender": "FEMALE",                 \
            "dataOfBirth": "14/01/1977",        \
            "mobileNumber": "1234567890",       \
            "email": "nominiee@gmail.com",      \ 
            "panNumber": "AJBDD12345G",         \
            "aadhaar": 123456789012,            \
            "addressLine1": "ketaki",           \
            "addressLine2": "maneklal",         \
            "addressLine3": "Ghatkopar",        \
            "city": "mumbai",                   \
            "pincode": 400086,                  \
            "latitude": 123223232,              \
            "Longitude": 12345643,              \
            "relationship": "self",             \
            "IsPrimary": false                  \
        }                                       \
    ]                                           \
}                                               \'
```

