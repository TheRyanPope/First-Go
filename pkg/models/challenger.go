package models

import "github.com/google/uuid"

type Challenger struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	CurrentRank  string    `json:"currentRank"`
	PreviousRank string    `json:"previousRank"`
	IsActive     bool      `json:"isActive"`
}
