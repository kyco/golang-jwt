package main

import (
	"encoding/json"
	"net/http"
)

// validateJWTAPIEndPoint decodes a JWT token
func validateJWTAPIEndPoint(w http.ResponseWriter, req *http.Request) {
	// Setup vars
	var (
		request = StaticStruct{}
	)

	// Try to decode the JSON request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		// Log the error
		log.Error(err.Error())
		response := map[string]string{"error": err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// If we didn't return, validate the strings
	isValid, err := validateJWT(&request.JWT)
	if err != nil {
		log.Error(err.Error())
		response := map[string]string{"error": err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	// If the JWT token is no longer valid
	if isValid == false {
		response := map[string]bool{"jwtValid": isValid}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// If we didn't return anywhere above, then return a successful response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"jwtValid": isValid})
}

// generateJWT decodes a JWT token
func generateJWT(w http.ResponseWriter, req *http.Request) {
	// Setup vars
	var (
		request = LoginDataStruct{}
	)

	// Try to decode the JSON request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		// Log the error
		log.Error(err.Error())
		response := map[string]string{"error": err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// If we didn't return, validate the strings
	isValid := validateLoginData(&request.Username, &request.Password)
	if isValid == false {
		log.Error("Username and password incorrect")
		response := map[string]string{"error": "Username and password incorrect"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate a JWT token
	jwtToken, err := generateJWTToken(&request.Username)
	if err != nil {
		log.Error(err.Error())
		response := map[string]string{"error": err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// If we didn't return anywhere above, then return a successful response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"jwt": jwtToken})
}
