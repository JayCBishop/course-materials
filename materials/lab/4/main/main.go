package main

import (
	"log"
	"net/http"
	"wyoassign/classes"
	"wyoassign/wyoassign"

	"github.com/gorilla/mux"
)

func main() {
	wyoassign.InitAssignments()
	classes.InitClasses()
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.GetAssignment).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.DeleteAssignment).Methods("DELETE")
	router.HandleFunc("/assignment", wyoassign.CreateAssignment).Methods("POST")
	router.HandleFunc("/assignments/{id}", wyoassign.UpdateAssignment).Methods("PUT")

	router.HandleFunc("/classes", classes.GetClasses).Methods("GET")
	router.HandleFunc("/class/{id}", classes.GetClass).Methods("GET")
	router.HandleFunc("/class/{id}", classes.DeleteClass).Methods("DELETE")
	router.HandleFunc("/class", classes.CreateClass).Methods("POST")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
