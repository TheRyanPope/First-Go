package mocks

import (
	"github.com/google/uuid"
	"github.com/malwhare/First-Go/pkg/models"
)

var Pyramid = models.Pyramid{
	Id: uuid.New(),
	Levels: []models.Level{
		{
			Id:   uuid.New(),
			Name: "1",
			Challengers: []models.Challenger{
				{
					Id:           uuid.New(),
					Name:         "Bony Tower",
					CurrentRank:  "1",
					PreviousRank: "6",
					IsActive:     true,
				},
			},
			Slots: 1,
		},
		{
			Id:   uuid.New(),
			Name: "2",
			Challengers: []models.Challenger{
				{
					Id:           uuid.New(),
					Name:         "Man Iunnerley",
					CurrentRank:  "2",
					PreviousRank: "3",
					IsActive:     true,
				},
				{
					Id:           uuid.New(),
					Name:         "Mim Tunnerley",
					CurrentRank:  "3",
					PreviousRank: "2",
					IsActive:     true,
				},
			},
			Slots: 2,
		},
		{
			Id:   uuid.New(),
			Name: "2",
			Challengers: []models.Challenger{
				{
					Id:           uuid.New(),
					Name:         "DatRick Pavis",
					CurrentRank:  "4",
					PreviousRank: "5",
					IsActive:     true,
				},
				{
					Id:           uuid.New(),
					Name:         "Dat Pavis",
					CurrentRank:  "5",
					PreviousRank: "4",
					IsActive:     true,
				},
				{
					Id:           uuid.New(),
					Name:         "Trick Davis",
					CurrentRank:  "6",
					PreviousRank: "1",
					IsActive:     true,
				},
			},
			Slots: 3,
		},
	},
}
