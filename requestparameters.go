package goapi

import "net/url"

// RequestParameters types contain the necessary information to return values
// from its URLValues method.
type RequestParameters interface {

	// URLValues returns the request parameter's fields as a URL values type. Use
	// the av parameter to specify additional parameter values.
	URLValues(av map[string]string) *url.Values
}
