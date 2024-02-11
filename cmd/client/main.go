package main

import (
	"flag"

	"github.com/hideyk/gossip/client"
)

func main() {
	host := flag.String("host", "localhost", "Socket Host")
	port := flag.Int("port", 9988, "Socket Port")
	protocol := flag.String("protocol", "tcp", "Socket connection protocol")
	flag.Parse()

	c := client.Client{
		Host:     *host,
		Port:     *port,
		Protocol: *protocol,
	}

	conn := c.Establish()

	client.Write(conn, "My message")
}
