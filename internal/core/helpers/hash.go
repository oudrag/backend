package helpers

import (
	"crypto/sha1"
	"fmt"
)

func HashSHA1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	bh := hash.Sum(nil)

	return fmt.Sprintf("%x", bh)
}
