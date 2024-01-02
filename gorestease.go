package gorestease

import (
	"io"
	"net/http"
	"time"
)

// Params represents the parameters needed to perform an HTTP request.
type Params struct {
	// BaseUrl is the base URL for the API endpoint.
	BaseUrl string

	// Path is the specific path of the API endpoint.
	Path string

	// Method is the HTTP method for the request (e.g., GET, POST).
	Method string

	// Body is the request body, typically an io.Reader.
	Body io.Reader

	// Headers is a map containing additional HTTP headers for the request.
	Headers map[string]string

	// BasicAuthCreds stores the basic authentication credentials (username and password).
	BasicAuthCreds BasicAuthCreds

	// Transport allows customization of the underlying HTTP transport.
	Transport http.RoundTripper

	// CookieJar stores cookies for the request, facilitating stateful interactions.
	CookieJar http.CookieJar

	// Timeout is the maximum duration for the HTTP request before timing out.
	Timeout time.Duration
}

// BasicAuthCreds represents the basic authentication credentials.
type BasicAuthCreds struct {
	// Username is the username for basic authentication.
	Username string

	// Password is the password for basic authentication.
	Password string
}

// SendRequest performs an HTTP request based on the provided parameters (Params).
// It takes a Params struct containing essential details such as the base URL, path, HTTP method, request body, headers,
// basic authentication credentials, custom transport, cookie jar, and timeout duration.
// The function returns an HTTP response (*http.Response) or an error if the request encounters issues.
// SendRequest offers flexibility and customization options, allowing users to tailor their HTTP requests to specific needs.
// NOTE: even though `body []byte` is nested in `http.Response` it is returned separately to provide flexibility and avoid
// issues related to deferring the closure of the response body. This design choice allows users to access and manipulate
// the response body independently, ensuring a more versatile and error-resistant usage of the `SendRequest` function."
func SendRequest(p Params) (httpResponse *http.Response, body []byte, err error) {
	client := &http.Client{
		Transport: p.Transport,
		Jar:       p.CookieJar,
		Timeout:   p.Timeout,
	}

	req, err := http.NewRequest(p.Method, p.BaseUrl+p.Path, p.Body)
	if err != nil {
		return nil, nil, err
	}

	for key, val := range p.Headers {
		req.Header.Set(key, val)
	}

	req.SetBasicAuth(p.BasicAuthCreds.Username, p.BasicAuthCreds.Password)

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res, b, nil
}
