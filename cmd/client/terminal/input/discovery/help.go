package discovery

import "strings"

type HelpDiscoverer struct {
	next MessageDiscoverer
}

func (hd *HelpDiscoverer) Execute(inputMessage string) (*InputMessage, error) {
	if strings.HasPrefix(inputMessage, "/help") {
		return &InputMessage{MessageDiscoveryTypeHelpCommand, nil}, nil
	}
	return hd.next.Execute(inputMessage)
}

func (hd *HelpDiscoverer) SetNext(inputMessage MessageDiscoverer) {
	hd.next = inputMessage
}
