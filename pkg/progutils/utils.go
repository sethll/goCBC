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
	"github.com/sethll/goCBC/pkg/hlcalc"
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

// RunHLCalculations performs half-life calculations for multiple substance intakes
// and returns the current body content and time when target amount will be reached.
func RunHLCalculations(results *Results, timesAndAmounts *[]TimeAndAmount, targetAmount, chemHL *float64) {
	slog.Debug("RunHLCalculations started",
		"targetAmount", *targetAmount,
		"chemHL", *chemHL,
		"entryCount", len(*timesAndAmounts))

	currentTime := GetCurrentTime() // timeOps.go
	theoreticalMode := false
	var timeMarker time.Time

	slog.Debug("Initial state",
		"currentTime", currentTime.Format("2006-01-02 15:04:05"),
		"theoreticalMode", theoreticalMode)

	for i, eachItem := range *timesAndAmounts {
		slog.Debug("Processing entry",
			"index", i,
			"timeString", eachItem.TimeString,
			"amount", eachItem.Amount,
			"timeObject", eachItem.TimeObject.Format("2006-01-02 15:04:05"))

		if !theoreticalMode && currentTime.Before(eachItem.TimeObject) {
			slog.Info(
				"Current time is before provided time. Activating theoretical mode",
				"currentTime", currentTime,
				"TimeObject", eachItem.TimeObject,
			)
			slog.Debug("Setting TBCC = BCC", "TBCC", (*results).TheoreticalBodyChemContent, "BCC", (*results).BodyChemContent)
			(*results).TheoreticalBodyChemContent = (*results).BodyChemContent
			theoreticalMode = true
			slog.Debug("Switched to theoretical mode", "theoreticalMode", theoreticalMode)
		} else if !theoreticalMode {
			(*results).LastRealTime = eachItem.TimeObject
			slog.Debug("Updated LastRealTime", "lastRealTime", results.LastRealTime.Format("2006-01-02 15:04:05"))
		}

		oldTotal := results.getChemIngestedTotal(theoreticalMode)
		newTotal := oldTotal + eachItem.Amount
		results.setChemIngestedTotal(theoreticalMode, newTotal)
		slog.Debug("Updated ingested total",
			"theoreticalMode", theoreticalMode,
			"oldTotal", oldTotal,
			"addedAmount", eachItem.Amount,
			"newTotal", newTotal)

		// logic for if first index item, only set initial values and don't process
		if i == 0 {
			timeMarker = eachItem.TimeObject
			results.setBodyChemContent(theoreticalMode, eachItem.Amount)
			slog.Debug("First entry processed",
				"timeMarker", timeMarker.Format("2006-01-02 15:04:05"),
				"initialBodyContent", eachItem.Amount,
				"theoreticalMode", theoreticalMode)
			continue
		}

		etHours := GetElapsedHours(&timeMarker, &eachItem.TimeObject) // timeOps.go
		localBCC := results.getBodyChemContent(theoreticalMode)
		slog.Debug("Before decay calculation", "bodyChemContent", localBCC)

		localBCC = hlcalc.CalcSubstanceInBody(&localBCC, &etHours, chemHL)
		slog.Debug("After decay calculation", "bodyChemContent", localBCC)

		localBCC += eachItem.Amount
		results.setBodyChemContent(theoreticalMode, localBCC)
		slog.Debug("After adding new amount",
			"addedAmount", eachItem.Amount,
			"finalBodyChemContent", localBCC,
			"theoreticalMode", theoreticalMode)

		timeMarker = eachItem.TimeObject
		slog.Debug("Updated timeMarker", "timeMarker", timeMarker.Format("2006-01-02 15:04:05"))
	}

	slog.Debug("Processing final calculations",
		"realIngestedTotal", (*results).ChemIngestedTotal,
		"theoreticalIngestedTotal", (*results).TheoreticalChemIngestedTotal)

	if (*results).ChemIngestedTotal > 0 {
		slog.Debug("Calculating real wearoff time")

		etHours := GetElapsedHours(&results.LastRealTime, &currentTime) // timeOps.go
		localBCC := (*results).BodyChemContent
		slog.Debug("Before final decay calculation", "bodyChemContent", localBCC)

		localBCC = hlcalc.CalcSubstanceInBody(&localBCC, &etHours, chemHL)
		(*results).BodyChemContent = localBCC
		slog.Debug("After final decay calculation", "bodyChemContent", localBCC)

		tValue := hlcalc.CalcTimeToGivenAmt(targetAmount, &localBCC, chemHL)
		(*results).WearoffTime = AddTime(&currentTime, &tValue) // timeOps.go
		slog.Debug("Real wearoff time calculated",
			"tValue", tValue,
			"wearoffTime", (*results).WearoffTime.Format("2006-01-02 15:04:05"))
	}

	if (*results).TheoreticalChemIngestedTotal > 0 {
		slog.Debug("Calculating theoretical wearoff time")
		//localBCC := (*results).TheoreticalBodyChemContent
		slog.Debug("Theoretical body chem content", "bodyChemContent", (*results).TheoreticalBodyChemContent)

		tValue := hlcalc.CalcTimeToGivenAmt(targetAmount, &(*results).TheoreticalBodyChemContent, chemHL)
		(*results).TheoreticalWearoffTime = AddTime(&timeMarker, &tValue) // timeOps.go
		slog.Debug("Theoretical wearoff time calculated",
			"tValue", tValue,
			"timeMarker", timeMarker.Format("2006-01-02 15:04:05"),
			"theoreticalWearoffTime", (*results).TheoreticalWearoffTime.Format("2006-01-02 15:04:05"))
	}

	slog.Debug("RunHLCalculations completed",
		"finalBodyChemContent", (*results).BodyChemContent,
		"finalTheoreticalBodyChemContent", (*results).TheoreticalBodyChemContent,
		"wearoffTime", (*results).WearoffTime.Format("2006-01-02 15:04:05"),
		"theoreticalWearoffTime", (*results).TheoreticalWearoffTime.Format("2006-01-02 15:04:05"))
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
