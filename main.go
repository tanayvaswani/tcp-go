package main

import "net"

type Server struct {
	listenAddr string
	ln net.Listener
	quitCh chan struct{}
}

func newServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) start() error {
	ln, err := net.Listen("tcp", s.listenAddr)

	if err != nil {
		return err
	}

	defer ln.Close()

	s.ln = ln
}

func main() {

}