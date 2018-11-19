package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func connect() {
	log.Println("Starting server at port :8180")
	http.ListenAndServe(":8180", nil)
}

type Response struct {
	Status string
}

func getPsql(w http.ResponseWriter, r *http.Request) {
	log.Println("getting data")
	resp := Response{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/getpsql", getPsql)
	log.Println(http.ListenAndServe(":8180", logger()))
}

func logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
	})
}
