package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/malwhare/First-Go/pkg/mocks"
)

func GetChallenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := uuid.Parse(vars["id"])

	if idError != nil {
		log.Fatalln((idError))
	}

	for _, challenger := range mocks.Challengers {
		if challenger.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(challenger)
			break
		}
	}

}
