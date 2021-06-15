package goapi

type RequestParameters interface {
	URLQueryParameters() map[string]string
}
