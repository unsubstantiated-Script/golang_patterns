package guns

type Ak47 struct {
	Gun
}

func (a Ak47) GetName() string {
	return a.name
}

func (a Ak47) GetPower() int {
	return a.power
}

func NewAK47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
