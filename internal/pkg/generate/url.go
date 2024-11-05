// Package generate contains help funcs to generate some stuff
package generate

import (
	"math/rand"
	"strings"
	"time"
)

var src = rand.NewSource(time.Now().Unix())

const (
	// ASCII runes
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 6 bits to represent a letter index
	letterIdxBits = 6
	// All 1-bits, as many as letterIdxBits
	letterIdxMask = 1<<letterIdxBits - 1
	// # of letter indices fitting in 63 bits
	letterIdxMax = 63 / letterIdxBits
)

// String generates a random ascii string with fixed length
func String(length int) string {
	builder := new(strings.Builder)
	builder.Grow(length)

	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letters) {
			builder.WriteByte(letters[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return builder.String()
}
