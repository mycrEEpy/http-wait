package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var (
	hostname string
)

func main() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		fmt.Print("Hostname not found\n")
		os.Exit(1)
	}
	r := mux.NewRouter()
	r.HandleFunc("/{millis}", handler)
	http.Handle("/", r)
	fmt.Printf("Listening on 8080\nHit me with curl %s:8080/<milliseconds>\n", hostname)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("HTTP listen and serve: %s\n", err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	millis, err := strconv.ParseInt(vars["millis"], 10, 64)
	if err != nil {
		fmt.Fprintf(w, "NaN")
		return
	}
	time.Sleep(time.Millisecond * time.Duration(millis))
	fmt.Fprintf(w, "Hello from %s! I have waited for %d milliseconds\n", hostname, millis)
}
