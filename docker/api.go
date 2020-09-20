package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type InvalidRC struct {
	Result struct {
		Valid bool `json:"valid"`
	} `json:"result"`
}

func isValidRC(rc string) bool {
	re := regexp.MustCompile(`^((\d\d)(\d\d)(\d\d)){1}[ /]*((\d\d\d)(\d?)){1}$`)
	if !re.MatchString(rc) {
		return false
	}

	rc = strings.ReplaceAll(rc, "/", "")
	rc = strings.ReplaceAll(rc, " ", "")
	rc = strings.TrimSpace(rc)

	year, _ := strconv.Atoi(rc[0:2])
	month, _ := strconv.Atoi(rc[2:4])
	day, _ := strconv.Atoi(rc[4:6])
	ext := rc[6:]
	c, _ := strconv.Atoi(rc[9:])

	// less than 3 digits or more than four digits is wrong
	if len(ext) < 3 || len(ext) > 4 {
		return false
	}

	if len(rc[9:]) == 0 {
		if year < 54 {
			year = year + 1900
		} else {
			year = year + 1800
		}
	} else {
		var full, _ = strconv.Atoi(rc[:9])
		mod := full % 11
		if mod == 10 {
			mod = 0
		}
		if mod != c {
			return false
		}

		if year < 54 {
			year = year + 2000
		} else {
			year = year + 1900
		}
	}

	// get the "real" month value
	if month > 70 && year > 2003 {
		month = month - 70
	} else if month > 50 {
		month = month - 50
	} else if month > 20 && year > 2003 {
		month = month - 20
	}

	// higher day number than possible
	if day > 29 && month == 2 {
		return false
	}
	if day > 30 && (month == 4 || month == 6 || month == 9 || month == 11) {
		return false
	}
	if day > 31 && (month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) {
		return false
	}

	return true
}

func RCChecker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	if isValidRC(string(vars["id"])) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			panic(err)
		}
	} else {
		_, err := w.Write([]byte("NOK"))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/validate/{id}", RCChecker)
	log.Fatal(http.ListenAndServe(":80", router))
}
