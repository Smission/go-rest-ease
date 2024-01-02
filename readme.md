
# gorestease

gorestease is a lightweight Golang package for simplifying HTTP requests with flexibility and customization options.

## Installation

```bash
go get -u github.com/Smission/gorestease
```

## Usage

```go
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"github.com/Smission/gorestease"
)

func main() {
    // Example: Sending a POST request with HTML content

    // Create Params for the request
    p := gorestease.Params{
        BaseUrl: "some-url",
        Method:  http.MethodPost,
        Body:    bytes.NewBuffer([]byte(htmlContent)),
        Headers: map[string]string{
            "Content-Type":  "text/html",
            "Cache-Control": "no-cache",
        },
    }

    // Send the request
    httpRes, body, err := gorestease.SendRequest(p)
    if err != nil {
        // Handle error
        fmt.Println("Error:", err)
        return
    }

    // Check the HTTP status code
    if httpRes.StatusCode != http.StatusOK {
        // Handle error
        fmt.Printf("Unexpected status code: %d\n", httpRes.StatusCode)
        return
    }

    // Use the response body as needed
    fmt.Println("Response Body:", string(body))
}
```
