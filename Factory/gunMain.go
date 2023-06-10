package Factory

import (
	"fmt"
	"golang_patterns/Factory/guns"
)

func GunMain() {
	ak47, _ := getGun("AK47")
	musket, _ := getGun("musket")
	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g guns.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
