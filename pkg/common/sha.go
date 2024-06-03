package common

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func EncryptSha(bytes []byte) [20]byte {
	return sha1.Sum(bytes)
}

func FileSha(bytes []byte) string {
	fmt.Println(string(bytes))
	sha := EncryptSha(bytes)
	return hex.EncodeToString(sha[:])
}
