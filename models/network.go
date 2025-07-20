// Package models contains shared data structures for the BSN.Cloud API client.
package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// PlayerNetworkStatus represents the network status of a player, including all interfaces.
type PlayerNetworkStatus struct {
	ExternalIp string                         `json:"externalIp"` // External IP address
	Interfaces []PlayerNetworkInterfaceStatus `json:"interfaces"` // Network interfaces
}

// PlayerNetworkInterfaceStatus is an interface for all network interface types.
type PlayerNetworkInterfaceStatus interface {
	// GetType returns the type of the network interface.
	GetType() PlayerNetworkInterfaceType
}

// PlayerNetworkInterfaceType is an enum for interface types.
type PlayerNetworkInterfaceType string

const (
	// PlayerNetworkInterfaceTypeEthernet represents an Ethernet interface.
	PlayerNetworkInterfaceTypeEthernet PlayerNetworkInterfaceType = "Ethernet"
	// PlayerNetworkInterfaceTypeWiFi represents a WiFi interface.
	PlayerNetworkInterfaceTypeWiFi PlayerNetworkInterfaceType = "WiFi"
	// PlayerNetworkInterfaceTypeVirtual represents a virtual interface.
	PlayerNetworkInterfaceTypeVirtual PlayerNetworkInterfaceType = "Virtual"
	// PlayerNetworkInterfaceTypeOther represents an interface of other type.
	PlayerNetworkInterfaceTypeOther PlayerNetworkInterfaceType = "Other"
	// PlayerNetworkInterfaceTypeCellular represents a cellular interface.
	PlayerNetworkInterfaceTypeCellular PlayerNetworkInterfaceType = "Cellular"
	// PlayerNetworkInterfaceTypeUnknown represents an unknown interface type.
	PlayerNetworkInterfaceTypeUnknown PlayerNetworkInterfaceType = "Unknown"
)

// NetworkConfigurationProtocol is an enum for network protocols.
type NetworkConfigurationProtocol string

const (
	// NetworkConfigurationProtocolStatic represents static configuration.
	NetworkConfigurationProtocolStatic NetworkConfigurationProtocol = "Static"
	// NetworkConfigurationProtocolDHCPv4 represents DHCPv4 configuration.
	NetworkConfigurationProtocolDHCPv4 NetworkConfigurationProtocol = "DHCPv4"
	// NetworkConfigurationProtocolDHCPv6 represents DHCPv6 configuration.
	NetworkConfigurationProtocolDHCPv6 NetworkConfigurationProtocol = "DHCPv6"
	// NetworkConfigurationProtocolNDP represents NDP configuration.
	NetworkConfigurationProtocolNDP NetworkConfigurationProtocol = "NDP"
	// NetworkConfigurationProtocolUnknown represents an unknown protocol.
	NetworkConfigurationProtocolUnknown NetworkConfigurationProtocol = "Unknown"
)

// NetworkInterfaceStatus represents a standard network interface.
type NetworkInterfaceStatus struct {
	Name    string                         `json:"name"`             // Interface name
	Type    PlayerNetworkInterfaceType     `json:"type"`             // Interface type
	Proto   []NetworkConfigurationProtocol `json:"proto"`            // Protocols
	Mac     string                         `json:"mac"`              // MAC address
	Ip      []string                       `json:"ip"`               // IP addresses
	Gateway string                         `json:"gateway"`          // Gateway
	Metric  *int                           `json:"metric,omitempty"` // Metric
}

// UnmarshalJSON handles proto as either a CSV string or array.
func (n *NetworkInterfaceStatus) UnmarshalJSON(data []byte) error {
	type Alias NetworkInterfaceStatus
	aux := &struct {
		Proto interface{} `json:"proto"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch v := aux.Proto.(type) {
	case string:
		// CSV string
		if v == "" {
			n.Proto = nil
		} else {
			parts := make([]NetworkConfigurationProtocol, 0)
			for _, s := range splitAndTrim(v, ",") {
				parts = append(parts, NetworkConfigurationProtocol(s))
			}
			n.Proto = parts
		}
	case []interface{}:
		parts := make([]NetworkConfigurationProtocol, 0, len(v))
		for _, s := range v {
			if str, ok := s.(string); ok {
				parts = append(parts, NetworkConfigurationProtocol(str))
			}
		}
		n.Proto = parts
	default:
		n.Proto = nil
	}

	n.Name = aux.Name
	n.Type = aux.Type
	n.Mac = aux.Mac
	n.Ip = aux.Ip
	n.Gateway = aux.Gateway
	n.Metric = aux.Metric
	return nil
}

// splitAndTrim splits a string by sep and trims whitespace from each part.
func splitAndTrim(s, sep string) []string {
	var out []string
	for _, part := range split(s, sep) {
		trimmed := trim(part)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func split(s, sep string) []string {
	return strings.Split(s, sep)
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

// GetType returns the type of the network interface.
func (n NetworkInterfaceStatus) GetType() PlayerNetworkInterfaceType { return n.Type }

// CellularInterfaceStatus represents a cellular network interface.
type CellularInterfaceStatus struct {
	Name    string                         `json:"name"`             // Interface name
	Type    PlayerNetworkInterfaceType     `json:"type"`             // Interface type
	Proto   []NetworkConfigurationProtocol `json:"proto"`            // Protocols
	Mac     string                         `json:"mac"`              // MAC address
	Ip      []string                       `json:"ip"`               // IP addresses
	Gateway string                         `json:"gateway"`          // Gateway
	Metric  *int                           `json:"metric,omitempty"` // Metric
	Modem   CellularModemInfo              `json:"modem"`            // Modem info
	Sims    []CellularSimInfo              `json:"sims"`             // SIMs info
}

// GetType returns the type of the network interface.
func (c CellularInterfaceStatus) GetType() PlayerNetworkInterfaceType { return c.Type }

// CellularModemInfo represents modem info for cellular interfaces.
type CellularModemInfo struct {
	IMEI         string `json:"imei"`         // IMEI number
	Manufacturer string `json:"manufacturer"` // Manufacturer
	Model        string `json:"model"`        // Model
	Revision     string `json:"revision"`     // Revision
}

// CellularSimInfo represents SIM info for cellular interfaces.
type CellularSimInfo struct {
	Status     string                `json:"status"`     // SIM status
	ICCID      string                `json:"iccid"`      // ICCID
	Connection CellularSimConnection `json:"connection"` // Connection info
}

// CellularSimConnection represents SIM connection info.
type CellularSimConnection struct {
	Network string `json:"network"` // Network name
	Signal  int16  `json:"signal"`  // Signal strength
}

// PlayerNetworkInterfaceStatusWrapper is used for custom unmarshalling of interface status.
type PlayerNetworkInterfaceStatusWrapper struct {
	PlayerNetworkInterfaceStatus
}

// UnmarshalJSON implements custom unmarshalling for PlayerNetworkInterfaceStatusWrapper.
// It determines the interface type by the "type" field and unmarshals into the correct struct.
func (w *PlayerNetworkInterfaceStatusWrapper) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	typeVal, ok := raw["type"].(string)
	if !ok {
		return fmt.Errorf("missing or invalid type field in PlayerNetworkInterfaceStatus")
	}

	switch PlayerNetworkInterfaceType(typeVal) {
	case PlayerNetworkInterfaceTypeCellular:
		var cellular CellularInterfaceStatus
		if err := json.Unmarshal(data, &cellular); err != nil {
			return err
		}
		w.PlayerNetworkInterfaceStatus = cellular
	default:
		var standard NetworkInterfaceStatus
		if err := json.Unmarshal(data, &standard); err != nil {
			return err
		}
		w.PlayerNetworkInterfaceStatus = standard
	}
	return nil
}

// UnmarshalJSON for PlayerNetworkStatus to handle interface slice.
func (p *PlayerNetworkStatus) UnmarshalJSON(data []byte) error {
	type Alias PlayerNetworkStatus
	aux := &struct {
		Interfaces []json.RawMessage `json:"interfaces"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p.ExternalIp = aux.ExternalIp
	p.Interfaces = make([]PlayerNetworkInterfaceStatus, len(aux.Interfaces))
	for i, raw := range aux.Interfaces {
		var wrapper PlayerNetworkInterfaceStatusWrapper
		if err := json.Unmarshal(raw, &wrapper); err != nil {
			return err
		}
		p.Interfaces[i] = wrapper.PlayerNetworkInterfaceStatus
	}
	return nil
}

// PlayerNetworkSettings represents network settings for a player.
type PlayerNetworkSettings struct {
	Hostname    string                           `json:"hostname"`    // Hostname
	ProxyServer string                           `json:"proxyServer"` // Proxy server
	ProxyBypass []string                         `json:"proxyBypass"` // Proxy bypass list
	TimeServers []string                         `json:"timeServers"` // Time servers
	Interfaces  []PlayerNetworkInterfaceSettings `json:"interfaces"`  // Network interface settings
}

// PlayerNetworkInterfaceSettings is now defined in network_interface_settings.go
