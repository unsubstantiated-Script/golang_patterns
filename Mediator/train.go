package Mediator

type Train interface {
	arrive()
	depart()
	permitArrival()
}
