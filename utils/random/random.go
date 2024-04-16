package random

import (
	"math/rand"
	"sync"
	"time"
)

var (
	r  *rand.Rand
	rl *sync.Mutex
)

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
	rl = &sync.Mutex{}
}

func Choice[T any](s []T) T {
	rl.Lock()
	defer rl.Unlock()
	l := len(s)
	return s[r.Intn(l)]
}

func RandInt(i, j int) int {
	rl.Lock()
	defer rl.Unlock()
	return i + r.Intn(j-i)
}

func Random() float64 {
	rl.Lock()
	defer rl.Unlock()
	return r.Float64()
}
