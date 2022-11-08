package domain

import (
	"Encrypter/database"
	"Encrypter/internal"
	"Encrypter/models"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

type EncryptionHandler struct{}

var re = regexp.MustCompile("^[0-9]+")

func sortString(message string) string {
	s := strings.Split(message, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func validateKey(key []byte) error {
	valid, err := regexp.MatchString("^[a-z0-9A-Z]+$", string(key))
	errMsg := fmt.Errorf("invalid key - should be unique, only letters and digits")
	if err != nil {
		return err
	}
	if !valid {
		return errMsg
	}

	for index := 0; index < len(key)-1; index++ {
		if key[index] == key[index+1] {
			return errMsg
		}
	}
	return nil
}

func (h *EncryptionHandler) constructor(obj models.EncryptData) (*models.EncryptDetails, error) {
	data := models.EncryptDetails{}
	data.ColumnSize = len(obj.Key)
	data.RowSize = int(math.Ceil(float64(len(obj.Value)) / float64(data.ColumnSize)))

	keyString := sortString(obj.Key)
	if err := validateKey([]byte(keyString)); err != nil {
		return nil, err
	}

	if len(obj.Value) < len(obj.Key) {
		return nil, fmt.Errorf("message cannot be shorter than key")
	}

	data.OriginalKey = make(map[byte]int)
	for index, val := range obj.Key {
		data.OriginalKey[byte(val)] = index
		data.NumericKey = append(data.NumericKey, index)
	}

	numIndex := re.FindAllSubmatchIndex([]byte(keyString), -1)[0][1]

	buffer := strings.Split(keyString, "")
	sortedKey := strings.Join(buffer[numIndex:], "") + strings.Join(buffer[:numIndex], "")

	for index := 0; index < len(sortedKey); index++ {
		if val, ok := data.OriginalKey[sortedKey[index]]; ok {
			data.TransPosKey = append(data.TransPosKey, val)
		}
	}
	return &data, nil
}

func (h *EncryptionHandler) initializeCipherMatrix(msg string, key []int, row, col int) map[int][]string {
	keyedCipherMatrix := make(map[int][]string, 0)
	var cipherMatrix [][]string
	message := strings.Split(msg, "")

	valuePosition := 0
	for index := 0; index < row; index++ {
		var matrixRow []string
		for j := 0; j < col; j++ {
			if valuePosition != len(msg) {
				matrixRow = append(matrixRow, message[valuePosition])
				valuePosition++
			} else {
				matrixRow = append(matrixRow, " ")
			}
		}
		cipherMatrix = append(cipherMatrix, matrixRow)
	}

	for keyIndex := 0; keyIndex < len(key); keyIndex++ {
		var columnSlice []string
		for index := 0; index < len(cipherMatrix); index++ {
			columnSlice = append(columnSlice, cipherMatrix[index][keyIndex])
		}
		keyedCipherMatrix[key[keyIndex]] = columnSlice
	}

	return keyedCipherMatrix
}

func (h *EncryptionHandler) Encode(obj models.EncryptData) (string, error) {
	data, err := h.constructor(obj)
	if err != nil {
		return "", err
	}

	cipherMatrix := h.initializeCipherMatrix(obj.Value, data.TransPosKey, data.RowSize, data.ColumnSize)
	//Get internal private key
	privateKey := internal.PrivateKey(obj.Key)
	keys := models.Internal{
		OriginalKey: obj.Key,
		InternalKey: privateKey,
	}

	encodedMessage := ""
	for index := 0; index < data.RowSize; index++ {
		for col := 0; col < len(cipherMatrix); col++ {
			char := []byte(cipherMatrix[col][index])[0] + 1
			encodedMessage += string(char)
		}
	}

	dbHandler := database.DbHandler{}
	dbHandler.AddInternalKey(keys)
	encodedMessage = privateKey + encodedMessage
	return encodedMessage, nil
}

func (h *EncryptionHandler) Decode(obj models.EncryptData) (string, error) {
	data, err := h.constructor(obj)
	if err != nil {
		println("error occurred - ", err)
		return "", err
	}

	dbHandler := database.DbHandler{}
	internalPrivateKey, err := dbHandler.GetInternalKey(obj.Key)
	if err != nil {
		return "", err
	}

	if internalPrivateKey == nil {
		return obj.Value, nil
	}

	privateKey := obj.Value[0:8]
	if privateKey != internalPrivateKey.InternalKey {
		return obj.Value, nil
	} else {
		obj.Value = obj.Value[8:]
	}

	cipheredMatrix := h.initializeCipherMatrix(obj.Value, data.NumericKey, data.RowSize, data.ColumnSize)

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
