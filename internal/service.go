package internal

import (
	"math/rand"
	"strings"
	"time"
)

func shuffleString(strBuffer []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(strBuffer), func(i, j int) {
		strBuffer[i], strBuffer[j] = strBuffer[j], strBuffer[i]
	})
	return strBuffer
}

func PrivateKey(key string) string {
	stringBuffer := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	for _, char := range key {
		stringBuffer = strings.ReplaceAll(stringBuffer, string(char), "")
	}

	internalKey := shuffleString([]byte(stringBuffer))[0:8]

	return string(internalKey)
}
