package write

import (
	"net/http"
)

type responseSender struct {
	rw 	http.ResponseWriter

	statusCode 		int

	// data to be json encoded
	data 			interface{}
}

// Send Must be called in order to send response
func (j responseSender) Send() {
	sendJson(j.rw, j.statusCode, j.data)
}

func newResponseSender(rw http.ResponseWriter, statusCode int, data interface{}) sender {
	return responseSender{
		rw: rw,
		statusCode:     statusCode,
		data:           data,
	}
}