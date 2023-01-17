package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
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

func PrintEnvByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	envVar := vars["envVar"]
	if os.Getenv(envVar) != "" {
		w.WriteHeader(http.StatusOK)
		enVarResult := fmt.Sprintf("%s=%s", envVar, os.Getenv(envVar))
		_, err := w.Write([]byte(enVarResult))
		if err != nil {
			panic(err)
		}
	}
}

func ReadFileAndPrintContent(w http.ResponseWriter, r *http.Request) {
	fileName := os.Getenv("FILE")
	readFile, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	_, err = w.Write(readFile)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/file/{fileName}", RandomFileGenerator)
	r.HandleFunc("/env/{envVar}", PrintEnvByName)
	r.HandleFunc("/read", ReadFileAndPrintContent)

	port := os.Getenv("PORT")
	if port == "" {
		port = Port
	}

	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "TEST") {
			log.Printf("%s", env)

		}
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
