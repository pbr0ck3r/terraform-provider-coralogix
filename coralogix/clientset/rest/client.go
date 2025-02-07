package rest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Client for Coralogix API
type Client struct {
	url    string
	apiKey string
	client *http.Client
}

func NewRestClient(url string, apiKey string) *Client {
	return &Client{url, apiKey, &http.Client{}}
}

// Request executes request to Coralogix API
func (c *Client) Request(ctx context.Context, method, path, contentType string, body interface{}) (string, error) {
	var request *http.Request
	if body != nil {
		bodyReader := bytes.NewBuffer([]byte(body.(string)))
		var err error
		request, err = http.NewRequest(method, c.url+path, bodyReader)
		if err != nil {
			return "", err
		}

		request.Header.Set("Content-Type", contentType)
	} else {
		request, _ = http.NewRequest(method, c.url+path, nil)
	}

	request = request.WithContext(ctx)
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Authorization", "Bearer "+c.apiKey)

	response, err := c.client.Do(request)
	if err != nil {
		return "", status.Convert(err).Err()
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusMultipleChoices {

		bodyResp, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}

		return string(bodyResp), nil
	}

	if response.StatusCode == http.StatusNotFound {
		return "", status.Error(codes.NotFound, "Not found")
	}

	responseBody, err := httputil.DumpResponse(response, true)
	if err != nil {
		return "", status.Convert(err).Err()
	}

	return "", fmt.Errorf("API Error: %s. Status code: %s", string(responseBody), response.Status)
}

// Get executes GET request to Coralogix API
func (c *Client) Get(ctx context.Context, path string) (string, error) {
	return c.Request(ctx, "GET", path, "", nil)
}

// Post executes POST request to Coralogix API
func (c *Client) Post(ctx context.Context, path, contentType, body string) (string, error) {
	return c.Request(ctx, "POST", path, contentType, body)
}

// Put executes PUT request to Coralogix API
func (c *Client) Put(ctx context.Context, path, contentType, body string) (string, error) {
	return c.Request(ctx, "PUT", path, contentType, body)
}

// Delete executes DELETE request to Coralogix API
func (c *Client) Delete(ctx context.Context, path string) (string, error) {
	return c.Request(ctx, "DELETE", path, "", nil)
}
