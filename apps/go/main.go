package main

import (
	"net/http"
)

func api(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`{"status":"OK"}`))
}

func main() {
	http.HandleFunc("/", api)
	http.ListenAndServe(":8080", nil)
}
