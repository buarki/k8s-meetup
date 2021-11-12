package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	fileToBeRead = "/usr/data/dumbtext"
)

func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(fileToBeRead)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Println("listening at " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
