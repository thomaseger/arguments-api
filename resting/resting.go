package resting

import (
	"io"
	"net/http"
	"net/url"
	"testing"
)

const (
	GET    = "GET"
	PUT    = "PUT"
	POST   = "POST"
	DELETE = "DELETE"
)

// Checks for an error in url syntax, GET execution and status code.
// Returns the io.ReadCloser of the http.Get on success.
func GetResource(t *testing.T, url string) io.ReadCloser {
	urlSyntaxError(t, url)
	response, err := http.Get(url)
	methodError(t, GET, err)
	statusCodeError(t, response.StatusCode, []int{200, 202, 203})
	return response.Body
}

func PostResource(t *testing.T, url string, mime string, data io.Reader) io.ReadCloser {
	urlSyntaxError(t, url)
	response, err := http.Post(url, mime, data)
	methodError(t, POST, err)
	statusCodeError(t, response.StatusCode, []int{200, 201, 202})
	return response.Body
}

func PutResource(t *testing.T, url string) {

}

func DeleteResource(t *testing.T, url string) {

}

func urlSyntaxError(t *testing.T, rawurl string) {
	_, err := url.ParseRequestURI(rawurl)

	if err != nil {
		t.Errorf("Error parsing request url: ", err)
	}
}

func methodError(t *testing.T, method string, err error) {
	if err != nil {
		t.Errorf("GET failed: ", err)
	}
}

func statusCodeError(t *testing.T, code int, allowed []int) {
	contains := false
	for _, c := range allowed {
		if c == code {
			contains = true
		}
	}

	if !contains {
		t.Errorf("Unexpected status code: %d", code)
	}
}
