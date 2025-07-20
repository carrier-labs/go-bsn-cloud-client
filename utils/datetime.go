// Package utils provides utility types and functions for the BSN.Cloud API client.
package utils

import (
	"fmt"
	"strings"
	"time"
)

// BsnTime is a custom time type for BSN.Cloud API that handles missing timezone.
type BsnTime struct {
	time.Time
}

// UnmarshalJSON parses time in RFC3339 or "2006-01-02T15:04:05.000" (no timezone, assume UTC).
func (bt *BsnTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	if s == "" || s == "null" {
		bt.Time = time.Time{}
		return nil
	}
	// Try RFC3339 with/without milliseconds
	layouts := []string{
		"2006-01-02T15:04:05.000Z07:00",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05",
	}
	var err error
	for _, layout := range layouts {
		bt.Time, err = time.Parse(layout, s)
		if err == nil {
			if layout == "2006-01-02T15:04:05.000" || layout == "2006-01-02T15:04:05" {
				bt.Time = bt.Time.UTC()
			}
			return nil
		}
	}
	return fmt.Errorf("BsnTime: could not parse time: %q", s)
}

// MarshalJSON outputs time in "2006-01-02T15:04:05.000" (no timezone, always UTC).
func (bt BsnTime) MarshalJSON() ([]byte, error) {
	if bt.Time.IsZero() {
		return []byte("null"), nil
	}
	s := bt.Time.UTC().Format("2006-01-02T15:04:05.000")
	return []byte("\"" + s + "\""), nil
}
