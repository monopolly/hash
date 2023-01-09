package hash

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	hashids "github.com/speps/go-hashids/v2"
)

func Salt(lenght int) (salt string) {
	salt = newHash(time.Now().String(), alphabetBcrypt, lenght).Encode(time.Now().UnixNano(), rand.Int63())
	if len(salt) > lenght {
		return salt[:lenght]
	}
	return
}

func newLower(salt string) (h *engine) {
	return newHash(salt, alphabetLower)
}

func letter() string {
	switch rand.Intn(2) {
	case 1:
		return "a"
	default:
		return "b"
	}
}

func newBcrypt(salt string) (h *engine) {
	//https://www.tunnelsup.com/hash-analyzer/
	return newHash(salt, alphabetBcrypt, 48)
}

func newUpper(salt string) (h *engine) {
	return newHash(salt, alphabetUpper)
}

func newFull(salt string) (h *engine) {
	return newHash(salt, alphabetFull)
}

func newSHA256(salt string) (h *engine) {
	return newHash(salt, "abcdef0123456789", 64)
}

func newSHA512(salt string) (h *engine) {
	return newHash(salt, "abcdef0123456789", 128)
}

func newSHA1(salt string) (h *engine) {
	return newHash(salt, "abcdef0123456789", 40)
}

func newHash(salt, alphabet string, min ...int) (h *engine) {
	mins := 1
	if len(min) > 0 {
		mins = min[0]
	}
	h = new(engine)
	h.data = &hashids.HashIDData{
		Alphabet:  alphabet,
		Salt:      salt,
		MinLength: mins,
	}

	var err error
	h.hash, err = hashids.NewWithData(h.data)
	if err != nil {
		println(err.Error())
	}
	return
}

type engine struct {
	data *hashids.HashIDData
	hash *hashids.HashID
}

func (a *engine) Min(len int) *engine {
	a.data.MinLength = len
	a.hash, _ = hashids.NewWithData(a.data)
	return a
}

func (a *engine) Bcrypt(nums ...interface{}) string {
	return fmt.Sprintf("$2a$10$%s", a.Encode(nums...))
}

func (a *engine) UnBcrypt(h string) ([]int64, error) {
	h = strings.TrimPrefix(h, "$2a$10$")
	return a.Decode(h)
}

func (a *engine) Encode(nums ...interface{}) (h string) {
	var m []int64
	for _, x := range nums {
		switch v := x.(type) {
		case int:
			m = append(m, int64(v))
		case int8:
			m = append(m, int64(v))
		case int16:
			m = append(m, int64(v))
		case int32:
			m = append(m, int64(v))
		case int64:
			m = append(m, int64(v))
		case uint:

			m = append(m, int64(v))
		case uint8:

			m = append(m, int64(v))
		case uint16:

			m = append(m, int64(v))
		case uint32:

			m = append(m, int64(v))
		case uint64:
			m = append(m, int64(v))
		case []uint64:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []uint32:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []uint16:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []uint8:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []uint:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []int64:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []int32:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []int16:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []int8:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case []int:
			for _, x := range v {
				m = append(m, int64(x))
			}
		case string:
			for _, y := range v {
				m = append(m, int64(y))
			}
		case bool:
			switch v {
			case true:
				m = append(m, 1)
			default:
				m = append(m, 0)
			}
		}
	}

	h, _ = a.hash.EncodeInt64(m)
	return
}

func (a *engine) Decode(code string) (d []int64, err error) {
	return a.hash.DecodeInt64WithError(code)
}

func (a *engine) Func() *hashids.HashID {
	return a.hash
}

func (a *engine) EncodeText(email string) (h string) {

	m := make([]int64, len(email))
	for i, x := range email {
		m[i] = int64(x)
	}

	h, _ = a.hash.EncodeInt64(m)
	return
}

func (a *engine) EncodeTextToNums(email string) (m []int64) {
	m = make([]int64, len(email))
	for i, x := range email {
		m[i] = int64(x)
	}
	return
}

func (a *engine) DecodeNumsToText(m []int64) (email string) {
	var b strings.Builder
	for _, x := range m {
		b.WriteRune(rune(x))
	}
	return b.String()
}

func (a *engine) DecodeText(hash string) (email string, err error) {

	list, err := a.Decode(hash)
	if err != nil {
		return
	}
	if len(list) == 0 {
		return
	}

	var s strings.Builder

	for _, x := range list {
		s.WriteRune(rune(x))
	}

	return s.String(), nil
}
