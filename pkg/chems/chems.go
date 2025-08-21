package chems

import "fmt"

type HalfLifeStruct struct {
	Caffeine float64
	Nicotine float64
}

var (
	Available = map[string]float64{
		"caffeine": HalfLife.Caffeine,
		"nicotine": HalfLife.Nicotine,
	}
	HalfLife = HalfLifeStruct{
		Caffeine: 5.7,
		Nicotine: 1.9,
	}
)

func ListAvailableChems() {
	fmt.Println("Available chem options:")
	for chemName, halfLife := range Available {
		fmt.Printf("  %s (half-life: %.1f hours)\n", chemName, halfLife)
	}
}
