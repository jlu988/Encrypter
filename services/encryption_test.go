package services

import (
	"Encrypter/models"
	"fmt"
	"testing"
)

func TestEncodeFunction(t *testing.T) {
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "secretinformation",
	}

	expectedEncodedMsg := "2e1Cafftdsjpuognjsbuo!p!!"
	mockInternal := "2e1Ca"

	actualEncodedData, _ := Encode(mockData, mockInternal)
	if expectedEncodedMsg != actualEncodedData {
		t.Errorf("expected encoded message: %+v, but got: %+v\n", expectedEncodedMsg, actualEncodedData)
	}
}

func TestEncodeFunction2(t *testing.T) {
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "ecrs",
	}
	mockInternalKey := "2e1Ca"
	_, actualErrorMsg := Encode(mockData, mockInternalKey)
	expectedErrorMsg := fmt.Errorf("message cannot be shorter than key")

	if expectedErrorMsg.Error() != actualErrorMsg.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedErrorMsg, actualErrorMsg)
	}
}
