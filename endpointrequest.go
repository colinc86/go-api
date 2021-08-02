package goapi

import (
	"fmt"
	"net/http"
	"net/url"
)

// An API request.
type EndpointRequest struct {

	// The request's HTTP method.
	Method string

	// The request's endpoint.
	Endpoint Endpoint
}

// NewEndpointRequest creates a new API request from the given HTTP method and
// endpoint.
func NewEndpointRequest(method string, endpoint Endpoint) EndpointRequest {
	return EndpointRequest{
		Method:   method,
		Endpoint: endpoint,
	}
}

// HTTPRequest creates and returns an HTTP request from the receiver with the
// given request parameters and additional request parameter.
func (r EndpointRequest) HTTPRequest(
	p RequestParameters,
	av map[string]string,
) (*http.Request, error) {
	u := &url.URL{
		Scheme: r.Endpoint.Scheme(),
		Host:   r.Endpoint.Host(),
		Path:   fmt.Sprintf("v%d/%s", r.Endpoint.Version(), r.Endpoint.Path()),
	}

	if p != nil {
		u.RawQuery = p.URLValues(av).Encode()
	}

	return http.NewRequest(r.Method, u.String(), nil)
}
