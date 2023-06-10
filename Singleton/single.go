package Singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating a single instance now!")
			singleInstance = &single{}
		} else {
			fmt.Println("Yo! The single instance has already been made!?")
		}
	} else {
		fmt.Println("Yo! The single instance has already been made!?")
	}
	return singleInstance
}
