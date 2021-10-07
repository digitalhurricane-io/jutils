package write

import "net/http"

// responseStatusSender differs from a responseSender in that it has a Status method
// Implements senderStatusSetter
type responseStatusSender struct {
	responseSender
}

func (j responseStatusSender) Status(statusCode int) sender {
	return newResponseStatusSender(j.rw, statusCode, j.data)
}

func newResponseStatusSender(rw http.ResponseWriter, statusCode int, data interface{}) responseStatusSender {
	return responseStatusSender{
		responseSender{rw: rw, statusCode: statusCode, data: data},
	}
}
