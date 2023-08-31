package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	greetv1 "github.com/danny-yamamoto/go-connect-greet/gen/greet/v1"
	"github.com/danny-yamamoto/go-connect-greet/gen/greet/v1/greetv1connect"
)

func main() {
	log.Print("starting client...")
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:8080"
	}
	log.Printf("defaulting to BASE_URL %s", baseUrl)
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		baseUrl,
		connect.WithGRPC(),
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
