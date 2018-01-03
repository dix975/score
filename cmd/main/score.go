package main

import (
	"dix975.com/database"
	"dix975.com/score/configuration"
	"dix975.com/score/score"
	"dix975.com/score/team"
	"github.com/dix975/logger"
	"github.com/dix975/www"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {

	logger.Init()

	logger.Info.Println("Initiating Mongo Connection")

	_, err := db.NewDB(configuration.Config().MongoDBConfig)

	if err != nil {
		logger.Error.Println(err)
	}

	router := mux.NewRouter()

	router.Methods("GET").Path("/").Handler(www.Handle{score.HandleRoot})
	router.Methods("POST").Path("/team").Handler(www.Handle{team.PostTeam})

	server := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info.Println("Server Listning")
	log.Fatal(server.ListenAndServe())
}
