package wu

import (
	"time"
	"wu/wtime"
)

// Global common functions

// use it to set global time check for app expires
func TimeBomb() {
	wtime.CheckTimeToExit(2025, 1, 1) // limit app run after 2025-01-01
}

// SetTimeBomb(2021,1,1)
func SetTimeBomb(year int, month time.Month, day int) {
	// limit app run after year-month-day
	wtime.CheckTimeToExit(year, month, day)
}
