package main

import (
	"fmt"
	"gophertalk/cmd/client/terminal"
	"gophertalk/internal/dto"
	"net"
	"os"
)

const (
	connHostClient = "localhost"
	connPortClient = "8080"
	connTypeClient = "tcp"
)

var user *dto.UserDto

func main() {
	fmt.Println("Connecting to " + connTypeClient + " server " + connHostClient + ":" + connPortClient)

	conn, err := net.Dial(connTypeClient, connHostClient+":"+connPortClient)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	user = terminal.DoLogin(conn)
	go terminal.StartReadMessages(conn)
	terminal.PrintHelp()
	terminal.StartStdinListener(conn, user)
}
