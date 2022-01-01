package server

import (
	"gophertalk/internal/dto"
	"sync"
)

type DefaultManager struct {
	clients *sync.Map
}

func (s *DefaultManager) Remove(client *Client, msg string) {
	defer client.CloseConn()
	s.clients.Delete(client.ClientId)
	if msg != "" {
		_ = client.NewMessageFromServer(msg)
	}
}

func (s *DefaultManager) SendMessage(dto dto.MessageDto) {
	if dto.To == "" {
		s.clients.Range(func(clientId, client interface{}) bool {
			if clientId == dto.From {
				return true
			}
			s.sendMessageTo(dto, client.(*Client))
			return true
		})
	} else {
		client, ok := s.clients.Load(dto.To)
		if ok {
			s.sendMessageTo(dto, client.(*Client))
		} else if dto.To != dto.From {
			newDto := dto
			newDto.To = "<server>"
			newDto.Msg = "Impossible send message to " + dto.To
			s.SendMessage(newDto)
		}
	}
}

func (s *DefaultManager) sendMessageTo(dto dto.MessageDto, client *Client) {
	err := client.NewMessage(dto)
	if err != nil {
		defer s.Remove(client, "")
	}
}

func (s *DefaultManager) GetAll() *sync.Map {
	return s.clients
}

func (s *DefaultManager) Add(client *Client) bool {
	_, ok := s.clients.Load(client.ClientId)
	if ok {
		return false
	}
	s.clients.Store(client.ClientId, client)
	return true
}

func NewManager() Manager {
	return &DefaultManager{clients: new(sync.Map)}
}
