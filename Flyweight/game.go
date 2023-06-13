package Flyweight

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) addTerrorist(skinType string) {
	player := newPlayer("T", skinType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) addCounterTerrorist(skinType string) {
	player := newPlayer("CT", skinType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}
