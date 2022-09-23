package models

type Challenger struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	CurrentRank  string `json:"currentRank"`
	PreviousRank string `json:"previousRank"`
	IsActive     bool   `json:"isActive"`
}
