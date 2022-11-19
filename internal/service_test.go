package internal

import (
	"testing"
)

func TestDecodeFunction2(t *testing.T) {
	mockKey := "2e1Ca"
	actualKey := GeneratePrivateKey(mockKey)

	if len(mockKey) != len(actualKey) {
		t.Errorf("expected key size: %+v, but got: %+v\n", len(mockKey), len(actualKey))
	}
}
