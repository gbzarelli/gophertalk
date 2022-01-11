package main

import (
	"bufio"
	"fmt"
	"gophertalk/cmd/client/terminal"
	discovery2 "gophertalk/cmd/client/terminal/input/discovery"
	"gophertalk/internal/conv"
	"log"
	"net"
	"os"
)

const (
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
	fmt.Println("Type the address server (Ex: 'localhost:8080'): ")
	reader := bufio.NewReader(os.Stdin)
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	server := conv.BytesToStringWithoutDelim(bytes)

	fmt.Println("Connecting to " + connTypeClient + " server " + server)
	conn, err := net.Dial(connTypeClient, server)
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
