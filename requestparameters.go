package goapi

import "net/url"

type RequestParameters interface {
	URLValues(av map[string]string) *url.Values
}
