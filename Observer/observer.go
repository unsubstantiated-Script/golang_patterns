package Observer

type Observer interface {
	update(string)
	getID() string
}
