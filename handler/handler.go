package handler

import (
	"Encrypter/database"
	"Encrypter/internal"
	"Encrypter/models"
	"Encrypter/services"
	"net/http"

	"github.com/gorilla/schema"
)

func Encryption(w http.ResponseWriter, r *http.Request) {
	var (
		data = models.EncryptData{}
	)

	decoder := schema.NewDecoder()
	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		Response(w, http.StatusInternalServerError, err.Error())
	}
	internalPrivateKey := internal.GeneratePrivateKey(data.Key)
	encodedMsg, err := services.Encode(data, internalPrivateKey)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, encodedMsg)
}

func Decryption(w http.ResponseWriter, r *http.Request) {
	var (
		data = models.EncryptData{}
	)
	decoder := schema.NewDecoder()
	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		Response(w, http.StatusInternalServerError, err.Error())
	}
	dbHandler := database.DbHandler{}
	internalPrivateKey, err := dbHandler.GetInternalKey(data.Key)
	if err != nil {
		Response(w, http.StatusInternalServerError, err.Error())
	}
	decodedMsg, err := services.Decrypt(data, internalPrivateKey.InternalKey)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponseOk(w, decodedMsg)
}
