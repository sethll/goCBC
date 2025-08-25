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
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

// TimeAndAmount represents a time entry with an associated substance amount.
type TimeAndAmount struct {
	TimeString string
	Amount     float64
	TimeObject time.Time
}

var (
	// LogLevelSelector maps verbosity levels to log levels.
	LogLevelSelector = map[int]log.Level{
		0: log.ErrorLevel,
		1: log.WarnLevel,
		2: log.InfoLevel,
		3: log.DebugLevel,
	}
)

// GetTimesAndAmounts parses command-line time:amount strings into TimeAndAmount structs.
// Input format should be "HHMM:amount" (e.g., "1100:300").
func GetTimesAndAmounts(inputs *[]string) (returnList []TimeAndAmount) {
	slog.Debug("Processing time and amount inputs", "inputCount", len(*inputs))
	for _, eachItem := range *inputs {
		splitStrings := strings.Split(eachItem, ":")
		var eachTAndA = TimeAndAmount{
			TimeString: splitStrings[0],
			Amount:     StringToFloat(&splitStrings[1]),
		}

		// Validate that time meets expected format before appending
		if minutesHours, err := ValidateTime(&eachTAndA); err == nil {
			timeObject := getTimeObject(&minutesHours)
			eachTAndA.TimeObject = timeObject
			returnList = append(returnList, eachTAndA) // SET RETURN VAR

			slog.Debug("Parsed and validated time entry", "timeString", eachTAndA.TimeString, "amount", eachTAndA.Amount, "timeObject", timeObject.Format("15:04"))
			slog.Info("Accepted input", "eachTAndA", eachTAndA)
		} else {
			slog.Error(err.Error())
			slog.Warn("Discarding entry", "TimeString", eachTAndA.TimeString)
		}
	}
	returnList = sortTimeEntries(returnList)
	slog.Debug("Completed processing inputs", "validEntries", len(returnList), "totalInputs", len(*inputs))
	return
}

// StringToTitleCase converts a string to title case (first letter uppercase, rest lowercase).
func StringToTitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

// StringToFloat converts a string to a float64, returning 0.0 if conversion fails.
func StringToFloat(input *string) float64 {
	if f64, err := strconv.ParseFloat(*input, 64); err == nil {
		slog.Debug("Float parsed from string", "input", *input, "f64", f64)
		return f64
	}

	defaultFloat := 0.0
	slog.Warn("Unable to parse input. Using default", "input", input, "default", defaultFloat)

	return defaultFloat
}
