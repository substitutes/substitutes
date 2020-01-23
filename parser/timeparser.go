package parser

import (
	"strings"
	"time"
)

// UntisTimeLayout defines the time format as used in the Untis export
const UntisTimeLayout = "2.1.2006 15:04"

// UntisDateLayout defines the date format as used in the Untis export
const UntisDateLayout = "2.1.2006"

// WeekdayLanguageMapping is a translation mapping for days of the week from German to English
var WeekdayLanguageMapping = map[string]string{
	"Montag":     "Monday",
	"Dienstag":   "Tuesday",
	"Mittwoch":   "Wednesday",
	"Donnerstag": "Thursday",
	"Freitag":    "Friday",
	"Samstag":    "Saturday", // Optional
	"Sonntag":    "Sunday",   // Optional
}

// ParseUntisTime parses an untis time stamp to a golang time
func ParseUntisTime(untisTime string) (time.Time, error) {
	return time.Parse(UntisTimeLayout, untisTime)
}

// ParseUntisDate parses an untis date to golang time
func ParseUntisDate(untisDate string) (time.Time, error) {
	t, err := time.Parse(UntisDateLayout, strings.Split(strings.Replace(untisDate, "Vertretungen  ", "", -1), " ")[0])
	if err != nil {
		return t, err
	}
	return t.AddDate(2019, 0, 0), nil
}
