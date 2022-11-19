package services

import (
	"Encrypter/database"
	"Encrypter/models"
)

func Encode(obj models.EncryptData, internalKey string) (string, error) {
	service := Service{}
	dbHandler := database.DbHandler{}
	keys := models.Internal{
		OriginalKey: obj.Key,
		InternalKey: internalKey,
	}
	err := dbHandler.AddInternalKey(keys)
	if err != nil {
		return "", err
	}
	obj.Key = internalKey

	data, err := service.constructor(obj)
	if err != nil {
		return "", err
	}

	cipherMatrix := service.initializeCipherMatrix(obj.Value, data.TransPosKey, data.RowSize, data.ColumnSize)

	encodedMessage := ""
	for index := 0; index < data.RowSize; index++ {
		for col := 0; col < len(cipherMatrix); col++ {
			char := []byte(cipherMatrix[col][index])[0] + 1
			encodedMessage += string(char)
		}
	}

	encodedMessage = internalKey + encodedMessage
	return encodedMessage, nil
}
