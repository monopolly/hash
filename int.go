package hash

import (
	"fmt"

	"github.com/OneOfOne/xxhash"
)

// Uint xxhash
func Uint(id interface{}) (hash uint) {
	return uint(Uint64(id))
}

// Int xxhash
func Int(id interface{}) (hash int) {
	switch id.(type) {
	case []byte:
		return int(xxhash.Checksum32(id.([]byte)))
	case string:
		return int(xxhash.Checksum32([]byte(id.(string))))
	default:
		return int(xxhash.Checksum32([]byte(fmt.Sprint(id))))
	}
}

// Uint32 xxhash
func Uint32(id interface{}) uint32 {
	switch id.(type) {
	case []byte:
		return xxhash.Checksum32(id.([]byte))
	case string:
		return xxhash.Checksum32([]byte(id.(string)))
	default:
		return xxhash.Checksum32([]byte(fmt.Sprint(id)))
	}
}

// Uint64 xxhash
func Uint64(id interface{}) uint64 {
	switch id.(type) {
	case []byte:
		return xxhash.Checksum64(id.([]byte))
	case string:
		return xxhash.Checksum64([]byte(id.(string)))
	default:
		return xxhash.Checksum64([]byte(fmt.Sprint(id)))
	}
}
