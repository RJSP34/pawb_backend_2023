package service

import "encoding/base64"

func DecodeBase64(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

func EncodeToBase64(decoded []byte) string {
	return base64.StdEncoding.EncodeToString(decoded)
}
