package services

import (
	"Encrypter/models"
	"fmt"
	"testing"
)

func TestDecodeFunction(t *testing.T) {
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "2e1Cafftdsjpuognjsbuo!p!!",
	}
	mockInternalKey := "2e1Ca"
	expectedDecodedMsg := "secretinformation"
	actualDecodedData, _ := Decrypt(mockData, mockInternalKey)
	if expectedDecodedMsg != actualDecodedData {
		t.Errorf("expected decoded message: %+v, but got: %+v\n", expectedDecodedMsg, actualDecodedData)
	}
}

func TestDecodeFunction2(t *testing.T) {
	mockData := models.EncryptData{
		Key:   "2e1Ca",
		Value: "ecrs",
	}
	mockInternalKey := "2e1Ca"
	_, actualErrorMsg := Decrypt(mockData, mockInternalKey)
	expectedErrorMsg := fmt.Errorf("message cannot be shorter than key")

	if expectedErrorMsg.Error() != actualErrorMsg.Error() {
		t.Errorf("expected error message: %+v, but got: %+v\n", expectedErrorMsg, actualErrorMsg)
	}
}
