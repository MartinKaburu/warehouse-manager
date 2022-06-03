package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type NewServer struct {
	Router *Router
}

//RunServer ...
func RunServer() (err error) {
	var (
		clientPort = flag.Int("client-port", 5000, "Port on which HTTP Client will listen on.")
		serverPort = flag.Int("server-grpc-port", 2000, "Port on which gRPC server is listening on")
	)

	flag.Parse()

	log.Printf("Starting HTTP server at localhost:%d", *clientPort)

	server := NewServer{NewRouter()}
	server.Router.InitializeRoutes(serverPort)

	if err := http.ListenAndServe(fmt.Sprintf("%v:%d", "localhost", *clientPort), *server.Router); err != nil {
		return err
	}

	return nil
}
