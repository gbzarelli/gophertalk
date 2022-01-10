package main

import (
	"fmt"
	"gophertalk/cmd/client/terminal"
	discovery2 "gophertalk/cmd/client/terminal/input/discovery"
	"net"
	"os"
)

const (
	connHostClient = "localhost"
	connPortClient = "8080"
	connTypeClient = "tcp"
)

var user *terminal.ConnectedUser

func main() {
	msgDiscoveryChain := buildDiscovery()
	conn := connectToServer()

	user = terminal.DoLogin(conn)

	go user.StartReadMessages()

	terminal.PrintHelp()

	user.StartInputListener(msgDiscoveryChain)
}

func connectToServer() net.Conn {
	fmt.Println("Connecting to " + connTypeClient + " server " + connHostClient + ":" + connPortClient)
	conn, err := net.Dial(connTypeClient, connHostClient+":"+connPortClient)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	return conn
}

func buildDiscovery() discovery2.MessageDiscoverer {
	help := &discovery2.HelpDiscoverer{}
	listUsers := &discovery2.ListUsersDiscoverer{}
	setUser := &discovery2.SetUserDiscoverer{}
	messageServer := &discovery2.MessageServerDiscoverer{}

	help.SetNext(listUsers)
	listUsers.SetNext(setUser)
	setUser.SetNext(messageServer)

	return help
}
