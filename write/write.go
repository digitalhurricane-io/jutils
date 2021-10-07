package write

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// jResponser Easily send json responses
type jResponser struct {
	rw 	http.ResponseWriter

	Templates Template
}

// Json 200 status. Data is encoded as json in the response
func (s jResponser) Json(data interface{}) senderStatusSetter {
	return newResponseStatusSender(s.rw, 200, data)
}

// Success 200 status {"success": true}
func (s jResponser) Success() senderStatusSetter {
	return newResponseStatusSender(s.rw, 200, map[string]bool{"success": true})
}

// Error Set custom error message {"error": message}
func (s jResponser) Error(statusCode int, message string) sender {
	return newResponseSender(s.rw, statusCode, fmtErr(message))
}

// Errorf Set custom formatted error message {"error": message}
func (s jResponser) Errorf(statusCode int, format string, a ...interface{}) sender {
	errMsg := fmtErr(fmt.Sprintf(format, a...))
	return newResponseSender(s.rw, statusCode, errMsg)
}

func Response(w http.ResponseWriter) jResponser {
	w.Header().Set("Content-Type", "application/json")
	return jResponser{rw: w, Templates: Template{rw: w}}
}

func fmtErr(msg string) map[string]string {
	return map[string]string{"error": msg}
}

func sendJson(w http.ResponseWriter, statusCode int, data interface{}) {
	if statusCode == 0 {
		statusCode = 400
	}

	w.WriteHeader(statusCode)

	if data == nil {
		// No response body
		return
	}

	json.NewEncoder(w).Encode(data)
}