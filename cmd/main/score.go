package main

import (
	"dix975.com/logger"
	"github.com/gorilla/mux"
	"dix975.com/www"
	"dix975.com/score/score"
	"time"
	"net/http"
	"log"
	"dix975.com/score/team"
)

func main(){

	logger.Init()

	logger.Info.Println("Booting")

	router := mux.NewRouter()

	router.Methods("GET").Path("/").Handler(www.Handle{score.HandleRoot})
	router.Methods("POST").Path("/team").Handler(www.Handle{team.PostTeam})

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
