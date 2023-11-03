package model

import (
	"encoding/json"
)

type Message struct {
	Id          string `json:"id"`
	MessageText string `json:"text"`
}

// MarshalBibnary to implement encoding.BinaryMarshaler need for redis
func (m *Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
