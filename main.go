package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitCh     chan struct{}
	msgCh      chan Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitCh:     make(chan struct{}),
		msgCh: make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitCh
	close(s.msgCh)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accepting ERR:", err)
			continue
		}

		fmt.Println("New Connection to Server:", conn.RemoteAddr())

		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read ERR:", err)
			continue
		}

		s.msgCh <- Message{
			from: conn.RemoteAddr().String(),
			payload: buf[:n],
		}

		conn.Write([]byte("Received a message, Thanks!\n"))
	}
}

func main() {
	server := NewServer(":3000")

	go func ()  {
		for msg := range server.msgCh{
			fmt.Printf("incoming message from (%s): %s\n", msg.from, string(msg.payload))
		}
	}()

	log.Fatal(server.Start())
}
