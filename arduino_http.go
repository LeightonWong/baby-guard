package main

import (
	"fmt"
	"log"

	"net/http"
)

func record(w http.ResponseWriter, r *http.Request) {
	temp := r.FormValue("temp")
	log.Println("Save temprature", temp)
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/record", record)
	http.ListenAndServe(":8080", nil)
}
