package main

import (
	"log"
	"net/http"
)

func main() {
	// Register the two new handler functions and corresponding route patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()

	// Prefix the route patterns with the required HTTP method (for now, we will
	// restrict all three routes to acting on GET requests).
	mux.HandleFunc("GET /{$}", home)                      // Restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // Add the {id} wildcard segment
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	// Create the new route, which is restricted to POST requests only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
