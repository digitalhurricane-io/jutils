package jutils

import (
	"fmt"
	"net/http"
)

// jSender Easily send common json responses
type jSender struct {
	rw 	http.ResponseWriter

	Templates Template
}

// Json The data to be encoded as json in the response
func (s jSender) Json(data interface{}) standardResponse {
	return newStandardResponse(s.rw, 200, data)
}

// Success 200 status {"success": true}
func (s jSender) Success() standardResponse {
	return newStandardResponse(s.rw, 200, map[string]bool{"success": true})
}

// Error Set custom error message {"error": message}
func (s jSender) Error(statusCode int, message string) standardResponse {
	return newStandardResponse(s.rw, statusCode, fmtErr(message))
}

// Errorf Set custom formatted error message {"error": message}
// If no status code is set upon calling Send(), status code
// will automatically be set to 400
func (s jSender) Errorf(statusCode int, format string, a ...interface{}) standardResponse {
	errMsg := fmtErr(fmt.Sprintf(format, a...))
	return newStandardResponse(s.rw, statusCode, errMsg)
}

func Response(w http.ResponseWriter) jSender {
	w.Header().Set("Content-Type", "application/json")
	return jSender{rw: w, Templates: Template{rw: w}}
}

func fmtErr(msg string) map[string]string {
	return map[string]string{"error": msg}
}