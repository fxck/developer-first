package main

import (
	"net/http"
)

func api(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(`{"status":"OK"}`))
}

func main() {
	http.HandleFunc("/", api)
	http.ListenAndServe(":8080", nil)
}
