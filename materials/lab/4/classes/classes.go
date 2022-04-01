package classes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Classes []Class `json:"classes"`
}

type Class struct {
	Id          string `json:"id"`
	Title       string `json:"title`
	Description string `json:"desc"`
	Department  string `json:"department"`
}

var Classes []Class

func InitClasses() {
	var class Class
	class.Id = "COSC-5010-03"
	class.Title = "Cyber Security"
	class.Description = "Learn the Go programming language through a series of practical cyber-security-oriented challenges."
	class.Department = "Computer Science"
	Classes = append(Classes, class)
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Classes = Classes

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func GetClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, class := range Classes {
		if class.Id == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(class)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func DeleteClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, class := range Classes {
		if class.Id == params["id"] {
			Classes = append(Classes[:index], Classes[index+1:]...)
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

func CreateClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var class Class
	r.ParseForm()

	var errors []string

	if r.FormValue("id") == "" {
		errors = append(errors, "An id is required for a class")
	}

	if r.FormValue("title") == "" {
		errors = append(errors, "A title is required for a class")
	}

	if r.FormValue("desc") == "" {
		errors = append(errors, "A description is required for a class")
	}

	if r.FormValue("department") == "" {
		errors = append(errors, "A department is required for a class")
	}

	if len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
	} else {
		class.Id = r.FormValue("id")
		class.Title = r.FormValue("title")
		class.Description = r.FormValue("desc")
		class.Department = r.FormValue("department")
		Classes = append(Classes, class)
		w.WriteHeader(http.StatusCreated)
	}
}
