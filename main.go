package main

import "net"

type Server struct {
	listenAddr string
	ln net.Listener
}

func newServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func main() {

}