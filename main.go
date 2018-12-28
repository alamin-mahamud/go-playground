package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kr/pretty"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/alamin/mahamud/api/url/{id}", simpleHandler)
	http.Handle("/", r)
	checkErr(http.ListenAndServe(":8080", nil))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(fmt.Sprintf("%#v", pretty.Formatter(r.URL)))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	w.Write(jsonData)
}
