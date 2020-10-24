package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// GenURLCode generates short code as url id for redirecting.
func GenURLCode() string {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	r := rand.New(source)

	b := make([]byte, 4)
	r.Read(b)
	s := hex.EncodeToString(b)
	return s
}
