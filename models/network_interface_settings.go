// Package models contains shared data structures for the BSN.Cloud API client.
// Package models contains shared data structures for the BSN.Cloud API client.
package models

// PlayerNetworkInterfaceSettings is the base interface for all network interface settings.
type PlayerNetworkInterfaceSettings interface {
	// GetType returns the type of the network interface settings.
	GetType() PlayerNetworkInterfaceType
}

// EthernetInterfaceSettings represents Ethernet interface settings.
type EthernetInterfaceSettings struct {
	Enabled                               bool                         `json:"enabled"`                                         // Whether the interface is enabled
	Name                                  string                       `json:"name"`                                            // Interface name
	Type                                  PlayerNetworkInterfaceType   `json:"type"`                                            // Interface type
	Proto                                 NetworkConfigurationProtocol `json:"proto"`                                           // Protocol
	IP                                    []string                     `json:"ip"`                                              // IP addresses
	Gateway                               string                       `json:"gateway"`                                         // Gateway
	DNS                                   []string                     `json:"dns"`                                             // DNS servers
	RateLimitDuringInitialDownloads       *int                         `json:"rateLimitDuringInitialDownloads,omitempty"`       // Rate limit during initial downloads
	RateLimitInsideContentDownloadWindow  *int                         `json:"rateLimitInsideContentDownloadWindow,omitempty"`  // Rate limit inside content download window
	RateLimitOutsideContentDownloadWindow *int                         `json:"rateLimitOutsideContentDownloadWindow,omitempty"` // Rate limit outside content download window
	ContentDownloadEnabled                bool                         `json:"contentDownloadEnabled"`                          // Content download enabled
	TextFeedsDownloadEnabled              bool                         `json:"textFeedsDownloadEnabled"`                        // Text feeds download enabled
	MediaFeedsDownloadEnabled             bool                         `json:"mediaFeedsDownloadEnabled"`                       // Media feeds download enabled
	HealthReportingEnabled                bool                         `json:"healthReportingEnabled"`                          // Health reporting enabled
	LogsUploadEnabled                     bool                         `json:"logsUploadEnabled"`                               // Logs upload enabled
}

// GetType returns the type of the network interface settings.
func (e EthernetInterfaceSettings) GetType() PlayerNetworkInterfaceType { return e.Type }

// VirtualInterfaceSettings represents Virtual interface settings.
type VirtualInterfaceSettings struct {
	Enabled                               bool                         `json:"enabled"`                                         // Whether the interface is enabled
	Name                                  string                       `json:"name"`                                            // Interface name
	Type                                  PlayerNetworkInterfaceType   `json:"type"`                                            // Interface type
	Parent                                string                       `json:"parent"`                                          // Parent interface
	VlanId                                uint16                       `json:"vlanId"`                                          // VLAN ID
	Proto                                 NetworkConfigurationProtocol `json:"proto"`                                           // Protocol
	IP                                    []string                     `json:"ip"`                                              // IP addresses
	Gateway                               string                       `json:"gateway"`                                         // Gateway
	DNS                                   []string                     `json:"dns"`                                             // DNS servers
	RateLimitDuringInitialDownloads       *int                         `json:"rateLimitDuringInitialDownloads,omitempty"`       // Rate limit during initial downloads
	RateLimitInsideContentDownloadWindow  *int                         `json:"rateLimitInsideContentDownloadWindow,omitempty"`  // Rate limit inside content download window
	RateLimitOutsideContentDownloadWindow *int                         `json:"rateLimitOutsideContentDownloadWindow,omitempty"` // Rate limit outside content download window
	ContentDownloadEnabled                bool                         `json:"contentDownloadEnabled"`                          // Content download enabled
	TextFeedsDownloadEnabled              bool                         `json:"textFeedsDownloadEnabled"`                        // Text feeds download enabled
	MediaFeedsDownloadEnabled             bool                         `json:"mediaFeedsDownloadEnabled"`                       // Media feeds download enabled
	HealthReportingEnabled                bool                         `json:"healthReportingEnabled"`                          // Health reporting enabled
	LogsUploadEnabled                     bool                         `json:"logsUploadEnabled"`                               // Logs upload enabled
}

// GetType returns the type of the network interface settings.
func (v VirtualInterfaceSettings) GetType() PlayerNetworkInterfaceType { return v.Type }

// WiFiInterfaceSettings represents WiFi interface settings.
type WiFiInterfaceSettings struct {
	Enabled                               bool                         `json:"enabled"`                                         // Whether the interface is enabled
	Name                                  string                       `json:"name"`                                            // Interface name
	Type                                  PlayerNetworkInterfaceType   `json:"type"`                                            // Interface type
	SSID                                  string                       `json:"ssid"`                                            // WiFi SSID
	Security                              WiFiSecuritySettings         `json:"security"`                                        // WiFi security settings
	Proto                                 NetworkConfigurationProtocol `json:"proto"`                                           // Protocol
	IP                                    []string                     `json:"ip"`                                              // IP addresses
	Gateway                               string                       `json:"gateway"`                                         // Gateway
	DNS                                   []string                     `json:"dns"`                                             // DNS servers
	RateLimitDuringInitialDownloads       *int                         `json:"rateLimitDuringInitialDownloads,omitempty"`       // Rate limit during initial downloads
	RateLimitInsideContentDownloadWindow  *int                         `json:"rateLimitInsideContentDownloadWindow,omitempty"`  // Rate limit inside content download window
	RateLimitOutsideContentDownloadWindow *int                         `json:"rateLimitOutsideContentDownloadWindow,omitempty"` // Rate limit outside content download window
	ContentDownloadEnabled                bool                         `json:"contentDownloadEnabled"`                          // Content download enabled
	TextFeedsDownloadEnabled              bool                         `json:"textFeedsDownloadEnabled"`                        // Text feeds download enabled
	MediaFeedsDownloadEnabled             bool                         `json:"mediaFeedsDownloadEnabled"`                       // Media feeds download enabled
	HealthReportingEnabled                bool                         `json:"healthReportingEnabled"`                          // Health reporting enabled
	LogsUploadEnabled                     bool                         `json:"logsUploadEnabled"`                               // Logs upload enabled
}

// GetType returns the type of the network interface settings.
func (w WiFiInterfaceSettings) GetType() PlayerNetworkInterfaceType { return w.Type }

// WiFiSecuritySettings represents WiFi security settings.
type WiFiSecuritySettings struct {
	Authentication WiFiAuthenticationSettings `json:"authentication"` // Authentication settings
	Encryption     WiFiEncryptionSettings     `json:"encryption"`     // Encryption settings
}

type WiFiAuthenticationSettings struct {
	Mode       string `json:"mode"`                 // Authentication mode
	Passphrase string `json:"passphrase,omitempty"` // Passphrase
}

type WiFiEncryptionSettings struct {
	Mode string `json:"mode"` // Encryption mode
}

// CellularInterfaceSettings represents Cellular interface settings.
type CellularInterfaceSettings struct {
	Enabled                               bool                              `json:"enabled"`                                         // Whether the interface is enabled
	Name                                  string                            `json:"name"`                                            // Interface name
	Type                                  PlayerNetworkInterfaceType        `json:"type"`                                            // Interface type
	Modems                                []PlayerCellularModemSettings     `json:"modems"`                                          // Modem settings
	Model                                 string                            `json:"model"`                                           // Cellular model
	USBDeviceIds                          []string                          `json:"usbDeviceIds"`                                    // USB device IDs
	SIMS                                  []PlayerCellularModuleSettings    `json:"sims"`                                            // SIM card settings
	MCC                                   string                            `json:"mcc"`                                             // Mobile country code
	MNC                                   string                            `json:"mnc"`                                             // Mobile network code
	Connection                            *PlayerCellularConnectionSettings `json:"connection,omitempty"`                            // Connection settings
	RateLimitDuringInitialDownloads       *int                              `json:"rateLimitDuringInitialDownloads,omitempty"`       // Rate limit during initial downloads
	RateLimitInsideContentDownloadWindow  *int                              `json:"rateLimitInsideContentDownloadWindow,omitempty"`  // Rate limit inside content download window
	RateLimitOutsideContentDownloadWindow *int                              `json:"rateLimitOutsideContentDownloadWindow,omitempty"` // Rate limit outside content download window
	ContentDownloadEnabled                bool                              `json:"contentDownloadEnabled"`                          // Content download enabled
	TextFeedsDownloadEnabled              bool                              `json:"textFeedsDownloadEnabled"`                        // Text feeds download enabled
	MediaFeedsDownloadEnabled             bool                              `json:"mediaFeedsDownloadEnabled"`                       // Media feeds download enabled
	HealthReportingEnabled                bool                              `json:"healthReportingEnabled"`                          // Health reporting enabled
	LogsUploadEnabled                     bool                              `json:"logsUploadEnabled"`                               // Logs upload enabled
}

// GetType returns the type of the network interface settings.
func (c CellularInterfaceSettings) GetType() PlayerNetworkInterfaceType { return c.Type }

// PlayerCellularModemSettings represents a cellular modem setting.
type PlayerCellularModemSettings struct {
	// TODO: Fill in fields from API docs if available
}

// PlayerCellularModuleSettings represents a SIM card setting.
type PlayerCellularModuleSettings struct {
	// TODO: Fill in fields from API docs if available
}

// PlayerCellularConnectionSettings represents cellular connection options.
type PlayerCellularConnectionSettings struct {
	// TODO: Fill in fields from API docs if available
}
