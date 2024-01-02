package gorestease

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendRequestSuccess(t *testing.T) {
	// Create a test server to simulate API responses
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with a specific status code and response body
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test Response"))
	}))
	defer server.Close()

	// Create Params for the test
	testParams := Params{
		BaseUrl: server.URL,
		Path:    "/test",
		Method:  "GET",
	}

	// Perform the HTTP request using SendRequest
	resp, body, err := SendRequest(testParams)

	// Check for errors
	if err != nil {
		t.Fatalf("SendRequest failed: %v", err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expectedBody := "Test Response"
	if string(body) != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, string(body))
	}
}

func TestSendRequestFail(t *testing.T) {
	// Create a test server to simulate API responses
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a successful response with a specific status code and response body
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	// Create Params for the test
	testParams := Params{
		BaseUrl: server.URL,
		Path:    "/test",
		Method:  "GET",
	}

	// Perform the HTTP request using SendRequest
	resp, _, err := SendRequest(testParams)

	// Check for errors
	if err != nil {
		t.Fatalf("SendRequest failed: %v", err)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
