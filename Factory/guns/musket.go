package guns

type musket struct {
	Gun
}

func (m musket) GetName() string {
	return m.name
}

func (m musket) GetPower() int {
	return m.power
}

func NewMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
