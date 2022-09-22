package models

import "github.com/google/uuid"

type Level struct {
	Id          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Challengers []Challenger `json:"challenger"`
	Slots       int          `json:"slots"`
}
