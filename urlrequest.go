package goapi

import (
	"net/http"
)

// An API request.
type URLRequest struct {

	// The request's HTTP method.
	Method string

	// The request's URL.
	URL string
}

// NewURLRequest creates a new API request from the given HTTP method and URL.
func NewURLRequest(method string, url string) URLRequest {
	return URLRequest{
		Method: method,
		URL:    url,
	}
}

// HTTPRequest creates and returns an HTTP request from the receiver with the
// given request parameters and additional request parameter.
func (r URLRequest) HTTPRequest(
	p RequestParameters,
	av map[string]string,
) (*http.Request, error) {
	return http.NewRequest(r.Method, r.URL, nil)
}
