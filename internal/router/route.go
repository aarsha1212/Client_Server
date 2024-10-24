package router

import (
	"fmt"
	"log"
	"net/http"
	handler "p1/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r:= mux.NewRouter()
	fmt.Printf("Created router \n")
	return r
}

func StarctClientServer(clienthandler *handler.ClientHandler, r *mux.Router) {
	fmt.Printf("starting client server \n")
	r.HandleFunc("/blacklist", clienthandler.HandleInitClient).Methods("POST")
}
func StartServer(router *mux.Router,) {
	// Start the HTTP server
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}