package helpers

import "crypto/sha1"

func HashSHA1(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))
	bh := hash.Sum(nil)

	return string(bh)
}
