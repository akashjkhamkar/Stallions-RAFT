package kvserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"stallionraft/src/raft"
	"sync"

	"github.com/gorilla/mux"
)

type Command struct {
	Operation string `json:"id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

type KVStore struct {
	rf *raft.Raft
	store map[string]string
	applyMsg chan raft.ApplyMsg
	store_mu sync.Mutex
}

func (kv *KVStore) getValueHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var command Command
	json.Unmarshal(reqBody, &command)

	// fetch data from in-memory map for key command.key and return val
	json.NewEncoder(w).Encode(command)

}

func (kv *KVStore) upsertValueHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	command := string(reqBody)

	// why string ? to call rf.Start() with command
	json.NewEncoder(w).Encode(command)

}

func StartKVserver(rf *raft.Raft, applyMsg chan raft.ApplyMsg) {
	// Initialize kv
	kv := &KVStore{}
	kv.rf = rf
	kv.applyMsg = applyMsg
	kv.store = make(map[string]string)

	// Initialize routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get-value/", kv.getValueHandler).Methods("POST")
	router.HandleFunc("/upsert-value/", kv.upsertValueHandler).Methods("POST")

	err := http.ListenAndServe(":8000", router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("KVserver up on port 8000")
	}
}
