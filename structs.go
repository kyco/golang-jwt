package main

import (
	"net/http"
	"os"

	logging "github.com/op/go-logging"
)

// Log file
var logFile *os.File
var log = logging.MustGetLogger("golang-jwt")

// JWTSecretKey is for JWT purposes
var JWTSecretKey = []byte("mySuperSecretSecret")

// LoggingResponseWriter struct for logging http responses and keeping status codes
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// StaticStruct cotains static jwt tokens
type StaticStruct struct {
	JWT string `json:"jwt"`
}

// LoginData contains username and passowrd for logins
type LoginDataStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
