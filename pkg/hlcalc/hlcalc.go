package hlcalc

import (
	"log/slog"
	"math"
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

func CalcSubstanceInBody(initialAmount, timeInHours, metabolicHalfLife float64) (mgSubstanceInBody float64) {
	slog.Debug("CalcSubstanceInBody called",
		"initialAmount", initialAmount,
		"timeInHours", timeInHours,
		"metabolicHalfLife", metabolicHalfLife)

	timeOverHalflife := timeInHours / metabolicHalfLife
	intermediaryCalc := math.Pow(1.0/2.0, timeOverHalflife)
	mgSubstanceInBody = initialAmount * intermediaryCalc

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

func CalcTimeToGivenAmt(givenAmt, initialAmt, metabolicHalfLife *float64) (timeToGivenAmt float64) {
	slog.Debug("CalcTimeToGivenAmt called",
		"givenAmt", givenAmt,
		"initialAmt", initialAmt,
		"metabolicHalfLife", metabolicHalfLife)

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
