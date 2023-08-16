package main

import (
	"net/http"

	"github.com/andrisutanto/golang-modal-crud/controllers/mahasiswacontroller"
)

func main() {
	http.HandleFunc("/", mahasiswacontroller.Index)
	http.ListenAndServe(":8000", nil)
}
