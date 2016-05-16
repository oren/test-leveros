package main

import (
	"log"
	"os"

	leverapi "github.com/leveros/leveros/api"
)

func main() {
	client, err := leverapi.NewClient()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	client.ForceHost = os.Getenv("LEVEROS_IP_PORT")
	leverService := client.Service("dev.lever", "helloService")
	var reply string
	err = leverService.Invoke(&reply, "SayHello", "world")
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	log.Printf("%s\n", reply) // Hello, world!
}
