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
	"sync"
	"time"
)

// DefaultBaseAPI is the default BSN.Cloud API base URL.
const DefaultBaseAPI = "https://api.bsn.cloud/v1"

// Client is the main struct for interacting with the BSN.Cloud API.
type Client struct {
	ClientID     string
	ClientSecret string
	BaseAPI      string
	NetworkName  string
	Token        string
	Expiry       time.Time
	mu           sync.Mutex
	httpClient   *http.Client
}

// New creates a new BSN.Cloud API client with the given credentials and optional base API URL and network name.
// If baseAPI is empty, it defaults to DefaultBaseAPI.
func New(clientID, clientSecret, baseAPI, networkName string) *Client {
	if baseAPI == "" {
		baseAPI = DefaultBaseAPI
	}
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseAPI:      baseAPI,
		NetworkName:  networkName,
		httpClient:   &http.Client{Timeout: 10 * time.Second},
	}
}

// Authenticate fetches and caches an access token for the BSN.Cloud API.
func (c *Client) Authenticate(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

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
	auth := base64.StdEncoding.EncodeToString([]byte(c.ClientID + ":" + c.ClientSecret))
	req.Header.Set("Authorization", "Basic "+auth)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("auth failed: %s -- %s", resp.Status, body)
	}

	var ar bsnAuthResp
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		return err
	}
	c.Token = ar.AccessToken
	c.Expiry = time.Now().Add(time.Duration(ar.ExpiresIn) * time.Second)
	return nil
}

// SelectNetwork sets the active network context for the client.
func (c *Client) SelectNetwork(ctx context.Context) error {
	url := c.BaseAPI + "/self/session/network"
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
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return fmt.Errorf("network select failed: %s", resp.Status)
	}
	return nil
}

// ListNetworks fetches available networks for the authenticated user.
func (c *Client) ListNetworks(ctx context.Context) ([]map[string]interface{}, error) {
	url := c.BaseAPI + "/networks"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("networks fetch failed: %s - %s", resp.Status, body)
	}
	var result struct {
		Items []map[string]interface{} `json:"items"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Items, nil
}

// HttpClient returns the underlying http.Client for making requests.
func (c *Client) HttpClient() *http.Client {
	return c.httpClient
}
