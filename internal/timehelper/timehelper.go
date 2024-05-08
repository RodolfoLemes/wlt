package timehelper

import "time"

const (
	// Date
	YYYYMMDD = "2006-01-02"

	// ISO 8601
	YYYYMMDDTHHMMSSZ = "2006-01-02T15:04:05Z"

	// Timestamp Postgres
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
)

// ParseDate parses a date string in YYYY-MM-DD format
func ParseDate(date string) (time.Time, error) {
	return time.Parse(YYYYMMDD, date)
}

// ParseDateTime parses a date string in YYYY-MM-DDTHH:MM:SSZ format
func ParseDateTime(date string) (time.Time, error) {
	return time.Parse(YYYYMMDDTHHMMSSZ, date)
}

// ParseDateTimePostgres parses a date string in YYYY-MM-DD HH:MM:SS format
func ParseDateTimePostgres(date string) (time.Time, error) {
	return time.Parse(YYYYMMDDHHMMSS, date)
}

// FormatDate formats a date in YYYY-MM-DD format
func FormatDate(date time.Time) string {
	return date.Format(YYYYMMDD)
}

// FormatDateTime formats a date in YYYY-MM-DDTHH:MM:SSZ format
func FormatDateTime(date time.Time) string {
	return date.Format(YYYYMMDDTHHMMSSZ)
}

// FormatDateTimePostgres formats a date in YYYY-MM-DD HH:MM:SS format
func FormatDateTimePostgres(date time.Time) string {
	return date.Format(YYYYMMDDHHMMSS)
}
