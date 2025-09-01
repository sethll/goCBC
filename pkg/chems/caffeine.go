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

var Caffeine = Chem{
	Name:         "caffeine",
	Halflife:     5.00,
	Description:  "unused",
	StandardUnit: "mg",
	CommonValues: caffeineCommonValuesTable,
}
var caffeineCommonValuesTable = `
| Source | Serving Size | Caffeine (mg) |
|--------|--------------|---------------|
| **COFFEE** | | |
| Espresso (single shot) | 1 oz (30 mL) | 63 |
| Coffee, drip brewed | 8 oz (237 mL) | 95 |
| Coffee, cold brew | 16 oz (473 mL) | 200 |
| Coffee, instant | 8 oz (237 mL) | 76 |
| Coffee, decaffeinated | 8 oz (237 mL) | 5 |
| Coffee beans, Arabica | 10 g | 120 |
| Coffee beans, Robusta | 10 g | 220 |
| **TEA** | | |
| Black tea | 8 oz (237 mL) | 42 |
| Green tea | 8 oz (237 mL) | 35 |
| White tea | 8 oz (237 mL) | 28 |
| Oolong tea | 8 oz (237 mL) | 37 |
| Matcha powder | 2 g | 68 |
| Yerba mate | 8 oz (237 mL) | 78 |
| Chai tea | 8 oz (237 mL) | 47 |
| **ENERGY DRINKS** | | |
| Red Bull | 8.4 oz (250 mL) | 80 |
| Monster Energy | 16 oz (473 mL) | 160 |
| Rockstar | 16 oz (473 mL) | 160 |
| Bang Energy | 16 oz (473 mL) | 300 |
| 5-hour Energy | 2 oz (59 mL) | 215 |
| **SOFT DRINKS** | | |
| Coca-Cola Classic | 12 oz (355 mL) | 34 |
| Pepsi | 12 oz (355 mL) | 38 |
| Mountain Dew | 12 oz (355 mL) | 54 |
| Dr Pepper | 12 oz (355 mL) | 42 |
| Diet Coke | 12 oz (355 mL) | 46 |
| Diet Dr Pepper | 12 oz (355 mL) | 44 |
| Pepsi Zero Sugar | 12 oz (355 mL) | 38 |
| **CHOCOLATE & CONFECTIONERY** | | |
| Dark chocolate (70-85% cacao) | 1 oz (28 g) | 23 |
| Milk chocolate | 1 oz (28 g) | 6 |
| Cocoa powder | 1 tablespoon | 12 |
| Chocolate milk | 8 oz (237 mL) | 2 |
| Coffee ice cream | 1/2 cup | 26 |
| **SUPPLEMENTS** | | |
| Caffeine pills | 1 tablet | 200 |
| Pre-workout supplement | 1 serving | 275 |
| Guarana supplement | 500 mg | 115 |

## Notes

- All values represent average caffeine content based on standard preparations
- Actual caffeine content may vary based on preparation method, brand, and specific product formulations
- Liquid measurements provided in both metric (mL) and US customary (oz) units
- Solid measurements provided in grams or US measurements as appropriate
- Coffee ice cream value represents midpoint of typical range (3-50mg)
- Pre-workout supplement value represents midpoint of typical range (150-400mg)
- Guarana supplement value represents midpoint of typical range (100-130mg)
`
