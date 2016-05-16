# Level example in Go

https://github.com/leveros/leveros/tree/master/api

```
go get
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./serve server.go
lever deploy
```


Hit the endpoint

```
curl -H "Content-Type: application/json" -X POST -d '["world"]' http://127.0.0.1:8080/helloService/SayHello?forceenv=dev.lever "Hello, world!"
lever invoke /helloService/SayHello '"world"'
LEVEROS_IP_PORT="127.0.0.1:8080" go run client.go
```
