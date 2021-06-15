# package go-api

Package goapi provides types to facilitate the creation of RESTful API wrappers.

## Usage

### Creating Requests

Create a request structure that implements the `Request` interface for the desired endpoint.

```go
// MyEndpoint will render as https://hostname.com/v1/my/endpoint
type MyEndpoint struct{}

func (e MyEndpoint) Scheme() string {
  return "https"
}

func (e MyEndpoint) Host() string {
  return "hostname.com"
}

func (e MyEndpoint) Version() int {
  return 1
}

func (e MyEndpoint) Path() string {
  return "my/endpoint"
}
```

### Request Parameters

If an endpoint accepts URL query parameters, specify them by creating a structure that implements the `RequestParameters` interface. It's your responsibility to do something with the additional values parameter, `av`.

```go
type MyEndpointParameters struct {
  Foo string
}

func (p MyEndpointParameters) URLValues(av map[string]string) *url.Values {
  v := new(url.Values)

  if len(p.Foo) > 0 {
    v.Add("foo", p.Foo)
  }

  for key, value := range av {
    v.Add(key, value)
  }

  return nil
}
```

### Performing Requests

Create an `APIClient` and give it a request and some, optional, request parameters.

```go
c := goapi.NewAPIClient()
request := goapi.NewRequest(http.MethodGet, MyEndpoint{})
requestParameters := MyEndpointParameters{Foo: "bar"}

httpResponse, err := c.api.PerformRequest(
  request,
  requestParameters,
  nil,
)
```

If you have a custom response type, you can use the `PerformUnmarshalRequest`.

```go
var responseData []string

httpResponse, err := c.api.PerformUnmarshalRequest(
  request,
  requestParameters,
  nil,
  &responseData,
)
```
