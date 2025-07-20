// Package service provides logical groupings of BSN.Cloud API endpoints.
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/carrier-labs/go-bsn-cloud-client/client"
	"github.com/carrier-labs/go-bsn-cloud-client/debug"
	"github.com/carrier-labs/go-bsn-cloud-client/models"
)

type DeviceService struct {
	Client *client.Client
}

// NewDeviceService creates a new DeviceService.
func NewDeviceService(c *client.Client) *DeviceService {
	return &DeviceService{Client: c}
}

// GetDevices fetches the list of devices from BSN.Cloud using the configured network name.
func (s *DeviceService) GetDevices(ctx context.Context) ([]models.Player, error) {
	if err := s.Client.Authenticate(ctx); err != nil {
		debug.Debug("DeviceService: authentication error", "error", err)
		return nil, fmt.Errorf("authentication error: %w", err)
	}

	if s.Client.NetworkName == "" {
		debug.Debug("DeviceService: missing network name")
		return nil, fmt.Errorf("network name must be configured in the client")
	}

	if err := s.Client.SelectNetwork(ctx); err != nil {
		debug.Debug("DeviceService: network selection error", "error", err)
		return nil, fmt.Errorf("network selection error: %w", err)
	}

	url := fmt.Sprintf("%s/Devices", s.Client.BaseAPI)
	debug.Debug("DeviceService: fetching devices", "url", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		debug.Debug("DeviceService: request creation error", "error", err)
		return nil, fmt.Errorf("creating device request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.Client.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := s.Client.HttpClient().Do(req)
	if err != nil {
		debug.Debug("DeviceService: send request error", "error", err)
		return nil, fmt.Errorf("sending device request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	debug.Debug("DeviceService: raw response body", "body", string(body))

	if resp.StatusCode != http.StatusOK {
		debug.Debug("DeviceService: API error response", "status", resp.Status, "body", string(body))
		return nil, fmt.Errorf("device fetch failed: %s - %s", resp.Status, body)
	}

	var result models.PlayerListResponse
	if err := json.Unmarshal(body, &result); err != nil {
		debug.Debug("DeviceService: decode error", "error", err)
		return nil, fmt.Errorf("parsing devices: %w", err)
	}
	pretty, _ := json.MarshalIndent(result, "", "  ")
	debug.Debug("DeviceService: API response", "data", string(pretty))

	return result.Items, nil
}
