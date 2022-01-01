package main

import (
	"fmt"
	"gophertalk/cmd/server"
	"log"
	"net"
	"sync"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	manager := server.NewManager()
	abortClientChan := make(chan *server.Client)

	listener := openServerListener()
	defer shutdownServer(listener, manager)

	fmt.Println("The GopherTalk Server was started and listener new clients at " + connType + "://" + connHost + ":" + connPort)

	go abortClientListener(manager, abortClientChan)
	acceptNewClient(listener, manager, abortClientChan)
}

func acceptNewClient(listener net.Listener, manager server.Manager, abortChan chan *server.Client) {
	fmt.Println("Waiting new client...")
	c, err := listener.Accept()
	if err == nil {
		fmt.Println("New Client " + c.RemoteAddr().String() + " connected.")
		go handlerConnection(c, manager, abortChan)
	} else {
		fmt.Println("Error to open new client connection:", err.Error())
	}
	acceptNewClient(listener, manager, abortChan)
}

func abortClientListener(manager server.Manager, channel chan *server.Client) {
	c := <-channel
	manager.Remove(c, "")
	fmt.Println("Abort client called by: " + c.ClientId)
	abortClientListener(manager, channel)
}

func shutdownServer(listener net.Listener, manager server.Manager) {
	var wg = new(sync.WaitGroup)
	manager.GetAll().Range(func(key, value interface{}) bool {
		wg.Add(1)
		go func() {
			defer wg.Done()
			manager.Remove(value.(*server.Client), "The server was shutdown")
		}()
		return true
	})
	wg.Wait()

	err := listener.Close()
	if err != nil {
		log.Fatal("Force shutdown server! Close failure:", err.Error())
	}
}

func openServerListener() net.Listener {
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatal("Error to open server:", err.Error())
	}
	return listener
}

func handlerConnection(conn net.Conn, manager server.Manager, channel chan *server.Client) {
	client, err := server.NewClient(conn)
	if err != nil {
		defer forceConnClose(conn)
		fmt.Println("Impossible to connect new client "+conn.RemoteAddr().String(), err.Error())
		return
	}

	if manager.Add(client) {
		err = client.NewMessageFromServer("Connection Successfully")
		if err == nil {
			go client.StartReadMessages(channel, manager)
		} else {
			manager.Remove(client, "")
			fmt.Println("New Client " + client.ClientId + " removed because dont receive the confirmation message")
		}
	} else {
		defer client.CloseConn()
		err = client.NewMessageFromServer("This user was exists! try with other user")
		fmt.Println("New Client " + client.ClientId + " was exists")
	}
}

func forceConnClose(conn net.Conn) {
	_ = conn.Close()
}
