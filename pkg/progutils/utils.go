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
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/sethll/goCBC/pkg/hlcalc"
	"github.com/sethll/goCBC/pkg/progmeta"
)

type StylesType struct {
	Bedtime  lipgloss.Style
	Caffeine lipgloss.Style
	Header   lipgloss.Style
}

type TimeAndAmount struct {
	TimeString string
	Amount     float64
	TimeObject time.Time
}

var (
	Styles = StylesType{
		Bedtime: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#87CEEB")),
		Caffeine: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500")),
		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()),
	}

	LogLevelSelector = map[int]log.Level{
		0: log.ErrorLevel,
		1: log.WarnLevel,
		2: log.InfoLevel,
		3: log.DebugLevel,
	}
)

func GenerateOutputTable(chemInBody *float64, bedTime *time.Time, sleepTarget *string, chem *string) *table.Table {
	slog.Debug("Generating output table", "chemInBody", *chemInBody, "bedTime", (*bedTime).Format("2006-01-02 15:04"), "sleepTarget", *sleepTarget)
	rows := [][]string{
		{
			fmt.Sprintf(
				"%s remaining in system:",
				Styles.Caffeine.Render(StringToTitleCase(*chem)),
			),
			Styles.Caffeine.Render(
				fmt.Sprintf(
					"~%.0fmg",
					math.Round(*chemInBody),
				),
			),
		},
		{
			fmt.Sprintf(
				"Reach target (%s) for %s at:",
				Styles.Caffeine.Render(
					fmt.Sprintf("%smg", *sleepTarget),
				),
				Styles.Bedtime.Render("sleep"),
			),
			Styles.Bedtime.Render(
				(*bedTime).Format("2006-01-02 15:04"),
			),
		},
	}
	generatedTable := table.New().Border(lipgloss.HiddenBorder()).Rows(rows...)
	slog.Debug("Output table generated successfully", "rowCount", len(rows))
	return generatedTable
}

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
	slog.Debug("Completed processing inputs", "validEntries", len(returnList), "totalInputs", len(*inputs))
	return
}

func PrintProgHeader() {
	slog.Debug("Printing program header", "progName", progmeta.ProgName, "version", progmeta.AllVersionBuildRuntimeInfo())
	fmt.Println(
		Styles.Header.Render(
			fmt.Sprintf(
				"%s\t© %s %s\n%s",
				progmeta.ProgName,
				progmeta.CopyrightYear,
				progmeta.Author,
				progmeta.AllVersionBuildRuntimeInfo(),
			),
		),
	)
}

func RunHLCalculations(timesAndAmounts *[]TimeAndAmount, targetAmount, chemHL *float64) (bodyChemContent float64, chemTargetReachedTime time.Time) {
	firstItem, remainingItems := (*timesAndAmounts)[0], (*timesAndAmounts)[1:]
	previousTimeMarker := firstItem.TimeObject
	bodyChemContent = firstItem.Amount

	for _, eachTAndA := range remainingItems {
		elapsedTime := eachTAndA.TimeObject.Sub(previousTimeMarker)
		bodyChemContent = hlcalc.CalcSubstanceInBody(bodyChemContent, elapsedTime.Hours(), *chemHL)
		bodyChemContent += eachTAndA.Amount
		previousTimeMarker = eachTAndA.TimeObject
	}

	currentTime := GetCurrentTime()
	elapsedTime := currentTime.Sub(previousTimeMarker)
	bodyChemContent = hlcalc.CalcSubstanceInBody(bodyChemContent, elapsedTime.Hours(), *chemHL)

	tValue := hlcalc.CalcTimeToGivenAmt(targetAmount, &bodyChemContent, chemHL)
	chemTargetReachedTime = currentTime.Add(time.Duration(tValue * float64(time.Hour)))
	return
}

func SortTimeEntries(timesAndAmounts []TimeAndAmount) []TimeAndAmount {
	// Sort the slice of TimeAndAmount structs by TimeString
	sort.Slice(timesAndAmounts, func(i, j int) bool {
		return timesAndAmounts[i].TimeString < timesAndAmounts[j].TimeString
	})
	slog.Info("Sorted timesAndAmounts by TimeAndAmount.TimeString", "timesAndAmounts", timesAndAmounts)
	return timesAndAmounts
}

func StringToTitleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func StringToFloat(input *string) float64 {
	if f64, err := strconv.ParseFloat(*input, 64); err == nil {
		slog.Debug("Float parsed from string", "input", *input, "f64", f64)
		return f64
	}

	defaultFloat := 0.0
	slog.Warn("Unable to parse input. Using default", "input", input, "default", defaultFloat)

	return defaultFloat
}
