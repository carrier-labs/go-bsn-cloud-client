// Package service provides logical groupings of BSN.Cloud API endpoints.
package service

import (
	"context"
	"encoding/json"
	"fmt"

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

// GetDevices fetches the list of devices from BSN.Cloud using the configured network context.
func (s *DeviceService) GetDevices(ctx context.Context) ([]models.Player, error) {
	if err := s.Client.Authenticate(ctx); err != nil {
		debug.Debug("DeviceService: authentication error", "error", err)
		return nil, fmt.Errorf("authentication error: %w", err)
	}

	url := "/Devices"
	respBody, err := s.Client.DoRequest(ctx, "GET", url, nil)
	if err != nil {
		debug.Debug("DeviceService: API error", "error", err)
		return nil, err
	}
	debug.Debug("DeviceService: raw response body", "body", string(respBody))

	var result models.PlayerListResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		debug.Debug("DeviceService: decode error", "error", err)
		return nil, fmt.Errorf("parsing devices: %w", err)
	}
	pretty, _ := json.MarshalIndent(result, "", "  ")
	debug.Debug("DeviceService: API response", "data", string(pretty))

	return result.Items, nil
}
