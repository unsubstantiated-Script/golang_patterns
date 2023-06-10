package Adapter

import "fmt"

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into dis Mac.")
}
