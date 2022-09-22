package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/malwhare/First-Go/pkg/mocks"
)

func ActivateChallenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := uuid.Parse(vars["id"])

	if idError != nil {
		log.Fatalln((idError))
	}

	for index, challenger := range mocks.Challengers {
		if challenger.Id == id {
			var response string

			if challenger.IsActive {
				response = mocks.Challengers[index].Name + " is already Activated"
			} else {
				mocks.Challengers[index].IsActive = true
				response = mocks.Challengers[index].Name + " Activated"
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
			break
		}
	}
}
