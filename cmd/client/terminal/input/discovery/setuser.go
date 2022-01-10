package discovery

import (
	"errors"
	"strings"
)

type SetUserDiscoverer struct {
	next MessageDiscoverer
}

func (hd *SetUserDiscoverer) Execute(inputMessage string) (*InputMessage, error) {
	if strings.HasPrefix(inputMessage, "/to") {
		after := strings.SplitAfter(inputMessage, " ")
		if len(after) > 1 {
			return &InputMessage{MessageDiscoveryTypeDefineSender, after[1]}, nil
		} else {
			return nil, errors.New("the set command need one argument, like: '/set username'")
		}

	} else if strings.HasPrefix(inputMessage, "/all") {
		return &InputMessage{MessageDiscoveryTypeDefineSender, ""}, nil
	}
	return hd.next.Execute(inputMessage)
}

func (hd *SetUserDiscoverer) SetNext(inputMessage MessageDiscoverer) {
	hd.next = inputMessage
}
