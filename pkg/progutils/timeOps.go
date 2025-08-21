package progutils

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

func GetBeginningOfToday(currentTime *time.Time) time.Time {
	slog.Debug("Getting beginning of today", "currentTime", (*currentTime).Format("2006-01-02 15:04:05"))
	beginningOfDay := time.Date(
		(*currentTime).Year(),
		(*currentTime).Month(),
		(*currentTime).Day(),
		0,
		0,
		0,
		0,
		time.Local)
	slog.Debug("Beginning of day calculated", "beginningOfDay", beginningOfDay.Format("2006-01-02 15:04:05"))
	return beginningOfDay
}

func GetCurrentTime() time.Time {
	currentTime := time.Now()
	slog.Debug("Retrieved current time", "currentTime", currentTime.Format("2006-01-02 15:04:05"))
	return currentTime
}

// ValidateTime validates that the time input is a proper 24-hour format time (HHMM).
func ValidateTime(tAndA *TimeAndAmount) ([]int, error) {
	var retIntList []int

	if tAndA == nil {
		slog.Error("ValidateTime Uninitialized TimeAndAmount struct pointer")
		return retIntList, fmt.Errorf("invalid operation attempted")
	}

	// Check if it's exactly 4 digits
	if len(tAndA.TimeString) != 4 {
		slog.Warn("Time string must be exactly 4 digits", "time", tAndA.TimeString, "length", len(tAndA.TimeString))
		return retIntList, fmt.Errorf("%s is not a valid time: must be 4 digits (HHMM)", tAndA.TimeString)
	}

	// Check if all characters are digits
	givenInt, parseIntErr := strconv.ParseInt(tAndA.TimeString, 10, 64)
	if parseIntErr != nil {
		slog.Warn("Time string contains non-numeric characters", "time", tAndA.TimeString, "error", parseIntErr)
		return retIntList, fmt.Errorf("%s is not a valid time: must contain only digits", tAndA.TimeString)
	}

	// Extract hours and minutes
	hours := int(givenInt / 100)
	minutes := int(givenInt % 100)

	// Validate hours (0-23)
	if hours < 0 || hours > 23 {
		slog.Warn("Invalid hours in time", "time", tAndA.TimeString, "hours", hours)
		return retIntList, fmt.Errorf("%s is not a valid time: hours must be 00-23", tAndA.TimeString)
	}

	// Validate minutes (0-59)
	if minutes < 0 || minutes > 59 {
		slog.Warn("Invalid minutes in time", "time", tAndA.TimeString, "minutes", minutes)
		return retIntList, fmt.Errorf("%s is not a valid time: minutes must be 00-59", tAndA.TimeString)
	}

	retIntList = []int{hours, minutes}
	slog.Debug("TimeString input is valid", "time", tAndA.TimeString, "hours", hours, "minutes", minutes)
	return retIntList, nil
}

func getTimeObject(input *[]int) time.Time {
	slog.Debug("Creating time object from hours and minutes", "hours", (*input)[0], "minutes", (*input)[1])
	currentTime := GetCurrentTime()

	timeObject := time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		(*input)[0],
		(*input)[1],
		0,
		0,
		time.Local,
	)
	slog.Debug("Time object created", "timeObject", timeObject.Format("2006-01-02 15:04:05"))
	return timeObject
}
