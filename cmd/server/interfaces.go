package server

import (
	"gophertalk/internal/dto"
	"sync"
)

type Manager interface {
	Add(client *Client) bool
	Remove(client *Client, msg string)
	SendMessage(dto dto.MessageDto)
	GetAll() *sync.Map
}
