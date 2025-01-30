# grpc with golang

Testing out gRPC with Golang.

In order to test it out run the server with 'go run .'. Then open Postman or other testing tool that can use gRPC, upload the proto file and send request with example such as:

```json
{
  "client": {
    "lastName": "ABC",
    "name": "DEF"
  },
  "ticketNumber": "66"
}
```
