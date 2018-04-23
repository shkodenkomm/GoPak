package main 

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

func main() {
	router := mux.NewRouter()
    router.HandleFunc("/", defaultHandler).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}
