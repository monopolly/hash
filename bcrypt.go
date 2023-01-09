package hash

func Bcrypt(salt string) (hash *bcrypt) {
	hash = new(bcrypt)
	hash.hash = newBcrypt(salt)
	return
}

// hashids
type bcrypt struct {
	hash *engine
}

func (a *bcrypt) Encode(any ...interface{}) string {
	return a.hash.Bcrypt(any...)
}

func (a *bcrypt) Decode(h string) ([]int64, error) {
	return a.hash.UnBcrypt(h)
}
