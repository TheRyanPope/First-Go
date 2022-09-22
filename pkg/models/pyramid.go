package models

import "github.com/google/uuid"

type Pyramid struct {
	Id     uuid.UUID `json:"id"`
	Levels []Level   `json:"level"`
}
