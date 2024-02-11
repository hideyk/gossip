package server

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	Host        string
	Port        int
	Protocol    string
	Listener    net.Listener
	wg          sync.WaitGroup
	shutdown    chan struct{}
	connections chan net.Conn
}

func CreateServer(host string, port int, protocol string) (*Server, error) {
	listener, err := net.Listen(protocol, host+":"+strconv.Itoa(port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen on address %s: %w", strconv.Itoa(port), err)
	}
	return &Server{
		Host:        host,
		Port:        port,
		Protocol:    protocol,
		Listener:    listener,
		shutdown:    make(chan struct{}),
		connections: make(chan net.Conn),
	}, nil
}

func (s *Server) acceptConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		default:
			conn, err := s.Listener.Accept()
			if err != nil {
				continue
			}
			s.connections <- conn
		}
	}
}

func (s *Server) handleConnections() {
	defer s.wg.Done()

	for {
		select {
		case <-s.shutdown:
			return
		case conn := <-s.connections:
			go processClient(conn)
		}
	}
}

func (s *Server) Start() {
	s.wg.Add(2)
	go s.acceptConnections()
	go s.handleConnections()
}

func (s *Server) Stop() {
	close(s.shutdown)
	s.Listener.Close()

	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Timed out waiting for connections to finish")
		return
	}
}

func processClient(conn net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received:", string(buffer[:mLen]), "from", conn.RemoteAddr())
	_, err = conn.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))

	if err != nil {
		fmt.Println("Error writing back to client:" + err.Error())
	}
	conn.Close()
}
