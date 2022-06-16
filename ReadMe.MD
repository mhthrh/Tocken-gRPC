**JWT authenticator with gRPC**
----

Introduction
============
Implement JWT with gRPC.

Golang: ver 1.18

PostgresSQL: ver 14.1

Main features
- gRPC JWT (180 seconds).
- Service time out support with context and channel
- ORM connection to DB



StartApplication
============
Server

- Run
    - go run ./Server/main.go
- Build
    - go build ./Server/main.go

Client

- Run
    - go run ./Client/main.go
- Build
    - go build ./Client/main.go


Services
============

SignIn
============

```json
  {
  "username":"mohsen",
  "password":"qaz@123456"
  }
```
- Response
    - Return Valid duration and sing key for reuse in other services

```json
{

    "ValidTill": "Sun Jun 5 13:04:27 BST 2022",
    "Key": "nzTTN4nD7U4kdar7KkHOx9wOPd8zTcoV7eDCZaiITTE/mD//E02QVIcWVno2oWIToCwZPxy/Vot/yqfNmwXr"
  }
```

## By: Mohsen Taheri
## June 2022