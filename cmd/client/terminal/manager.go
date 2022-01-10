package terminal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gophertalk/cmd/client/terminal/input/discovery"
	"gophertalk/internal/conv"
	"gophertalk/internal/dto"
	"log"
	"net"
	"os"
)

const clearCurrentLinePattern = "\\33[2K\r"

type ConnectedUser struct {
	toUser      string
	currentUser *dto.UserDto
	conn        net.Conn
}

func DoLogin(conn net.Conn) *ConnectedUser {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter with your nickname:")
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	clientId := conv.BytesToStringWithoutDelim(bytes)
	userDto := dto.UserDto{ClientId: clientId}

	err = conv.MarshalAndWrite(conn, userDto)
	if err != nil {
		log.Fatal(err)
	}

	return &ConnectedUser{toUser: "", currentUser: &userDto, conn: conn}
}

func (usr *ConnectedUser) StartInputListener(msgDiscovery discovery.MessageDiscoverer) {
	reader := bufio.NewReader(os.Stdin)
	bytes, err := reader.ReadBytes('\n')

	if err == nil {
		inputString := conv.BytesToStringWithoutDelim(bytes)
		usr.executeInputCommand(msgDiscovery, inputString)
	} else {
		fmt.Println("*** Failed to read your input. Try again ***")
	}

	printYourTime(usr.toUser)
	usr.StartInputListener(msgDiscovery)
}

func (usr *ConnectedUser) executeInputCommand(msgDiscovery discovery.MessageDiscoverer, inputString string) {
	msg, err := msgDiscovery.Execute(inputString)

	if err == nil {
		switch msg.Type {
		case discovery.MessageDiscoveryTypeSendMessage:
			messageToSend := msg.Value.(dto.MessageDto)
			messageToSend.To = usr.toUser
			messageToSend.From = usr.currentUser.ClientId
			usr.sendMessage(err, messageToSend)
			break
		case discovery.MessageDiscoveryTypeDefineSender:
			usr.toUser = msg.Value.(string)
			if usr.toUser == "" {
				fmt.Println("**Defined messages for all")
			} else {
				fmt.Println("**Defined messages to: " + usr.toUser)
			}
			break
		case discovery.MessageDiscoveryTypeHelpCommand:
			PrintHelp()
			break
		}

	} else {
		fmt.Println(err, "\ntry again...")
	}
}

func (usr *ConnectedUser) sendMessage(err error, msg dto.MessageDto) {
	err = conv.MarshalAndWrite(usr.conn, msg)
	if err != nil {
		fmt.Println("*** Failed to write message: ", err, "\ntry again...")
	}
}

func (usr *ConnectedUser) StartReadMessages() {
	bytes, err := bufio.NewReader(usr.conn).ReadBytes('\n')
	var obj = dto.MessageDto{}
	err = json.Unmarshal(conv.BytesWithoutDelim(bytes), &obj)

	if err != nil {
		fmt.Println(clearCurrentLinePattern, "***Error to unmarshal message from server***")
	} else {
		printNewMessage(obj)
	}

	printYourTime(usr.toUser)
	usr.StartReadMessages()
}
