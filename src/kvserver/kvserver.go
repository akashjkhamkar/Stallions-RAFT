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

func (kv *KVStore) getValueHandler(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	var command Command
	json.Unmarshal(reqBody, &command)
	json.NewEncoder(res).Encode(kv.store[command.Key])
}

func (kv *KVStore) upsertValueHandler(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	json_string_command := string(reqBody)
	kv.rf.Start(json_string_command)
	json.NewEncoder(res).Encode("OK!")
}

func (kv *KVStore) getAllValueHandler(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(kv.store)
}

func (kv *KVStore) apply_channel_listener() {
	for {
		select{
			case msg := <- kv.applyMsg:
				kv.store_mu.Lock()

				data := msg.Command
				jsonstr, _ := data.(string)
				var cmd Command
				json.Unmarshal([]byte(jsonstr), &cmd)
				kv.store[cmd.Key] = cmd.Value

				kv.store_mu.Unlock()
				fmt.Print("stored", cmd)
		}
	}
}

func StartKVserver(rf *raft.Raft, applyMsg chan raft.ApplyMsg, id int) {
	// Initialize kv
	kv := &KVStore{}
	kv.rf = rf
	kv.applyMsg = applyMsg
	kv.store = make(map[string]string)

	// Initialize routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get/", kv.getValueHandler).Methods("POST")
	router.HandleFunc("/upsert/", kv.upsertValueHandler).Methods("POST")
	router.HandleFunc("/all/", kv.getAllValueHandler).Methods("GET")

	port := 8000 + id
	go kv.apply_channel_listener() 

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("KVserver up on port 8000")
	}
}
