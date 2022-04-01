package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"scrape/logger"
	"strconv"

	"github.com/gorilla/mux"
)

//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and
// if responsewriter is passed, outputs to http

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		for _, r := range regexes {
			if r.MatchString(path) {
				var tfile FileInfo
				dir, filename := filepath.Split(path)
				tfile.Filename = string(filename)
				tfile.Location = string(dir)

				if containsFile(tfile) {
					continue
				}

				Files = append(Files, tfile)

				if w != nil && len(Files) > 0 {
					w.Write([]byte(`"` + (strconv.FormatInt(1+FilesAdded, 10)) + `":  `))
					json.NewEncoder(w).Encode(tfile)
					w.Write([]byte(`,`))
				}

				logger.Log(logger.Debug, "[+] HIT: %s\n", path)
				FilesAdded++
			}
		}
		return nil
	}
}

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		regex, regErr := regexp.Compile(query)

		if regErr != nil {
			logger.Logln(logger.Info, "`%s` is not a valid regular expression", query)
			w.WriteHeader(401)
			return nil
		}

		if regex.MatchString(path) {
			var tfile FileInfo
			dir, filename := filepath.Split(path)
			tfile.Filename = string(filename)
			tfile.Location = string(dir)

			if containsFile(tfile) {
				return nil
			}

			Files = append(Files, tfile)

			if w != nil && len(Files) > 0 {
				w.Write([]byte(`"` + (strconv.FormatInt(1+FilesAdded, 10)) + `":  `))
				json.NewEncoder(w).Encode(tfile)
				w.Write([]byte(`,`))
			}

			logger.Log(logger.Debug, "[+] HIT: %s\n", path)
			FilesAdded++
		}
		return nil
	}
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	logger.Log(logger.Info, "Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status" : "API is up and running ",`))
	var regexstrings []string

	for _, regex := range regexes {
		regexstrings = append(regexstrings, regex.String())
	}

	w.Write([]byte(` "regexs" :`))
	json.NewEncoder(w).Encode(regexstrings)
	w.Write([]byte(`}`))
	logger.Logln(logger.Debug, regexes)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	logger.Log(logger.Info, "Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "<html><body><H1>This API walks the local file system, searching for files that match regular expressions</H1></body>")
}

func FindFile(w http.ResponseWriter, r *http.Request) {
	logger.Log(logger.Info, "Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	q, ok := r.URL.Query()["q"]

	w.WriteHeader(http.StatusOK)
	if ok && len(q[0]) > 0 {
		logger.Log(logger.Info, "Entering search with query=%s", q[0])

		// ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
		// e.g., func finder(query string) []string { ... }

		var FOUND = false
		for _, File := range Files {
			if File.Filename == q[0] {
				json.NewEncoder(w).Encode(File.Location)
				FOUND = true
			}
		}

		if !FOUND {
			var notFoundMessage = fmt.Sprintf(`"status": "The file %s was not found"`, q[0])
			logger.Logln(logger.Debug, notFoundMessage)
			w.Write([]byte(notFoundMessage))
			return
		}

	} else {
		// didn't pass in a search term, show all that you've found
		w.Write([]byte(`"files":`))
		json.NewEncoder(w).Encode(Files)
	}
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	logger.Log(logger.Info, "Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	location, locOK := r.URL.Query()["location"]

	var rootDir string = "C:/Grad School/"

	if locOK && len(location[0]) > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(`{ "parameters" : {"required": "location",`))
		w.Write([]byte(`"optional": "regex"},`))
		w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
		w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
		return
	}

	var dir = rootDir + location[0]

	//wrapper to make "nice json"
	w.Write([]byte(`{ `))

	regex, regOK := r.URL.Query()["regex"]
	if regOK && len(regex[0]) > 0 {
		if err := filepath.Walk(dir, walkFn2(w, regex[0])); err != nil {
			logger.Panicln(logger.Debug, err)
		}
	} else if err := filepath.Walk(dir, walkFn(w)); err != nil {
		logger.Panicln(logger.Debug, err)
	}

	//wrapper to make "nice json"
	w.Write([]byte(` "status": "completed"} `))
}

func ClearFiles(w http.ResponseWriter, r *http.Request) {
	logger.Log(logger.Info, "Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resetRegEx()
	Files = nil
	FilesAdded = 0
	w.Write([]byte(`"status": "completed"`))
}

func ClearRegex(w http.ResponseWriter, r *http.Request) {
	logger.Log(logger.Info, "Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	clearRegEx()
	w.Write([]byte(`"status": "completed"`))
}

func AddRegex(w http.ResponseWriter, r *http.Request) {
	logger.Log(logger.Info, "Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)
	if len(params["regex"]) > 0 {
		regex := "(?i)" + params["regex"]
		logger.Log(logger.Info, "Entering add with regex=%s", regex)
		addRegEx(regex)
	}

	w.Write([]byte(`"status": "completed"`))
}

func containsFile(fileInfo FileInfo) bool {
	for _, file := range Files {
		if file == fileInfo {
			return true
		}
	}
	return false
}
