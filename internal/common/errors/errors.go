package errors

import (
	"encoding/json"
	"io"
	"net/http"
)

// HTTPError is the universal interface for all HTTPError 's
// it includes their StatusCode()
//and a StatusMessage() which contains a short description of the error
type HTTPError interface {
	StatusCode() int
	StatusMessage() string
	ToJSON(io.Writer) error
	mustEmbedGeneralHTTPError()
	error
}

// ErrorToJSON converts an error type to JSON and writes to provided writer
func ErrorToJSON(e HTTPError, w *io.Writer) error {
	data := map[string]interface{}{
		"statusCode":    e.StatusCode(),
		"statusMessage": e.StatusMessage(),
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// write to writer
	_, err = (*w).Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

// generalHTTPError is the default HTTPError struct
type generalHTTPError struct {
	error
}

// StatusCode is 500 for unimplemented error status codes
func (generalHTTPError) StatusCode() int {
	return http.StatusInternalServerError
}

func (generalHTTPError) mustEmbedGeneralHTTPError() {}

// ToJSON is the default implementation for converting an error type to
// JSON to be sent in the response body
func (g generalHTTPError) ToJSON(w io.Writer) error {
	return ErrorToJSON(g, &w)
}

// StatusMessage looks up the already defined map[int]string from the http package containing
// maps of common http errors and their texts
func (g generalHTTPError) StatusMessage() string {
	return http.StatusText(g.StatusCode())
}

// NotFound - HTTP error for 404 error code
type NotFound struct {
	generalHTTPError
}

func (e NotFound) StatusCode() int {
	return http.StatusNotFound
}
func (e NotFound) StatusMessage() string {
	return http.StatusText(e.StatusCode())
}
func (e NotFound) ToJSON(w io.Writer) error {
	return ErrorToJSON(&e, &w)
}

// BadRequest - HTTP error for 400 error code
type BadRequest struct {
	generalHTTPError
}

func (e *BadRequest) StatusMessage() string {
	return http.StatusText(e.StatusCode())
}
func (e *BadRequest) StatusCode() int {
	return http.StatusBadRequest
}
func (e *BadRequest) ToJSON(w io.Writer) error {
	return ErrorToJSON(e, &w)
}

// InternalServerError - HTTP error for 500 error code
type InternalServerError struct {
	generalHTTPError
}
