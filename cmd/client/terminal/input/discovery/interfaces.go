package discovery

const (
	MessageDiscoveryTypeSendMessage  = 0
	MessageDiscoveryTypeDefineSender = 1
	MessageDiscoveryTypeHelpCommand  = 2
)

type InputMessage struct {
	Type  int8
	Value interface{}
}

type MessageDiscoverer interface {
	Execute(string) (*InputMessage, error)
	SetNext(MessageDiscoverer)
}
