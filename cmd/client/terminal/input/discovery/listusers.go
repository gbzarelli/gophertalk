package discovery

import (
	"gophertalk/internal/dto"
	"strings"
)

type ListUsersDiscoverer struct {
	next MessageDiscoverer
}

func (hd *ListUsersDiscoverer) Execute(inputMessage string) (*InputMessage, error) {
	if strings.HasPrefix(inputMessage, "/users") {
		return &InputMessage{MessageDiscoveryTypeSendMessage,
			dto.NewMessageToListUsersDto()}, nil
	}
	return hd.next.Execute(inputMessage)
}

func (hd *ListUsersDiscoverer) SetNext(inputMessage MessageDiscoverer) {
	hd.next = inputMessage
}
