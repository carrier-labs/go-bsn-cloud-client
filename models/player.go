// Package models contains shared data structures for the BSN.Cloud API client.
package models

import (
	"encoding/json"

	"github.com/carrier-labs/go-bsn-cloud-client/utils"
)

// Player represents a player (device) in BSN.Cloud.
type Player struct {
	Id               int                `json:"id"`               // Unique player ID
	Serial           string             `json:"serial"`           // Serial number
	Model            PlayerModel        `json:"model"`            // Player model
	Family           PlayerFamily       `json:"family"`           // Player family
	RegistrationDate utils.BsnTime      `json:"registrationDate"` // Registration date
	LastModifiedDate utils.BsnTime      `json:"lastModifiedDate"` // Last modification date
	Settings         PlayerSettings     `json:"settings"`         // Player settings
	Status           PlayerFullStatus   `json:"status"`           // Player status
	Subscription     PlayerSubscription `json:"subscription"`     // Player subscription
	TaggedGroups     []TaggedGroupInfo  `json:"taggedGroups"`     // Tagged groups
	Permissions      []Permission       `json:"permissions"`      // Permissions
}

// PlayerListResponse is a list response for players.
type PlayerListResponse struct {
	Items []Player `json:"items"` // List of players
}

// PlayerSettings represents the settings entity for a player in BSN.Cloud.
type PlayerSettings struct {
	Name                string                         `json:"name"`                       // Player name
	Description         string                         `json:"description"`                // Player description
	ConcatNameAndSerial bool                           `json:"concatNameAndSerial"`        // Whether to concatenate name and serial
	SetupType           DeviceSetupType                `json:"setupType"`                  // Setup type
	Group               *GroupInfo                     `json:"group,omitempty"`            // Group info
	BrightWall          *BrightWallScreenInfo          `json:"brightWall,omitempty"`       // BrightWall info
	Timezone            string                         `json:"timezone"`                   // Timezone
	Screen              *DeviceScreenSettings          `json:"screen,omitempty"`           // Screen settings
	Synchronization     *PlayerSynchronizationSettings `json:"synchronization,omitempty"`  // Synchronization settings
	Network             *PlayerNetworkSettings         `json:"network,omitempty"`          // Network settings
	Beacons             []DeviceBeacon                 `json:"beacons,omitempty"`          // Beacons
	Location            *DeviceLocation                `json:"location,omitempty"`         // Location
	Screenshots         *PlayerScreenshotsSettings     `json:"screenshots,omitempty"`      // Screenshot settings
	Logging             *DeviceLogsSettings            `json:"logging,omitempty"`          // Logging settings
	LWS                 *LocalWebServerSettings        `json:"lws,omitempty"`              // Local web server settings
	LDWS                *DiagnosticWebServerSettings   `json:"ldws,omitempty"`             // Diagnostic web server settings
	LastModifiedDate    *utils.BsnTime                 `json:"lastModifiedDate,omitempty"` // Last modification date
}

// DeviceSetupType is an enum for player setup types.
type DeviceSetupType string

const (
	// DeviceSetupTypeStandalone represents standalone setup.
	DeviceSetupTypeStandalone DeviceSetupType = "Standalone"
	// DeviceSetupTypeBSN represents BSN setup.
	DeviceSetupTypeBSN DeviceSetupType = "BSN"
	// DeviceSetupTypeLFN represents LFN setup.
	DeviceSetupTypeLFN DeviceSetupType = "LFN"
	// DeviceSetupTypeSFN represents SFN setup.
	DeviceSetupTypeSFN DeviceSetupType = "SFN"
	// DeviceSetupTypePartnerApplication represents partner application setup.
	DeviceSetupTypePartnerApplication DeviceSetupType = "PartnerApplication"
	// DeviceSetupTypeUnknown represents an unknown setup type.
	DeviceSetupTypeUnknown DeviceSetupType = "Unknown"
)

// DeviceScreenSettings represents screen settings for a player.
type DeviceScreenSettings struct {
	IdleColor string `json:"idleColor"` // Idle color
	SplashUrl string `json:"splashUrl"` // Splash screen URL
}

// PlayerSynchronizationSettings represents synchronization settings for a player.
type PlayerSynchronizationSettings struct {
	Status   *PlayerStatusSynchronizationSettings   `json:"status,omitempty"`   // Status sync settings
	Settings *PlayerSettingsSynchronizationSettings `json:"settings,omitempty"` // Settings sync settings
	Schedule *PlayerScheduleSynchronizationSettings `json:"schedule,omitempty"` // Schedule sync settings
	Content  *PlayerContentSynchronizationSettings  `json:"content,omitempty"`  // Content sync settings
}

// PlayerStatusSynchronizationSettings represents status sync settings.
type PlayerStatusSynchronizationSettings struct {
	Period TimeSpan `json:"period"` // Sync period
}

// PlayerSettingsSynchronizationSettings represents settings sync settings.
type PlayerSettingsSynchronizationSettings struct {
	Period TimeSpan `json:"period"` // Sync period
}

// PlayerScheduleSynchronizationSettings represents schedule sync settings.
type PlayerScheduleSynchronizationSettings struct {
	Period TimeSpan `json:"period"` // Sync period
}

// PlayerContentSynchronizationSettings represents content sync settings.
type PlayerContentSynchronizationSettings struct {
	Start TimeSpan `json:"start"` // Sync start time
	End   TimeSpan `json:"end"`   // Sync end time
}

// DeviceLocation represents the location of a player.
type DeviceLocation struct {
	PlaceId                 string   `json:"placeId"`                 // Place ID
	GPSLatitude             *float64 `json:"gpsLatitude,omitempty"`   // GPS latitude
	GPSLongitude            *float64 `json:"gpsLongitude,omitempty"`  // GPS longitude
	Country                 string   `json:"country"`                 // Country code
	CountryLongName         string   `json:"countryLongName"`         // Country name
	AdminAreaLevel1         string   `json:"adminAreaLevel1"`         // Admin area level 1 code
	AdminAreaLevel1LongName string   `json:"adminAreaLevel1LongName"` // Admin area level 1 name
	AdminAreaLevel2         string   `json:"adminAreaLevel2"`         // Admin area level 2 code
	AdminAreaLevel2LongName string   `json:"adminAreaLevel2LongName"` // Admin area level 2 name
	Locality                string   `json:"locality"`                // Locality code
	LocalityLongName        string   `json:"localityLongName"`        // Locality name
	Path                    string   `json:"path"`                    // Path code
	PathLongName            string   `json:"pathLongName"`            // Path name
}

// PlayerScreenshotsSettings represents screenshot settings for a player.
type PlayerScreenshotsSettings struct {
	Interval    TimeSpan          `json:"interval"`    // Screenshot interval
	CountLimit  uint16            `json:"countLimit"`  // Max screenshot count
	Quality     byte              `json:"quality"`     // Screenshot quality
	Orientation ScreenOrientation `json:"orientation"` // Screenshot orientation
}

// ScreenOrientation is an enum for screenshot orientation.
type ScreenOrientation string

const (
	// ScreenOrientationUnknown represents an unknown orientation.
	ScreenOrientationUnknown ScreenOrientation = "Unknown"
	// ScreenOrientationLandscape represents landscape orientation.
	ScreenOrientationLandscape ScreenOrientation = "Landscape"
	// ScreenOrientationPortraitBottomLeft represents portrait (bottom left) orientation.
	ScreenOrientationPortraitBottomLeft ScreenOrientation = "PortraitBottomLeft"
	// ScreenOrientationPortraitBottomRight represents portrait (bottom right) orientation.
	ScreenOrientationPortraitBottomRight ScreenOrientation = "PortraitBottomRight"
)

// DeviceLogsSettings represents logging settings for a player.
type DeviceLogsSettings struct {
	EnableDiagnosticLog bool      `json:"enableDiagnosticLog"`  // Enable diagnostic log
	EnableEventLog      bool      `json:"enableEventLog"`       // Enable event log
	EnablePlaybackLog   bool      `json:"enablePlaybackLog"`    // Enable playback log
	EnableStateLog      bool      `json:"enableStateLog"`       // Enable state log
	EnableVariableLog   bool      `json:"enableVariableLog"`    // Enable variable log
	UploadAtBoot        bool      `json:"uploadAtBoot"`         // Upload logs at boot
	UploadTime          *TimeSpan `json:"uploadTime,omitempty"` // Log upload time
}

// LocalWebServerSettings represents local web server settings for a player.
type LocalWebServerSettings struct {
	Username                  string `json:"username"`                  // Web server username
	Password                  string `json:"password"`                  // Web server password
	EnableUpdateNotifications bool   `json:"enableUpdateNotifications"` // Enable update notifications
}

// DiagnosticWebServerSettings represents diagnostic web server settings for a player.
type DiagnosticWebServerSettings struct {
	Password string `json:"password"` // Web server password
}

// PlayerFullStatus represents the full status of a player.
type PlayerFullStatus struct {
	Group                    GroupInfo                   `json:"group"`                      // Group info
	BrightWall               *BrightWallScreenInfo       `json:"brightWall,omitempty"`       // BrightWall info
	Presentation             []PresentationInfo          `json:"presentation"`               // Presentations
	Script                   PlayerScript                `json:"script"`                     // Player script
	Firmware                 FirmwareInfo                `json:"firmware"`                   // Firmware info
	Storage                  []StorageStatus             `json:"storage"`                    // Storage status
	Network                  PlayerNetworkStatus         `json:"network"`                    // Network status
	Uptime                   TimeSpan                    `json:"uptime"`                     // Uptime
	CurrentSettingsTimestamp utils.BsnTime               `json:"currentSettingsTimestamp"`   // Current settings timestamp
	CurrentScheduleTimestamp utils.BsnTime               `json:"currentScheduleTimestamp"`   // Current schedule timestamp
	Timezone                 string                      `json:"timezone"`                   // Timezone
	Health                   PlayerHealthStatus          `json:"health"`                     // Health status
	LastModifiedDate         *utils.BsnTime              `json:"lastModifiedDate,omitempty"` // Last modification date
	Synchronization          PlayerSynchronizationStatus `json:"synchronization"`            // Synchronization status
}

// UnmarshalJSON handles presentation as either an object, array, or null.
func (p *PlayerFullStatus) UnmarshalJSON(data []byte) error {
	type Alias PlayerFullStatus
	aux := &struct {
		Presentation json.RawMessage `json:"presentation"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// Handle presentation as array, object, or null
	var arr []PresentationInfo
	if len(aux.Presentation) == 0 || string(aux.Presentation) == "null" {
		p.Presentation = nil
	} else if aux.Presentation[0] == '{' {
		var single PresentationInfo
		if err := json.Unmarshal(aux.Presentation, &single); err != nil {
			return err
		}
		p.Presentation = []PresentationInfo{single}
	} else if aux.Presentation[0] == '[' {
		if err := json.Unmarshal(aux.Presentation, &arr); err != nil {
			return err
		}
		p.Presentation = arr
	} else {
		p.Presentation = nil
	}
	// Copy other fields
	p.Group = aux.Group
	p.BrightWall = aux.BrightWall
	p.Script = aux.Script
	p.Firmware = aux.Firmware
	p.Storage = aux.Storage
	p.Network = aux.Network
	p.Uptime = aux.Uptime
	p.CurrentSettingsTimestamp = aux.CurrentSettingsTimestamp
	p.CurrentScheduleTimestamp = aux.CurrentScheduleTimestamp
	p.Timezone = aux.Timezone
	p.Health = aux.Health
	p.LastModifiedDate = aux.LastModifiedDate
	p.Synchronization = aux.Synchronization
	return nil
}

// GroupInfo represents group information.
type GroupInfo struct {
	Id   int    `json:"id"`   // Group ID
	Name string `json:"name"` // Group name
}

// BrightWallScreenInfo represents BrightWall screen info.
type BrightWallScreenInfo struct {
	Id     *int   `json:"id,omitempty"` // BrightWall ID
	Name   string `json:"name"`         // BrightWall name
	Screen byte   `json:"screen"`       // Screen index
	Link   string `json:"link"`         // BrightWall link
}

// PresentationInfo represents a presentation on the player.
type PresentationInfo struct {
	Id   int    `json:"id"`   // Presentation ID
	Name string `json:"name"` // Presentation name
	Link string `json:"link"` // Presentation link
}

// PlayerScript represents the script running on the player.
type PlayerScript struct {
	Type    ScriptType         `json:"type"`    // Script type
	Version string             `json:"version"` // Script version
	Plugins []ScriptPluginInfo `json:"plugins"` // Script plugins
}

// ScriptType is an enum for script types.
type ScriptType string

const (
	// ScriptTypeSetup represents a setup script.
	ScriptTypeSetup ScriptType = "Setup"
	// ScriptTypeAutorun represents an autorun script.
	ScriptTypeAutorun ScriptType = "Autorun"
	// ScriptTypeRecovery represents a recovery script.
	ScriptTypeRecovery ScriptType = "Recovery"
	// ScriptTypeCustom represents a custom script.
	ScriptTypeCustom ScriptType = "Custom"
	// ScriptTypeUnknown represents an unknown script type.
	ScriptTypeUnknown ScriptType = "Unknown"
)

// ScriptPluginInfo represents a script plugin.
type ScriptPluginInfo struct {
	FileName string `json:"fileName"` // Plugin file name
	FileSize uint   `json:"fileSize"` // Plugin file size
	FileHash string `json:"fileHash"` // Plugin file hash
}

// FirmwareInfo represents firmware information.
type FirmwareInfo struct {
	Version string `json:"version"` // Firmware version
}

// StorageStatus represents the storage status entity.
type StorageStatus struct {
	Interface StorageInterface `json:"interface"` // Storage interface
	System    FileSystem       `json:"system"`    // File system
	Access    []AccessMode     `json:"access"`    // Access modes
	Stats     StorageStats     `json:"stats"`     // Storage stats
}

// UnmarshalJSON handles access as either a CSV string or array.
func (s *StorageStatus) UnmarshalJSON(data []byte) error {
	type Alias StorageStatus
	aux := &struct {
		Access interface{} `json:"access"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.Interface = aux.Interface
	s.System = aux.System
	s.Stats = aux.Stats

	s.Access = nil
	switch v := aux.Access.(type) {
	case string:
		if v == "" {
			s.Access = nil
		} else {
			for _, part := range splitAndTrim(v, ",") {
				s.Access = append(s.Access, AccessMode(part))
			}
		}
	case []interface{}:
		for _, item := range v {
			if str, ok := item.(string); ok {
				s.Access = append(s.Access, AccessMode(str))
			}
		}
	}
	return nil
}

// StorageInterface is an enum for storage interfaces.
type StorageInterface string

const (
	// StorageInterfaceInternal represents internal storage.
	StorageInterfaceInternal StorageInterface = "Internal"
	// StorageInterfaceTmp represents temporary storage.
	StorageInterfaceTmp StorageInterface = "Tmp"
	// StorageInterfaceFlash represents flash storage.
	StorageInterfaceFlash StorageInterface = "Flash"
	// StorageInterfaceSD1 represents SD1 storage.
	StorageInterfaceSD1 StorageInterface = "SD1"
	// StorageInterfaceUSB1 represents USB1 storage.
	StorageInterfaceUSB1 StorageInterface = "USB1"
	// StorageInterfaceUnknown represents an unknown storage interface.
	StorageInterfaceUnknown StorageInterface = "Unknown"
)

// FileSystem is an enum for file systems.
type FileSystem string

const (
	// FileSystemExFAT represents exFAT file system.
	FileSystemExFAT FileSystem = "exFAT"
	// FileSystemExt3 represents ext3 file system.
	FileSystemExt3 FileSystem = "ext3"
	// FileSystemExt4 represents ext4 file system.
	FileSystemExt4 FileSystem = "ext4"
	// FileSystemFAT12 represents FAT12 file system.
	FileSystemFAT12 FileSystem = "FAT12"
	// FileSystemFAT16 represents FAT16 file system.
	FileSystemFAT16 FileSystem = "FAT16"
	// FileSystemFAT32 represents FAT32 file system.
	FileSystemFAT32 FileSystem = "FAT32"
	// FileSystemHFS represents HFS file system.
	FileSystemHFS FileSystem = "HFS"
	// FileSystemHFSplus represents HFS+ file system.
	FileSystemHFSplus FileSystem = "HFSplus"
	// FileSystemNTFS represents NTFS file system.
	FileSystemNTFS FileSystem = "NTFS"
	// FileSystemUnknown represents an unknown file system.
	FileSystemUnknown FileSystem = "Unknown"
)

// AccessMode is an enum for access modes.
type AccessMode string

const (
	// AccessModeRead represents read access.
	AccessModeRead AccessMode = "Read"
	// AccessModeWrite represents write access.
	AccessModeWrite AccessMode = "Write"
	// AccessModeUnknown represents an unknown access mode.
	AccessModeUnknown AccessMode = "Unknown"
)

// StorageStats represents diagnostic storage status information. Accepts any fields from the server.
type StorageStats map[string]any

// PlayerHealthStatus is an enum for health status.
type PlayerHealthStatus string

const (
	// PlayerHealthStatusNormal represents normal health status.
	PlayerHealthStatusNormal PlayerHealthStatus = "Normal"
	// PlayerHealthStatusWarning represents warning health status.
	PlayerHealthStatusWarning PlayerHealthStatus = "Warning"
	// PlayerHealthStatusError represents error health status.
	PlayerHealthStatusError PlayerHealthStatus = "Error"
	// PlayerHealthStatusUnknown represents an unknown health status.
	PlayerHealthStatusUnknown PlayerHealthStatus = "Unknown"
)

// PlayerSynchronizationStatus is a placeholder for synchronization status.
type PlayerSynchronizationStatus struct {
	Settings PlayerSettingsSynchronizationStatus `json:"settings"` // Settings synchronization status
	Schedule PlayerScheduleSynchronizationStatus `json:"schedule"` // Schedule synchronization status
	Content  PlayerContentSynchronizationStatus  `json:"content"`  // Content synchronization status
}

type PlayerSettingsSynchronizationStatus struct {
	Enabled bool `json:"enabled"` // Whether settings synchronization is enabled
}

type PlayerScheduleSynchronizationStatus struct {
	Enabled bool `json:"enabled"` // Whether schedule synchronization is enabled
}

type PlayerContentSynchronizationStatus struct {
	Enabled bool `json:"enabled"` // Whether content synchronization is enabled
}

// PlayerSubscription represents a player subscription entity.
type PlayerSubscription struct {
	Id               int                      `json:"id"`                       // Subscription ID
	Device           DeviceInfo               `json:"device"`                   // Device info
	Type             PlayerSubscriptionType   `json:"type"`                     // Subscription type
	ActivityPeriod   TimeSpan                 `json:"activityPeriod"`           // Activity period
	Status           DeviceSubscriptionStatus `json:"status"`                   // Subscription status
	CreationDate     utils.BsnTime            `json:"creationDate"`             // Creation date
	ActivationDate   *utils.BsnTime           `json:"activationDate,omitempty"` // Activation date
	SuspensionDate   *utils.BsnTime           `json:"suspensionDate,omitempty"` // Suspension date
	ExpirationDate   *utils.BsnTime           `json:"expirationDate,omitempty"` // Expiration date
	LastModifiedDate utils.BsnTime            `json:"lastModifiedDate"`         // Last modified date
}

// DeviceInfo represents minimal device info for a subscription.
type DeviceInfo struct {
	Id     int    `json:"id"`     // Device ID
	Serial string `json:"serial"` // Device serial number
}

// PlayerSubscriptionType is an enum for subscription types.
type PlayerSubscriptionType string

const (
	// PlayerSubscriptionTypeContent represents a content subscription.
	PlayerSubscriptionTypeContent PlayerSubscriptionType = "Content"
	// PlayerSubscriptionTypeControl represents a control subscription.
	PlayerSubscriptionTypeControl PlayerSubscriptionType = "Control"
	// PlayerSubscriptionTypeUnknown represents an unknown subscription type.
	PlayerSubscriptionTypeUnknown PlayerSubscriptionType = "Unknown"
)

// DeviceSubscriptionStatus is an enum for subscription status.
type DeviceSubscriptionStatus string

const (
	// DeviceSubscriptionStatusActive represents an active subscription.
	DeviceSubscriptionStatusActive DeviceSubscriptionStatus = "Active"
	// DeviceSubscriptionStatusSuspending represents a suspending subscription.
	DeviceSubscriptionStatusSuspending DeviceSubscriptionStatus = "Suspending"
	// DeviceSubscriptionStatusSuspended represents a suspended subscription.
	DeviceSubscriptionStatusSuspended DeviceSubscriptionStatus = "Suspended"
	// DeviceSubscriptionStatusUnknown represents an unknown subscription status.
	DeviceSubscriptionStatusUnknown DeviceSubscriptionStatus = "Unknown"
)

// TaggedGroupInfo represents a tagged group for a player.
type TaggedGroupInfo struct {
	Id   int               `json:"id"`   // Group ID
	Name string            `json:"name"` // Group name
	Tags map[string]string `json:"tags"` // Tags
}

// Permission represents a permission entity for a player.
type Permission struct {
	EntityId     *int          `json:"entityId,omitempty"` // Entity ID
	OperationUID string        `json:"operationUID"`       // Operation UID
	Principal    Principal     `json:"principal"`          // Principal
	User         string        `json:"user,omitempty"`     // User
	Role         string        `json:"role,omitempty"`     // Role
	IsFixed      bool          `json:"isFixed"`            // Whether fixed
	IsInherited  bool          `json:"isInherited"`        // Whether inherited
	IsAllowed    bool          `json:"isAllowed"`          // Whether allowed
	CreationDate utils.BsnTime `json:"creationDate"`       // Creation date
}

// Principal represents the principal entity for permissions.
type Principal struct {
	Name     string        `json:"name"`     // Principal name
	IsCustom bool          `json:"isCustom"` // Whether custom
	Type     PrincipalType `json:"type"`     // Principal type
	Id       int           `json:"id"`       // Principal ID
}

// PrincipalType is an enum for principal types.
type PrincipalType string

const (
	// PrincipalTypeUser represents a user principal.
	PrincipalTypeUser PrincipalType = "User"
	// PrincipalTypeRole represents a role principal.
	PrincipalTypeRole PrincipalType = "Role"
	// PrincipalTypeUnknown represents an unknown principal type.
	PrincipalTypeUnknown PrincipalType = "Unknown"
)

// Network represents a network in BSN.Cloud.
type Network struct {
	Id   string `json:"id"`   // Network ID
	Name string `json:"name"` // Network name
}

// NetworkListResponse is a list response for networks.
type NetworkListResponse struct {
	Items []Network `json:"items"` // List of networks
}

// TimeSpan represents a period of time as a string.
type TimeSpan string
