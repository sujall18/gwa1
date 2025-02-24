package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving homepage at /")
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func main() {

	http.HandleFunc("/", homePage)

	log.Println("Server is starting at 8080...")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

