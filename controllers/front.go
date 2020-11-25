package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers register http handlers
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// encodeResponseAsJSON encode response as JSON
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
