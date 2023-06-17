package ChainOfResponsibility

type Department interface {
	execute(*Patient)
	setNext(Department)
}
