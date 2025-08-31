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

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/sethll/goCBC/pkg/chems"
	"github.com/sethll/goCBC/pkg/progmeta"
)

// GenerateOutputTableV1 creates a formatted table displaying current substance levels
// and anticipated wearoff time based on the target amount.
func GenerateOutputTableV1(r *Results, wearoffTarget *string, chem *string) *table.Table {
	slog.Debug("Generating output table", "wearoffTarget", *wearoffTarget, "chem", *chem, "results", (*r).String())
	rows := [][]string{
		{
			fmt.Sprintf(
				"%s remaining in system:",
				Styles.Chem.Render(StringToTitleCase(*chem)),
			),
			Styles.Chem.Render(
				fmt.Sprintf(
					"~%.0fmg",
					math.Round((*r).BodyChemContent),
				),
			),
		},
		{
			fmt.Sprintf(
				"Reach target (%s) for %s at:",
				Styles.Chem.Render(
					fmt.Sprintf("%smg", *wearoffTarget),
				),
				Styles.Wearoff.Render("wear-off"),
			),
			Styles.Wearoff.Render(
				(*r).WearoffTime.Format("2006-01-02 15:04"),
			),
		},
	}
	if (*r).TheoreticalChemIngestedTotal > 0 {
		rows = append(rows,
			[]string{
				fmt.Sprintf(
					"Future %s intake total:",
					Styles.Chem.Render(*chem),
				),
				Styles.Chem.Render(
					fmt.Sprintf(
						"%.0fmg",
						math.Round((*r).TheoreticalChemIngestedTotal),
					),
				),
			},
		)
		rows = append(rows,
			[]string{
				fmt.Sprintf(
					"With future intake reach %s target (%s) at:",
					Styles.Wearoff.Render("wear-off"),
					Styles.Chem.Render(
						fmt.Sprintf("%smg", *wearoffTarget),
					),
				),
				Styles.Wearoff.Render(
					(*r).TheoreticalWearoffTime.Format("2006-01-02 15:04"),
				),
			},
		)
	}
	generatedTable := table.New().Border(lipgloss.HiddenBorder()).Rows(rows...)
	slog.Debug("Output table generated successfully", "rowCount", len(rows))
	return generatedTable
}

// ListAvailableChems prints a formatted list of all available substances and their half-lives.
func ListAvailableChems() {
	fmt.Println("Available chem options:")
	fmt.Println(genChemOutputTable())
}

// PrintProgHeader displays the program header with name, version, and copyright information.
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

func genChemOutputTable() *table.Table {
	header := []string{
		"Chem", "Half-life",
	}
	
	var rows [][]string
	for _, chemName := range chems.ListAvailable() {
		chemPointer, err := chems.GetChem(&chemName)
		if err != nil {
			slog.Error(err.Error())
			panic("Something is very wrong, this should never throw an error")
		}

		if chemName == chems.DefaultChem {
			chemName = fmt.Sprintf("%s (default)", chemName)
		}
		rows = append(rows, []string{chemName, fmt.Sprintf("%.2f hours", (*chemPointer).Halflife)})
	}

	//chemTable := table.New().Border(lipgloss.HiddenBorder()).BorderHeader(true).Rows(rows...).Headers(header...)
	chemTable := table.New().
		Headers(header...).Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return Styles.TableHeader
			default:
				return Styles.TableEvenRow
				//saving these lines for when there are more chems
				//case row%2 == 0:
				//	return progutils.Styles.TableEvenRow
				//default:
				//	return progutils.Styles.TableOddRow
			}
		}).
		BorderHeader(false).
		BorderColumn(false).
		BorderRow(false).
		BorderLeft(false).
		BorderRight(false).
		BorderTop(false).
		BorderBottom(false)

	return chemTable
}
