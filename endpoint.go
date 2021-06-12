package goapi

import (
	"net/url"
)

// Endpoint types contain the information necessary to form a URL.
type Endpoint interface {

	// The endpoint's scheme.
	Scheme() string

	// The endpoint's hostname.
	Host() string

	// The endpoint's version.
	Version() int

	// The endpoint's path.
	Path() string

	// The endpoint's query parameters.
	Values() *url.Values
}
