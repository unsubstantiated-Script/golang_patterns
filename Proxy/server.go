package Proxy

type server interface {
	handleRequest(string, string) (int, string)
}
