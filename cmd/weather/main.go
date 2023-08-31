package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	weatherv1 "github.com/danny-yamamoto/go-connect-greet/gen/weather/v1"
	"github.com/danny-yamamoto/go-connect-greet/gen/weather/v1/weatherv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type WeatherServer struct{}

func (s *WeatherServer) Weather(
	ctx context.Context,
	req *connect.Request[weatherv1.WeatherRequest],
) (*connect.Response[weatherv1.WeatherResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&weatherv1.WeatherResponse{
		Text: fmt.Sprintf("Today's weather is %s.", req.Msg.Condition),
	})
	res.Header().Set("Weather-Version", "v1")
	return res, nil
}

func main() {
	log.Print("starting server...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	weather := &WeatherServer{}
	mux := http.NewServeMux()
	path, handler := weatherv1connect.NewWeatherServiceHandler(weather)
	mux.Handle(path, handler)
	//	log.Printf("listening on port %s", port)
	http.ListenAndServe(
		":"+port,
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
