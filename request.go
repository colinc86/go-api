package goapi

import (
	"fmt"
	"net/http"
	"net/url"
)

// An API request.
type Request struct {

	// The request's HTTP method.
	Method string

	// The request's endpoint.
	Endpoint Endpoint
}

// NewRequest creates a new API request from the given HTTP method and endpoint.
func NewRequest(method string, endpoint Endpoint) Request {
	return Request{
		Method:   method,
		Endpoint: endpoint,
	}
}

// HTTPRequest creates and returns an HTTP request from the receiver with the
// given request parameters.
func (r Request) HTTPRequest(p RequestParameters) (*http.Request, error) {
	u := &url.URL{
		Scheme: r.Endpoint.Scheme(),
		Host:   r.Endpoint.Host(),
		Path:   fmt.Sprintf("v%d/%s", r.Endpoint.Version(), r.Endpoint.Path()),
	}

	if p != nil {
		u.RawQuery = p.URLValues().Encode()
	}

	return http.NewRequest(r.Method, u.String(), nil)
}
