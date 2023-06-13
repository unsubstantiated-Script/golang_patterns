package Flyweight

import "fmt"

func MainFlyweight() {
	fmt.Println("Hello Flyweight!")

	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristSkinType)
	game.addTerrorist(TerroristSkinType)
	game.addTerrorist(TerroristSkinType)
	game.addTerrorist(TerroristSkinType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerroristSkinType)
	game.addCounterTerrorist(CounterTerroristSkinType)
	game.addCounterTerrorist(CounterTerroristSkinType)

	skinFactoryInstance := getSkinFactorySingleInstance()

	for skinType, skin := range skinFactoryInstance.skinMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", skinType, skin.getColor())
	}
}
