package Iterator

type Iterator interface {
	hasNext() bool
	getNext() *User
}
