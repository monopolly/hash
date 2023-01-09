package hash

// hashids
func Upper(salt string, min ...int) (hash *upper) {
	hash = new(upper)
	hash.hash = newUpper(salt)
	if min != nil {
		hash.hash.Min(min[0])
	}
	return
}

// hashids
type upper struct {
	hash *engine
}

func (a *upper) Encode(any ...interface{}) string {
	return a.hash.Encode(any...)
}

func (a *upper) Decode(h string) ([]int64, error) {
	return a.hash.Decode(h)
}
