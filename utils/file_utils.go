package utils

import (
	"crypto/rand"
	"os"
)

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	rand.Read(b)
	for i := range b {
		b[i] = letters[b[i]%byte(len(letters))]
	}
	return string(b)
}

func SafeDeleteFile(path string) {
	if path == "" {
		return
	}
	if _, err := os.Stat(path); err == nil {
		os.Remove(path)
	}
}

