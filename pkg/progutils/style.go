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

import "github.com/charmbracelet/lipgloss"

// StylesType defines the styling configuration for different UI elements.
type StylesType struct {
	Bedtime      lipgloss.Style
	Caffeine     lipgloss.Style
	Header       lipgloss.Style
	TableHeader  lipgloss.Style
	TableEvenRow lipgloss.Style
	TableOddRow  lipgloss.Style
}

var (
	// Styles contains the styling configuration for different UI elements.
	Styles = StylesType{
		Bedtime: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#87CEEB")),
		Caffeine: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500")),
		Header: lipgloss.NewStyle().
			Bold(true).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()),
		TableHeader: lipgloss.NewStyle().
			Bold(true).
			Padding(0, 2),
		TableEvenRow: lipgloss.NewStyle().
			Padding(0, 2),
		TableOddRow: lipgloss.NewStyle().
			Padding(0, 2).
			Background(lipgloss.Color("#888888")),
	}
)
