package internal

import (
	"math/rand"
	"strings"
	"time"
)

func randomizeString(str []byte) []byte {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(str), func(i, j int) {
		str[i], str[j] = str[j], str[i]
	})
	return str
}

func GeneratePrivateKey(key string) string {
	stringBuffer := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for _, char := range key {
		stringBuffer = strings.ReplaceAll(stringBuffer, string(char), "")
	}
	keySize := len(key)
	newKey := string(randomizeString([]byte(stringBuffer))[0:keySize])

	return newKey
}
