package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/malwhare/First-Go/pkg/mocks"
	"github.com/malwhare/First-Go/pkg/models"
)

func UpdateChallengerRank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := uuid.Parse(vars["id"])

	if idError != nil {
		log.Fatalln((idError))
	}

	defer r.Body.Close()
	body, bodyError := io.ReadAll(r.Body)

	if bodyError != nil {
		log.Fatalln(bodyError)
	}

	var updatedChallenger models.Challenger
	json.Unmarshal(body, &updatedChallenger)

	if updatedChallenger.CurrentRank == "" {
		log.Fatalln("Current Rank is a required field to update a challengers rank.")
	}

	for index, challenger := range mocks.Challengers {
		if challenger.Id == id {
			var response string

			if challenger.CurrentRank == updatedChallenger.CurrentRank {
				response = challenger.Name + " is already rank " + challenger.CurrentRank
			} else {
				challenger.PreviousRank = challenger.CurrentRank
				challenger.CurrentRank = updatedChallenger.CurrentRank

				mocks.Challengers[index] = challenger

				response = challenger.Name + " rank has been updated from " + challenger.PreviousRank + " to " + updatedChallenger.CurrentRank
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
			break
		}
	}
}
