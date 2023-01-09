package hash

import (
	"strings"
)

// Email — sha1, sha256, sha512, bcrypt
func Email(salt string, types string) (hash *email) {

	hash = new(email)
	switch strings.ToLower(types) {
	case "sha1":
		hash.hash = newSHA1(salt)
	case "sha256":
		hash.hash = newSHA256(salt)
	case "sha512":
		hash.hash = newSHA512(salt)
	case "bcrypt":
		hash.hash = newBcrypt(salt)
	case "lower":
		hash.hash = newLower(salt)
	default:
		hash.hash = newFull(salt)
	}
	return
}

// hashids
type email struct {
	hash *engine
}

func (a *email) Encode(email string) (hash string) {
	return a.hash.EncodeText(email)
}

func (a *email) Decode(hash string) (email string, err error) {
	return a.hash.DecodeText(hash)
}
