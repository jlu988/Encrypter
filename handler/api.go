package handler

import (
	"github.com/gorilla/mux"
)

func RouterInitialize() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/encrypt", Encryption).Methods("GET")
	r.HandleFunc("/decrypt", Decryption).Methods("GET")
	return r
}
