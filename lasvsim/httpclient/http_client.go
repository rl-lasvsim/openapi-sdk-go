package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func IsApiError(err error) (*APIError, bool) {
	e, ok := err.(*APIError)
	return e, ok
}

// APIError represents an error returned by the API
type APIError struct {
	StatusCode int
	Message    interface{}
	URL        string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError %s %d: %v", e.URL, e.StatusCode, e.Message)
}

// HttpConfig holds configuration for the client
type HttpConfig struct {
	Token    string
	Endpoint string
}

// Option is a function type that modifies HttpClient
type Option func(*HttpClient)

// WithHeaders sets custom headers for the client
func WithHeaders(headers map[string]string) Option {
	return func(c *HttpClient) {
		for k, v := range headers {
			c.Headers[k] = v
		}
	}
}

// HttpClient is the HTTP client for the API
type HttpClient struct {
	config  *HttpConfig
	Headers map[string]string
	client  *http.Client
}

// NewHttpClient creates a new HTTP client
func NewHttpClient(config *HttpConfig, opts ...Option) *HttpClient {
	client := &HttpClient{
		config:  config,
		Headers: make(map[string]string),
		client:  &http.Client{},
	}

	// Set default Authorization header
	client.Headers["Authorization"] = "Bearer " + config.Token

	// Apply options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *HttpClient) Clone() *HttpClient {
	clone := &HttpClient{
		config:  c.config,
		client:  &http.Client{},
		Headers: c.Headers,
	}

	return clone
}

// Get performs a GET request
func (c *HttpClient) Get(path string, params map[string]string, out any) error {
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Add(k, v)
	}

	reqURL := c.config.Endpoint + path
	if len(urlValues) > 0 {
		reqURL += "?" + urlValues.Encode()
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	return c.doRequest(req, out)
}

// Post performs a POST request
func (c *HttpClient) Post(path string, data, out any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.config.Endpoint+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	return c.doRequest(req, out)
}

func (c *HttpClient) doRequest(req *http.Request, out any) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var errResp interface{}
		if err := json.Unmarshal(body, &errResp); err != nil {
			errResp = string(body)
		}
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    errResp,
			URL:        req.Method + "," + req.URL.String(),
		}
	}
	if out == nil {
		return nil
	}

	err = json.Unmarshal(body, out)

	return err
}
