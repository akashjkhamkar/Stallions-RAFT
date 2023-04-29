package kvserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Command struct {
	Operation string `json:"id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func getValueHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var command Command
	json.Unmarshal(reqBody, &command)

	// fetch data from in-memory map for key command.key and return val

	json.NewEncoder(w).Encode(command)

}

func upsertValueHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	command := string(reqBody)

	// why string ? to call rf.Start() with command

	json.NewEncoder(w).Encode(command)

}

func StartKVserver() {

	// initialize routes
	fmt.Println("Added")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/get-value/", getValueHandler).Methods("POST")
	router.HandleFunc("/upsert-value/", upsertValueHandler).Methods("POST")

	// need to start another go-routine to track the applychannel
	// add KVserver struct <- which will hold rf instance

	fmt.Printf("KVserver up on port 8000")

	err := http.ListenAndServe(":8000", router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
