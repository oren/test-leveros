package main

import (
	"fmt"
	"log"

	leverapi "github.com/leveros/leveros/api"
)

func main() {
	server, err := leverapi.NewServer()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	err = server.RegisterHandlerObject(new(Handler))
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	server.Serve()
}

type Handler struct {
}

func (*Handler) SayHello(name string) (result string, err error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
