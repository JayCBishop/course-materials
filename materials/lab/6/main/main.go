package main

import (
	"net/http"
	"scrape/logger"
	"scrape/scrape"

	"github.com/gorilla/mux"
)

// I separated the logging in to the scrape/logger module

func main() {

	logger.Logln(logger.Info, "starting API server")
	//create a new router
	router := mux.NewRouter()
	logger.Logln(logger.Info, "creating routes")
	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")
	//TODO_2 router.HandleFunc("/addsearch/{regex}", scrape.TODOREPLACE).Methods("GET")
	//TODO_3 router.HandleFunc("/clear", scrape.TODOREPLACE).Methods("GET")
	//TODO_4 router.HandleFunc("/reset", scrape.TODOREPLACE).Methods("GET")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
