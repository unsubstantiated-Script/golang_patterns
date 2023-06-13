package Flyweight

type TerroristSkin struct {
	color string
}

func (t *TerroristSkin) getColor() string {
	return t.color
}

func newTerroristSkin() *TerroristSkin {
	return &TerroristSkin{color: "red"}
}
