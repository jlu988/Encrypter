package services

import (
	"Encrypter/models"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

var (
	re = regexp.MustCompile("^[0-9]+")
)

type Service struct {
	encryptData    models.EncryptData
	encryptDetails models.EncryptDetails
}

func (s Service) sortKey() string {
	str := strings.Split(s.encryptData.Key, "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func (s Service) keyValidation() error {
	invalidMsg := fmt.Errorf("key should be unique, only letters and digits")
	//Check for any non letters and digits
	key := s.encryptData.Key
	if valid, err := regexp.MatchString("^[a-z0-9A-Z]+$", key); err != nil {
		return err
	} else if !valid {
		return invalidMsg
	}

	if len(s.encryptData.Value) < len(s.encryptData.Key) {
		return fmt.Errorf("message cannot be shorter than key")
	}

	for charPos := 0; charPos < len(key)-1; charPos++ {
		if key[charPos] == key[charPos+1] {
			return invalidMsg
		}
	}
	return nil
}

func (s Service) getMatrixSize() (int, int) {
	columnSize := len(s.encryptData.Key)
	rowSize := int(math.Ceil(float64(len(s.encryptData.Value)) / float64(columnSize)))
	return columnSize, rowSize
}

func (s Service) constructor(obj models.EncryptData) (*models.EncryptDetails, error) {
	s.encryptData = obj
	s.encryptDetails = models.EncryptDetails{}
	s.encryptDetails.ColumnSize, s.encryptDetails.RowSize = s.getMatrixSize()

	if err := s.keyValidation(); err != nil {
		return nil, err
	}

	s.encryptDetails.KeyMap = make(map[byte]int)
	for index, val := range obj.Key {
		s.encryptDetails.KeyMap[byte(val)] = index
		s.encryptDetails.NumericKey = append(s.encryptDetails.NumericKey, index)
	}

	sortedKey := ""
	valid, err := regexp.MatchString("^[0-9]+$", obj.Key)
	if err != nil {
		return nil, err
	}
	if valid {
		lastNumIndex := re.FindAllSubmatchIndex([]byte(s.sortKey()), -1)[0][1]
		stringBuffer := strings.Split(s.sortKey(), "")
		sortedKey = strings.Join(stringBuffer[lastNumIndex:], "") + strings.Join(stringBuffer[:lastNumIndex], "")
	} else {
		sortedKey = s.sortKey()
	}
	for index := 0; index < len(sortedKey); index++ {
		if val, ok := s.encryptDetails.KeyMap[sortedKey[index]]; ok {
			s.encryptDetails.TransPosKey = append(s.encryptDetails.TransPosKey, val)
		}
	}

	return &s.encryptDetails, nil
}

func (s Service) initializeCipherMatrix(msg string, key []int, rowSize, colSize int) map[int][]string {
	keyedCipherMatrix := make(map[int][]string, 0)
	var cipherMatrix [][]string
	message := strings.Split(msg, "")

	valuePosition := 0
	for rowPos := 0; rowPos < rowSize; rowPos++ {
		var matrixRow []string
		for colPos := 0; colPos < colSize; colPos++ {
			if valuePosition != len(msg) {
				matrixRow = append(matrixRow, message[valuePosition])
				valuePosition++
			} else {
				matrixRow = append(matrixRow, " ")
			}
		}
		cipherMatrix = append(cipherMatrix, matrixRow)
	}

	//key pair slice
	for keyIndex := 0; keyIndex < len(key); keyIndex++ {
		var columnSlice []string
		for index := 0; index < len(cipherMatrix); index++ {
			columnSlice = append(columnSlice, cipherMatrix[index][keyIndex])
		}
		keyedCipherMatrix[key[keyIndex]] = columnSlice
	}
	return keyedCipherMatrix
}
