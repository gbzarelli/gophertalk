package dto

const (
	MessageDtoTypeSendMessage = "MESSAGE"
	MessageDtoTypeListUsers   = "LIST_USERS"
)

type UserDto struct {
	ClientId string `json:"clientId"`
}

type MessageDto struct {
	Type string `json:"type"`
	To   string `json:"to"`
	From string `json:"from"`
	Msg  string `json:"msg"`
}

func NewMessageDto(to string, from string, msg string) MessageDto {
	return MessageDto{To: to, From: from, Msg: msg, Type: MessageDtoTypeSendMessage}
}

func NewMessageToListUsersDto() MessageDto {
	return MessageDto{Type: MessageDtoTypeListUsers}
}
