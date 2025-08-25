package progutils

/*
	goCBC
	Copyright (C) 2025  Seth L

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"
)

// GetBeginningOfToday returns a time.Time representing the start of the day (00:00:00)
// for the given current time.
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

// AddTime returns the duration (float, hours) plus the start
func AddTime(start *time.Time, duration *float64) time.Time {
	return (*start).Add(time.Duration(*duration * float64(time.Hour)))
}

// GetCurrentTime returns the current local time.
func GetCurrentTime() time.Time {
	currentTime := time.Now()
	slog.Debug("Retrieved current time", "currentTime", currentTime.Format("2006-01-02 15:04:05"))
	return currentTime
}

// GetElapsedHours returns the number of hours passed from input1 to input2
func GetElapsedHours(start, end *time.Time) float64 {
	elapsedTime := (*end).Sub(*start)
	etHours := elapsedTime.Hours()
	slog.Debug("Time calculations",
		"elapsedTime", elapsedTime,
		"etHours", etHours,
		"previousTimeMarker", (*start).Format("2006-01-02 15:04:05"),
		"currentTimeMarker", (*end).Format("2006-01-02 15:04:05"))
	return etHours
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
