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

var Nicotine = Chem{
	Name:         "nicotine",
	Halflife:     4.25,
	Description:  "unused",
	CommonValues: nicotineCommonValuesTable,
}
var nicotineCommonValuesTable = `
| Source | Measurement | Total Nicotine Content | Absorbed Nicotine |
|--------|-------------|------------------------|-------------------|
| **TRADITIONAL TOBACCO** | | | |
| Cigarette (regular) | 1 cigarette | 10-14mg | 1-1.5mg |
| Cigarette (light) | 1 cigarette | 10-12mg | 0.8-1.2mg |
| Small cigar/cigarillo | 1 cigar | 5.9-30mg | 1-2mg |
| Medium cigar (corona) | 1 cigar | 50-150mg | 2-4mg |
| Large cigar (churchill) | 1 cigar | 200-335mg | 3-4.5mg |
| Pipe tobacco | 1g | 30-51mg | Variable | Most users don't inhale |
| Pipe tobacco (typical bowl) | 2.5-3g | 75-150mg | Variable |
| **SMOKELESS TOBACCO** | | | |
| Moist snuff | 1g | 4.4-25mg | 3-5mg |
| Moist snuff (typical portion) | 2.5g | 11-63mg | 7.5-15mg |
| Chewing tobacco | 1g | 3.4-39.7mg | 3-4mg |
| Swedish snus (General) | 1g pouch | 8.5mg | 1-2mg |
| American snus | 1g pouch | 4-6mg | 0.5-1mg |
| **HOOKAH/SHISHA** | | | |
| Virginia tobacco shisha | 10g bowl | 5mg | 2-3mg |
| Dark leaf shisha | 10g bowl | 30-40mg | 8-11mg |
| Typical hookah session | 10-20g | 5-80mg | 2-11.4mg |
| **MODERN NICOTINE PRODUCTS** | | | |
| ZYN pouch (3mg) | 1 pouch | 3mg | 1.5mg |
| ZYN pouch (6mg) | 1 pouch | 6mg | 3mg |
| VELO pouch (2mg) | 1 pouch | 2mg | 1mg |
| VELO pouch (4mg) | 1 pouch | 4mg | 2mg |
| VELO pouch (7mg) | 1 pouch | 7mg | 3.5mg |
| Rogue pouch (6mg) | 1 pouch | 6mg | 3mg |
| FRĒ pouch (12mg) | 1 pouch | 12mg | 6mg |
| On! pouch (2-8mg) | 1 pouch | 2-8mg | 1-4mg |
| **E-CIGARETTES/VAPES** | | | |
| E-liquid (freebase) | 1mL | 3-24mg | Variable |
| E-liquid (freebase) | 1 fl oz (30mL) | 90-720mg | Variable |
| E-liquid (salt nicotine) | 1mL | 20-50mg | Variable |
| E-liquid (salt nicotine) | 1 fl oz (30mL) | 600-1500mg | Variable |
| JUUL pod | 0.7mL pod | ~40mg | Variable |
| Elf Bar BC5000 | 13mL device | ~650mg | Variable |
| Typical vaping session | 10 puffs | 2-6mg | 0.5-3mg |
| **HEATED TOBACCO** | | | |
| IQOS HEETS (advertised) | 1 stick | 0.5mg | 1-2mg |
| IQOS HEETS (actual) | 1 stick | 4.3-6mg | 1-2mg |
| glo tobacco stick | 1 stick | 4-5mg | 1-1.5mg |
| **PHARMACEUTICAL NRT** | | | |
| Nicotine patch (7mg) | 24 hours | 7mg | 5.7mg |
| Nicotine patch (14mg) | 24 hours | 14mg | 11.5mg |
| Nicotine patch (21mg) | 24 hours | 21mg | 17-19mg |
| Nicotine gum (2mg) | 1 piece | 2mg | 0.6mg |
| Nicotine gum (4mg) | 1 piece | 4mg | 1.3mg |
| Nicotine lozenge (2mg) | 1 lozenge | 2mg | 1.5mg |
| Nicotine lozenge (4mg) | 1 lozenge | 4mg | 2-3mg |
| Nicotine nasal spray | 1 spray (0.5mg/nostril) | 1mg | 0.8mg |

## Important Notes on Bioavailability

- **Absorption rates vary significantly** based on pH, individual metabolism, and usage technique
- **Cigarette absorption** typically 10-20% of total nicotine content
- **Oral products** absorption heavily influenced by pH (optimal 8.0-9.0)
- **Transdermal patches** provide most consistent and complete absorption (76-82%)
- **Individual variation** can be 4-fold due to genetic factors (CYP2A6 polymorphisms)
- **Peak plasma levels** range from 2-5 minutes (cigarettes) to 6-12 hours (patches)
`
