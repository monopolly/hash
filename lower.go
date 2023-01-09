package hash

// hashids
func Lower(salt string) (hash *lower) {
	hash = new(lower)
	hash.hash = newLower(salt)
	return
}

// hashids
type lower struct {
	hash *engine
}

func (a *lower) Min(len int) *lower {
	a.hash.Min(len)
	return a
}

func (a *lower) Encode(any ...interface{}) string {
	return a.hash.Encode(any...)
}

func (a *lower) Decode(h string) ([]int64, error) {
	return a.hash.Decode(h)
}
