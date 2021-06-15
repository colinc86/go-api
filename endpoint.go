package goapi

// Endpoint types contain the information necessary to form an endpoint's URL.
//
// Requests use an endpoint to form a URL of the form
//
// [SCHEME]://[HOST]/v[VERSION]/[PATH]
type Endpoint interface {

	// The endpoint's scheme.
	Scheme() string

	// The endpoint's hostname.
	Host() string

	// The endpoint's version.
	Version() int

	// The endpoint's path.
	Path() string
}
