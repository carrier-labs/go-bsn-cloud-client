// Package service provides logical groupings of BSN.Cloud API endpoints.
package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/carrier-labs/go-bsn-cloud-client/client"
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
		return nil, fmt.Errorf("authentication error: %w", err)
	}

	if s.Client.NetworkName == "" {
		return nil, fmt.Errorf("network name must be configured in the client")
	}

	if err := s.Client.SelectNetwork(ctx); err != nil {
		return nil, fmt.Errorf("network selection error: %w", err)
	}

	url := fmt.Sprintf("%s/Devices", s.Client.BaseAPI)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating device request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.Client.Token)
	req.Header.Set("Accept", "application/json")

	resp, err := s.Client.HttpClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending device request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("device fetch failed: %s - %s", resp.Status, body)
	}

	var result models.PlayerListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("parsing devices: %w", err)
	}

	return result.Items, nil
}
