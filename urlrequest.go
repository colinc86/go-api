package goapi

import (
	"net/http"
	"net/url"
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
	u, err := url.Parse(r.URL)
	if err != nil {
		return nil, err
	}

	var newValues url.Values
	if p == nil {
		newValues = make(url.Values)

		for key, value := range av {
			newValues.Add(key, value)
		}
	} else {
		newValues = p.URLValues(av)
	}

	currentValues := u.Query()
	for key, values := range currentValues {
		for _, value := range values {
			newValues.Add(key, value)
		}
	}

	u.RawQuery = newValues.Encode()
	return http.NewRequest(r.Method, u.String(), nil)
}
