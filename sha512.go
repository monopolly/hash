package hash

// SHA512
func SHA512(salt string) (hash *sha512) {
	hash = new(sha512)
	hash.hash = newSHA512(salt)
	return
}

// hashids
type sha512 struct {
	hash *engine
}

func (a *sha512) Encode(any ...interface{}) string {
	return a.hash.Encode(any...)
}

func (a *sha512) Decode(h string) ([]int64, error) {
	return a.hash.Decode(h)
}
