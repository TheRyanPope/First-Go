package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/malwhare/First-Go/internal"
	"github.com/malwhare/First-Go/pkg/controllers"
)

func main() {
	_ = internal.ConnectDB()

	createRoutes()
}

func createRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/pyramids/getPyramid", controllers.GetPyramid).Methods(http.MethodGet)

	router.HandleFunc("/challengers/create", controllers.CreateChallenger).Methods(http.MethodPost)
	router.HandleFunc("/challengers/get/{id}", controllers.GetChallenger).Methods(http.MethodGet)
	router.HandleFunc("/challengers/updateRank/{id}", controllers.UpdateChallengerRank).Methods(http.MethodPut)
	router.HandleFunc("/challengers/updateName/{id}", controllers.UpdateChallengerName).Methods(http.MethodPut)
	router.HandleFunc("/challengers/deactivate/{id}", controllers.DeactivateChallenger).Methods(http.MethodPut)
	router.HandleFunc("/challengers/activate/{id}", controllers.ActivateChallenger).Methods(http.MethodPut)

	http.ListenAndServe(":4000", router)
}
