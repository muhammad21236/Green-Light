package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	fmt.Fprint(w, `"status":"available"`, "\n")
	fmt.Fprint(w, `"system_time":"`, time.Now().Format(time.RFC3339), `"`, "\n")

	fmt.Fprint(w, "environment: ", app.config.env, "\n")
	fmt.Fprint(w, "version: ", version, "\n")
}
