package hash

// hashids
func Full(salt string, min ...int) (hash *full) {
	hash = new(full)
	hash.hash = newFull(salt)
	if len(min) > 0 {
		hash.hash.Min(min[0])
	}
	return
}

// hashids
type full struct {
	hash *engine
}

func (a *full) Encode(any ...interface{}) string {
	return a.hash.Encode(any...)
}

func (a *full) Decode(h string) ([]int64, error) {
	return a.hash.Decode(h)
}
