package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const (
	// Port is the port the server will listen on
	Port = "8080"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!!\n"))
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = Port
	}
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
