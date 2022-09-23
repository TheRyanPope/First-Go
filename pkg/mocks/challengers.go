package mocks

import "github.com/malwhare/First-Go/pkg/models"

var Challengers = []models.Challenger{
	{
		Id:           int64(1),
		Name:         "Doofus McGee",
		CurrentRank:  "Last",
		PreviousRank: "First",
		IsActive:     true,
	},
}
