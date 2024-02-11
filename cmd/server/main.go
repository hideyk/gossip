package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/hideyk/gossip/server"
)

func main() {
	host := flag.String("host", "localhost", "Socket Host")
	port := flag.Int("port", 9988, "Socket Port")
	protocol := flag.String("protocol", "tcp", "Socket connection protocol")
	flag.Parse()

	s, err := server.CreateServer(*host, *port, *protocol)

	fmt.Println("Server created on ", *host+":"+strconv.Itoa(*port))

	if err != nil {
		fmt.Println(err)
	}

	s.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Shutting down server...")
	s.Stop()
	fmt.Println("Server stopped")
}
