package Visitor

type Shape interface {
	getType() string
	accept(Visitor)
}
