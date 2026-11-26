package common

import (
	"encoding/json"
	"net/http"
)

// WriteJson function that writes a JSON response to the http.ResponseWriter with the given status code and data
func WriteJson(w http.ResponseWriter, status int, data any) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")
	// Write the status code to the response
	w.WriteHeader(status)
	// Encode the data as JSON and write it to the response body
	json.NewEncoder(w).Encode(data)
}

// ReadJson function that reads a JSON request body from the http.Request and decodes it into the provided data structure
func ReadJson(r *http.Request, data any) error {
	// Decode the JSON request body into the provided data structure
	return json.NewDecoder(r.Body).Decode(data)
}

// WriteError function that writes an error response as JSON to the http.ResponseWriter with the given status code and error message
func WriteError(w http.ResponseWriter, status int, message string) {
	// Call the WriteJson function to write the error response as JSON with the provided status code and error message
	WriteJson(w, status, map[string]string{"error": message})
}
