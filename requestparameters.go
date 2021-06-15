package goapi

import "net/url"

type RequestParameters interface {
	URLValues() *url.Values
}
