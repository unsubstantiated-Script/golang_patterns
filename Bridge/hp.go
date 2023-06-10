package Bridge

import "fmt"

type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by an Hp Printer")
}
