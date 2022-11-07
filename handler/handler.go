package handler

import (
	"Encrypter/domain"
	"Encrypter/models"
	"net/http"

	"github.com/gorilla/schema"
)

func Encoder(w http.ResponseWriter, r *http.Request) {
	var (
		encryptionService = domain.EncryptionHandler{}
		data              = models.EncryptData{}
	)

	decoder := schema.NewDecoder()
	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		Response(w, http.StatusInternalServerError, err.Error())
	}

	encodedMsg, err := encryptionService.Encode(data)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, encodedMsg)
}

func Decoder(w http.ResponseWriter, r *http.Request) {
	var (
		encryptionService = domain.EncryptionHandler{}
		data              = models.EncryptData{}
	)
	decoder := schema.NewDecoder()
	err := decoder.Decode(&data, r.URL.Query())
	if err != nil {
		Response(w, http.StatusInternalServerError, err.Error())
	}

	decodedMsg, err := encryptionService.Decode(data)
	if err != nil {
		Response(w, http.StatusBadRequest, err.Error())
		return
	}

	ResponseOk(w, decodedMsg)
}
