package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/malwhare/First-Go/pkg/mocks"
	"github.com/malwhare/First-Go/pkg/models"
)

func ActivateChallenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := strconv.ParseInt(vars["id"], 10, 64)

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

	//challenger.Id = uuid.New()
	challenger.IsActive = true
	mocks.Challengers = append(mocks.Challengers, challenger)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created: " + strconv.Itoa(int(challenger.Id)))
}

func DeactivateChallenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := strconv.ParseInt(vars["id"], 10, 64)

	if idError != nil {
		log.Fatalln((idError))
	}

	for index, challenger := range mocks.Challengers {
		if challenger.Id == id {
			var response string

			if !challenger.IsActive {
				response = mocks.Challengers[index].Name + " is already Deactivated"
			} else {
				mocks.Challengers[index].IsActive = false
				response = mocks.Challengers[index].Name + " Deactivated"
			}

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
			break
		}
	}
}

func GetChallenger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := strconv.ParseInt(vars["id"], 10, 64)

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

func UpdateChallengerName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := strconv.ParseInt(vars["id"], 10, 64)

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

func UpdateChallengerRank(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, idError := strconv.ParseInt(vars["id"], 10, 64)

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
