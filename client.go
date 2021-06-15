// Package goapi provides types to facilitate the creation of RESTful API
// wrappers.
package goapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// An application programming interface.
type APIClient struct {

	// The Client's HTTP Client.
	Client *http.Client
}

// NewAPIClient creates a new API.
func NewAPIClient() *APIClient {
	return NewAPIClientFromHTTPClient(&http.Client{})
}

// NewAPIClientFromHTTPClient creates a new API from the given HTTP client.
func NewAPIClientFromHTTPClient(client *http.Client) *APIClient {
	return &APIClient{
		Client: client,
	}
}

// PerformRequest performs the API request with the given query parameters and
// returns an HTTP response.
func (a APIClient) PerformRequest(
	r Request,
	p RequestParameters,
	av map[string]string,
) (*http.Response, error) {
	// Create the HTTP request
	request, err := r.HTTPRequest(p, av)
	if err != nil {
		return nil, err
	}

	// Get the response
	return a.Client.Do(request)
}

// PerformUnmarshalRequest performs the API request with the given query
// parameters and unmarshals the response data in to v.
func (a APIClient) PerformUnmarshalRequest(
	r Request,
	p RequestParameters,
	av map[string]string,
	v interface{},
) (*http.Response, error) {
	response, err := a.PerformRequest(r, p, av)

	// Unmarshal objects
	if response.Body != nil && v != nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return response, err
		}

		if err = json.Unmarshal(data, v); err != nil {
			return response, err
		}
	}

	// Return objects
	return response, err
}
