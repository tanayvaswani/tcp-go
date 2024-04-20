package main

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln net.Listener
	quitCh chan struct{}
}

func newServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitCh: make(chan struct{}),
	}
}

func (s *Server) start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	<- s.quitCh

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accepting ERR:", err)
			continue
		}
	}
}

func main() {

}