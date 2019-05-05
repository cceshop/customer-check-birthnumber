package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type InvalidRC struct {
	Result struct {
		Valid    bool   `json:"valid"`
	} `json:"result"`
}

func isValidRC(rc string) bool {
	url := "https://rcapi.abalin.net/validate/" + rc
	var contents []byte
	var id InvalidRC

	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	contents, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(contents, &id)
	if err != nil {
		panic(err)
	}

	if id.Result.Valid {
		return true
	}

	return false
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
	router.HandleFunc("/validate/cardid/{id}", RCChecker)
	log.Fatal(http.ListenAndServe(":80", router))
}
