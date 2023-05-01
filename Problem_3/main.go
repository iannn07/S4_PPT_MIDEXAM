package main

import (
	"Data"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/nama", single_data)
	mux.HandleFunc("/semuadata", all_data)
	http.ListenAndServe(":5050", mux)
}

func single_data(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Nama")
	w.Header().Set("Content-Type", "application/json")
	str, err := json.Marshal(Data.SingleData(name))
	if err != nil {
		fmt.Print(err)
	} else {

		io.WriteString(w, string(str))
	}
}

func all_data(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	str, err := json.Marshal(Data.AllData())
	if err != nil {
		fmt.Print(err)
	} else {

		io.WriteString(w, string(str))
	}
}
