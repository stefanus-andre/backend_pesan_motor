// Modify your DTO package to handle date parsing properly
package dto

import (
	"strings"
	"time"
)

// Custom time type to handle just dates without time components
type Date struct {
	time.Time
}

// UnmarshalJSON custom unmarshaler for Date type
func (d *Date) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string
	dateStr := strings.Trim(string(data), "\"")

	// If empty, return early
	if dateStr == "" || dateStr == "null" {
		return nil
	}

	// Parse date-only format
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}

	d.Time = parsedTime
	return nil
}
