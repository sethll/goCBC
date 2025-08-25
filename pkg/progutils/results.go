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
	"time"
)

type Results struct {
	ChemIngestedTotal            float64
	BodyChemContent              float64
	TheoreticalChemIngestedTotal float64
	TheoreticalBodyChemContent   float64
	LastRealTime                 time.Time
	WearoffTime                  time.Time
	TheoreticalWearoffTime       time.Time
}

// NewResults initializes a new Results
func NewResults() Results {
	return Results{
		ChemIngestedTotal:            0,
		BodyChemContent:              0,
		TheoreticalChemIngestedTotal: 0,
		TheoreticalBodyChemContent:   0,
		LastRealTime:                 time.Time{},
		WearoffTime:                  time.Time{},
		TheoreticalWearoffTime:       time.Time{},
	}
}

func (r *Results) String() string {
	return fmt.Sprintf("Results{ChemIngestedTotal: %.2f, BodyChemContent: %.2f, TheoreticalChemIngestedTotal: %.2f, TheoreticalBodyChemContent: %.2f, LastRealTime: %s, WearoffTime: %s, TheoreticalWearoffTime: %s}",
		r.ChemIngestedTotal,
		r.BodyChemContent,
		r.TheoreticalChemIngestedTotal,
		r.TheoreticalBodyChemContent,
		r.LastRealTime.Format("2006-01-02 15:04:05"),
		r.WearoffTime.Format("2006-01-02 15:04:05"),
		r.TheoreticalWearoffTime.Format("2006-01-02 15:04:05"))
}

func (r *Results) getChemIngestedTotal(theoreticalMode bool) float64 {
	if theoreticalMode {
		slog.Debug("Returning theoretical", "TCIT", r.TheoreticalChemIngestedTotal)
		return r.TheoreticalChemIngestedTotal
	} else {
		slog.Debug("Returning real", "CIT", r.ChemIngestedTotal)
		return r.ChemIngestedTotal
	}
}

func (r *Results) setChemIngestedTotal(theoreticalMode bool, input float64) {
	if theoreticalMode {
		r.TheoreticalChemIngestedTotal = input
		slog.Debug("Setting theoretical", "input", input, "TCIT", r.TheoreticalChemIngestedTotal)
	} else {
		r.ChemIngestedTotal = input
		slog.Debug("Setting real", "input", input, "CIT", r.ChemIngestedTotal)
	}
}

func (r *Results) getBodyChemContent(theoreticalMode bool) float64 {
	if theoreticalMode {
		slog.Debug("Returning theoretical", "TBCC", r.TheoreticalBodyChemContent)
		return r.TheoreticalBodyChemContent
	} else {
		slog.Debug("Returning theoretical", "BCC", r.BodyChemContent)
		return r.BodyChemContent
	}
}

func (r *Results) setBodyChemContent(theoreticalMode bool, input float64) {
	if theoreticalMode {
		r.TheoreticalBodyChemContent = input
		slog.Debug("Setting theoretical", "input", input, "TBCC", r.TheoreticalBodyChemContent)
	} else {
		r.BodyChemContent = input
		slog.Debug("Setting real", "input", input, "BCC", r.BodyChemContent)
	}
}
