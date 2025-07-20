// Package client provides the core HTTP client, authentication, and configuration for BSN.Cloud API.
package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/carrier-labs/go-bsn-cloud-client/debug"
)

const DefaultBaseAPI = "https://api.bsn.cloud/2022/06/REST"

// Config holds configuration for the BSN.Cloud API client.
type Config struct {
	ClientID     string
	ClientSecret string
	BaseAPI      string        // Optional; if empty, DefaultBaseAPI is used
	Timeout      time.Duration // Optional; if zero, 10s is used
	NetworkName  string        // Optional; if set, network context is selected after auth
}

type Client struct {
	clientID     string
	clientSecret string
	baseAPI      string
	httpClient   *http.Client
	Token        string
	Expiry       time.Time
	NetworkName  string
}

// New creates a new BSN.Cloud API client using the provided Config.
func New(cfg Config) *Client {
	baseAPI := cfg.BaseAPI
	if baseAPI == "" {
		baseAPI = DefaultBaseAPI
	}
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	c := &Client{
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		baseAPI:      baseAPI,
		httpClient:   &http.Client{Timeout: timeout},
		NetworkName:  cfg.NetworkName,
	}
	debug.Debug("Client initialized", "clientID", cfg.ClientID, "baseAPI", baseAPI, "networkName", cfg.NetworkName)
	return c
}

// Authenticate fetches and caches an access token for the BSN.Cloud API.
func (c *Client) Authenticate(ctx context.Context) error {
	if time.Now().Before(c.Expiry.Add(-30 * time.Second)) {
		return nil
	}

	type bsnAuthResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	url := "https://auth.bsn.cloud/realms/bsncloud/protocol/openid-connect/token"
	form := []byte("grant_type=client_credentials")
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(form))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	auth := base64.StdEncoding.EncodeToString([]byte(c.clientID + ":" + c.clientSecret))
	req.Header.Set("Authorization", "Basic "+auth)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	debug.Debug("API call", "method", req.Method, "url", req.URL.String())
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		debug.Debug("API error response", "status", resp.Status, "body", string(body))
		return fmt.Errorf("auth failed: %s -- %s", resp.Status, body)
	}

	var ar bsnAuthResp
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		return err
	}
	c.Token = ar.AccessToken
	c.Expiry = time.Now().Add(time.Duration(ar.ExpiresIn) * time.Second)

	// Set network context if configured
	if c.NetworkName != "" {
		if err := c.SelectNetwork(ctx); err != nil {
			return fmt.Errorf("network selection error: %w", err)
		}
	}
	return nil
}

// DoRequest performs an HTTP request with context and returns the response body.
func (c *Client) DoRequest(ctx context.Context, method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	var bodyBytes []byte
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyBytes = b
		reqBody = bytes.NewBuffer(b)
	}
	url := c.baseAPI + endpoint
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Log request details
	reqHeaders := map[string][]string{}
	for k, v := range req.Header {
		reqHeaders[k] = v
	}
	debug.Debug("DoRequest: request", "method", req.Method, "url", req.URL.String(), "headers", reqHeaders, "body", string(bodyBytes))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Log response headers
	respHeaders := map[string][]string{}
	for k, v := range resp.Header {
		respHeaders[k] = v
	}
	debug.Debug("DoRequest: response status", "status", resp.StatusCode, "headers", respHeaders)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(respBody) == 0 {
		return nil, fmt.Errorf("API returned empty response body (status %d)", resp.StatusCode)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", respBody)
	}
	return respBody, nil
}

// HttpClient returns the underlying http.Client for advanced use.
func (c *Client) HttpClient() *http.Client {
	return c.httpClient
}

// SelectNetwork sets the active network context for the client.
func (c *Client) SelectNetwork(ctx context.Context) error {
	if c.NetworkName == "" {
		return nil
	}
	url := c.baseAPI + "/self/session/network"
	body := map[string]string{"name": c.NetworkName}
	buf, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	debug.Debug("API call", "method", req.Method, "url", req.URL.String(), "body", string(buf), "response_status", resp.StatusCode)
	for k, v := range resp.Header {
		debug.Debug("API response header", "key", k, "value", v)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		debug.Debug("API error response", "status", resp.Status, "body", string(body))
		return fmt.Errorf("network select failed: %s", resp.Status)
	}
	return nil
}
