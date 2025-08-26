package hlcalc

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
	"math"
	"time"

	"github.com/sethll/goCBC/pkg/progutils"
)

//// Substance-in-body calculation ////
//
// Substance-in-body, milligrams = X
// Initial amount = Xo
// time in hours since consumption = T
// metabolic half-life = M
//
// X = Xo * (1/2)^(T/M)
//

// CalcSubstanceInBody calculates the remaining amount of substance in the body
// using exponential decay based on half-life. Uses the formula:
// X = Xo * (1/2)^(T/M)
// where X is remaining amount, Xo is initial amount, T is time elapsed, M is metabolic half-life.
func CalcSubstanceInBody(initialAmount, timeInHours, metabolicHalfLife *float64) (mgSubstanceInBody float64) {
	slog.Debug("CalcSubstanceInBody called",
		"initialAmount", *initialAmount,
		"timeInHours", *timeInHours,
		"metabolicHalfLife", metabolicHalfLife)

	timeOverHalflife := *timeInHours / *metabolicHalfLife
	intermediaryCalc := math.Pow(1.0/2.0, timeOverHalflife)
	mgSubstanceInBody = *initialAmount * intermediaryCalc

	slog.Debug("CalcSubstanceInBody calculation",
		"timeOverHalflife", timeOverHalflife,
		"intermediaryCalc", intermediaryCalc,
		"result", mgSubstanceInBody)

	return mgSubstanceInBody
}

//// Time-to-certain-amount calculation ////
//   (in hours)
// Certain amount = Ca
// Initial amount = Xo
// Metabolic half-life (in hours) = M
//
// T = M * (ln(Ca/Xo)/ln(1/2))
//

// CalcTimeToGivenAmt calculates the time in hours required for a substance to decay
// from an initial amount to a target amount using half-life. Uses the formula:
// T = M * (ln(Ca/Xo)/ln(1/2))
// where T is time, M is metabolic half-life, Ca is target amount, Xo is initial amount.
func CalcTimeToGivenAmt(givenAmt, initialAmt, metabolicHalfLife *float64) (timeToGivenAmt float64) {
	slog.Debug("CalcTimeToGivenAmt called",
		"givenAmt", *givenAmt,
		"initialAmt", *initialAmt,
		"metabolicHalfLife", *metabolicHalfLife)

	lnOfOneHalf := math.Log(1.0 / 2.0)
	givenAmtOverInit := *givenAmt / *initialAmt
	lnOfGivenOverInit := math.Log(givenAmtOverInit)
	dividedLns := lnOfGivenOverInit / lnOfOneHalf

	timeToGivenAmt = *metabolicHalfLife * dividedLns

	slog.Debug("CalcTimeToGivenAmt calculation",
		"lnOfOneHalf", lnOfOneHalf,
		"givenAmtOverInit", givenAmtOverInit,
		"lnOfGivenOverInit", lnOfGivenOverInit,
		"dividedLns", dividedLns,
		"result", timeToGivenAmt)

	return timeToGivenAmt
}

// RunHLCalculations performs half-life calculations for multiple substance intakes
// and returns the current body content and time when target amount will be reached.
func RunHLCalculations(results *progutils.Results, timesAndAmounts *[]progutils.TimeAndAmount, targetAmount, chemHL *float64) {
	slog.Debug("RunHLCalculations started",
		"targetAmount", *targetAmount,
		"chemHL", *chemHL,
		"entryCount", len(*timesAndAmounts))

	currentTime := progutils.GetCurrentTime() // timeOps.go
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

		oldTotal := results.GetChemIngestedTotal(theoreticalMode)
		newTotal := oldTotal + eachItem.Amount
		results.SetChemIngestedTotal(theoreticalMode, newTotal)
		slog.Debug("Updated ingested total",
			"theoreticalMode", theoreticalMode,
			"oldTotal", oldTotal,
			"addedAmount", eachItem.Amount,
			"newTotal", newTotal)

		// logic for if first index item, only set initial values and don't process
		if i == 0 {
			timeMarker = eachItem.TimeObject
			results.SetBodyChemContent(theoreticalMode, eachItem.Amount)
			slog.Debug("First entry processed",
				"timeMarker", timeMarker.Format("2006-01-02 15:04:05"),
				"initialBodyContent", eachItem.Amount,
				"theoreticalMode", theoreticalMode)
			continue
		}

		etHours := progutils.GetElapsedHours(&timeMarker, &eachItem.TimeObject) // timeOps.go
		localBCC := results.GetBodyChemContent(theoreticalMode)
		slog.Debug("Before decay calculation", "bodyChemContent", localBCC)

		localBCC = CalcSubstanceInBody(&localBCC, &etHours, chemHL)
		slog.Debug("After decay calculation", "bodyChemContent", localBCC)

		localBCC += eachItem.Amount
		results.SetBodyChemContent(theoreticalMode, localBCC)
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

		etHours := progutils.GetElapsedHours(&results.LastRealTime, &currentTime) // timeOps.go
		localBCC := (*results).BodyChemContent
		slog.Debug("Before final decay calculation", "bodyChemContent", localBCC)

		localBCC = CalcSubstanceInBody(&localBCC, &etHours, chemHL)
		(*results).BodyChemContent = localBCC
		slog.Debug("After final decay calculation", "bodyChemContent", localBCC)

		tValue := CalcTimeToGivenAmt(targetAmount, &localBCC, chemHL)
		(*results).WearoffTime = progutils.AddTime(&currentTime, &tValue) // timeOps.go
		slog.Debug("Real wearoff time calculated",
			"tValue", tValue,
			"wearoffTime", (*results).WearoffTime.Format("2006-01-02 15:04:05"))
	}

	if (*results).TheoreticalChemIngestedTotal > 0 {
		slog.Debug("Calculating theoretical wearoff time")
		//localBCC := (*results).TheoreticalBodyChemContent
		slog.Debug("Theoretical body chem content", "bodyChemContent", (*results).TheoreticalBodyChemContent)

		tValue := CalcTimeToGivenAmt(targetAmount, &(*results).TheoreticalBodyChemContent, chemHL)
		(*results).TheoreticalWearoffTime = progutils.AddTime(&timeMarker, &tValue) // timeOps.go
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
