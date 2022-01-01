package terminal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gophertalk/internal/conv"
	"gophertalk/internal/dto"
	"log"
	"net"
	"os"
)

const clearCurrentLinePattern = "\\33[2K\r"

var toUser = ""

func DoLogin(conn net.Conn) *dto.UserDto {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter with your nickname:")
	bytes, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	clientId := conv.BytesToStringWithoutDelim(bytes)
	userDto := dto.UserDto{ClientId: clientId}

	err = marshalAndWrite(conn, userDto)
	if err != nil {
		log.Fatal(err)
	}
	return &userDto
}

func StartStdinListener(conn net.Conn, from *dto.UserDto) {
	reader := bufio.NewReader(os.Stdin)

	bytes, err := reader.ReadBytes('\n')
	if err == nil {
		msg := conv.BytesToStringWithoutDelim(bytes)
		messageDto := dto.NewMessageDto("", from.ClientId, msg)
		err = marshalAndWrite(conn, messageDto)
		if err != nil {
			fmt.Println("*** Failed to write message: ", err, "\ntry again...")
		}
	} else {
		fmt.Println("*** Failed to read your input. Try again ***")
	}

	printYourTime(toUser)
	StartStdinListener(conn, from)
}

func StartReadMessages(conn net.Conn) {
	bytes, err := bufio.NewReader(conn).ReadBytes('\n')
	var obj = dto.MessageDto{}
	err = json.Unmarshal(conv.BytesWithoutDelim(bytes), &obj)

	if err != nil {
		fmt.Println(clearCurrentLinePattern, "***Error to unmarshal message from server***")
	} else {
		printNewMessage(obj)
	}

	printYourTime(toUser)
	StartReadMessages(conn)
}

func marshalAndWrite(conn net.Conn, value interface{}) error {
	dataToSend, _ := json.Marshal(value)
	_, err := conn.Write(append(dataToSend, '\n'))
	return err
}
