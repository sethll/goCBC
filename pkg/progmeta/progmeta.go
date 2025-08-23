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
	"runtime/debug"
	"strings"
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
	build = setBuild()
	// ProgName is the name of the program.
	ProgName = "goCBC"
	// Author is the program author's name.
	Author = "Seth L"
	// CopyrightYear is the copyright year for the program.
	CopyrightYear = "2025"
	// ProgVersion contains the current version information for the program.
	ProgVersion = Version{
		Major: "0",
		Minor: "1",
		Patch: "3",
		Build: build,
	}
	// ShortDesc is a brief description of the program.
	ShortDesc = "A Go CLI tool for calculating substance metabolism and optimal sleep timing"
	// LongDesc is a detailed description of the program's functionality.
	LongDesc = `goCBC calculates when substances like caffeine and nicotine drop to target
levels for restful sleep using pharmacokinetic half-life modeling. Supports
multiple daily intakes with precise exponential decay calculations.`
	// Usage shows the command-line usage pattern.
	Usage = "goCBC [flags] <target> '<time:amount>' ['<time:amount>' ...]"
	// UsageExample provides an example of how to use the program.
	UsageExample = "goCBC 50 '1100:300' '1500:150'"
)

// readFromBuildInfo attempts to get version information from Go's build info
func readFromBuildInfo() (string, string, bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", "", false
	}

	version := info.Main.Version

	revision := ""

	// Look for VCS revision in build settings
	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			if len(setting.Value) >= 7 {
				revision = setting.Value[:7] // Short SHA
			} else {
				revision = setting.Value
			}
			return version, revision, true
		}
	}

	return version, revision, false
}

// setBuild returns the current Version information with fallback logic
func setBuild() string {
	// Try to get version from Go's build info first (works with go install)
	version, revision, ok := readFromBuildInfo()
	if !ok {
		if strings.HasPrefix(version, "v0.0.0-") {
			revision = "dev"
		} else {
			revision = "release"
		}
	}
	return revision
}

// String returns a formatted string representation of the version information.
func (v Version) String() string {
	return fmt.Sprintf("Version: %s.%s.%s Build: %s", v.Major, v.Minor, v.Patch, v.Build)
}

// RuntimeVersion returns the Go runtime version information.
func RuntimeVersion() string {
	return fmt.Sprintf("Runtime: %s", runtime.Version())
}

// AllVersionBuildRuntimeInfo returns a combined string with version and runtime information.
func AllVersionBuildRuntimeInfo() string {
	return fmt.Sprintf("%s %s", ProgVersion.String(), RuntimeVersion())
}
