package main

import (
	"fmt"
	"http_server/server"
	"http_server/validator"
	"log"
	"net/http"
	"time"
)

const port = ":8080"

func main() {
	fmt.Printf("Starting server on port %s...\n", port)

	constraints := []validator.Constraint{validator.User_Agent, validator.Accept_Header}
	validator := validator.New(constraints)

	s := &http.Server{
		Addr:         port,
		Handler:      server.New(validator),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
