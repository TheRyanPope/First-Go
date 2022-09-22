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

func UpdateChallengerName(w http.ResponseWriter, r *http.Request) {
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

	if updatedChallenger.Name == "" {
		log.Fatalln("Name is a required field to update a challengers name.")
	}

	for index, challenger := range mocks.Challengers {
		if challenger.Id == id {
			var response string
			previousName := challenger.Name

			if challenger.Name == updatedChallenger.Name {
				response = challenger.Name + " is already " + updatedChallenger.Name
			} else {
				mocks.Challengers[index].Name = updatedChallenger.Name
				response = "Challenger name updated from " + previousName + " to " + updatedChallenger.Name
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
			break
		}
	}
}
