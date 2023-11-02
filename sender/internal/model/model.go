package model

import "github.com/google/uuid"

type Message struct {
	Id          uuid.UUID `json:"id"`
	MessageText string    `json:"text"`
}
