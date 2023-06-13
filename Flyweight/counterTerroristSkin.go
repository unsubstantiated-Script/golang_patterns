package Flyweight

type CounterTerroristSkin struct {
	color string
}

func (c *CounterTerroristSkin) getColor() string {
	return c.color
}

func newCounterTerroristSkin() *CounterTerroristSkin {
	return &CounterTerroristSkin{color: "black"}
}
