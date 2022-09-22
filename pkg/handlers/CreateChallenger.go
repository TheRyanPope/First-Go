package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/malwhare/First-Go/pkg/mocks"
	"github.com/malwhare/First-Go/pkg/models"
)

func CreateChallenger(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, bodyError := io.ReadAll(r.Body)

	if bodyError != nil {
		log.Fatalln(bodyError)
	}

	var challenger models.Challenger
	json.Unmarshal(body, &challenger)

	if challenger.Name == "" {
		log.Fatalln("Name is a required field to create a challenger.")
	}

	challenger.Id = uuid.New()
	challenger.IsActive = true
	mocks.Challengers = append(mocks.Challengers, challenger)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created: " + challenger.Id.String())
}
