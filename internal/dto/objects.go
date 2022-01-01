package dto

type UserDto struct {
	ClientId string `json:"clientId"`
}

type MessageDto struct {
	To   string `json:"to"`
	From string `json:"from"`
	Msg  string `json:"msg"`
}

func NewMessageDto(to string, from string, msg string) MessageDto {
	return MessageDto{To: to, From: from, Msg: msg}
}
