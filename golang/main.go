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
	_, err := w.Write([]byte("Hello World!!!\n"))
	if err != nil {
		panic(err)
	}
}

func RandomFileGenerator(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	file, err := os.CreateTemp(os.TempDir(), fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bigBuff := make([]byte, 750000000)
	_, err = file.Write(bigBuff)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/file/{fileName}", RandomFileGenerator)

	port := os.Getenv("PORT")
	if port == "" {
		port = Port
	}
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
