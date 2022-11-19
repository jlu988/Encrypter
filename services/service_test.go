package services

import (
	"Encrypter/models"
	"fmt"
	"testing"
)

func compareMap(a, b interface{}) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestConstructor(t *testing.T) {
	encrypter := Service{}
	mockData := &models.EncryptDetails{
		KeyMap: map[byte]int{
			49:  2,
			50:  0,
			67:  3,
			97:  4,
			101: 1,
		},
		NumericKey:  []int{0, 1, 2, 3, 4},
		TransPosKey: []int{2, 0, 3, 4, 1},
		ColumnSize:  5,
		RowSize:     4,
	}

	data := models.EncryptData{
		Key:   "2e1Ca",
		Value: "secretinformation",
	}
	actualData, _ := encrypter.constructor(data)

	if !compareMap(mockData, actualData) {
		t.Errorf("\nexpected: %+v,\ngot: %+v", mockData, actualData)
	}
}

func TestConstructor2(t *testing.T) {
	encrypter := Service{}
	data := models.EncryptData{
		Key:   "2e1b8",
		Value: "secr",
	}
	expectedErrorMsg := fmt.Errorf("message cannot be shorter than key")
	_, actualErrorMsg := encrypter.constructor(data)
	if actualErrorMsg != nil && expectedErrorMsg.Error() != actualErrorMsg.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedErrorMsg, actualErrorMsg)
	}
}

func TestInitialCipherMatrixForEncoding(t *testing.T) {
	encrypter := Service{}
	mockData := models.EncryptData{
		Key:   "2e1Cab",
		Value: "secretinformation",
	}
	data := models.EncryptDetails{
		TransPosKey: []int{4, 5, 2, 3, 1},
		ColumnSize:  5,
		RowSize:     4,
	}
	expectMockData := map[int][]string{
		1: {"e", "o", "i", " "},
		2: {"c", "n", "a", " "},
		3: {"r", "f", "t", " "},
		4: {"s", "t", "r", "o"},
		5: {"e", "i", "m", "n"},
	}

	actualData := encrypter.initializeCipherMatrix(mockData.Value, data.TransPosKey, data.RowSize, data.ColumnSize)

	if !compareMap(expectMockData, actualData) {
		t.Errorf("\nexpected: %+v,\ngot: %+v", expectMockData, actualData)
	}
}

func TestInitialCipherMatrixForDecoding(t *testing.T) {
	encrypter := Service{}
	mockData := models.EncryptData{
		Key:   "2e1Cab",
		Value: "ecrseonftiiatrm   on",
	}

	data := models.EncryptDetails{
		NumericKey: []int{1, 2, 3, 4, 5},
		ColumnSize: 5,
		RowSize:    4,
	}

	expectMockData := map[int][]string{
		1: {"e", "o", "i", " "},
		2: {"c", "n", "a", " "},
		3: {"r", "f", "t", " "},
		4: {"s", "t", "r", "o"},
		5: {"e", "i", "m", "n"},
	}

	actualData := encrypter.initializeCipherMatrix(mockData.Value, data.NumericKey, data.RowSize, data.ColumnSize)
	if !compareMap(expectMockData, actualData) {
		t.Errorf("\nexpected: %+v,\ngot: %+v", expectMockData, actualData)
	}
}

func TestValidateKey(t *testing.T) {
	mockData := Service{
		encryptData: models.EncryptData{
			Key: "2e1Ca",
		},
	}

	actualError := mockData.keyValidation()
	var expectedError interface{} = nil

	if actualError != expectedError {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, actualError)
	}
}

func TestValidateKey2(t *testing.T) {
	mockData := Service{
		encryptData: models.EncryptData{
			Key: "2e1Ca#^!",
		},
	}

	actualError := mockData.keyValidation()
	expectedError := fmt.Errorf("invalid key - should be unique, only letters and digits")

	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, actualError)
	}
}

func TestValidateKey3(t *testing.T) {
	mockData := Service{
		encryptData: models.EncryptData{
			Key: "2e1Caa!",
		},
	}

	actualError := mockData.keyValidation()
	expectedError := fmt.Errorf("invalid key - should be unique, only letters and digits")

	if actualError.Error() != expectedError.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, actualError)
	}
}

func TestValidateKey4(t *testing.T) {
	mockData := Service{
		encryptData: models.EncryptData{
			Key: "1234567",
		},
	}

	actualError := mockData.keyValidation()
	expectedError := fmt.Errorf("invalid key - should be unique, only letters and digits")

	if actualError.Error() == expectedError.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, actualError)
	}
}

func TestValidateKey5(t *testing.T) {
	mockData := Service{
		encryptData: models.EncryptData{
			Key: "ArwWfVz",
		},
	}

	actualError := mockData.keyValidation()
	expectedError := fmt.Errorf("invalid key - should be unique, only letters and digits")

	if actualError.Error() == expectedError.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, actualError)
	}
}
