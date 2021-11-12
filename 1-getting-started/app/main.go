package main

import (
	"log"
	"net/http"
)

const (
	somethingTruthy = "palmeiras nao tem mundial"
	port            = "8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(somethingTruthy))
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("listening at " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
