package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	// address for the server. e.g. port 3000
	addr string
	// repository with all connection to the database. We're gonna inject this dependency.
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

/*
Initialize the routher and then register all of services & dependencies
At the end, listen for the server
*/
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// TODO: register our serivces

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
