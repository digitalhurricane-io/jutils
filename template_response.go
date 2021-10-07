package jutils

import (
	"encoding/json"
	"net/http"
)

type templateResponse struct {
	rw 	http.ResponseWriter

	statusCode 		int

	// data to be json encoded
	json 			interface{}
}

// Send Must be called in order to send response
func (j templateResponse) Send() {

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

func newTemplateResponse(rw http.ResponseWriter, statusCode int, json interface{}) templateResponse {
	return templateResponse{
		rw: rw,
		statusCode:     statusCode,
		json:           json,
	}
}