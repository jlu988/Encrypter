package main

import (
	"Encrypter/handler"
	"fmt"
	"net/http"
)

func main() {
	router := handler.RouterInitialize()
	fmt.Println("Listen and serve")
	http.ListenAndServe(":8080", router)
}
