package services

import (
	"Encrypter/models"
	"fmt"
	"strings"
)

func Decrypt(obj models.EncryptData, internalKey string) (string, error) {
	service := Service{}
	if len(obj.Value) < len(obj.Key) {
		return "", fmt.Errorf("message cannot be shorter than key")
	}
	privateKey := obj.Value[0:len(obj.Key)]
	if privateKey != internalKey {
		return obj.Value, nil
	} else {
		obj.Value = obj.Value[len(obj.Key):]
	}
	obj.Key = internalKey
	data, err := service.constructor(obj)
	if err != nil {
		println("error occurred - ", err)
		return "", err
	}

	cipheredMatrix := service.initializeCipherMatrix(obj.Value, data.NumericKey, data.RowSize, data.ColumnSize)

	var matrixBuffer [][]string
	for _, val := range data.TransPosKey {
		if char, ok := cipheredMatrix[val]; ok {
			matrixBuffer = append(matrixBuffer, char)
		}
	}

	msgBuffer := ""
	for index := 0; index < data.RowSize; index++ {
		for j := 0; j < len(matrixBuffer); j++ {
			if matrixBuffer[j][index] != " " {
				char := []byte(matrixBuffer[j][index])[0] - 1
				msgBuffer += string(char)
			}
		}
	}
	decodedMsg := strings.TrimRight(msgBuffer, " ")
	return decodedMsg, nil
}
