package domain

import (
	"Encrypter/models"
	"fmt"
	"testing"
)

func compareMap(a, b interface{}) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestConstructor(t *testing.T) {
	encrypter := EncryptionHandler{}

	mockData := &models.EncryptDetails{
		OriginalKey: map[byte]int{
			49:  2,
			50:  0,
			56:  4,
			98:  3,
			101: 1,
		},
		NumericKey:  []int{0, 1, 2, 3, 4},
		TransPosKey: []int{3, 1, 2, 0, 4},
		ColumnSize:  5,
		RowSize:     4,
	}

	data := models.EncryptData{
		Key:   "2e1b8",
		Value: "secretinformation",
	}

	expectedErrorMsg := fmt.Errorf("message cannot be shorter than key")
	actualData, actualErrorMsg := encrypter.constructor(data)
	if actualErrorMsg != nil {
		if expectedErrorMsg.Error() != actualErrorMsg.Error() {
			t.Errorf("expected error message: %+v, but got: %+v\n", expectedErrorMsg, actualErrorMsg)
		}
	}

	if !compareMap(mockData, actualData) {
		t.Errorf("\nexpected: %+v,\ngot: %+v", mockData, actualData)
	}
}

func TestConstructor2(t *testing.T) {
	encrypter := EncryptionHandler{}

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
	encrypter := EncryptionHandler{}
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
	encrypter := EncryptionHandler{}
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

func TestEncodeFunction(t *testing.T) {
	encryptionService := EncryptionHandler{}
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "secretinformation",
	}

	expectedEncodedMsg := "ecrseonftiiatrm   on"
	actualEncodedData, _ := encryptionService.Encode(mockData)
	if expectedEncodedMsg != actualEncodedData {
		t.Errorf("expected encoded message: %+v, but got: %+v\n", expectedEncodedMsg, actualEncodedData)
	}
}

func TestDecodeFunction(t *testing.T) {
	encryptionService := EncryptionHandler{}
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "ecrseonftiiatrm   on",
	}

	expectedDecodedMsg := "secretinformation"
	actualDecodedData, _ := encryptionService.Decode(mockData)
	if expectedDecodedMsg != actualDecodedData {
		t.Errorf("expected encoded message: %+v, but got: %+v\n", expectedDecodedMsg, actualDecodedData)
	}
}

func TestDecodeFunction2(t *testing.T) {
	encryptionService := EncryptionHandler{}
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "ecrs",
	}

	_, actualErrorMsg := encryptionService.Decode(mockData)
	expectedErrorMsg := fmt.Errorf("message cannot be shorter than key")

	if expectedErrorMsg.Error() != actualErrorMsg.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedErrorMsg, actualErrorMsg)
	}
}

func TestEncodeFunction2(t *testing.T) {
	encryptionService := EncryptionHandler{}
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "ecrs",
	}

	_, actualErrorMsg := encryptionService.Encode(mockData)
	expectedErrorMsg := fmt.Errorf("message cannot be shorter than key")

	if expectedErrorMsg.Error() != actualErrorMsg.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedErrorMsg, actualErrorMsg)
	}
}

func TestValidateKey(t *testing.T) {
	mockData := models.EncryptData{
		Key: "2e1Ca",
	}

	acutalError := validateKey([]byte(mockData.Key))
	var expectedError interface{} = nil

	if acutalError != expectedError {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, acutalError)
	}
}

func TestValidateKey2(t *testing.T) {
	mockData := models.EncryptData{
		Key: "2e1Ca#^!",
	}

	acutalError := validateKey([]byte(mockData.Key))
	expectedError := fmt.Errorf("invalid key - should be unique, only letters and digits")

	if acutalError.Error() != expectedError.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, acutalError)
	}
}

func TestValidateKey3(t *testing.T) {
	mockData := models.EncryptData{
		Key: "2e1Caa",
	}

	acutalError := validateKey([]byte(mockData.Key))
	expectedError := fmt.Errorf("invalid key - should be unique, only letters and digits")

	if acutalError.Error() != expectedError.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedError, acutalError)
	}
}
