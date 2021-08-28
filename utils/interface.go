package utils

import "time"

// IDateHelper interface
type IDateHelper interface {
	Now() time.Time
	AgeFromDate(time time.Time) (int, error)
	Parse(layout, date string) (time.Time, error)
	Hours(sec int) int
	Days(sec int) int
}