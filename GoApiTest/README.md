# RESTful API with golang
This is restful api server using **gorilla/mux**.  
This repo were using package from **kkamdooong/go-restful-api-example**.

## Install and Run
Clone this repo,
Run this command:
```shell
$ cd $GOPATH/GoApiTest
$ go build
$ go run .
```

## API Endpoint
- http://localhost:3000/api/v1/financings
    - `GET`: get list of financings
    - `POST`: create financing

## Data Structure
```json
{
  "name": "Productive Invoice",
  "count": 15,
  "sub": "invoice"
}
```