package Mediator

import "fmt"

func MediatorMain() {
	fmt.Println("Hello mediator")

	stationManager := newStationManager()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
