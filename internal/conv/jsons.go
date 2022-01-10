package conv

import (
	"encoding/json"
	"net"
)

func MarshalAndWrite(conn net.Conn, value interface{}) error {
	dataToSend, _ := json.Marshal(value)
	_, err := conn.Write(append(dataToSend, '\n'))
	return err
}
