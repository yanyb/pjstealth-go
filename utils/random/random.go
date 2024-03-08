package random

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

func Choice[T any](s []T) T {
	l := len(s)
	return s[r.Intn(l)]
}

func RandInt(i, j int) int {
	return i + r.Intn(j-i)
}

func Random() float64 {
	return r.Float64()
}
