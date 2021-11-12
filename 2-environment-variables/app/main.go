package main

import (
	"log"
	"net/http"
	"os"
)

const (
	envVarName = "DATABASE_PASSWORD"
)

func handler(w http.ResponseWriter, r *http.Request) {
	valueFromEnv := os.Getenv(envVarName)
	if valueFromEnv == "" {
		valueFromEnv = "NOT-DEFINED-BY-ENV-VAR"
	}
	w.Write([]byte(valueFromEnv))
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
