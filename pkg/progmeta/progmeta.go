package progmeta

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
	"runtime"
)

// Version represents the current version of goCBC.
type Version struct {
	Major string
	Minor string
	Patch string
	Build string
}

// These variables can be set at build time via -ldflags
var (
	build         = "dev"
	ProgName      = "goCBC"
	Author        = "Seth L"
	CopyrightYear = "2025"
	ProgVersion   = Version{
		Major: "0",
		Minor: "1",
		Patch: "0",
		Build: build,
	}
	ShortDesc = "A Go CLI tool for calculating substance metabolism and optimal sleep timing"
	LongDesc  = `goCBC calculates when substances like caffeine and nicotine drop to target
levels for restful sleep using pharmacokinetic half-life modeling. Supports
multiple daily intakes with precise exponential decay calculations.`
	Usage        = "goCBC [flags] <target> <time_amount...>"
	UsageExample = "goCBC 75 '1100:300' '1500:5'"
)

// Get returns the current Version information
func Get() Version {
	return ProgVersion
}

func (v Version) String() string {
	ver := fmt.Sprintf("Version: %s.%s.%s", v.Major, v.Minor, v.Patch)
	if v.Build != "" && v.Build != "dev" {
		return fmt.Sprintf("%s Build: %s", ver, v.Build)
	}
	return ver
}

func RuntimeVersion() string {
	return fmt.Sprintf("Runtime: %s", runtime.Version())
}

func AllVersionBuildRuntimeInfo() string {
	return fmt.Sprintf("%s %s", ProgVersion.String(), RuntimeVersion())
}
