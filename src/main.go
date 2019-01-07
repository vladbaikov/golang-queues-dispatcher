package main

import (
	d "dispatcher"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	maxWorkers = 4
	port       = ":8086"
)

var (
	dispatcher = d.NewDispatcher(maxWorkers)
)

type Request struct {
	Data string `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	request := Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	dispatcher.AddJob(d.Job{Data: request.Data})
	_, err = w.Write([]byte("Data added to queue"))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	dispatcher.Run()

	fmt.Println("Server is listening on", port)
	http.HandleFunc("/handler", Handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
