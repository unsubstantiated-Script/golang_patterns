package Mediator

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}
