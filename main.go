package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/op/go-logging"
)

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05.999} %{shortfunc} ▶%{color:reset} %{message}`,
)

// Wrapper for http calls to use logging
func wrapHandlerWithLogging(wrappedHandler http.Handler) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		start := time.Now()

		lrw := NewLoggingResponseWriter(c.Writer)
		wrappedHandler.ServeHTTP(lrw, c.Request)

		statusCode := lrw.statusCode
		elapsed := time.Since(start)

		if statusCode == 201 || statusCode == 200 {
			log.Noticef("%s <-- %s %s | HTTP %d %s in %s", c.Request.Header.Get("X-Forwarded-For"), c.Request.Method, c.Request.URL.Path, statusCode, http.StatusText(statusCode), elapsed)
		} else if statusCode == 500 {
			log.Errorf("%s <-- %s %s | HTTP %d %s in %s", c.Request.Header.Get("X-Forwarded-For"), c.Request.Method, c.Request.URL.Path, statusCode, http.StatusText(statusCode), elapsed)
		} else {
			log.Warningf("%s <-- %s %s | HTTP %d %s in %s", c.Request.Header.Get("X-Forwarded-For"), c.Request.Method, c.Request.URL.Path, statusCode, http.StatusText(statusCode), elapsed)
		}

	})
}

// NewLoggingResponseWriter Returns a logging response for the Response Writer
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

// WriteHeader Writes the status code to the header of the resoponse
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// configureLogging configures logging globally
func configureLogging() {

	// Check if it is already open
	logFile.Close()

	// Configure logging
	logFile, err := os.OpenFile("golang-jwt.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Error("error opening file: %v", err.Error())
	}

	// For demo purposes, create two backend for os.Stderr.
	loggingFile := logging.NewLogBackend(logFile, "", 0)

	// For messages written to loggingFile we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	loggingFileFormatter := logging.NewBackendFormatter(loggingFile, format)

	// Set the backends to be used.
	logging.SetBackend(loggingFileFormatter)
}

func main() {

	// Configure logging
	configureLogging()

	// don't forget to close the log file
	defer logFile.Close()

	// staticJWTToken
	validateJWTRouter := wrapHandlerWithLogging(http.HandlerFunc(validateJWTAPIEndPoint))
	loginRoute := wrapHandlerWithLogging(http.HandlerFunc(generateJWT))

	gin.SetMode(gin.ReleaseMode) // Enables release configs
	router := gin.New()

	// Setup route for jwtExample
	jwtExample := router.Group("/jwt")
	{
		jwtExample.POST("/validate/", validateJWTRouter)
		jwtExample.POST("/login/", loginRoute)
	}

	log.Notice("Running on port :8080")

	// Start Server
	router.Run(":8080") // listen and serve
}
