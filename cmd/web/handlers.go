package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Use the Header().Add() method to add a 'Server: Go' header to the
	// response header map. The first parameter is the header name, and
	// the second parameter is the header value.
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// and try to convert it to an integer using the strconv.Atoi() function. If
	// it can't be converted to an integer, or the value is less than 1, we
	// return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the w.WriteHeader() method to send a 201 status code.
	w.WriteHeader(http.StatusCreated)

	// Use the fmt.Sprintf() function to interpolate the id value with a
	// message, then write it as the HTTP response.
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Add a snippetCreatePost handler function.
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet..."))
}
