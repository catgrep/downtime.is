package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseDowntimeDuration(path string) (time.Duration, error) {
	// Handle numeric values directly, assume seconds
	if value, err := strconv.ParseFloat(path, 64); err == nil {
		return time.Duration(value * float64(time.Second)), nil
	}

	// Handle days format (e.g., "5d23h")
	if strings.Contains(path, "d") {
		parts := strings.SplitN(path, "d", 2)
		days, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, fmt.Errorf("invalid days format: %s", path)
		}

		var remaining time.Duration
		if len(parts) > 1 && parts[1] != "" {
			remaining, err = time.ParseDuration(parts[1])
			if err != nil {
				return 0, err
			}
		}

		return (time.Duration(days) * 24 * time.Hour) + remaining, nil
	}

	// Try standard duration parsing
	duration, err := time.ParseDuration(path)
	if err != nil {
		return 0, err
	}
	return duration, nil
}

func formatDuration(seconds float64) string {
	d := time.Duration(seconds * float64(time.Second))
	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	secs := int(d.Seconds()) % 60

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}
	if secs > 0 {
		parts = append(parts, fmt.Sprintf("%ds", secs))
	}
	if len(parts) == 0 {
		return "0s"
	}
	return strings.Join(parts, " ")
}
