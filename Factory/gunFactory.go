package Factory

import (
	"fmt"
	"golang_patterns/Factory/guns"
)

func getGun(gunType string) (guns.IGun, error) {
	if gunType == "AK47" {
		return guns.NewAK47(), nil

	}
	if gunType == "musket" {
		return guns.NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
