package goapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// An application programming interface.
type API struct {

	// The Client's HTTP Client.
	Client *http.Client
}

// Creates a new API.
func NewAPI() *API {
	return NewAPIFromHTTPClient(&http.Client{})
}

// Creates a new API from the given HTTP client.
func NewAPIFromHTTPClient(client *http.Client) *API {
	return &API{
		Client: client,
	}
}

// PerformRequest performs the API request with the given query parameters and
// returns an HTTP response.
func (a API) PerformRequest(r Request, p RequestParameters, av map[string]string) (*http.Response, error) {
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
func (a API) PerformUnmarshalRequest(r Request, p RequestParameters, av map[string]string, v interface{}) (*http.Response, error) {
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
