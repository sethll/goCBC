package chems

import "fmt"

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

// Chem defines the structure for individual chems
type Chem struct {
	Name         string
	Halflife     float64
	Description  string
	StandardUnit string
	CommonValues string
}

var (
	// DefaultChem sets the default chem
	DefaultChem = "caffeine"
	registry    = []*Chem{
		&Caffeine,
		&Nicotine,
	}
)

// ListAvailable returns a list of available chems in the registry
func ListAvailable() (retList []string) {
	for _, eachChem := range registry {
		retList = append(retList, eachChem.Name)
	}
	return
}

// GetChem returns the Chem struct with a matching inputName
func GetChem(inputName *string) (*Chem, error) {
	for _, eachChem := range registry {
		if (*eachChem).Name == *inputName {
			return eachChem, nil
		}
	}
	getChemError := fmt.Errorf("no such available chem: %s", *inputName)
	return nil, getChemError
}
