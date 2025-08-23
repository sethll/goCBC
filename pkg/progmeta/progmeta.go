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
	Usage        = "goCBC [options] <target> '<time:amount>' ['<time:amount>' ...]"
	UsageExample = "goCBC 50 '1100:300' '1500:150'"
)

// getVersionFromBuildInfo attempts to get version information from Go's build info
func getVersionFromBuildInfo() (string, string, bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", "", false
	}

	version := info.Main.Version
	var revision string

	// Look for VCS revision in build settings
	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			if len(setting.Value) >= 7 {
				revision = setting.Value[:7] // Short SHA
			} else {
				revision = setting.Value
			}
			break
		}
	}

	// If we have a proper version tag (not "(devel)" or "v0.0.0-..."), use it
	if version != "(devel)" && !strings.HasPrefix(version, "v0.0.0-") {
		return version, revision, true
	}

	return "", revision, false
}

// parseVersion parses a semantic version string like "v1.2.3" into components
func parseVersion(version string) (major, minor, patch string) {
	// Remove 'v' prefix if present
	version = strings.TrimPrefix(version, "v")

	parts := strings.Split(version, ".")
	if len(parts) >= 1 {
		major = parts[0]
	}
	if len(parts) >= 2 {
		minor = parts[1]
	}
	if len(parts) >= 3 {
		// Handle pre-release suffixes like "3-beta.1"
		patchParts := strings.Split(parts[2], "-")
		patch = patchParts[0]
	}

	// Default to "0" if any component is empty
	if major == "" {
		major = "0"
	}
	if minor == "" {
		minor = "0"
	}
	if patch == "" {
		patch = "0"
	}

	return major, minor, patch
}

// Get returns the current Version information with fallback logic
func Get() Version {
	// Try to get version from Go's build info first (works with go install)
	if version, revision, ok := getVersionFromBuildInfo(); ok {
		major, minor, patch := parseVersion(version)
		buildInfo := revision
		if buildInfo == "" {
			buildInfo = "release"
		}
		return Version{
			Major: major,
			Minor: minor,
			Patch: patch,
			Build: buildInfo,
		}
	}

	// Check if we have VCS info even without a proper version tag
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				revision := setting.Value
				if len(revision) >= 7 {
					revision = revision[:7] // Short SHA
				}
				return Version{
					Major: ProgVersion.Major,
					Minor: ProgVersion.Minor,
					Patch: ProgVersion.Patch,
					Build: revision,
				}
			}
		}
	}

	// Fall back to build-time injected version (works with make build)
	if build != "dev" && build != "" {
		return Version{
			Major: ProgVersion.Major,
			Minor: ProgVersion.Minor,
			Patch: ProgVersion.Patch,
			Build: build,
		}
	}

	// Default fallback
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
