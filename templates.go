package jutils

import (
	"net/http"
)


// Template Includes methods that have preset status code and json messages
type Template struct {
	rw 	http.ResponseWriter
}

// MissingParams 400 status {"error": "missing or invalid params"}
func (t Template) MissingParams() templateResponse {
	return newTemplateResponse(t.rw, 400, fmtErr("Missing or invalid params"))
}

// NotFound 404 status {"error": "Resource not found"}
func (t Template) NotFound() templateResponse {
	return newTemplateResponse(t.rw, 404, fmtErr("Resource not found"))
}

// InternalError Set generic error message {"error": "An error occurred"}
func (t Template) InternalError() templateResponse {
	return newTemplateResponse(t.rw, 500, fmtErr("An error occurred"))
}

// DBError 500 status. {"error": "DB Error"}
func (t Template) DBError() templateResponse {
	return newTemplateResponse(t.rw, 500, fmtErr("DB Error"))
}

// JsonParseError 400 status {"error": "JSON badly formatted"}
func (t Template) JsonParseError() templateResponse {
	return newTemplateResponse(t.rw, 500, fmtErr("JSON badly formatted"))
}

