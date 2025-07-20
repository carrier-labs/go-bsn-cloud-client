// Package models contains shared data structures for the BSN.Cloud API client.
package models

import (
	"encoding/json"
	"fmt"
)

// PlayerBeaconMode is an enum for beacon modes supported by BSN.Cloud.
type PlayerBeaconMode string

const (
	// PlayerBeaconModeIBeacon represents the iBeacon mode.
	PlayerBeaconModeIBeacon PlayerBeaconMode = "iBeacon"
	// PlayerBeaconModeEddystoneUid represents the Eddystone UID mode.
	PlayerBeaconModeEddystoneUid PlayerBeaconMode = "EddystoneUid"
	// PlayerBeaconModeEddystoneUrl represents the Eddystone URL mode.
	PlayerBeaconModeEddystoneUrl PlayerBeaconMode = "EddystoneUrl"
	// PlayerBeaconModeUnknown represents an unknown beacon mode.
	PlayerBeaconModeUnknown PlayerBeaconMode = "Unknown"
)

// DeviceBeacon is an interface for all beacon types supported by BSN.Cloud.
type DeviceBeacon interface {
	// GetMode returns the beacon mode for this beacon.
	GetMode() PlayerBeaconMode
}

// IBeacon represents an iBeacon entity in BSN.Cloud.
type IBeacon struct {
	Name  string           `json:"name"`  // Name of the beacon
	Mode  PlayerBeaconMode `json:"mode"`  // Beacon mode (should be iBeacon)
	Major uint16           `json:"major"` // iBeacon major value
	Minor uint16           `json:"minor"` // iBeacon minor value
	UUID  string           `json:"uuid"`  // iBeacon UUID
	Power int16            `json:"power"` // Transmission power
}

// GetMode returns the beacon mode for IBeacon (always iBeacon).
func (b IBeacon) GetMode() PlayerBeaconMode { return b.Mode }

// EddystoneUidBeacon represents an Eddystone UID beacon entity in BSN.Cloud.
type EddystoneUidBeacon struct {
	Name        string           `json:"name"`        // Name of the beacon
	Mode        PlayerBeaconMode `json:"mode"`        // Beacon mode (should be EddystoneUid)
	NamespaceID []byte           `json:"namespaceId"` // Eddystone namespace ID
	InstanceID  []byte           `json:"instanceId"`  // Eddystone instance ID
	Power       int16            `json:"power"`       // Transmission power
}

// GetMode returns the beacon mode for EddystoneUidBeacon (always EddystoneUid).
func (b EddystoneUidBeacon) GetMode() PlayerBeaconMode { return b.Mode }

// EddystoneUrlBeacon represents an Eddystone URL beacon entity in BSN.Cloud.
type EddystoneUrlBeacon struct {
	Name  string           `json:"name"`  // Name of the beacon
	Mode  PlayerBeaconMode `json:"mode"`  // Beacon mode (should be EddystoneUrl)
	URL   string           `json:"url"`   // Eddystone URL
	Power int16            `json:"power"` // Transmission power
}

// GetMode returns the beacon mode for EddystoneUrlBeacon (always EddystoneUrl).
func (b EddystoneUrlBeacon) GetMode() PlayerBeaconMode { return b.Mode }

// DeviceBeaconWrapper is used for custom unmarshalling of DeviceBeacon sum types.
type DeviceBeaconWrapper struct {
	DeviceBeacon
}

// UnmarshalJSON implements custom unmarshalling for DeviceBeaconWrapper.
// It determines the beacon type by the "mode" field and unmarshals into the correct struct.
func (w *DeviceBeaconWrapper) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	typeVal, ok := raw["mode"].(string)
	if !ok {
		return fmt.Errorf("missing or invalid mode field in DeviceBeacon")
	}

	switch PlayerBeaconMode(typeVal) {
	case PlayerBeaconModeIBeacon:
		var beacon IBeacon
		if err := json.Unmarshal(data, &beacon); err != nil {
			return err
		}
		w.DeviceBeacon = beacon
	case PlayerBeaconModeEddystoneUid:
		var beacon EddystoneUidBeacon
		if err := json.Unmarshal(data, &beacon); err != nil {
			return err
		}
		w.DeviceBeacon = beacon
	case PlayerBeaconModeEddystoneUrl:
		var beacon EddystoneUrlBeacon
		if err := json.Unmarshal(data, &beacon); err != nil {
			return err
		}
		w.DeviceBeacon = beacon
	default:
		return fmt.Errorf("unknown beacon mode: %s", typeVal)
	}
	return nil
}

// UnmarshalDeviceBeacons unmarshals a JSON array of beacons into a slice of DeviceBeacon interfaces.
func UnmarshalDeviceBeacons(data []byte) ([]DeviceBeacon, error) {
	var rawList []json.RawMessage
	if err := json.Unmarshal(data, &rawList); err != nil {
		return nil, err
	}
	beacons := make([]DeviceBeacon, len(rawList))
	for i, raw := range rawList {
		var wrapper DeviceBeaconWrapper
		if err := json.Unmarshal(raw, &wrapper); err != nil {
			return nil, err
		}
		beacons[i] = wrapper.DeviceBeacon
	}
	return beacons, nil
}
