package goapi

import "net/http"

// An API request.
type Request interface {

	// HTTPRequest returns an HTTP request.
	HTTPRequest(p RequestParameters, av map[string]string) (*http.Request, error)
}
