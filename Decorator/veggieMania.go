package Decorator

type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 15
}
