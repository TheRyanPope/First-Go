package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/malwhare/First-Go/pkg/handlers"
)

func main() {
	createRoutes()
}

func createRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/pyramids/getPyramid", handlers.GetPyramid).Methods(http.MethodGet)

	router.HandleFunc("/challengers/create", handlers.CreateChallenger).Methods(http.MethodPost)
	router.HandleFunc("/challengers/get/{id}", handlers.GetChallenger).Methods(http.MethodGet)
	router.HandleFunc("/challengers/updateRank/{id}", handlers.UpdateChallengerRank).Methods(http.MethodPut)
	router.HandleFunc("/challengers/updateName/{id}", handlers.UpdateChallengerName).Methods(http.MethodPut)
	router.HandleFunc("/challengers/deactivate/{id}", handlers.DeactivateChallenger).Methods(http.MethodPut)
	router.HandleFunc("/challengers/activate/{id}", handlers.ActivateChallenger).Methods(http.MethodPut)

	http.ListenAndServe(":4000", router)
}
