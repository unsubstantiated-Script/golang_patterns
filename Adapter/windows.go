package Adapter

import "fmt"

type Windows struct{}

func (m *Windows) insertIntoUSBPort() {
	fmt.Println("Lightning connector is plugged into dis PC.")
}
