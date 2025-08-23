package chems

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

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/sethll/goCBC/pkg/progutils"
)

// HalfLifeStruct represents the half-life values for different substances in hours.
type HalfLifeStruct struct {
	Caffeine float64
	Nicotine float64
}

var (
	// Available contains the mapping of substance names to their half-lives in hours.
	Available = map[string]float64{
		"caffeine": HalfLife.Caffeine,
		"nicotine": HalfLife.Nicotine,
	}
	// HalfLife contains the half-life values for all supported substances.
	HalfLife = HalfLifeStruct{
		Caffeine: 5.00,
		Nicotine: 4.25,
	}
	DefaultChem = "caffeine"
)

// ListAvailableChems prints a formatted list of all available substances and their half-lives.
func ListAvailableChems() {
	fmt.Println("Available chem options:")
	fmt.Println(genChemOutputTable())
}

func genChemOutputTable() *table.Table {
	header := []string{
		"Chem", "Half-life",
	}
	var rows [][]string
	for chemName, halfLife := range Available {
		if chemName == DefaultChem {
			chemName = fmt.Sprintf("%s (default)", chemName)
		}
		rows = append(rows, []string{chemName, fmt.Sprintf("%.2f hours", halfLife)})
	}

	//chemTable := table.New().Border(lipgloss.HiddenBorder()).BorderHeader(true).Rows(rows...).Headers(header...)
	chemTable := table.New().
		Headers(header...).Rows(rows...).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return progutils.Styles.TableHeader
			default:
				return progutils.Styles.TableEvenRow
				//saving these lines for when there are more chems
				//case row%2 == 0:
				//	return progutils.Styles.TableEvenRow
				//default:
				//	return progutils.Styles.TableOddRow
			}
		}).
		BorderHeader(true)

	return chemTable
}
