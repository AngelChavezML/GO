package web

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
)

func RequestJSON(r *http.Request, ptr any) (err error) {
	// check the content type
	if r.Header.Get("Content-Type") != "application/json" {
		err = ErrInvalidRequestBody
		return
	}

	// decode the request body
	if err = json.NewDecoder(r.Body).Decode(ptr); err != nil {
		return
	}

	return
}
