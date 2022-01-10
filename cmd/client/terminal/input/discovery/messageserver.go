package discovery

import "gophertalk/internal/dto"

type MessageServerDiscoverer struct {
	next MessageDiscoverer
}

func (hd *MessageServerDiscoverer) Execute(inputMessage string) (*InputMessage, error) {
	return &InputMessage{MessageDiscoveryTypeSendMessage,
		dto.NewMessageDto("", "", inputMessage)}, nil
}

func (hd *MessageServerDiscoverer) SetNext(inputMessage MessageDiscoverer) {
	hd.next = inputMessage
}
