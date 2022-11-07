package handler

import (
	"github.com/gorilla/mux"
)

func RouterInitialize() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/encoder", Encoder).Methods("GET")
	r.HandleFunc("/decoder", Decoder).Methods("GET")
	return r
}
