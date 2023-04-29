package kvserver

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getValueHandler(w http.ResponseWriter, r *http.Request) {
	// query params : /get-value?key=<key>
	// construcy a command object and call relevant method
	fmt.Printf("GET request\n")
	io.WriteString(w, "Get Request Received\n")
}

func upsertValueHandler(w http.ResponseWriter, r *http.Request) {
	// query params : /get-value?key=<key?value=<value>
	fmt.Printf("UPSERT request\n")
	io.WriteString(w, "Upsert Request Received!\n")
}

func StartKVserver() {
	// initialize routes
	http.HandleFunc("/get-value", getValueHandler)
	http.HandleFunc("/upsert-value", upsertValueHandler)

	// need to start another go-routine to track the applychannel
	// add KVserver struct <- which will hold rf instance

	fmt.Printf("KVserver up on port 8000")

	err := http.ListenAndServe(":8000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
