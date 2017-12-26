package main

import (
	"github.com/dix975/database"
	"github.com/dix975/logger"
	"github.com/gorilla/mux"
	"github.com/dix975/www"
	"time"
	"net/http"
	"log"
	"github.com/dix975/score/score"
	"github.com/dix975/score/team"
)

func main(){

	logger.Init()

	logger.Info.Println("Booting")


	logger.Info.Println("Initiating Mongo Connection with error handling")

	_, err := db.NewDB(
		db.MongoServerConfig{
			AuthDatabaseName: "score",
			URL:"mongodb://mongo:27017",
		})

	if err != nil { panic(err)}
	router := mux.NewRouter()

	router.Methods("GET").Path("/").Handler(www.Handle{score.HandleRoot})
	router.Methods("POST").Path("/team").Handler(www.Handle{team.PostTeam})

	server := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info.Println("Server Listning")
	log.Fatal(server.ListenAndServe())
}
