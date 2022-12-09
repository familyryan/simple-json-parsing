package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Ric", 24},
		{"Jason", 22},
	}

	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func main() {
	argumen := []struct {
		Name string
		Age  int
	}{
		{"Ria", 22},
		{"Jaso", 21},
	}

	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		jsonInBytes, err := json.Marshal(argumen)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)
	})
	fmt.Println("server at localhost:9000 ")
	http.ListenAndServe(":9000", nil)
}
