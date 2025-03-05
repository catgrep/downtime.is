package main

import (
	"fmt"
	"strings"
)

// Constants for various time units.
// Borrowed from https://uptime.is/uptime-lsp.txt.
const (
	// Base units remain the same
	hoursInDay   = 24
	daysInWeek   = 7
	monthsInYear = 12
	daysInYear   = 365.2425 // Gregorian calendar average

	// Derived durations in seconds
	secondsInHour    = 3600.0
	secondsInDay     = secondsInHour * hoursInDay
	secondsInWeek    = secondsInDay * daysInWeek
	secondsInMonth   = secondsInDay * (daysInYear / monthsInYear)
	secondsInQuarter = secondsInDay * (daysInYear / 4.0)
	secondsInYear    = secondsInDay * daysInYear
)

func formatSLAPeriod(downtimeSec float64, periodSec float64) string {
	ratio := downtimeSec / periodSec
	uptimePercentage := (1.0 - ratio) * 100.0
	if uptimePercentage < 0 {
		uptimePercentage = 0
	}

	// Format with 4 decimal places then trim trailing zeros
	formatted := fmt.Sprintf("%.4f", uptimePercentage)
	formatted = strings.TrimRight(formatted, "0")

	// Remove trailing decimal point if no decimals remain
	return fmt.Sprintf("%s %%", strings.TrimSuffix(formatted, "."))
}
