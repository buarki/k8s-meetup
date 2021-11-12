package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	pissedOffPeople = "VIDEO_URL"
)

func handler(w http.ResponseWriter, r *http.Request) {
	databaseHost := os.Getenv("DATABASE_HOST")
	if databaseHost == "" {
		databaseHost = "DATABASE_HOST_NOT-DEFINED-BY-ENV-VAR"
	}
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	if databasePassword == "" {
		databasePassword = "DATABASE_PASSWORD_NOT-DEFINED-BY-ENV-VAR"
	}
	databaseUser := os.Getenv("DATABASE_USER")
	if databaseUser == "" {
		databaseUser = "DATABASE_USER_NOT-DEFINED-BY-ENV-VAR"
	}
	dataToSend := fmt.Sprintf("DATABASE_HOST: %s\nDATABASE_PASSWORD: %s\nDATABASE_USER: %s", databaseHost, databasePassword, databaseUser)
	w.Write([]byte(dataToSend))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT env var must be defined")
	}
	http.HandleFunc("/", handler)
	log.Println("listening at " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
