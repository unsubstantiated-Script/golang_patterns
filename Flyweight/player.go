package Flyweight

type Player struct {
	skin       Skin
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, skinType string) *Player {
	skin, _ := getSkinFactorySingleInstance().getSkinByType(skinType)
	return &Player{
		playerType: playerType,
		skin:       skin,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
