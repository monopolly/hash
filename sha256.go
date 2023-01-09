package hash

// hashids
func SHA256(salt string) (hash *sha256) {
	hash = new(sha256)
	hash.hash = newSHA256(salt)
	return
}

// hashids
type sha256 struct {
	hash *engine
}

func (a *sha256) Encode(any ...interface{}) string {
	return a.hash.Encode(any...)
}

func (a *sha256) Decode(h string) ([]int64, error) {
	return a.hash.Decode(h)
}
