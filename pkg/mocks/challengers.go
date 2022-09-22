package mocks

import (
	"github.com/google/uuid"
	"github.com/malwhare/First-Go/pkg/models"
)

var Challengers = []models.Challenger{
	{
		Id:           uuid.New(),
		Name:         "Doofus McGee",
		CurrentRank:  "Last",
		PreviousRank: "First",
		IsActive:     true,
	},
}
