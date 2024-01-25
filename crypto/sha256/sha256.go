package sha256

import (
	"crypto/sha256"
	"fmt"
)

const Salt string = "*$salt@*"

// Sha 算签名
func New(text string) string {
	hash := sha256.New()
	text = Salt + text + Salt
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
