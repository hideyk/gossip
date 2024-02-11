package client

import (
	"fmt"
	"net"
	"strconv"
)

type Client struct {
	Host     string
	Port     int
	Protocol string
}

func (c *Client) Establish() net.Conn {
	conn, err := net.Dial(c.Protocol, c.Host+":"+strconv.Itoa(c.Port))
	if err != nil {
		panic(err)
	}
	return conn
}

func Write(conn net.Conn, s string) {
	_, err := conn.Write([]byte(s))
	if err != nil {
		fmt.Println("Failed to write message to connection:", err.Error())
	}
	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println("Received: ", string(buffer[:mLen]))
}
