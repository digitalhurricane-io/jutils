package jutils

import (
	"encoding/json"
	"net/http"
)

// standardResponse differs from a TemplateResponse in that it has a Status method
type standardResponse struct {
	rw 	http.ResponseWriter

	statusCode 		int

	// data to be json encoded
	json 			interface{}
}

// Send Must be called in order to send response
func (j standardResponse) Send() {

	if j.statusCode == 0 {
		j.statusCode = 400
	}

	j.rw.WriteHeader(j.statusCode)

	if j.json == nil {
		// No response body
		return
	}

	json.NewEncoder(j.rw).Encode(j.json)
}

func (j standardResponse) Status(statusCode int) standardResponse {
	return newStandardResponse(j.rw, statusCode, j.json)
}

func newStandardResponse(rw http.ResponseWriter, statusCode int, json interface{}) standardResponse {
	return standardResponse{
		rw: rw,
		statusCode:     statusCode,
		json:           json,
	}
}