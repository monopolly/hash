package hash

import "strings"

// SHA512
func JWT(salt string) (hash *sha512) {
	hash = new(sha512)
	hash.hash = newSHA512(salt)
	return
}

// hashids
type jwt struct {
	hash *engine
}

func (a *jwt) Encode(any ...interface{}) string {
	p := a.hash.Encode(any...)
	return strings.Join([]string{p[0:40], p[40:80], p[80:]}, ".")
}

func (a *jwt) Decode(h string) ([]int64, error) {
	h = strings.ReplaceAll(h, ".", "")
	return a.hash.Decode(h)
}
