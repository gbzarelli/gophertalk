package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gophertalk/internal/conv"
	"gophertalk/internal/dto"
	"net"
)

type Client struct {
	conn     net.Conn
	ClientId string
}

func NewClient(conn net.Conn) (*Client, error) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("Error in first read")
		return nil, err
	}

	obj := dto.UserDto{}
	err = json.Unmarshal(buffer[:len(buffer)-1], &obj)

	if err != nil {
		return nil, err
	}

	return &Client{conn, obj.ClientId}, nil
}

func (c *Client) NewMessageFromServer(msg string) error {
	return c.NewMessage(dto.NewMessageDto("", "<server>", msg))
}

func (c *Client) NewMessage(msg dto.MessageDto) error {
	err := conv.MarshalAndWrite(c.conn, msg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) StartReadMessages(abortChannel chan *Client, manager Manager) {
	buffer, err := bufio.NewReader(c.conn).ReadBytes('\n')

	if err != nil {
		abortChannel <- c
		return
	}

	messageDto := dto.MessageDto{}
	err = json.Unmarshal(buffer[:len(buffer)-1], &messageDto)

	if err == nil {
		switch messageDto.Type {
		case dto.MessageDtoTypeSendMessage:
			go manager.SendMessage(messageDto)
			break
		case dto.MessageDtoTypeListUsers:
			go manager.SendListOfUsers(c)
		}

	} else {
		_ = c.NewMessageFromServer("Payload invalid")
	}

	c.StartReadMessages(abortChannel, manager)
}

func (c *Client) CloseConn() {
	_ = c.conn.Close()
}
