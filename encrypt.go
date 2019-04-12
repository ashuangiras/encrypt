package encrypt

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"log"
)

// EncryptData : Encrypt the values with special singing secret known only to
// auth-proxy, can be changed through environment variables
func EncryptData(value string, secret []byte) string {

	log.Println("Encrypting the components : ")
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(value))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
