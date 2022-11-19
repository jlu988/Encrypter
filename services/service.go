package services

import (
	"Encrypter/models"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

type EncryptionServices struct{}

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

func constructor(obj models.EncryptData) (*models.EncryptDetails, error) {
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

func initializeCipherMatrix(msg string, key []int, row, col int) map[int][]string {
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
