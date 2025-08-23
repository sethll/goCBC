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

import "fmt"

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
		Caffeine: 5.0,
		Nicotine: 4.2,
	}
)

// ListAvailableChems prints a formatted list of all available substances and their half-lives.
func ListAvailableChems() {
	fmt.Println("Available chem options:")
	for chemName, halfLife := range Available {
		fmt.Printf("  %s (half-life: %.1f hours)\n", chemName, halfLife)
	}
}
