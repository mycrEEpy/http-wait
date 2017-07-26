package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{seconds}", handler)
	http.Handle("/", r)
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seconds, err := strconv.Atoi(vars["seconds"])
	if err != nil {
		fmt.Fprintf(w, "NaN")
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Hostname not found")
		return
	}
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Fprintf(w, "Hello from %v! I have waited for %v seconds", hostname, seconds)
}
