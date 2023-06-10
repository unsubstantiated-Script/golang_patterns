package Bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (m *Windows) Print() {
	fmt.Println("Print request for Windows")
	m.printer.PrintFile()
}

func (m *Windows) SetPrinter(p Printer) {
	m.printer = p
}
