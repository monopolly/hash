package hash

// hashids
func SHA1(salt string, min ...int) (hash *sha1) {
	hash = new(sha1)
	hash.hash = newSHA1(salt)
	if len(min) > 0 {
		hash.hash.Min(min[0])
	}
	return
}

// hashids
type sha1 struct {
	hash *engine
}

func (a *sha1) Encode(any ...interface{}) string {
	return a.hash.Encode(any...)
}

func (a *sha1) Decode(h string) ([]int64, error) {
	return a.hash.Decode(h)
}
