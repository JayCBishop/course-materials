package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id          string `json:"id"`
	Title       string `json:"title`
	Description string `json:"desc"`
	Points      int    `json:"points"`
}

var Assignments []Assignment

const Valkey string = "FooKey"

func InitAssignments() {
	var assignment Assignment
	assignment.Id = "Mike1A"
	assignment.Title = "Lab 4"
	assignment.Description = "Some lab this guy made yesterday?"
	assignment.Points = 20
	Assignments = append(Assignments, assignment)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(assignment)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
		if assignment.Id == params["id"] {
			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			response["status"] = "Success"
			break
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	params := mux.Vars(r)

	for index, assignment := range Assignments {
		if assignment.Id == params["id"] {
			Assignments = append(Assignments[:index], Assignments[index+1:]...)

			if r.FormValue("title") != "" {
				assignment.Title = r.FormValue("title")
			}

			if r.FormValue("desc") != "" {
				assignment.Description = r.FormValue("desc")
			}

			if r.FormValue("points") != "" {
				points, err := strconv.Atoi(r.FormValue("points"))
				if err == nil {
					assignment.Points = points
				}
			}

			Assignments = append(Assignments, assignment)
			break
		}
	}

	var response Response
	response.Assignments = Assignments

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignment Assignment
	r.ParseForm()

	var errors []string

	if r.FormValue("id") == "" {
		errors = append(errors, "An id is required for an assignment")
	}

	if r.FormValue("title") == "" {
		errors = append(errors, "A title is required for an assignment")
	}

	if r.FormValue("desc") == "" {
		errors = append(errors, "A description is required for an assignment")
	}

	if r.FormValue("points") == "" {
		errors = append(errors, "A point value is required for an assignment")
	} else {
		points, err := strconv.Atoi(r.FormValue("points"))
		if err != nil {
			errors = append(errors, "An integer value is required for an assignment's points")
		} else {
			assignment.Points = points
		}
	}

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
	} else {
		assignment.Id = r.FormValue("id")
		assignment.Title = r.FormValue("title")
		assignment.Description = r.FormValue("desc")
		Assignments = append(Assignments, assignment)
		w.WriteHeader(http.StatusCreated)
	}
}
