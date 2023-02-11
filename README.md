# goCompany
Welcome to the goCompany wiki!

GoCompany is web application.

Backend consists of:
1. REST Middleware (gorilla/mux)
2. Postgres DB (gorm)
3. Payload validator using validator10
4. JWT Authenticator 
5. Kafka event publishing mechanism.
6. TestSuit 


**Pre-Requisites:**

1. Go 1.16 or newer 
2. Postgres DB(with role user/user)
3. Kafka producer with topic: go-company


**Bootstrap infrastructure and run application**
1. start server: 
 go run main.go


After starting server RESTful API server will be running at http://127.0.0.1:8083. It provides the following endpoints:

Company related APIs:
* GET /company/?name=company1         returns a detailed company object (takes query param)
* POST /company                       creates a new company (need to pass request body payload)
* PATCH /company                      updates an existing company( need to pass request body payload)
* DELETE /company/?name=company1      deletes a company (takes query param)

* JWT Auth API:
GET /getJWT                         returns JWT token( takes the Access token in the header)

For POST,PUT,Delete request excepts JWT token in the header, else request will be considered as un-authorized.

Sample payload for create:
    {
        "name"  : "com1",
        "amount" : 113,
        "registered": true  ,
        "type" : "NonProfit",
        "Description": "company1 India"
    }


Sample Outputs:

1. Create API(with token):
<img width="898" alt="Screen Shot 2023-02-11 at 9 57 32 AM" src="https://user-images.githubusercontent.com/23651691/218243076-e236ca01-0722-439a-b735-22424dbb40b4.png">

curl --location --request POST '127.0.0.1:8083/company/' \
--header 'token: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.e30.VHolxFRHrjSRVoPsu9_wFlJ_qH6bXH9ig3J1mL_P5ztovyV90aAluKCWTA3uYAwTXIYxZNBRIuz-21Cke_y6BA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "com1",
    "amount": 113,
    "registered": true,
    "type": "NonProfit",
    "Description": "company1 India"
}'

2. Get Company:

<img width="870" alt="Screen Shot 2023-02-11 at 10 13 05 AM" src="https://user-images.githubusercontent.com/23651691/218243602-7af02468-5f69-4cf0-b2f3-1adbd97cc20b.png">

3. Delete Company:

<img width="917" alt="Screen Shot 2023-02-11 at 10 14 54 AM" src="https://user-images.githubusercontent.com/23651691/218243665-cbfbf7e4-6a34-4983-9d54-46cf81a8c653.png">





****KAFKA published message:****
1. {"EventType":**"company_created"**,"Payload":{"ID":"0cc444a4-d6d1-4793-95a0-387c911f400f","Name":"com1","Description":"company1 India","Amount":113,"Registered":true,"Type":"NonProfit"},"Time":"2023-02-11T09:53:50.838038+04:00"}

2. "EventType":"**company_deleted",**"Payload":{"ID":"","Name":"com1","Description":"","Amount":0,"Registered":false,"Type":""},"Time":"2023-02-11T10:14:42.063397+04:00"}











